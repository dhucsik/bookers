-- +goose Up
-- +goose StatementBegin
INSERT INTO categories (name)
VALUES
('Проза'),
('Фантастика'),
('Фэнтези'),
('Любовные романы'),
('Детективы'),
('Мистика'),
('Психологические романы'),
('Триллеры'),
('Исторические романы'),
('Повести, рассказы'),
('Биографии людей'),
('Мемуары'),
('Манга, комиксы и артбуки (издания для взрослых)');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
