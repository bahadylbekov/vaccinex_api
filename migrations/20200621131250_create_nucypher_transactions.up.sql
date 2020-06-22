BEGIN; 

CREATE TABLE nucypher_transactions (
    id bigserial not null primary key,
    txId text not null,
    timestamp TIMESTAMP not null,
    sender_id int REFERENCES organizations (organization_id) not null,
    sender_account_id int REFERENCES nucypher_accounts (account_id) not null,
    sender_account_address text REFERENCES nucypher_accounts (address) not null,
    recipient_id int REFERENCES organizations (organization_id) not null,
    recipient_account_id int REFERENCES nucypher_accounts (account_id) not null,
    recipient_account_address text REFERENCES nucypher_accounts (address) not null,
    value DECIMAL(10,2) not null,
    currency text not null,
    txHash text not null unique,
    txStatus text not null,
    confirmations int,
    created_by text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
);

COMMIT;