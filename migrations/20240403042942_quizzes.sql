-- +goose Up
-- +goose StatementBegin
CREATE TABLE quizzes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    book_id INT NOT NULL,
    title VARCHAR(100) NOT NULL,
    rating NUMERIC(2, 2) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE quizzes;
-- +goose StatementEnd
