-- +goose Up
-- +goose StatementBegin
CREATE TABLE quiz_ratings (
    id SERIAL PRIMARY KEY,
    quiz_id INT NOT NULL,
    user_id INT NOT NULL,
    rating INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE quiz_comments (
    id SERIAL PRIMARY KEY,
    quiz_id INT NOT NULL,
    user_id INT NOT NULL,
    comment TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE quiz_comments;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE quiz_ratings;
-- +goose StatementEnd
