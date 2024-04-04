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

	insertCommentStmt = `INSERT INTO book_comments (book_id, user_id, comment) VALUES ($1, $2, $3)`

	updateCommentStmt = `UPDATE book_comments SET comment = $2, updated_at = NOW() WHERE id = $1`

	listCommentsStmt = `SELECT id, book_id, user_id, comment, created_at FROM book_comments WHERE book_id = $1`

	insertBookRatingStmt = `INSERT INTO book_ratings (book_id, user_id, rating) VALUES ($1, $2, $3)
	ON CONFLICT (book_id, user_id) DO UPDATE SET rating = $3`

	getAvgRatingStmt = `SELECT AVG(rating) FROM book_ratings WHERE book_id = $1`

	updateBookRatingStmt = `UPDATE books SET rating = $1 WHERE id = $2`

	deleteCommentStmt = `DELETE FROM book_comments WHERE id = $1`

	getCommentStmt = `SELECT id, book_id, user_id, comment, created_at FROM book_comments WHERE id = $1`
)
