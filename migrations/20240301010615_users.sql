-- +goose Up
-- +goose StatementBegin

CREATE TYPE role AS ENUM ('admin', 'user');

-- +goose StatementEnd

-- +goose StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role role NOT NULL DEFAULT 'user',
    city VARCHAR(100)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TYPE role;
-- +goose StatementEnd
