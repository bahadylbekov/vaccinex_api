BEGIN; 

CREATE TABLE nucypher_policies (
    policy_id bigserial not null primary key,
    alice_sig_pubkey text not null,
    label text not null,
    policy_pubkey text not null,
    created_by text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by text,
    updated_at TIMESTAMP
);

COMMIT;