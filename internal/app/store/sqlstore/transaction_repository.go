package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
)

// TransactionRepository ...
type TransactionRepository struct {
	store *Store
}

// Create ...
func (r *TransactionRepository) Create(t *model.Transaction, now time.Time) error {
	if err := t.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		`INSERT INTO transactions (txid, 
			timestamp,
			sender_id,
			sender_account_id,
			sender_account_address, 
			recipient_id,
			recipient_account_id, 
			recipient_account_address,
			value, 
			currency, 
			txhash, 
			txstatus, 
			confirmations, 
			created_by, 
			created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $2) 
		RETURNING id`,
		t.TxID,
		now,
		t.SenderID,
		t.SenderAccountID,
		t.SenderAddress,
		t.RecipientID,
		t.RecipientAccountID,
		t.RecipientAddress,
		t.Value,
		t.Currency,
		t.TxHash,
		t.TxStatus,
		t.Confirmations,
		t.CreatedBy,
	).Scan(&t.ID)
}

// GetTransactions ...
func (r *TransactionRepository) GetTransactions(createdBy string) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	if err := r.store.db.Select(&transactions,
		`SELECT * FROM transactions 
		WHERE sender_id=(SELECT (organization_id) FROM organizations WHERE organizations.created_by=$1)
		UNION SELECT * FROM transactions
		WHERE recipient_id=(SELECT (organization_id) FROM organizations WHERE organizations.created_by=$1)`,
		createdBy,
	); err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetSendTransactions returns array of all transactions that was send by user or error
func (r *TransactionRepository) GetSendTransactions(createdBy string) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	if err := r.store.db.Select(&transactions,
		`SELECT * FROM transactions 
		WHERE sender_account_id=(SELECT (account_id) FROM accounts WHERE accounts.organization_id=(SELECT (organizations.organization_id) FROM organizations WHERE organizations.created_by=$1))`,
		createdBy,
	); err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetRecievedTransactions returns array of all transactions that was recieved by user or error
func (r *TransactionRepository) GetRecievedTransactions(createdBy string) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	if err := r.store.db.Select(&transactions,
		`SELECT * FROM transactions
		WHERE recipient_account_id=(SELECT (account_id) FROM accounts WHERE organization_id=(SELECT (organizations.organization_id) FROM organizations WHERE organizations.created_by=$1))`,
		createdBy,
	); err != nil {
		return nil, err
	}

	return transactions, nil
}
