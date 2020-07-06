BEGIN; 

CREATE TABLE requested_grants (
    grant_id bigserial not null primary key,
    alice_ethereum_account text not null,
    alice_nucypher_account_address text not null,
    alice_nucypher_account_name text not null,    
    bob_ethereum_account text not null,
    bob_nucypher_account_address text not null,
    bob_nucypher_account_name text not null,    
    token_id int not null,
    label text,
    hash_key text,
    is_active bool not null DEFAULT TRUE,
    policy_id bigserial not null REFERENCES nucypher_policies (policy_id),
    receipt_id bigserial not null REFERENCES nucypher_receipts (receipt_id),
    created_by text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by text,
    updated_at TIMESTAMP
);

COMMIT;