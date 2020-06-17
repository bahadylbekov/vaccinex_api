package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/shopspring/decimal"
)

// Account structure
type Account struct {
	AccountID      int             `json:"account_id" db:"account_id"`
	Name           string          `json:"name" db:"name"`
	OrganizationID int             `json:"organization_id" db:"organization_id"`
	Address        string          `json:"address" db:"address"`
	Balance        decimal.Decimal `json:"balance" db:"balance"`
	Tokens         decimal.Decimal `json:"tokens" db:"tokens"`
	OpenBalance    decimal.Decimal `json:"openBalance" db:"openbalance"`
	CloseBalance   decimal.Decimal `json:"closeBalance" db:"closebalance"`
	IsActive       bool            `json:"is_active" db:"is_active"`
	CreatedBy      string          `json:"created_by" db:"created_by"`
	CreatedAt      time.Time       `json:"created_at"  db:"created_at"`
	UpdatedBy      string          `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt      time.Time       `json:"updated_at,omitempty" db:"updated_at"`
	IsPrivate      bool            `json:"is_private" db:"is_private"`
}

// Accounts is an array of Account objects
type Accounts []Account

// Validate ...
func (a *Account) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Address, validation.Required),
		validation.Field(&a.Balance, validation.Required),
	)
}
