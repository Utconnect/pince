-- +goose Up
-- +goose StatementBegin
CREATE TABLE files
(
    id               SERIAL PRIMARY KEY,
    name             VARCHAR(255)             NOT NULL,
    type             VARCHAR(255)             NOT NULL,
    storage_platform VARCHAR(255)             NOT NULL,
    location         TEXT                     NOT NULL,
    size             INTEGER                  NOT NULL,
    created_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at       TIMESTAMP WITH TIME ZONE
) WITH (OIDS = FALSE);

CREATE INDEX idx_files_deleted_at ON files (deleted_at);
CREATE INDEX idx_files_name ON files (name);
-- OIDS are not typically used in PostgreSQL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
-- +goose StatementEnd
