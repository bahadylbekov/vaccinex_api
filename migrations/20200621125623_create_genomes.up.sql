BEGIN; 

CREATE TABLE genomes (
    genome_id bigserial not null primary key,
    name text not null,
    organization_name text REFERENCES organizations (name),
    file_url text not null,
    virus_name text REFERENCES viruses (name),
    simularity_rate text,
    origin text,
    is_active bool not null DEFAULT true,
    is_sold bool not null DEFAULT false,
    created_by text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_by text,
    updated_at TIMESTAMP,
);

COMMIT;