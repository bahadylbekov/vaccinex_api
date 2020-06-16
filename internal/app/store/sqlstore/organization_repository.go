package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vacinex_api/internal/app/model"
)

// OrganizationRepository ...
type OrganizationRepository struct {
	store *Store
}

// Create ...
func (r *OrganizationRepository) Create(c *model.Organization, now time.Time) error {
	if err := c.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO organizations (name, email, phone, website, country, city, street, postcode, regNum, regDate, is_active, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING organization_id",
		c.Name,
		c.Email,
		c.Phone,
		c.Website,
		c.Country,
		c.City,
		c.Street,
		c.Postcode,
		c.IsActive,
		c.CreatedBy,
		now,
	).Scan(&c.OrganizationID)
}

// GetMyOrganization ...
func (r *OrganizationRepository) GetMyOrganization(createdBy string) (*model.Organization, error) {
	var organization model.Organization

	if err := r.store.db.QueryRowx("SELECT organization_id, name, email, phone, website, country, city, street, postcode, regnum, regdate, is_active, created_by, created_at FROM organizations WHERE created_by=$1 LIMIT 1",
		createdBy,
	).StructScan(&organization); err != nil {
		return nil, err
	}

	return &organization, nil
}

// Update changes organization record in database
func (r *OrganizationRepository) Update(c *model.Organization, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE organizations 
	SET name=:new_name, email=:new_email, phone=:new_phone, website=:new_website, country=:new_country, city=:new_city, street=:new_street, postcode=:new_postcode, regnum=:new_regnum, regdate=:new_regdate, is_active=:new_isActive, updated_by=:created, updated_at=:time
	WHERE (created_by=:created AND organization_id=:id)`,
		map[string]interface{}{
			"id":           c.OrganizationID,
			"new_name":     c.Name,
			"new_email":    c.Email,
			"new_phone":    c.Phone,
			"new_website":  c.Website,
			"new_country":  c.Country,
			"new_city":     c.City,
			"new_street":   c.Street,
			"new_postcode": c.Postcode,
			"new_isActive": c.IsActive,
			"created":      c.CreatedBy,
			"time":         now,
		})
	if err != nil {
		return err
	}

	return nil

}
