package books

const (
	createBookStmt = `INSERT INTO books (title, pub_date, edition, language, rating, image, description)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	createBookAuthorStmt = `INSERT INTO books_authors (book_id, author_id) VALUES ($1, $2)`

	createBookCategoryStmt = `INSERT INTO books_categories (book_id, category_id) VALUES ($1, $2)`

	countBooksStmt = `SELECT COUNT(*) FROM books`

	listBooksStmt = `SELECT b.id, b.title, b.pub_date, b.edition, b.language, b.rating, b.image, b.description, COUNT(*) OVER()
	FROM books b 
	WHERE b.title ILIKE '%' || $1 || '%'
	ORDER BY b.id
	LIMIT $2 OFFSET $3`

	getBookStmt = `SELECT b.id, b.title, b.pub_date, b.edition, b.language, b.rating, b.image, b.description
	FROM books b WHERE b.id = $1`

	getBooksByIDsStmt = `SELECT b.id, b.title, b.pub_date, b.edition, b.language, b.rating, b.image, b.description
	FROM books b WHERE b.id = ANY($1)`

	insertCommentStmt = `INSERT INTO book_comments (book_id, user_id, comment) VALUES ($1, $2, $3) RETURNING id`

	updateCommentStmt = `UPDATE book_comments SET comment = $2, updated_at = NOW() WHERE id = $1`

	listCommentsStmt = `SELECT id, book_id, user_id, comment, created_at FROM book_comments WHERE book_id = $1`

	insertBookRatingStmt = `INSERT INTO book_ratings (book_id, user_id, rating) VALUES ($1, $2, $3)
	ON CONFLICT (book_id, user_id) DO UPDATE SET rating = $3`

	getAvgRatingStmt = `SELECT ROUND(AVG(rating), 2) as average_rating FROM book_ratings WHERE book_id = $1`

	updateBookRatingStmt = `UPDATE books SET rating = $1 WHERE id = $2`

	deleteCommentStmt = `DELETE FROM book_comments WHERE id = $1`

	getCommentStmt = `SELECT id, book_id, user_id, comment, created_at FROM book_comments WHERE id = $1`

	uploadStockBookStmt = `INSERT INTO stock_books (user_id, book_id) VALUES ($1, $2) RETURNING id`

	getStockBookStmt = `SELECT id, user_id, book_id FROM stock_books WHERE id = $1`

	getBooksByStockIDsStmt = `SELECT b.id, b.title, b.pub_date, b.edition, b.language, b.rating, b.image, b.description, sb.id
	FROM books b JOIN stock_books sb ON b.id = sb.book_id WHERE sb.id = ANY($1)`

	getStockBooksByUserStmt = `SELECT id, user_id, book_id FROM stock_books WHERE user_id = $1`

	getStockByBookStmt = `SELECT id, user_id, book_id FROM stock_books WHERE book_id = $1`

	getUserStockCountStmt = `SELECT COUNT(*) FROM stock_books WHERE user_id = $1`

	deleteStockBookStmt = `DELETE FROM stock_books WHERE id = $1`

	createNewRequestStmt = `INSERT INTO share_requests (sender_id, receiver_id, sender_book_id, receiver_book_id, sender_status, receiver_status)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	updateRequestStmt = `UPDATE share_requests 
	SET sender_status = $2, 
	receiver_status = $3, 
	sender_book_id = $4,
	updated_at = NOW() WHERE id = $1`

	updateStockBookUserStmt = `UPDATE stock_books SET user_id = $2 WHERE id = $1`

	getRequestStmt = `SELECT id, sender_id, receiver_id, sender_book_id, receiver_book_id, sender_status, receiver_status, created_at, updated_at
	FROM share_requests WHERE id = $1`

	getRequestsStmt = `SELECT id, sender_id, receiver_id, sender_book_id, receiver_book_id, sender_status, receiver_status, created_at, updated_at
	FROM share_requests WHERE sender_id = $1 OR receiver_id = $1`

	getSuccessRequestCountStmt = `SELECT COUNT(*) FROM share_requests WHERE (sender_id = $1 OR receiver_id = $1) AND sender_status = 'sender_proved' AND receiver_status = 'receiver_proved'`
)
