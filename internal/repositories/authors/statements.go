package authors

const (
	listAuthorsStmt = `SELECT id, name FROM authors 
	ORDER BY id ASC 
	WHERE name ILIKE '%' || $1 || '%'
	LIMIT $2 OFFSET $3`

	createAuthorStmt = `INSERT INTO authors (name) VALUES ($1)`

	deleteAuthorStmt = `DELETE FROM authors WHERE id = $1`

	listAuthorsByBookIDStmt = `SELECT a.id, a.name FROM authors a
	JOIN book_authors ba ON a.id = ba.author_id
	WHERE ba.book_id = $1`

	listAuthorsByBookIDsStmt = `SELECT ba.book_id, a.id, a.name FROM authors a
	JOIN book_authors ba ON a.id = ba.author_id
	WHERE ba.book_id = ANY($1)`
)
