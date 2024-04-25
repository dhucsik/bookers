-- +goose Up
-- +goose StatementBegin
ALTER TABLE book_ratings ADD CONSTRAINT unique_book_user_rating UNIQUE (book_id, user_id);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE quiz_ratings ADD CONSTRAINT unique_quiz_user_rating UNIQUE (quiz_id, user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE book_ratings DROP CONSTRAINT unique_book_user_rating;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE quiz_ratings DROP CONSTRAINT unique_quiz_user_rating;
-- +goose StatementEnd