package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleAccountCreate ...
func (s *Server) HandleNucypherAccountCreate(c *gin.Context) {
	var a *model.NucypherAccount
	if err := c.BindJSON(&a); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()
	// a.CreatedAt = now

	a.CreatedBy = c.Value("userID").(string)
	if err := s.store.NucypherAccount().Create(a, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"organization_id": a.OrganizationID,
		"address":         a.Address,
		"verifying_key":   a.VerifyingKey,
		"balance":         a.Balance,
		"tokens":          a.Tokens,
		"is_active":       a.IsActive,
		"is_private":      a.IsPrivate,
		"created_by":      a.CreatedBy,
		"updated_by":      a.UpdatedBy,
	})
}

// HandleGetNucypherAccounts returns all tezos accounts for user
func (s *Server) HandleGetNucypherAccounts(c *gin.Context) {
	createdBy := c.Value("userID").(string)
	accounts, err := s.store.NucypherAccount().GetAccounts(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, accounts)
}

// HandleUpdateNucypherAccount allow user update account name
func (s *Server) HandleUpdateNucypherAccount(c *gin.Context) {
	type repsonse struct {
		Name      string    `json:"name"`
		AccountID int       `json:"account_id"`
		UpdatedBy string    `json:"updated_by"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	var a *repsonse

	if err := c.BindJSON(&a); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	a.UpdatedAt = now
	a.UpdatedBy = c.Value("userID").(string)

	if err := s.store.NucypherAccount().UpdateName(a.Name, a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":       a.Name,
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleDeactivateNucypherAccount deactivates accounts
func (s *Server) HandleDeactivateNucypherAccount(c *gin.Context) {
	type response struct {
		AccountID int       `json:"account_id"`
		CreatedBy string    `json:"created_by"`
		UpdatedBy string    `json:"updated_by"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	var a *response

	if err := c.BindJSON(&a); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	a.UpdatedAt = now
	a.UpdatedBy = c.Value("userID").(string)

	if a.CreatedBy != a.UpdatedBy {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	if err := s.store.NucypherAccount().Deactivate(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleReactivateNucypherAccount activates accounts
func (s *Server) HandleReactivateNucypherAccount(c *gin.Context) {
	type response struct {
		AccountID int       `json:"account_id"`
		CreatedBy string    `json:"created_by"`
		UpdatedBy string    `json:"updated_by"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	var a *response

	if err := c.BindJSON(&a); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	a.UpdatedAt = now
	a.UpdatedBy = c.Value("userID").(string)

	if a.CreatedBy != a.UpdatedBy {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	if err := s.store.NucypherAccount().Reactivate(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleMakeNucypherAccountPrivate makes tezos accounts private
func (s *Server) HandleMakeNucypherAccountPrivate(c *gin.Context) {
	type response struct {
		AccountID int       `json:"account_id"`
		CreatedBy string    `json:"created_by"`
		UpdatedBy string    `json:"updated_by"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	var a *response

	if err := c.BindJSON(&a); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	a.UpdatedAt = now
	a.UpdatedBy = c.Value("userID").(string)

	if a.CreatedBy != a.UpdatedBy {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	if err := s.store.NucypherAccount().Private(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleMakeNucypherAccountUnprivate makes tezos accounts public
func (s *Server) HandleMakeNucypherAccountUnprivate(c *gin.Context) {
	type response struct {
		AccountID int       `json:"account_id"`
		CreatedBy string    `json:"created_by"`
		UpdatedBy string    `json:"updated_by"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	var a *response

	if err := c.BindJSON(&a); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	a.UpdatedAt = now
	a.UpdatedBy = c.Value("userID").(string)

	if a.CreatedBy != a.UpdatedBy {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	if err := s.store.NucypherAccount().Unprivate(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}
