package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Organization structure the same as into database
type Organization struct {
	OrganizationID int            `json:"organization_id" db:"organization_id"`
	Name           string         `json:"name" db:"name"`
	Email          string         `json:"email" db:"email"`
	Phone          string         `json:"phone" db:"phone"`
	Website        string         `json:"website" db:"website"`
	Country        string         `json:"country" db:"country"`
	City           string         `json:"city" db:"city"`
	Street         string         `json:"street" db:"street"`
	Postcode       string         `json:"postcode" db:"postcode"`
	IsActive       bool           `json:"isActive" db:"is_active"`
	CreatedBy      string         `json:"created_by" db:"created_by"`
	CreatedAt      time.Time      `json:"created_at,omitempty"  db:"created_at"`
	UpdatedBy      string         `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt      time.Time      `json:"updated_at,omitempty" db:"updated_at"`
	Accounts       []*Account     `json:"accounts,omitempty"`
	Transactions   []*Transaction `json:"transactions,omitempty"`
}

// Validate ...
func (c *Organization) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.Country, validation.Required),
	)
}

// ConnectedOrganizations ...
type ConnectedOrganizations struct {
	ID             string    `json:"id" db:"id"`
	OrganizationID string    `json:"organization_id" db:"organization_id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	CreatedBy      string    `json:"created_by" db:"created_by"`
}

// ValidateInvite ...
func (c *ConnectedOrganizations) ValidateInvite() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.OrganizationID, validation.Required),
		validation.Field(&c.CreatedAt, validation.Required),
		validation.Field(&c.CreatedBy, validation.Required),
	)
}
