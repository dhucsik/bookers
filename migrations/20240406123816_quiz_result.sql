-- +goose Up
-- +goose StatementBegin
CREATE TABLE quiz_results (
    id SERIAL PRIMARY KEY,
    quiz_id INT NOT NULL,
    user_id INT NOT NULL,
    coorect INT NOT NULL,
    total INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE question_results (
    id SERIAL PRIMARY KEY,
    quiz_result_id INT NOT NULL,
    quiestion_id INT NOT NULL,
    user_answer VARCHAR(255) NOT NULL,
    is_correct BOOLEAN NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE question_results;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE quiz_results;
-- +goose StatementEnd