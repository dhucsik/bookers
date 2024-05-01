-- +goose Up
-- +goose StatementBegin
CREATE TABLE liked_books (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    book_id INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE liked_books;
-- +goose StatementEnd
