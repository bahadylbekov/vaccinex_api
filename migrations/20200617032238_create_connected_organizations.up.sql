BEGIN; 

CREATE TABLE connected_organizations (
    id bigserial REFERENCES organizations (organization_id) not null,
    organization_id bigserial REFERENCES organizations (organization_id) not null,
    created_by text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id, organization_id),
    CHECK (id < organization_id)
);

COMMIT;