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
	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/bahadylbekov/vaccinex_api/internal/app/store"
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
			aud := "http://localhost:8000"
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAud {
				return token, errors.New("Invalid audience")
			}
			// Verify 'iss' claim
			iss := "https://vaccinex.us.auth0.com/"
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
	config.AllowOrigins = []string{"*", "http://localhost:4200"}
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

		private.GET("/organizations", s.HandleGetOrganizations)
		private.DELETE("/organizations", s.HandleDeleteOrganization)
		private.GET("/organizations/:id", s.HandleGetOrganization)
		private.GET("invite", s.HandleAddOrganizationToMyList)
		private.GET("/search", s.HandleFindOrganizations)

		private.GET("/viruses", s.HandleGetViruses)
		private.GET("/viruses/:id", s.HandleGetVirusByID)
		private.POST("/viruses", s.HandleVirusCreate)
		private.PUT("/viruses", s.HandleUpdateVirus)

		private.GET("/genomes", s.HandleGetGenomes)
		private.GET("/genome", s.HandleGetMyGenomes)
		private.POST("/genomes", s.HandleGenomeCreate)
		private.GET("/genomes/virus/:id", s.HandleGetGenomesByVirus)
		private.GET("/genomes/organization/:id", s.HandleGetGenomesByOrganization)
		private.GET("/genomes/vaccines/:id", s.HandleGetGenomesByVaccine)

		private.POST("/transactions", s.HandleTransactionCreate)
		private.GET("/transactions", s.HandleGetTransactions)
		private.GET("/transactions/send", s.HandleGetSendTransactions)
		private.GET("/transactions/recieved", s.HandleGetRecievedTransactions)

		private.POST("/accounts/ethereum", s.HandleEthereumAccountCreate)
		private.GET("/accounts/ethereum", s.HandleGetEthereumAccounts)
		private.GET("/organization/accounts/ethereum/:id", s.HandleGetEthereumAccountForOrganization)
		private.PUT("/accounts/ethereum", s.HandleUpdateEthereumAccount)
		private.PUT("/accounts/ethereum/address", s.HandleUpdateEthereumAccountAddress)
		private.POST("/accounts/ethereum/deactivate", s.HandleDeactivateEthereumAccount)
		private.POST("/accounts/ethereum/reactivate", s.HandleReactivateEthereumAccount)
		private.POST("/accounts/ethereum/private", s.HandleMakeEthereumAccountPrivate)
		private.POST("/accounts/ethereum/unprivate", s.HandleMakeEthereumAccountUnprivate)

		private.POST("/accounts/nucypher", s.HandleNucypherAccountCreate)
		private.GET("/accounts/nucypher", s.HandleGetNucypherAccounts)
		private.GET("/organization/accounts/nucypher/:id", s.HandleGetNucypherAccountForOrganization)
		private.PUT("/accounts/nucypher", s.HandleUpdateNucypherAccount)
		private.PUT("/accounts/nucypher/address", s.HandleUpdateNucypherAccountAddress)
		private.PUT("/accounts/nucypher/verifyingkey", s.HandleUpdateNucypherAccountVerifyingKey)
		private.POST("/accounts/nucypher/deactivate", s.HandleDeactivateNucypherAccount)
		private.POST("/accounts/nucypher/reactivate", s.HandleReactivateNucypherAccount)
		private.POST("/accounts/nucypher/private", s.HandleMakeNucypherAccountPrivate)
		private.POST("/accounts/nucypher/unprivate", s.HandleMakeNucypherAccountUnprivate)

		private.GET("/vaccines", s.HandleGetVaccines)
		private.GET("/vaccines/:id", s.HandleGetVaccineByID)
		private.POST("/vaccines", s.HandleCreateVaccine)
		private.PUT("/vaccines/name", s.HandleUpdateVaccineName)
		private.PUT("/vaccines/description", s.HandleUpdateVaccineDescription)
		private.PUT("/vaccines/amount", s.HandleUpdateVaccineAmount)

		private.POST("/policies", s.HandleCreatePolicy)
		private.GET("/policies/:id", s.HandleGetPolicyByID)

		private.POST("/receipts", s.HandleCreateReceipt)
		private.GET("/receipts/:id", s.HandleGetReceiptByID)

		private.POST("/grants", s.HandleCreateGrant)
		private.GET("/grants/:id", s.HandleGetGrantsForMe)
		private.GET("/completed/grants", s.HandleGetCompletedGrantsForMe)
		private.PUT("/grants", s.HandleSubmitGrant)
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
			_, err := c.Writer.Write([]byte("Unauthorized"))
			if err != nil {
				fmt.Println(err)
				return
			}
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
	if err := c.BindJSON(&u); err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}

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
			time.Since(start),
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
