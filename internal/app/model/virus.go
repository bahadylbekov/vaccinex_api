package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Virus structure the same as into database
type Virus struct {
	VirusID      int       `json:"virus_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	PhotoUrl     string    `json:"photo_url"`
	Family       string    `json:"family"`
	FatalityRate string    `json:"fatality_rate"`
	Spread       string    `json:"spread"`
	IsActive     bool      `json:"is_active"`
	IsVaccine    bool      `json:"is_vaccine"`
	CreatedBy    string    `json:"created_by" db:"created_by"`
	CreatedAt    time.Time `json:"created_at"  db:"created_at"`
	UpdatedBy    string    `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Validate ...
func (v *Virus) Validate() error {
	return validation.ValidateStruct(
		v,
		validation.Field(&v.Name, validation.Required),
	)
}
