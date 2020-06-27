package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Vaccine structure the same as into database
type Vaccine struct {
	VaccineID       int       `json:"vaccine_id" db:"vaccine_id"`
	Name            string    `json:"vaccine_name" db:"vaccine_name"`
	VirusID         string    `json:"virus_id" db:"virus_id"`
	VirusName       string    `json:"virus_name" db:"virus_name"`
	Description     string    `json:"description" db:"description"`
	RequestedAmount string    `json:"requested_amount" db:"requested_amount"`
	FundedAmount    string    `json:"funded_amount" db:"funded_amount"`
	IsActive        bool      `json:"is_active" db:"is_active"`
	CreatedBy       string    `json:"created_by" db:"created_by"`
	CreatedAt       time.Time `json:"created_at"  db:"created_at"`
	UpdatedBy       string    `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Validate ...
func (g *Vaccine) Validate() error {
	return validation.ValidateStruct(
		g,
		validation.Field(&g.Name, validation.Required),
		validation.Field(&g.VirusID, validation.Required),
		validation.Field(&g.VirusName, validation.Required),
		validation.Field(&g.Description, validation.Required),
	)
}
