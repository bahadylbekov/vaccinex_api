package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Genome structure the same as into database
type Genome struct {
	GenomeID         int       `json:"genome_id" db:"genome_id"`
	Name             string    `json:"genome_name" db:"genome_name"`
	OrganizationName string    `json:"organization_name" db:"organization_name"`
	FileUrl          string    `json:"file_url" db:"file_url"`
	VirusName        string    `json:"virus_name" db:"virus_name"`
	SimularityRate   string    `json:"simularity_rate" db:"simularity_rate"`
	Origin           string    `json:"origin" db:"origin"`
	IsActive         bool      `json:"is_active" db:"is_active"`
	IsSold           bool      `json:"is_sold,omitempty" db:"is_sold"`
	CreatedBy        string    `json:"created_by" db:"created_by"`
	CreatedAt        time.Time `json:"created_at"  db:"created_at"`
	UpdatedBy        string    `json:"updated_by,omitempty" db:"updated_by"`
	UpdatedAt        time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Validate ...
func (g *Genome) Validate() error {
	return validation.ValidateStruct(
		g,
		validation.Field(&g.Name, validation.Required),
		validation.Field(&g.FileUrl, validation.Required),
	)
}
