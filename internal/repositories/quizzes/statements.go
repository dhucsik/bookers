package quizzes

const (
	getQuizStmt = `SELECT q.id, q.user_id, q.book_id, q.title, q.rating, q.created_at as total_quizzes,
	(SELECT COUNT(*) FROM questions WHERE quiz_id = q.id) as total_questions
	FROM quizzes q
	WHERE q.id = $1`

	getQuizzesStmt = `SELECT q.id, q.user_id, q.book_id, q.title, q.rating, q.created_at as total_quizzes,
	(SELECT COUNT(*) FROM questions WHERE quiz_id = q.id) as total_questions
	FROM quizzes q
	WHERE q.id = ANY($1)`

	getQuestionsStmt = `SELECT id, quiz_id, question, options, answer
	FROM questions WHERE quiz_id = $1`

	listQuizzesStmt = `SELECT q.id, q.user_id, q.book_id, q.title, q.rating, q.created_at, COUNT(*) OVER() as total_quizzes,
						(SELECT COUNT(*) FROM questions WHERE quiz_id = q.id) as total_questions
						FROM quizzes q
						ORDER BY q.id
						LIMIT $1 OFFSET $2
`

	listQuizzesByBookIDStmt = `SELECT q.id, q.user_id, q.book_id, q.title, q.rating, q.created_at as total_quizzes,
	(SELECT COUNT(*) FROM questions WHERE quiz_id = q.id) as total_questions
	FROM quizzes q
	WHERE q.book_id = $1`

	listQuizzesByUserIDStmt = `SELECT q.id, q.user_id, q.book_id, q.title, q.rating, q.created_at as total_quizzes,
	(SELECT COUNT(*) FROM questions WHERE quiz_id = q.id) as total_questions
	FROM quizzes q
	WHERE q.user_id = $1`

	createQuizStmt = `INSERT INTO quizzes (user_id, book_id, title, rating)
	VALUES ($1, $2, $3, $4) RETURNING id`

	updateQuizTitleStmt = `UPDATE quizzes SET title = $2 WHERE id = $1`

	deleteQuizStmt = `DELETE FROM quizzes WHERE id = $1`

	createQuestionStmt = `INSERT INTO questions (quiz_id, question, options, answer)
	VALUES ($1, $2, $3, $4) RETURNING id`

	updateQuestionStmt = `UPDATE questions SET question = $2, options = $3, answer = $4
	WHERE id = $1`

	deleteQuestionStmt = `DELETE FROM questions WHERE id = $1`

	insertCommentStmt = `INSERT INTO quiz_comments (quiz_id, user_id, comment) VALUES ($1, $2, $3) RETURNING id`

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

	insertQuizResultStmt = `INSERT INTO quiz_results (quiz_id, user_id, coorect, total) VALUES ($1, $2, $3, $4) RETURNING id`

	insertQuestionResultStmt = `INSERT INTO question_results (quiz_result_id, quiestion_id, user_answer, is_correct)
	VALUES ($1, $2, $3, $4)`

	getQuizResultsStmt = `SELECT id, quiz_id, user_id, coorect, total FROM quiz_results WHERE user_id = $1`

	getQuizResultStmt = `SELECT id, quiz_id, user_id, coorect, total FROM quiz_results WHERE id = $1`

	getQuestionResultsStmt = `SELECT id, quiz_result_id, quiestion_id, user_answer, is_correct
	FROM question_results WHERE quiz_result_id = $1`
)
