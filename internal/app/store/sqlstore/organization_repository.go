package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
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

// FindOrganizationsByEmail ...
func (r *OrganizationRepository) FindOrganizationsByEmail(search string) ([]*model.Organization, error) {
	var organizations []*model.Organization

	if err := r.store.db.Select(&organizations,
		"SELECT * FROM organizations WHERE email LIKE $1",
		"%"+search+"%",
	); err != nil {
		return nil, err
	}
	return organizations, nil
}

// AddOrganizationToMyList ...
func (r *OrganizationRepository) AddOrganizationToMyList(myID string, organizationID string, createdBy string, now time.Time) error {
	// if err := c.ValidateInvite(); err != nil {
	// 	return err
	// }

	if myID < organizationID {
		return r.store.db.QueryRow(
			"INSERT INTO connected_organizations (id, organization_id, created_by, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
			myID,
			organizationID,
			createdBy,
			now,
		).Scan(&myID)
	} else {
		return r.store.db.QueryRow(
			"INSERT INTO connected_organizations (id, organization_id, created_by, created_at) VALUES ($2, $1, $3, $4) RETURNING id",
			myID,
			organizationID,
			createdBy,
			now,
		).Scan(&myID)
	}
}

// GetConnectedOrganizations ...
func (r *OrganizationRepository) GetConnectedOrganizations(createdBy string) ([]*model.Organization, error) {
	var organizations []*model.Organization
	if err := r.store.db.Select(&organizations,
		`SELECT organizations.organization_id, name, email, phone, website, country, city, street, postcode, regnum, regdate, is_active, organizations.created_by, organizations.created_at
		FROM organizations
		INNER JOIN connected_organizations 
		ON organizations.organization_id=connected_organizations.id 
		WHERE connected_organizations.organization_id=(SELECT organizations.organization_id FROM organizations WHERE organizations.created_by=$1)
		UNION 
		SELECT organizations.organization_id, name, email, phone, website, country, city, street, postcode, regnum, regdate, is_active, organizations.created_by, organizations.created_at
		FROM organizations 
		INNER JOIN connected_organizations 
		ON organizations.organization_id=connected_organizations.organization_id 
		WHERE connected_organizations.id=(SELECT organizations.organization_id FROM organizations WHERE organizations.created_by=$1);`,
		createdBy); err != nil {
		return nil, err
	}

	return organizations, nil
}

// GetOrganization returns organization with accounts and transactions list
func (r *OrganizationRepository) GetOrganization(organizationID string) (*model.Organization, error) {
	var organization model.Organization
	var accounts []*model.Account
	var transactions []*model.Transaction

	if err := r.store.db.QueryRowx("SELECT organization_id, name, email, phone, website, country, city, street, postcode, regnum, regdate, is_active, created_by, created_at FROM organizations WHERE organization_id=$1 LIMIT 1",
		organizationID,
	).StructScan(&organization); err != nil {
		return nil, err
	}

	if err := r.store.db.Select(&accounts,
		"SELECT account_id, organization_id, address, balance, tokens, openbalance, closebalance, created_by, name, is_active, created_at, is_private FROM accounts WHERE organization_id=$1",
		organizationID,
	); err != nil {
		return &organization, err
	}

	organization.Accounts = accounts

	if err := r.store.db.Select(&transactions,
		`SELECT * FROM transactions 
		WHERE sender_account_id=(SELECT (account_id) FROM accounts WHERE organization_id=$1)
		UNION SELECT * FROM transactions
		WHERE recipient_account_id=(SELECT (account_id) FROM accounts WHERE organization_id=$1)`,
		organizationID,
	); err != nil {
		return &organization, err
	}

	organization.Transactions = transactions

	return &organization, nil
}

// Delete removes connection between 2 organizations in database
func (r *OrganizationRepository) Delete(myID string, secondID string) error {

	if myID < secondID {
		_, err := r.store.db.NamedExec(`DELETE FROM connected_organizations 
		WHERE (id=:my_id AND organization_id=:second_id)`,
			map[string]interface{}{
				"my_id":     myID,
				"second_id": secondID,
			})
		if err != nil {
			return err
		}
		return nil

	} else {
		_, err := r.store.db.NamedExec(`DELETE FROM connected_organizations 
		WHERE (id=:second_id AND organization_id=:my_id)`,
			map[string]interface{}{
				"my_id":     myID,
				"second_id": secondID,
			})
		if err != nil {
			return err
		}

		return nil
	}
}

// ID returns organization_id from database using created_by user_id
func (r *OrganizationRepository) ID(createdBy string) (*string, error) {
	var organizationID string
	if err := r.store.db.Get(&organizationID,
		"SELECT (organization_id) FROM organizations WHERE created_by=$1",
		createdBy,
	); err != nil {
		return nil, err
	}

	return &organizationID, nil
}
