-- +goose Up
-- +goose StatementBegin

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    is_approved BOOLEAN NOT NULL,
    user_email TEXT NOT NULL,
    user_phone TEXT NOT NULL,
    rating INT NOT NULL,
    bank_id INT NOT NULL,
    user_name TEXT NOT NULL,
    date TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reviews;
-- +goose StatementEnd
