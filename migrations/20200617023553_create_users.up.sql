BEGIN;

CREATE TABLE users (
    id bigserial not null primary key,
    email text not null unique,
    encrypted_password varchar not null,
    created_by text unique,
    created_at TIMESTAMP not null,
    updated_by text,
    updated_at TIMESTAMP,
    is_deleted bool,
    deleted_by text,
    deleted_at TIMESTAMP
);

COMMIT;