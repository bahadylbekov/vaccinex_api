BEGIN; 

CREATE TABLE nucypher_receipts (
    receipt_id bigserial not null primary key,
    data_source_public_key text not null,
    hash_key text not null UNIQUE,
    created_by text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by text,
    updated_at TIMESTAMP
);

COMMIT;