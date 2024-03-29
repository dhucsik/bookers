-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    pub_date DATE NOT NULL,
    edition VARCHAR(50) NOT NULL,
    language VARCHAR(50) NOT NULL,
    rating NUMERIC(2, 2) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE books_authors (
    book_id INT NOT NULL,
    author_id INT NOT NULL,
    PRIMARY KEY (book_id, author_id),
    FOREIGN KEY (book_id) REFERENCES books (id),
    FOREIGN KEY (author_id) REFERENCES authors (id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE books_categories (
    book_id INT NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY (book_id, category_id),
    FOREIGN KEY (book_id) REFERENCES books (id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books_categories;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE books_authors;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
