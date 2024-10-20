-- +goose Up
-- +goose StatementBegin

CREATE TABLE translations (
    id SERIAL PRIMARY KEY,
    lexeme TEXT NOT NULL,
    translated_lexeme TEXT,
    source_language TEXT NOT NULL,
    target_language TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS translations;
-- +goose StatementEnd
