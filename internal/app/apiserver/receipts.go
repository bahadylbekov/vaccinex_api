package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleCreateReceipt creates receipt by provided information
func (s *Server) HandleCreateReceipt(c *gin.Context) {
	var receipt *model.NucypherReceipt
	if err := c.BindJSON(&receipt); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()

	receipt.CreatedBy = c.Value("userID").(string)
	if err := s.store.NucypherReceipt().Create(receipt, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"receipt_id":             receipt.ReceiptID,
		"data_source_public_key": receipt.EnricoPublicKey,
		"hash_key":               receipt.HashKey,
		"created_by":             receipt.CreatedBy,
		"updated_by":             receipt.UpdatedBy,
	})
}

// HandleGetReceiptByID returns specific receipt data
func (s *Server) HandleGetReceiptByID(c *gin.Context) {
	id := c.Param("id")
	policy, err := s.store.NucypherPolicy().GetByID(id)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, policy)
}

// HandleGetReceiptByHash returns specific receipt data
func (s *Server) HandleGetReceiptByHash(c *gin.Context) {
	hash := c.Param("hash")
	policy, err := s.store.NucypherReceipt().GetReceiptByHash(hash)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, policy)
}
