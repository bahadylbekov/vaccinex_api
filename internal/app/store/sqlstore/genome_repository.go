package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
)

// GenomeRepository ...
type GenomeRepository struct {
	store *Store
}

// Create ...
func (r *GenomeRepository) Create(g *model.Genome, now time.Time) error {
	if err := g.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO genomes (name, organization_name, file_url, virus_name, simularity_rate, origin, is_active, is_sold, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING organization_id",
		g.Name,
		g.OrganizationName,
		g.FileUrl,
		g.VirusName,
		g.SimularityRate,
		g.Origin,
		g.IsActive,
		g.IsSold,
		g.CreatedBy,
		now,
	).Scan(&g.GenomeID)
}

// GetMyGenomes returns genomes made by user ...
func (r *GenomeRepository) GetMyGenomes(createdBy string) ([]*model.Genome, error) {
	var genomes []*model.Genome
	if err := r.store.db.Select(&genomes,
		"SELECT genome_id, name, organization_name, file_url, virus_name, simularity_rate, origin, is_active, is_sold, created_by, created_at FROM genomes WHERE created_by=$1",
		createdBy,
	); err != nil {
		return nil, err
	}

	return genomes, nil
}

// GetGenomes returns all genomes from database...
func (r *GenomeRepository) GetGenomes() ([]*model.Genome, error) {
	var genomes []*model.Genome
	if err := r.store.db.Select(&genomes,
		"SELECT * from genomes",
	); err != nil {
		return nil, err
	}

	return genomes, nil
}
