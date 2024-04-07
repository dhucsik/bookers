-- +goose Up
-- +goose StatementBegin
CREATE TABLE share_requests (
    id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL,
    receiver_id INT NOT NULL,
    sender_book_id INT NOT NULL,
    receiver_book_id INT NOT NULL,
    sender_status VARCHAR(100) NOT NULL,
    receiver_status VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE share_requests;
-- +goose StatementEnd
