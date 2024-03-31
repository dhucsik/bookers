package books

const (
	createBookStmt = `INSERT INTO books (title, pub_date, edition, language, rating)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	createBookAuthorStmt = `INSERT INTO books_authors (book_id, author_id) VALUES ($1, $2)`

	createBookCategoryStmt = `INSERT INTO books_categories (book_id, category_id) VALUES ($1, $2)`

	listBooksStmt = `SELECT b.id, b.title, b.pub_date, b.edition, b.language, b.rating
	FROM books b 
	WHERE b.title ILIKE '%' || $1 || '%'
	ORDER BY b.id
	LIMIT $2 OFFSET $3`

	getBookStmt = `SELECT b.id, b.title, b.pub_date, b.edition, b.language, b.rating
	FROM books b WHERE b.id = $1`
)
