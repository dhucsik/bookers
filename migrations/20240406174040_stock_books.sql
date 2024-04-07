-- +goose Up
-- +goose StatementBegin
CREATE TABLE stock_books (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    book_id INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stock_books;
-- +goose StatementEnd
