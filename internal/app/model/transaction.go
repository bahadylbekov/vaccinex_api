package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/shopspring/decimal"
)

// Transaction structure the same as into database
type Transaction struct {
	ID                 int             `json:"id" db:"id"`
	TxID               int             `json:"tx_id" db:"txid"`
	Timestamp          time.Time       `json:"timestamp" db:"timestamp"`
	SenderID           int             `json:"sender_id" db:"sender_id"`
	SenderAccountID    int             `json:"sender_account_id" db:"sender_account_id"`
	SenderAddress      string          `json:"sender_address" db:"sender_account_address"`
	RecipientID        int             `json:"recipient_id" db:"recipient_id"`
	RecipientAccountID int             `json:"recipient_account_id" db:"recipient_account_id"`
	RecipientAddress   string          `json:"recipient_address" db:"recipient_account_address"`
	Value              decimal.Decimal `json:"value" db:"value"`
	Currency           string          `json:"currency" db:"currency"`
	TxHash             string          `json:"txHash" db:"txhash"`
	TxStatus           string          `json:"txStatus" db:"txstatus"`
	Confirmations      int             `json:"confirmations" db:"confirmations"`
	CreatedBy          string          `json:"created_by" db:"created_by"`
	CreatedAt          time.Time       `json:"created_at" db:"created_at"`
}

// Validate ...
func (t *Transaction) Validate() error {
	return validation.ValidateStruct(
		t,
		validation.Field(&t.TxID, validation.Required),
		validation.Field(&t.SenderID, validation.Required),
		validation.Field(&t.SenderAccountID, validation.Required),
		validation.Field(&t.SenderAddress, validation.Required),
		validation.Field(&t.RecipientID, validation.Required),
		validation.Field(&t.RecipientAccountID, validation.Required),
		validation.Field(&t.RecipientAddress, validation.Required),
		validation.Field(&t.Value, validation.Required),
		validation.Field(&t.Currency, validation.Required),
		validation.Field(&t.TxHash, validation.Required),
	)
}
