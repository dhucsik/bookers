package quizzes

const (
	getQuizStmt = `SELECT id, user_id, book_id, title, rating FROM quizzes WHERE id = $1`

	getQuestionsStmt = `SELECT id, quiz_id, question, options, answer
	FROM questions WHERE quiz_id = $1`

	createQuizStmt = `INSERT INTO quizzes (user_id, book_id, title, rating)
	VALUES ($1, $2, $3, $4)`

	updateQuizTitleStmt = `UPDATE quizzes SET title = $2 WHERE id = $1`

	deleteQuizStmt = `DELETE FROM quizzes WHERE id = $1`

	createQuestionStmt = `INSERT INTO questions (quiz_id, question, options, answer)
	VALUES ($1, $2, $3, $4)`

	updateQuestionStmt = `UPDATE questions SET question = $2, options = $3, answer = $4
	WHERE id = $1`

	deleteQuestionStmt = `DELETE FROM questions WHERE id = $1`

	insertCommentStmt = `INSERT INTO quiz_comments (quiz_id, user_id, comment) VALUES ($1, $2, $3)`

	updateCommentStmt = `UPDATE quiz_comments SET comment = $2, updated_at = NOW() WHERE id = $1`

	listCommentsStmt = `SELECT id, quiz_id, user_id, comment, created_at FROM quiz_comments WHERE quiz_id = $1`

	insertQuizRatingStmt = `INSERT INTO quiz_ratings (quiz_id, user_id, rating) VALUES ($1, $2, $3)
	ON CONFLICT (quiz_id, user_id) DO UPDATE SET rating = $3`

	getAvgRatingStmt = `SELECT AVG(rating) FROM quiz_ratings WHERE quiz_id = $1`

	updateQuizRatingStmt = `UPDATE quizzes SET rating = $1 WHERE id = $2`

	deleteCommentStmt = `DELETE FROM quiz_comments WHERE id = $1`

	getCommentStmt = `SELECT id, quiz_id, user_id, comment, created_at FROM quiz_comments WHERE id = $1`

	getQuizByQuestionStmt = `SELECT q.id, q.user_id, q.book_id, q.title, q.rating
	FROM quizzes q JOIN questions qu ON q.id = qu.quiz_id WHERE qu.id = $1`
)
