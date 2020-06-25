BEGIN;

CREATE TABLE organizations (
    organization_id bigserial not null primary key,
    organization_name varchar(255) not null unique,
    email text not null unique,
    photo_url text,
    website text,
    country text,
    city text,
    description text,
    specialization text,
    deals text,
    genomes_amount text,
    funded_amount text,
    is_active boolean,
    created_by text unique,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by text,
    updated_at TIMESTAMP
);

COMMIT;