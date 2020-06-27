BEGIN; 

CREATE TABLE vaccines (
    vaccine_id bigserial not null primary key,
    vaccine_name text not null,
    virus_id text not null REFERENCES viruses (virus_id),
    virus_name text not null REFERENCES viruses (virus_name),
    description text,
    requested_amount text not null,
    funded_amount text not null,
    is_active bool not null DEFAULT true,
    created_by text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by text,
    updated_at TIMESTAMP
);

COMMIT;