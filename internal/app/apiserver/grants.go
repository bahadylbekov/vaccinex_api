package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleCreateGrant creates grant by provided information
func (s *Server) HandleCreateGrant(c *gin.Context) {
	var grant *model.RequestedGrant
	if err := c.BindJSON(&grant); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()

	grant.CreatedBy = c.Value("userID").(string)
	if err := s.store.Grants().Create(grant, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"grant_id":                       grant.GrantID,
		"alice_nucypher_account_address": grant.AliceNucypherAccountAddress,
		"alice_nucypher_account_name":    grant.AliceNucypherAccountName,
		"bob_ethereum_account":           grant.BobEthereumAccount,
		"bob_nucypher_account_address":   grant.BobNucypherAccountAddress,
		"bob_nucypher_account_name":      grant.BobNucypherAccountName,
		"token_id":                       grant.TokenID,
		"label":                          grant.Label,
		"hash_key":                       grant.HashKey,
		"is_active":                      grant.IsActive,
		"created_by":                     grant.CreatedBy,
		"updated_by":                     grant.UpdatedBy,
	})
}

// GetGrantsForMe returns grants available for submit
func (s *Server) HandleGetGrantsForMe(c *gin.Context) {
	id := c.Param("id")
	grants, err := s.store.Grants().GetGrantsForMe(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, grants)
}

// HandleGetCompletedGrantsForMe returns completed grants
func (s *Server) HandleGetCompletedGrantsForMe(c *gin.Context) {
	var created_by string
	created_by = c.Value("userID").(string)
	grants, err := s.store.Grants().GetCompletedGrantsForMe(created_by)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, grants)
}

// HandleSubmitGrant allow user submit grant
func (s *Server) HandleSubmitGrant(c *gin.Context) {
	type response struct {
		IsActive bool   `json:"is_active"`
		HashKey  string `json:"hash_key"`
	}
	var g *response

	if err := c.BindJSON(&g); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	now := time.Now()
	UserID := c.Value("userID").(string)

	if err := s.store.Grants().Submit(g.IsActive, UserID, g.HashKey, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_active":  g.IsActive,
		"hash_key":   g.HashKey,
		"updated_by": UserID,
	})
}
