BEGIN; 

CREATE TABLE accounts (
    account_id bigserial not null primary key,
    name text not null,
    organization_id int not null,
    address varchar(255) not null unique,
    balance DECIMAL(10,2) not null,
    tokens DECIMAL(10,2) not null,
    openBalance DECIMAL(10,2) not null,
    closeBalance DECIMAL(10,2) not null,
    is_active bool not null DEFAULT true,
    is_private bool not null DEFAULT false,
    created_by text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by text,
    updated_at TIMESTAMP,
    FOREIGN KEY (organization_id) REFERENCES organizations (organization_id)
);

COMMIT;