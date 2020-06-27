package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Organization structure the same as into database
type Organization struct {
	OrganizationID int       `json:"organization_id" db:"organization_id"`
	Name           string    `json:"organization_name" db:"organization_name"`
	Email          string    `json:"email" db:"email"`
	PhotoUrl       string    `json:"photo_url" db:"photo_url"`
	Website        string    `json:"website" db:"website"`
	Country        string    `json:"country" db:"country"`
	City           string    `json:"city" db:"city"`
	Description    string    `json:"description" db:"description"`
	Specialization string    `json:"specialization" db:"specialization"`
	Deals          string    `json:"deals" db:"deals"`
	GenomesAmount  string    `json:"genomes_amount" db:"genomes_amount"`
	FundedAmount   string    `json:"funded_amount" db:"funded_amount"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	CreatedBy      string    `json:"created_by" db:"created_by"`
	CreatedAt      time.Time `json:"created_at,omitempty"  db:"created_at"`
	UpdatedBy      string    `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Organizations is an array of Organization objects
type Organizations []Organization

// Validate ...
func (o *Organization) Validate() error {
	return validation.ValidateStruct(
		o,
		validation.Field(&o.Name, validation.Required),
		validation.Field(&o.Email, validation.Required, is.Email),
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
