package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
)

type RequestedGrantsRepository struct {
	store *Store
}

// Create ...
func (r *RequestedGrantsRepository) Create(grant *model.RequestedGrant, now time.Time) error {
	if err := grant.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		`INSERT INTO requested_grants (alice_ethereum_account, alice_nucypher_account_address, alice_nucypher_account_name, bob_ethereum_account, bob_nucypher_account_address, bob_nucypher_account_name, token_id, 
			label, hash_key, is_active, policy_id, receipt_id, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING grant_id`,
		grant.AliceEthereumAccount,
		grant.AliceNucypherAccountAddress,
		grant.AliceNucypherAccountName,
		grant.BobEthereumAccount,
		grant.BobNucypherAccountAddress,
		grant.BobNucypherAccountName,
		grant.TokenID,
		grant.Label,
		grant.HashKey,
		grant.IsActive,
		grant.PolicyID,
		grant.ReceiptID,
		grant.CreatedBy,
		now,
	).Scan(&grant.GrantID)
}

// Submit changes is_active value to false
func (r *RequestedGrantsRepository) Submit(is_active bool, created_by string, hash_key string, updated_by string, now time.Time) error {

	_, err := r.store.db.NamedExec(`UPDATE requested_grants 
	SET is_active=:is_active, updated_by=:updated_by, updated_at=:updated_at
	WHERE (created_by=:created_by AND hash_key=:hash_key)`,
		map[string]interface{}{
			"is_active":  is_active,
			"updated_by": updated_by,
			"created_by": created_by,
			"hash_key":   hash_key,
			"updated_at": now,
		})
	if err != nil {
		return err
	}

	return nil

}

func (r *RequestedGrantsRepository) GetGrantsForMe(alice_nucypher_address string) ([]*model.RequestedGrant, error) {
	var grants []*model.RequestedGrant

	if err := r.store.db.Select(&grants,
		"SELECT grant_id, alice_ethereum_account, alice_nucypher_account_address, alice_nucypher_account_name, bob_ethereum_account, bob_nucypher_account_address, bob_nucypher_account_name, token_id, hash_key, label, is_active, policy_id, receipt_id, created_by, created_at FROM requested_grants WHERE alice_nucypher_account_address=$1 AND is_active=True",
		alice_nucypher_address,
	); err != nil {
		return nil, err
	}

	return grants, nil

}

func (r *RequestedGrantsRepository) GetCompletedGrantsForMe(created_by string) ([]*model.RequestedGrant, error) {
	var grants []*model.RequestedGrant

	if err := r.store.db.Select(&grants,
		"SELECT grant_id, alice_ethereum_account, alice_nucypher_account_address, alice_nucypher_account_name, bob_ethereum_account, bob_nucypher_account_address, bob_nucypher_account_name, token_id, hash_key, label, is_active, policy_id, receipt_id, created_by, created_at FROM requested_grants WHERE created_by=$1 AND is_active=False",
		created_by,
	); err != nil {
		return nil, err
	}

	return grants, nil

}
