package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/shopspring/decimal"
)

// NucypherAccount structure
type NucypherAccount struct {
	AccountID      int             `json:"account_id" db:"account_id"`
	Name           string          `json:"name" db:"name"`
	OrganizationID int             `json:"organization_id" db:"organization_id"`
	Address        string          `json:"address" db:"address"`
	SigningKey     string          `json:"signing_key" db:"signing_key"`
	EncryptingKey  string          `json:"encrypting_key" db:"encrypting_key"`
	Balance        decimal.Decimal `json:"balance" db:"balance"`
	Tokens         decimal.Decimal `json:"tokens" db:"tokens"`
	IsActive       bool            `json:"is_active" db:"is_active"`
	IsPrivate      bool            `json:"is_private" db:"is_private"`
	CreatedBy      string          `json:"created_by" db:"created_by"`
	CreatedAt      time.Time       `json:"created_at"  db:"created_at"`
	UpdatedBy      string          `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt      time.Time       `json:"updated_at,omitempty" db:"updated_at"`
}

// NucypherAccounts is an array of NucypherAccount objects
type NucypherAccounts []NucypherAccount

// ValidateNuCypher ...
func (a *NucypherAccount) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Address, validation.Required),
		validation.Field(&a.Balance, validation.Required),
		validation.Field(&a.SigningKey, validation.Required),
		validation.Field(&a.EncryptingKey, validation.Required),
	)
}

type NucypherPolicy struct {
	PolicyID              int       `json:"policy_id" db:"policy_id"`
	AliceSigningPublicKey string    `json:"alice_sig_pubkey" db:"alice_sig_pubkey"`
	Label                 string    `json:"label" db:"label"`
	PolicyPublicKey       string    `json:"policy_pubkey" db:"policy_pubkey"`
	CreatedBy             string    `json:"created_by" db:"created_by"`
	CreatedAt             time.Time `json:"created_at"  db:"created_at"`
	UpdatedBy             string    `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt             time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// ValidatePolicy validates policy for required data
func (p *NucypherPolicy) ValidatePolicy() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AliceSigningPublicKey, validation.Required),
		validation.Field(&p.Label, validation.Required),
		validation.Field(&p.PolicyPublicKey, validation.Required),
	)
}

type NucypherReceipt struct {
	ReceiptID       int       `json:"receipt_id" db:"receipt_id"`
	EnricoPublicKey string    `json:"data_source_public_key" db:"data_source_public_key"`
	HashKey         string    `json:"hash_key" db:"hash_key"`
	CreatedBy       string    `json:"created_by" db:"created_by"`
	CreatedAt       time.Time `json:"created_at"  db:"created_at"`
	UpdatedBy       string    `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// ValidatePolicy validates policy for required data
func (r *NucypherReceipt) ValidateReceipt() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.EnricoPublicKey, validation.Required),
		validation.Field(&r.HashKey, validation.Required),
	)
}
