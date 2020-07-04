package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
)

type NucypherReceiptRepository struct {
	store *Store
}

// Create function for creating new receipt record in database
func (r *NucypherReceiptRepository) Create(receipt *model.NucypherReceipt, now time.Time) error {
	if err := receipt.ValidateReceipt(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		`INSERT INTO nucypher_accounts (data_source_public_key, hash_key, created_by, 
			created_at) VALUES ($1, $2, $3, $4) RETURNING receipt_id`,
		receipt.EnricoPublicKey,
		receipt.HashKey,
		receipt.CreatedBy,
		now,
	).Scan(&receipt.ReceiptID)
}

func (r *NucypherReceiptRepository) GetByID(receipt_id string) (*model.NucypherReceipt, error) {
	var receipt *model.NucypherReceipt

	if err := r.store.db.QueryRowx("SELECT receipt_id, data_source_public_key, hash_key, created_by, created_at FROM nucypher_receipts WHERE receipt_id=$1 LIMIT 1",
		receipt_id,
	).StructScan(receipt); err != nil {
		return nil, err
	}

	return receipt, nil

}

func (r *NucypherReceiptRepository) GetReceiptByHash(hash string) (*model.NucypherReceipt, error) {
	var receipt *model.NucypherReceipt

	if err := r.store.db.QueryRowx("SELECT receipt_id, data_source_public_key, hash_key, created_by, created_at FROM nucypher_receipts WHERE hash_key=$1 LIMIT 1",
		hash,
	).StructScan(receipt); err != nil {
		return nil, err
	}

	return receipt, nil

}
