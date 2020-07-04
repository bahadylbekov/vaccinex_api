package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleCreatePolicy creates policy by provided information
func (s *Server) HandleCreatePolicy(c *gin.Context) {
	var policy *model.NucypherPolicy
	if err := c.BindJSON(&policy); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()

	policy.CreatedBy = c.Value("userID").(string)
	if err := s.store.NucypherPolicy().Create(policy, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"policy_id":        policy.PolicyID,
		"alice_sig_pubkey": policy.AliceSigningPublicKey,
		"label":            policy.Label,
		"policy_pubkey":    policy.PolicyPublicKey,
		"created_by":       policy.CreatedBy,
		"updated_by":       policy.UpdatedBy,
	})
}

// HandleGetPolicyByID returns specific policy information
func (s *Server) HandleGetPolicyByID(c *gin.Context) {
	id := c.Param("id")
	policy, err := s.store.NucypherPolicy().GetByID(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, policy)
}
