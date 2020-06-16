BEGIN;

CREATE TABLE organizations (
    organization_id bigserial not null primary key,
    name text not null,
    email text not null unique,
    phone text,
    website text,
    country text,
    city text,
    street text,
    postcode text,
    isActive boolean,
    created_by text unique,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by text,
    updated_at TIMESTAMP,
    is_deleted bool,
    deleted_by text,
    deleted_at TIMESTAMP
);

COMMIT;