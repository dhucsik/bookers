package authors

const (
	listAuthorsStmt = `SELECT id, name, COUNT(*) OVER() FROM authors 
	WHERE name ILIKE '%' || $1 || '%'
	ORDER BY id ASC 
	LIMIT $2 OFFSET $3`

	createAuthorStmt = `INSERT INTO authors (name) VALUES ($1) RETURNING id`

	deleteAuthorStmt = `DELETE FROM authors WHERE id = $1`

	listAuthorsByBookIDStmt = `SELECT a.id, a.name FROM authors a
	JOIN books_authors ba ON a.id = ba.author_id
	WHERE ba.book_id = $1`

	listAuthorsByBookIDsStmt = `SELECT ba.book_id, a.id, a.name FROM authors a
	JOIN books_authors ba ON a.id = ba.author_id
	WHERE ba.book_id = ANY($1)`
)
