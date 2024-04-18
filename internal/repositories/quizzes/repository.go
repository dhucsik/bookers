package quizzes

import (
	"context"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type Repository interface {
	CreateQuiz(ctx context.Context, quiz *models.Quiz) (int, error)
	GetQuiz(ctx context.Context, quizID int) (*models.Quiz, error)
	GetQuizzes(ctx context.Context, quizIDs []int) ([]*models.Quiz, error)
	UpdateQuizTitle(ctx context.Context, quizID int, title string) error
	DeleteQuiz(ctx context.Context, quizID int) error
	InsertQuestion(ctx context.Context, question *models.Question) (int, error)
	UpdateQuestion(ctx context.Context, question *models.Question) error
	DeleteQuestion(ctx context.Context, questionID int) error
	GetQuestions(ctx context.Context, quizID int) ([]*models.Question, error)
	SetRating(ctx context.Context, quizID, userID, rating int) error
	InsertComment(ctx context.Context, quizID, userID int, comment string) (int, error)
	UpdateComment(ctx context.Context, id int, comment string) error
	ListComments(ctx context.Context, quizID int) ([]*models.QuizComment, error)
	DeleteComment(ctx context.Context, id int) error
	GetComment(ctx context.Context, id int) (*models.QuizComment, error)
	GetQuizByQuestion(ctx context.Context, questionID int) (*models.Quiz, error)
	SaveResults(ctx context.Context, result *models.QuizWithQuestionResults) error
	GetQuizResults(ctx context.Context, userID int) ([]*models.QuizResult, error)
	GetQuizResult(ctx context.Context, id int) (*models.QuizWithQuestionResults, error)
	ListQuizzes(ctx context.Context, limit, offset int) ([]*models.Quiz, int, error)
	ListQuizzesByBookID(ctx context.Context, bookID int) ([]*models.Quiz, error)
	ListQuizzesByUserID(ctx context.Context, userID int) ([]*models.Quiz, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetQuiz(ctx context.Context, quizID int) (*models.Quiz, error) {
	var quiz models.Quiz
	err := r.db.QueryRow(ctx, getQuizStmt, quizID).Scan(&quiz.ID, &quiz.UserID, &quiz.BookID, &quiz.Title, &quiz.Rating, &quiz.CreatedAt, &quiz.QuestionsCount)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.ErrQuizNotFound
		}

		return nil, err
	}
	return &quiz, nil
}

func (r *repository) GetQuizzes(ctx context.Context, quizIDs []int) ([]*models.Quiz, error) {
	rows, err := r.db.Query(ctx, getQuizzesStmt, pq.Array(quizIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models.Quiz
	for rows.Next() {
		quiz := &models.Quiz{}
		if err := rows.Scan(&quiz.ID, &quiz.UserID, &quiz.BookID, &quiz.Title, &quiz.Rating, &quiz.CreatedAt, &quiz.QuestionsCount); err != nil {
			return nil, err
		}

		out = append(out, quiz)
	}

	return out, nil
}

func (r *repository) CreateQuiz(ctx context.Context, quiz *models.Quiz) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, createQuizStmt, quiz.UserID, quiz.BookID, quiz.Title, quiz.Rating).Scan(&id)
	return id, err
}

func (r *repository) UpdateQuizTitle(ctx context.Context, quizID int, title string) error {
	_, err := r.db.Exec(ctx, updateQuizTitleStmt, title, quizID)
	return err
}

func (r *repository) DeleteQuiz(ctx context.Context, quizID int) error {
	_, err := r.db.Exec(ctx, deleteQuizStmt, quizID)
	return err
}

func (r *repository) InsertQuestion(ctx context.Context, question *models.Question) (int, error) {
	model, err := newQuestionModel(question)
	if err != nil {
		return 0, err
	}

	var id int
	err = r.db.QueryRow(ctx, createQuestionStmt, model.QuizID, model.Question, string(model.Options), model.Answer).Scan(&id)
	return id, err
}

func (r *repository) UpdateQuestion(ctx context.Context, question *models.Question) error {
	model, err := newQuestionModel(question)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, updateQuestionStmt, model.ID, model.Question, string(model.Options), model.Answer)
	return err
}

func (r *repository) DeleteQuestion(ctx context.Context, questionID int) error {
	_, err := r.db.Exec(ctx, deleteQuestionStmt, questionID)
	return err
}

func (r *repository) GetQuestions(ctx context.Context, quizID int) ([]*models.Question, error) {
	out := make([]*models.Question, 0)

	rows, err := r.db.Query(ctx, getQuestionsStmt, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var model questionModel
		err := rows.Scan(&model.ID, &model.QuizID, &model.Question, &model.Options, &model.Answer)
		if err != nil {
			return nil, err
		}

		outModel, err := model.convert()
		if err != nil {
			return nil, err
		}
		out = append(out, outModel)
	}

	return out, nil
}

func (r *repository) InsertComment(ctx context.Context, quizID, userID int, comment string) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, insertCommentStmt, quizID, userID, comment).Scan(&id)
	return id, err
}

func (r *repository) UpdateComment(ctx context.Context, id int, comment string) error {
	_, err := r.db.Exec(ctx, updateCommentStmt, id, comment)
	return err
}

func (r *repository) SetRating(ctx context.Context, quizID, userID, rating int) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, insertQuizRatingStmt, quizID, userID, rating)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	var avgRating float64
	err = tx.QueryRow(ctx, getAvgRatingStmt, quizID).Scan(&avgRating)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, updateQuizRatingStmt, avgRating, quizID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (r *repository) ListComments(ctx context.Context, quizID int) ([]*models.QuizComment, error) {
	rows, err := r.db.Query(ctx, listCommentsStmt, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models.QuizComment
	for rows.Next() {
		comment := &quizComment{}
		if err := rows.Scan(&comment.ID, &comment.QuizID, &comment.UserID, &comment.Comment, &comment.CreatedAt); err != nil {
			return nil, err
		}

		out = append(out, comment.convert())
	}

	return out, nil
}

func (r *repository) DeleteComment(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, deleteCommentStmt, id)
	return err
}

func (r *repository) GetComment(ctx context.Context, id int) (*models.QuizComment, error) {
	comment := &quizComment{}
	err := r.db.QueryRow(ctx, getCommentStmt, id).Scan(&comment.ID, &comment.QuizID, &comment.UserID, &comment.Comment, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return comment.convert(), nil
}

func (r *repository) GetQuizByQuestion(ctx context.Context, questionID int) (*models.Quiz, error) {
	var quiz models.Quiz
	err := r.db.QueryRow(ctx, getQuizByQuestionStmt, questionID).Scan(&quiz.ID, &quiz.UserID, &quiz.BookID, &quiz.Title, &quiz.Rating)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &quiz, nil
}

func (r *repository) ListQuizzes(ctx context.Context, limit, offset int) ([]*models.Quiz, int, error) {
	var totalCount int

	rows, err := r.db.Query(ctx, listQuizzesStmt, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var out []*models.Quiz
	for rows.Next() {
		quiz := &models.Quiz{}
		if err := rows.Scan(&quiz.ID, &quiz.UserID, &quiz.BookID, &quiz.Title, &quiz.Rating, &quiz.CreatedAt, &totalCount, &quiz.QuestionsCount); err != nil {
			return nil, 0, err
		}

		out = append(out, quiz)
	}

	return out, totalCount, nil
}

func (r *repository) ListQuizzesByBookID(ctx context.Context, bookID int) ([]*models.Quiz, error) {
	rows, err := r.db.Query(ctx, listQuizzesByBookIDStmt, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models.Quiz
	for rows.Next() {
		quiz := &models.Quiz{}
		if err := rows.Scan(&quiz.ID, &quiz.UserID, &quiz.BookID, &quiz.Title, &quiz.Rating, &quiz.CreatedAt, &quiz.QuestionsCount); err != nil {
			return nil, err
		}

		out = append(out, quiz)
	}

	return out, nil
}

func (r *repository) ListQuizzesByUserID(ctx context.Context, userID int) ([]*models.Quiz, error) {
	rows, err := r.db.Query(ctx, listQuizzesByUserIDStmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models.Quiz
	for rows.Next() {
		quiz := &models.Quiz{}
		if err := rows.Scan(&quiz.ID, &quiz.UserID, &quiz.BookID, &quiz.Title, &quiz.Rating, &quiz.CreatedAt, &quiz.QuestionsCount); err != nil {
			return nil, err
		}

		out = append(out, quiz)
	}

	return out, nil
}
