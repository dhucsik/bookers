-- +goose Up
-- +goose StatementBegin
ALTER TABLE liked_books 
    ADD CONSTRAINT unique_liked_books UNIQUE (user_id, book_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
