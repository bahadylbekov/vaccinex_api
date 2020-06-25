package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleAccountCreate ...
func (s *Server) HandleTezosAccountCreate(c *gin.Context) {
	var a *model.TezosAccount
	if err := c.BindJSON(&a); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()
	// a.CreatedAt = now

	a.CreatedBy = c.Value("userID").(string)
	if err := s.store.TezosAccount().Create(a, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"organization_id": a.OrganizationID,
		"address":         a.Address,
		"balance":         a.Balance,
		"tokens":          a.Tokens,
		"is_active":       a.IsActive,
		"is_private":      a.IsPrivate,
		"created_by":      a.CreatedBy,
		"updated_by":      a.UpdatedBy,
	})
}

// HandleGetTezosAccounts returns all tezos accounts for user
func (s *Server) HandleGetTezosAccounts(c *gin.Context) {
	createdBy := c.Value("userID").(string)
	accounts, err := s.store.TezosAccount().GetAccounts(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, accounts)
}

// HandleGetTezosAccounts returns all tezos accounts for user
func (s *Server) HandleGetTezosAccountForOrganization(c *gin.Context) {
	id := c.Param("id")
	accounts, err := s.store.TezosAccount().GetAccountOrganization(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, accounts)
}

// HandleUpdateTezosAccount allow user update account name
func (s *Server) HandleUpdateTezosAccount(c *gin.Context) {
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

	if err := s.store.TezosAccount().UpdateName(a.Name, a.UpdatedBy, a.AccountID, now); err != nil {
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

// HandleDeactivateTezosAccount deactivates accounts
func (s *Server) HandleDeactivateTezosAccount(c *gin.Context) {
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

	if err := s.store.TezosAccount().Deactivate(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleReactivateTezosAccount activates accounts
func (s *Server) HandleReactivateTezosAccount(c *gin.Context) {
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

	if err := s.store.TezosAccount().Reactivate(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleMakeTezosAccountPrivate makes tezos accounts private
func (s *Server) HandleMakeTezosAccountPrivate(c *gin.Context) {
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

	if err := s.store.TezosAccount().Private(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleMakeTezosAccountUnprivate makes tezos accounts public
func (s *Server) HandleMakeTezosAccountUnprivate(c *gin.Context) {
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

	if err := s.store.TezosAccount().Unprivate(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}
