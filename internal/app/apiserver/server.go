package apiserver

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/bahadylbekov/vacinex_api/internal/app/model"
	"github.com/bahadylbekov/vacinex_api/internal/app/store"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	sessionName        = "go"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

var (
	errIncorrectEmailOrPassword = "Incorrect email or password"
	errInternalServerError      = "Internal server error"
	errNotAuthenticated         = "Not authenticated"
	errBadRequest               = "Bad request"
	errBearerFormat             = "Bearer token not in proper format"
)

// Server ...
type Server struct {
	router        *gin.Engine
	logger        *logrus.Logger
	store         store.Store
	sessionStore  sessions.Store
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	IdleTimeout   time.Duration
	TLSConfig     *tls.Config
	jwtMiddleware *jwtmiddleware.JWTMiddleware
}

type ctxKey int8

// NewServer ...
func NewServer(store store.Store, sessionStore sessions.Store) *Server {

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			// aud := "http://localhost:8000"
			// checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			// if !checkAud {
			// 	return token, errors.New("Invalid audience")
			// }
			// Verify 'iss' claim
			iss := "https://vacinex.auth0.com/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer")
			}

			cert, err := getPemCert(token)
			if err != nil {
				panic(err.Error())
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})

	tlsConfig := &tls.Config{
		// Causes servers to use Go's default cipher suite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},

		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	s := &Server{
		router:        gin.Default(),
		logger:        logrus.New(),
		store:         store,
		sessionStore:  sessionStore,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   120 * time.Second,
		TLSConfig:     tlsConfig,
		jwtMiddleware: jwtMiddleware,
	}

	s.configureRouter()

	return s
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// configureRouter ..
func (s *Server) configureRouter() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://vacinex.io", "http://localhost:8080"}
	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Length", "Content-Type"}

	s.router.Use(s.SetRequestID())
	s.router.Use(s.logRequest())
	s.router.Use(cors.New(config))
	s.router.POST("/users", s.handleUsersCreate)

	private := s.router.Group("/private")
	private.Use(s.authMiddleware())
	{
		private.GET("/whoami", s.getMyUserID)
		private.POST("/profile", s.HandleProfileCreate)
		private.GET("/profile", s.HandleGetMyProfile)
		private.PUT("/profile", s.HandleUpdateProfile)
	}
}

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client secret key
		err := s.jwtMiddleware.CheckJWT(c.Writer, c.Request)
		if err != nil {
			// Token not found
			fmt.Println(err)
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Unauthorized"))
			return
		}
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")

		if len(splitToken) != 2 {
			respondWithError(c, http.StatusUnauthorized, errBearerFormat)
		}

		reqToken = strings.TrimSpace(splitToken[1])
		userID := extractIDFromToken(reqToken)
		c.Set("userID", userID)
	}
}

func extractIDFromToken(tokenStr string) interface{} {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		subClaim := claims["sub"]
		return subClaim
	} else {
		log.Printf("Invalid JWT Token")
		return nil
	}
}

// handleUsersCreate ...
func (s *Server) handleUsersCreate(c *gin.Context) {
	var u *model.User
	c.BindJSON(&u)

	if err := s.store.User().Create(u); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	u.Sanitize()
	c.JSON(http.StatusOK, gin.H{
		"email":              u.Email,
		"encrypted_password": u.EncryptedPassword,
	})
}

func (s *Server) logRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": c.Request.RemoteAddr,
			"request_id":  c.Value("ctxKeyRequestID"),
		})
		logger.Infof("started %s %s", c.Request.Method, c.Request.RequestURI)
		start := time.Now()
		c.Next()

		var level logrus.Level
		switch {
		case c.Writer.Status() >= 500:
			level = logrus.ErrorLevel
		case c.Writer.Status() >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			c.Writer.Status(),
			http.StatusText(c.Writer.Status()),
			time.Now().Sub(start),
		)
	}
}

// SetRequestID ...
func (s *Server) SetRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()
		c.Header("X-Request-ID", id)
		c.Set("ctxKeyRequestID", id)
		c.Next()
	}
}

// handleUsersCreate ...
func (s *Server) getMyUserID(c *gin.Context) {
	c.JSON(http.StatusOK, c.Value("userID"))
}

// respondWithError ...
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
