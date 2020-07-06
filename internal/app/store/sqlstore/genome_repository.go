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
		`INSERT INTO genomes (genome_name, organization_id, organization_name, vaccine_id, vaccine_name, file_url, price, virus_name, simularity_rate, origin, owner_account, nucypher_account, filename, token_id, receipt_id, policy_id, ethereum_address, is_active, is_sold, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21) 
		RETURNING genome_id`,
		g.Name,
		g.OrganizationID,
		g.OrganizationName,
		g.VaccineID,
		g.VaccineName,
		g.FileUrl,
		g.Price,
		g.VirusName,
		g.SimularityRate,
		g.Origin,
		g.OwnerAccount,
		g.NucypherAccount,
		g.Filename,
		g.TokenID,
		g.ReceiptID,
		g.PolicyID,
		g.EthereumAddress,
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
		"SELECT genome_id, genome_name, organization_id, organization_name, vaccine_id, vaccine_name, file_url, price, virus_name, simularity_rate, origin, owner_account, nucypher_account, filename, token_id, receipt_id, policy_id, ethereum_address, is_active, is_sold, created_by, created_at FROM genomes WHERE created_by=$1",
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
		"SELECT genome_id, genome_name, organization_id, organization_name, vaccine_id, vaccine_name, file_url, price, virus_name, simularity_rate, origin, owner_account, nucypher_account, filename, token_id, receipt_id, policy_id, ethereum_address, is_active, is_sold, created_by, created_at FROM genomes",
	); err != nil {
		return nil, err
	}

	return genomes, nil
}

// GetGenomesByVirus returns all genomes for specific virus
func (r *GenomeRepository) GetGenomesByVirus(virusID string) ([]*model.Genome, error) {
	var genomes []*model.Genome
	if err := r.store.db.Select(&genomes,
		`SELECT genome_id, genome_name, organization_id, organization_name, vaccine_id, vaccine_name, file_url, price, virus_name, simularity_rate, origin, owner_account, nucypher_account, filename, token_id, receipt_id, policy_id, ethereum_address, is_active, is_sold, created_by, created_at FROM genomes WHERE virus_name=(SELECT (virus_name) FROM viruses WHERE virus_id=$1)`,
		virusID,
	); err != nil {
		return nil, err
	}

	return genomes, nil
}

// GetGenomesByOrganization returns all genomes for specific organization
func (r *GenomeRepository) GetGenomesByOrganization(virusID string) ([]*model.Genome, error) {
	var genomes []*model.Genome
	if err := r.store.db.Select(&genomes,
		`SELECT genome_id, genome_name, organization_id, organization_name, vaccine_id, vaccine_name, file_url, price, virus_name, simularity_rate, origin, owner_account, nucypher_account, filename, token_id, receipt_id, policy_id, ethereum_address, is_active, is_sold, created_by, created_at FROM genomes WHERE organization_name=(SELECT (organization_name) FROM organizations WHERE organization_id=$1)`,
		virusID,
	); err != nil {
		return nil, err
	}

	return genomes, nil
}

// GetGenomesByOrganization returns all genomes for specific organization
func (r *GenomeRepository) GetGenomesByVaccine(vaccineID string) ([]*model.Genome, error) {
	var genomes []*model.Genome
	if err := r.store.db.Select(&genomes,
		`SELECT genome_id, genome_name, organization_id, organization_name, vaccine_id, vaccine_name, file_url, price, virus_name, simularity_rate, origin, owner_account, nucypher_account, filename, token_id, receipt_id, policy_id, ethereum_address, is_active, is_sold, created_by, created_at FROM genomes WHERE vaccine_name=(SELECT (vaccine_name) FROM vaccines WHERE vaccine_id=$1)`,
		vaccineID,
	); err != nil {
		return nil, err
	}

	return genomes, nil
}
