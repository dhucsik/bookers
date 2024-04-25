-- +goose Up
-- +goose StatementBegin
ALTER TABLE books
ALTER COLUMN rating SET DATA TYPE numeric(5,2);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE quizzes
ALTER COLUMN rating SET DATA TYPE numeric(5,2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
