package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
)

// VaccineRepository ...
type VaccineRepository struct {
	store *Store
}

// Create ...
func (r *VaccineRepository) Create(v *model.Vaccine, now time.Time) error {
	if err := v.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO viruses (vaccine_name, virus_id, virus_name, description, requested_amount, funded_amount, is_active, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING vaccine_id",
		v.Name,
		v.VirusID,
		v.VirusName,
		v.Description,
		v.RequestedAmount,
		v.FundedAmount,
		v.IsActive,
		v.CreatedBy,
		now,
	).Scan(&v.VirusID)
}

// GetMyVaccines returns all vaccines made by user
func (r *VaccineRepository) GetMyVaccines(createdBy string) ([]*model.Vaccine, error) {
	var vaccines []*model.Vaccine
	if err := r.store.db.Select(&vaccines,
		"SELECT vaccine_id, vaccine_name, virus_id, virus_name, description, requested_amount, funded_amount, is_active, created_by, created_at FROM vaccines WHERE created_by=$1",
		createdBy,
	); err != nil {
		return nil, err
	}

	return vaccines, nil
}

// GetVaccines returns all vaccines from database
func (r *VaccineRepository) GetVaccines() ([]*model.Vaccine, error) {
	var vaccines []*model.Vaccine
	if err := r.store.db.Select(&vaccines,
		"SELECT vaccine_id, vaccine_name, virus_id, virus_name, description, requested_amount, funded_amount, is_active, created_by, created_at FROM vaccines",
	); err != nil {
		return nil, err
	}

	return vaccines, nil
}

// GetVaccineByID returns vaccine by specific ID
func (r *VaccineRepository) GetVaccineByID(virusID string) (*model.Vaccine, error) {
	var vaccine model.Vaccine

	if err := r.store.db.QueryRowx("SELECT vaccine_id, vaccine_name, virus_id, virus_name, description, requested_amount, funded_amount, is_active, created_by, created_at FROM viruses WHERE virus_id=$1 LIMIT 1",
		virusID,
	).StructScan(&vaccine); err != nil {
		return nil, err
	}

	return &vaccine, nil
}

// UpdateAmount changes vaccine's funded amount in database
func (r *VaccineRepository) UpdateAmount(funded_amount string, updatedBy string, vaccine_id int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE vaccines 
	SET funded_amount=:funded_amount, updated_by=:updated_by, updated_at=:updated_at
	WHERE (created_by=:updated_by AND vaccine_id=:vaccine_id)`,
		map[string]interface{}{
			"funded_amount": funded_amount,
			"updated_by":    updatedBy,
			"vaccine_id":    vaccine_id,
			"updated_at":    now,
		})
	if err != nil {
		return err
	}

	return nil

}

// UpdateName changes vaccine's name in database
func (r *VaccineRepository) UpdateName(vaccine_name string, updatedBy string, vaccine_id int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE vaccines 
	SET vaccine_name=:vaccine_name, updated_by=:updated_by, updated_at=:updated_at
	WHERE (created_by=:updated_by AND vaccine_id=:vaccine_id)`,
		map[string]interface{}{
			"vaccine_name": vaccine_name,
			"updated_by":   updatedBy,
			"vaccine_id":   vaccine_id,
			"updated_at":   now,
		})
	if err != nil {
		return err
	}

	return nil

}

// UpdateDescription changes vaccine's description record in database
func (r *VaccineRepository) UpdateDescription(description string, updatedBy string, vaccine_id int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE vaccines 
	SET description=:description, updated_by=:updated_by, updated_at=:updated_at
	WHERE (created_by=:updated_by AND vaccine_id=:vaccine_id)`,
		map[string]interface{}{
			"description": description,
			"updated_by":  updatedBy,
			"vaccine_id":  vaccine_id,
			"updated_at":  now,
		})
	if err != nil {
		return err
	}

	return nil

}
