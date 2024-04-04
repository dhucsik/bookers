-- +goose Up
-- +goose StatementBegin
ALTER TABLE books ADD COLUMN image VARCHAR(255) DEFAULT '';
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE books ADD COLUMN description TEXT DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
