-- +goose Up
-- +goose StatementBegin
CREATE TABLE friends (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    friend_id INT NOT NULL,
    status VARCHAR(50) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE friends ADD CONSTRAINT unique_friendship UNIQUE (user_id, friend_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE friends;
-- +goose StatementEnd
