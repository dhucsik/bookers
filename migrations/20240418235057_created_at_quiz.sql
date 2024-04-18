-- +goose Up
-- +goose StatementBegin
ALTER TABLE quizzes ADD COLUMN created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE quizzes DROP COLUMN created_at;
-- +goose StatementEnd
