package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/shopspring/decimal"
)

// EthereumAccount structure
type EthereumAccount struct {
	AccountID      int             `json:"account_id" db:"account_id"`
	Name           string          `json:"name" db:"name"`
	OrganizationID int             `json:"organization_id" db:"organization_id"`
	Address        string          `json:"address" db:"address"`
	Balance        decimal.Decimal `json:"balance" db:"balance"`
	Tokens         decimal.Decimal `json:"tokens" db:"tokens"`
	IsActive       bool            `json:"is_active" db:"is_active"`
	IsPrivate      bool            `json:"is_private" db:"is_private"`
	CreatedBy      string          `json:"created_by" db:"created_by"`
	CreatedAt      time.Time       `json:"created_at"  db:"created_at"`
	UpdatedBy      string          `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt      time.Time       `json:"updated_at,omitempty" db:"updated_at"`
}

// EthereumAccounts is an array of EthereumAccount objects
type EthereumAccounts []EthereumAccount

// ValidateEthereum ...
func (a *EthereumAccount) ValidateEthereum() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Address, validation.Required),
		validation.Field(&a.Balance, validation.Required),
	)
}

// NucypherAccount structure
type NucypherAccount struct {
	AccountID      int             `json:"account_id" db:"account_id"`
	Name           string          `json:"name" db:"name"`
	OrganizationID int             `json:"organization_id" db:"organization_id"`
	Address        string          `json:"address" db:"address"`
	VerifyingKey   string          `json:"verifying_key" db:"verifying_key"`
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
func (a *NucypherAccount) ValidateNuCypher() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Address, validation.Required),
		validation.Field(&a.Balance, validation.Required),
		validation.Field(&a.VerifyingKey, validation.Required),
	)
}

// TezosAccount structure
type TezosAccount struct {
	AccountID      int             `json:"account_id" db:"account_id"`
	Name           string          `json:"name" db:"name"`
	OrganizationID int             `json:"organization_id" db:"organization_id"`
	Address        string          `json:"address" db:"address"`
	Balance        decimal.Decimal `json:"balance" db:"balance"`
	Tokens         decimal.Decimal `json:"tokens" db:"tokens"`
	IsActive       bool            `json:"is_active" db:"is_active"`
	IsPrivate      bool            `json:"is_private" db:"is_private"`
	CreatedBy      string          `json:"created_by" db:"created_by"`
	CreatedAt      time.Time       `json:"created_at"  db:"created_at"`
	UpdatedBy      string          `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt      time.Time       `json:"updated_at,omitempty" db:"updated_at"`
}

// TezosAccounts is an array of TezosAccount objects
type TezosAccounts []TezosAccount

// ValidateTezos ...
func (a *TezosAccount) ValidateTezos() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Address, validation.Required),
		validation.Field(&a.Balance, validation.Required),
	)
}
