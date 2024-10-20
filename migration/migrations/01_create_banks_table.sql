-- +goose Up
-- +goose StatementBegin

CREATE TABLE banks (
    id SERIAL PRIMARY KEY,
    external_id INT NOT NULL,
    external_legacy_id INT NOT NULL,
    name TEXT NOT NULL,
    logo TEXT NOT NULL,
    url TEXT NOT NULL,
    bank_data JSONB NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS banks;
-- +goose StatementEnd
