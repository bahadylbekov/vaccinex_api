BEGIN; 

CREATE TABLE viruses (
    virus_id bigserial not null primary key,
    name text not null,
    description text,
    photo_url text,
    family text,
    fatality_rate text,
    spread text,
    is_active bool not null DEFAULT true,
    is_vaccine bool not null DEFAULT false,
    created_by text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by text,
    updated_at TIMESTAMP,
);

COMMIT;