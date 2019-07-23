CREATE TABLE IF NOT EXISTS ishare_migrations (
    id              SERIAL,
    applied         TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    migration_name  VARCHAR(100),
    PRIMARY KEY (id)
);
