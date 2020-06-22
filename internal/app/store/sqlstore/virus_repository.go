package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
)

// VirusRepository ...
type VirusRepository struct {
	store *Store
}

// Create ...
func (r *VirusRepository) Create(v *model.Virus, now time.Time) error {
	if err := v.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO viruses (name, description, photo_url, family, fatality_rate, spread, is_active, is_vaccine, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING organization_id",
		v.Name,
		v.Description,
		v.PhotoUrl,
		v.Family,
		v.FatalityRate,
		v.Spread,
		v.IsActive,
		v.IsVaccine,
		v.CreatedBy,
		now,
	).Scan(&v.VirusID)
}

// GetMyViruses returns viruses made by user ...
func (r *VirusRepository) GetMyViruses(createdBy string) ([]*model.Virus, error) {
	var viruses []*model.Virus
	if err := r.store.db.Select(&viruses,
		"SELECT virus_id, name, description, photo_url, family, fatality_rate, spread, is_active, is_vaccine, created_by, created_at FROM genomes WHERE created_by=$1",
		createdBy,
	); err != nil {
		return nil, err
	}

	return viruses, nil
}

// GetViruses returns all viruses from database...
func (r *VirusRepository) GetViruses() ([]*model.Virus, error) {
	var viruses []*model.Virus
	if err := r.store.db.Select(&viruses,
		"SELECT * from viruses",
	); err != nil {
		return nil, err
	}

	return viruses, nil
}

// GetVirus returns virus with genomes list
func (r *VirusRepository) GetVirusByID(virusID string) (*model.Virus, error) {
	var virus model.Virus

	if err := r.store.db.QueryRowx("SELECT * FROM viruses WHERE virus_id=$1 LIMIT 1",
		virusID,
	).StructScan(&virus); err != nil {
		return nil, err
	}

	return &virus, nil
}

// Update changes virus record in database
func (r *VirusRepository) Update(v *model.Virus, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE viruses 
	SET name=:new_name, description=:new_description, photo_url=:new_photo_url, family=:new_family, fatality_rate=:new_fatality_rate, spread=:new_spread, is_active=:new_isActive, is_vaccine=:new_isVaccine, updated_by=:created, updated_at=:time
	WHERE (created_by=:created AND virus_id=:id)`,
		map[string]interface{}{
			"id":                v.VirusID,
			"new_name":          v.Name,
			"new_description":   v.Description,
			"new_photo_url":     v.PhotoUrl,
			"new_family":        v.Family,
			"new_fatality_rate": v.FatalityRate,
			"new_spread":        v.Spread,
			"new_isActive":      v.IsActive,
			"new_isVaccine":     v.IsVaccine,
			"created":           v.CreatedBy,
			"time":              now,
		})
	if err != nil {
		return err
	}

	return nil

}
