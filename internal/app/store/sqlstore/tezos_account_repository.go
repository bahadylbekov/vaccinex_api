package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
)

// TezosAccountRepository ...
type TezosAccountRepository struct {
	store *Store
}

// Create ...
func (r *TezosAccountRepository) Create(a *model.TezosAccount, now time.Time) error {
	if err := a.ValidateTezos(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		`INSERT INTO tezos_accounts (organization_id, 
			address, 
			balance, 
			tokens, 
			created_by, 
			name,
			created_at) VALUES ((
				SELECT organization_id 
				FROM organizations 
				WHERE created_by=$4), 
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
func (r *TezosAccountRepository) GetAccounts(createdBy string) ([]*model.TezosAccount, error) {
	var accounts []*model.TezosAccount

	if err := r.store.db.Select(&accounts,
		"SELECT account_id, organization_id, address, balance, tokens, created_by, name, is_active, created_at, is_private  FROM tezos_accounts WHERE created_by=$1",
		createdBy,
	); err != nil {
		return nil, err
	}

	return accounts, nil
}

// GetAccounts ...
func (r *TezosAccountRepository) GetAccountOrganization(id string) ([]*model.TezosAccount, error) {
	var accounts []*model.TezosAccount

	if err := r.store.db.Select(&accounts,
		"SELECT account_id, organization_id, address, balance, tokens, created_by, name, is_active, created_at, is_private FROM tezos_accounts WHERE organization_id=$1",
		id,
	); err != nil {
		return nil, err
	}

	return accounts, nil
}

// UpdateName updates name of organization
func (r *TezosAccountRepository) UpdateName(name string, updatedBy string, accountID int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE tezos_accounts 
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
func (r *TezosAccountRepository) Deactivate(updatedBy string, accountID int, now time.Time) error {

	_, err := r.store.db.Exec(`UPDATE tezos_accounts 
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
func (r *TezosAccountRepository) Reactivate(updatedBy string, accountID int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE tezos_accounts 
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
func (r *TezosAccountRepository) Private(updatedBy string, accountID int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE tezos_accounts 
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
func (r *TezosAccountRepository) Unprivate(updatedBy string, accountID int, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE tezos_accounts 
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
