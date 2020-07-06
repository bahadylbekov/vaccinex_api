package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestedGrant struct {
	GrantID                     int       `json:"grant_id" db:"grant_id"`
	AliceEthereumAccount        string    `json:"alice_ethereum_account" db:"alice_ethereum_account"`
	AliceNucypherAccountAddress string    `json:"alice_nucypher_account_address" db:"alice_nucypher_account_address"`
	AliceNucypherAccountName    string    `json:"alice_nucypher_account_name" db:"alice_nucypher_account_name"`
	BobEthereumAccount          string    `json:"bob_ethereum_account" db:"bob_ethereum_account"`
	BobNucypherAccountAddress   string    `json:"bob_nucypher_account_address" db:"bob_nucypher_account_address"`
	BobNucypherAccountName      string    `json:"bob_nucypher_account_name" db:"bob_nucypher_account_name"`
	TokenID                     int       `json:"token_id" db:"token_id"`
	Label                       string    `json:"label" db:"label"`
	HashKey                     string    `json:"hash_key" db:"hash_key"`
	IsActive                    bool      `json:"is_active" db:"is_active"`
	PolicyID                    int       `json:"policy_id" db:"policy_id"`
	ReceiptID                   int       `json:"receipt_id" db:"receipt_id"`
	CreatedBy                   string    `json:"created_by" db:"created_by"`
	CreatedAt                   time.Time `json:"created_at"  db:"created_at"`
	UpdatedBy                   string    `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt                   time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Validate validates requested grant for required data
func (r *RequestedGrant) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.AliceEthereumAccount, validation.Required),
		validation.Field(&r.AliceNucypherAccountAddress, validation.Required),
		validation.Field(&r.AliceNucypherAccountName, validation.Required),
		validation.Field(&r.BobEthereumAccount, validation.Required),
		validation.Field(&r.BobNucypherAccountAddress, validation.Required),
		validation.Field(&r.BobNucypherAccountName, validation.Required),
		validation.Field(&r.TokenID, validation.Required),
		validation.Field(&r.Label, validation.Required),
		validation.Field(&r.HashKey, validation.Required),
		validation.Field(&r.PolicyID, validation.Required),
		validation.Field(&r.ReceiptID, validation.Required),
	)
}
