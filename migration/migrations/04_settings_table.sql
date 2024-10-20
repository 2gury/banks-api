-- +goose Up
-- +goose StatementBegin

CREATE TABLE settings (
    id SERIAL PRIMARY KEY,
    automoderation_strategy BOOLEAN NOT NULL
);

INSERT INTO settings (id, automoderation_strategy) VALUES (1, false)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS settings;
-- +goose StatementEnd
