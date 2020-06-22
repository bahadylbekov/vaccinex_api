package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
)

// EthereumAccountRepository ...
type EthereumAccountRepository struct {
	store *Store
}

// Create ...
func (r *EthereumAccountRepository) Create(a *model.EthereumAccount, now time.Time) error {
	if err := a.ValidateEthereum(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		`INSERT INTO accounts (organization_id, 
			address, 
			balance, 
			tokens, 
			created_by, 
			name,
			created_at) VALUES ((
				SELECT organization_id 
				FROM organizations 
				WHERE createdBy=$4), 
				$1, 
				$2, 
				$3, 
				$4, 
				$5, 
				$6) 
				RETURNING account_id`,
		a.Address,
		a.Balance,
		a.Tokens,
		a.CreatedBy,
		a.Name,
		now,
	).Scan(&a.AccountID)
}

// GetAccounts ...
func (r *EthereumAccountRepository) GetAccounts(createdBy string) ([]*model.EthereumAccount, error) {
	var accounts []*model.EthereumAccount

	if err := r.store.db.Select(&accounts,
		"SELECT account_id, organization_id, address, balance, tokens, created_by, name, is_active, created_at, is_private  FROM ethereum_accounts WHERE created_by=$1",
		createdBy,
	); err != nil {
		return nil, err
	}

	return accounts, nil
}

// UpdateName updates name of organization
func (r *EthereumAccountRepository) UpdateName(name string, updatedBy string, accountID int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE ethereum_accounts 
	SET name=:name, updated_by=:created, updated_at=:time
	WHERE (created_by=:created AND account_id=:account_id)`,
		map[string]interface{}{
			"name":       name,
			"created":    updatedBy,
			"account_id": accountID,
			"time":       now,
		})
	if err != nil {
		return err
	}

	return nil

}

// Deactivate updates is_active of account to false
func (r *EthereumAccountRepository) Deactivate(updatedBy string, accountID int, now time.Time) error {

	_, err := r.store.db.Exec(`UPDATE ethereum_accounts 
	SET is_active=false, updated_by=$1, updated_at=$3 
	WHERE (created_by=$1 AND account_id=$2)`,
		updatedBy,
		accountID,
		now,
	)
	if err != nil {
		return err
	}

	return nil

}

// Reactivate updates is_active of account to true
func (r *EthereumAccountRepository) Reactivate(updatedBy string, accountID int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE ethereum_accounts 
	SET is_active=true, updated_by=:created, updated_at=:time 
	WHERE (created_by=:created AND account_id=:account_id AND is_active=false)`,
		map[string]interface{}{
			"created":    updatedBy,
			"account_id": accountID,
			"time":       now,
		})
	if err != nil {
		return err
	}

	return nil

}

// Private updated is_private of account to true
func (r *EthereumAccountRepository) Private(updatedBy string, accountID int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE ethereum_accounts 
	SET is_private=true, updated_by=:created, updated_at=:time 
	WHERE (created_by=:created AND account_id=:account_id AND is_private=false)`,
		map[string]interface{}{
			"created":    updatedBy,
			"account_id": accountID,
			"time":       now,
		})
	if err != nil {
		return err
	}

	return nil
}

// Unprivate updated is_private of account to true
func (r *EthereumAccountRepository) Unprivate(updatedBy string, accountID int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE ethereum_accounts 
	SET is_private=false, updated_by=:created, updated_at=:time 
	WHERE (created_by=:created AND account_id=:account_id AND is_private=true)`,
		map[string]interface{}{
			"created":    updatedBy,
			"account_id": accountID,
			"time":       now,
		})
	if err != nil {
		return err
	}

	return nil
}
