package apiserver

import (
	"net/http"
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
	"github.com/gin-gonic/gin"
)

// HandleTransactionCreate ...
func (s *Server) HandleTransactionCreate(c *gin.Context) {
	var t *model.Transaction
	if err := c.BindJSON(&t); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}
	now := time.Now()
	t.Timestamp = now
	t.CreatedBy = c.Value("userID").(string)
	if err := s.store.Transaction().Create(t, now); err != nil {
		respondWithError(c, http.StatusBadRequest, errBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"txId":              t.TxID,
		"timestamp":         t.Timestamp,
		"sender_address":    t.SenderAddress,
		"recipient_address": t.RecipientAddress,
		"value":             t.Value,
		"currency":          t.Currency,
		"txHash":            t.TxHash,
		"txStatus":          t.TxStatus,
		"confirmations":     t.Confirmations,
		"created_by":        t.CreatedBy,
	})
}

// HandleGetTransactions ...
func (s *Server) HandleGetTransactions(c *gin.Context) {
	createdBy := c.Value("userID").(string)
	transactions, err := s.store.Transaction().GetTransactions(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, transactions)
}

// HandleGetSendTransactions ...
func (s *Server) HandleGetSendTransactions(c *gin.Context) {
	createdBy := c.Value("userID").(string)
	transactions, err := s.store.Transaction().GetSendTransactions(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, transactions)
}

// HandleGetRecievedTransactions ...
func (s *Server) HandleGetRecievedTransactions(c *gin.Context) {
	createdBy := c.Value("userID").(string)
	transactions, err := s.store.Transaction().GetRecievedTransactions(createdBy)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, errInternalServerError)
		return
	}
	c.JSON(http.StatusOK, transactions)
}
