package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleAccountCreate ...
func (s *Server) HandleEthereumAccountCreate(c *gin.Context) {
	var a *model.EthereumAccount
	if err := c.BindJSON(&a); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()
	// a.CreatedAt = now

	a.CreatedBy = c.Value("userID").(string)
	if err := s.store.EthereumAccount().Create(a, now); err != nil {
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

// HandleGetEthereumAccounts returns all tezos accounts for user
func (s *Server) HandleGetEthereumAccounts(c *gin.Context) {
	createdBy := c.Value("userID").(string)
	accounts, err := s.store.EthereumAccount().GetAccounts(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, accounts)
}

// HandleGetEthereumAccountForOrganization returns all nucypher accounts for organization
func (s *Server) HandleGetEthereumAccountForOrganization(c *gin.Context) {
	id := c.Param("id")
	accounts, err := s.store.EthereumAccount().GetAccountByOrganization(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, accounts)
}

// HandleUpdateEthereumAccount allow user update account name
func (s *Server) HandleUpdateEthereumAccount(c *gin.Context) {
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

	if err := s.store.EthereumAccount().UpdateName(a.Name, a.UpdatedBy, a.AccountID, now); err != nil {
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

// HandleUpdateEthereumAccountAddress allow user update account address
func (s *Server) HandleUpdateEthereumAccountAddress(c *gin.Context) {
	type repsonse struct {
		Address   string    `json:"address"`
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

	if err := s.store.EthereumAccount().UpdateAddress(a.Address, a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address":    a.Address,
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleDeactivateEthereumAccount deactivates accounts
func (s *Server) HandleDeactivateEthereumAccount(c *gin.Context) {
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

	if err := s.store.EthereumAccount().Deactivate(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleReactivateEthereumAccount activates accounts
func (s *Server) HandleReactivateEthereumAccount(c *gin.Context) {
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

	if err := s.store.EthereumAccount().Reactivate(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleMakeEthereumAccountPrivate makes tezos accounts private
func (s *Server) HandleMakeEthereumAccountPrivate(c *gin.Context) {
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

	if err := s.store.EthereumAccount().Private(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}

// HandleMakeEthereumAccountUnprivate makes tezos accounts public
func (s *Server) HandleMakeEthereumAccountUnprivate(c *gin.Context) {
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

	if err := s.store.EthereumAccount().Unprivate(a.UpdatedBy, a.AccountID, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id": a.AccountID,
		"updated_by": a.UpdatedBy,
		"updated_at": a.UpdatedAt,
	})
}
