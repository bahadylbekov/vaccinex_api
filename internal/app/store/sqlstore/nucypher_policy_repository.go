package sqlstore

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
)

type NucypherPolicyRepository struct {
	store *Store
}

// Create ...
func (r *NucypherPolicyRepository) Create(policy *model.NucypherPolicy, now time.Time) error {
	if err := policy.ValidatePolicy(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		`INSERT INTO nucypher_policies (alice_sig_pubkey, label, policy_pubkey, created_by, 
			created_at) VALUES ($1, $2, $3, $4, $5) RETURNING policy_id`,
		policy.AliceSigningPublicKey,
		policy.Label,
		policy.PolicyPublicKey,
		policy.CreatedBy,
		now,
	).Scan(&policy.PolicyID)
}

func (r *NucypherPolicyRepository) GetByID(policy_id string) (*model.NucypherPolicy, error) {
	var policy []*model.NucypherPolicy

	if err := r.store.db.Select(&policy,
		"SELECT policy_id, alice_sig_pubkey, label, policy_pubkey , created_by, created_at FROM nucypher_policies WHERE policy_id=$1 LIMIT 1",
		policy_id,
	); err != nil {
		return nil, err
	}

	return policy[0], nil

}

func (r *NucypherPolicyRepository) GetByLabel(label string) (*model.NucypherPolicy, error) {
	var policy []*model.NucypherPolicy

	if err := r.store.db.Select(&policy,
		"SELECT policy_id, alice_sig_pubkey, label, policy_pubkey , created_by, created_at FROM nucypher_policies WHERE label=$1 LIMIT 1",
		label,
	); err != nil {
		return nil, err
	}

	return policy[0], nil

}
