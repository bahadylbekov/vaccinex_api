BEGIN; 

CREATE TABLE genomes (
    genome_id bigserial not null primary key,
    genome_name text not null UNIQUE,
    organization_id bigserial not null REFERENCES organizations (organization_id),
    organization_name text not null REFERENCES organizations (organization_name),
    vaccine_id bigserial not null REFERENCES vaccines (vaccine_id),
    vaccine_name text not null REFERENCES vaccines (vaccine_name),
    policy_id bigserial not null REFERENCES nucypher_policies (policy_id),
    file_url text not null,
    price text not null,
    virus_name text not null REFERENCES viruses (virus_name),
    simularity_rate text,
    origin text,
    owner_account text not null REFERENCES nucypher_accounts (address),
    is_active bool not null DEFAULT true,
    is_sold bool not null DEFAULT false,
    created_by text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by text,
    updated_at TIMESTAMP
);

COMMIT;