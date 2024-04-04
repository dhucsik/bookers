package quizzes

import (
	"context"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/repositories/books"
	"github.com/dhucsik/bookers/internal/repositories/quizzes"
	"github.com/dhucsik/bookers/internal/repositories/users"
)

type Service interface {
	CreateQuiz(ctx context.Context, quiz *models.Quiz) error
	UpdateQuizTitle(ctx context.Context, userID, quizID int, title string) error
	DeleteQuiz(ctx context.Context, userID, quizID int) error
	AddQuestion(ctx context.Context, userID int, question *models.Question) error
	UpdateQuestion(ctx context.Context, userID int, question *models.Question) error
	DeleteQuestion(ctx context.Context, userID, questionID int) error
	GetQuiz(ctx context.Context, quizID, userID int) (*models.QuizWithFields, error)
	GetQuizWithoutAnswers(ctx context.Context, quizID int) (*models.QuizWithFields, error)
	ListComments(ctx context.Context, quizID int) ([]*models.QuizComment, error)
	DeleteComment(ctx context.Context, commentID, userID int) error
	UpdateComment(ctx context.Context, comment *models.QuizComment) error
	SetRating(ctx context.Context, rating *models.QuizRating) error
	AddComment(ctx context.Context, comment *models.QuizComment) error
}

type service struct {
	quizRepo quizzes.Repository
	userRepo users.Repository
	bookRepo books.Repository
}

func NewService(quizRepo quizzes.Repository) Service {
	return &service{
		quizRepo: quizRepo,
	}
}

func (s *service) CreateQuiz(ctx context.Context, quiz *models.Quiz) error {
	return s.quizRepo.CreateQuiz(ctx, quiz)
}

func (s *service) UpdateQuizTitle(ctx context.Context, userID, quizID int, title string) error {
	quiz, err := s.quizRepo.GetQuiz(ctx, quizID)
	if err != nil {
		return nil
	}

	if quiz.UserID != userID {
		return errors.ErrForbidden
	}

	return s.quizRepo.UpdateQuizTitle(ctx, quizID, title)
}

func (s *service) DeleteQuiz(ctx context.Context, userID, quizID int) error {
	quiz, err := s.quizRepo.GetQuiz(ctx, quizID)
	if err != nil {
		return nil
	}

	if quiz.UserID != userID {
		return errors.ErrForbidden
	}

	return s.quizRepo.DeleteQuiz(ctx, quizID)
}

func (s *service) AddQuestion(ctx context.Context, userID int, question *models.Question) error {
	quiz, err := s.quizRepo.GetQuiz(ctx, question.QuizID)
	if err != nil {
		return nil
	}

	if quiz.UserID != userID {
		return errors.ErrForbidden
	}

	return s.quizRepo.InsertQuestion(ctx, question)
}

func (s *service) UpdateQuestion(ctx context.Context, userID int, question *models.Question) error {
	quiz, err := s.quizRepo.GetQuizByQuestion(ctx, question.ID)
	if err != nil {
		return nil
	}

	if quiz.UserID != userID {
		return errors.ErrForbidden
	}

	return s.quizRepo.UpdateQuestion(ctx, question)
}

func (s *service) DeleteQuestion(ctx context.Context, userID, questionID int) error {
	quiz, err := s.quizRepo.GetQuizByQuestion(ctx, questionID)
	if err != nil {
		return nil
	}

	if quiz.UserID != userID {
		return errors.ErrForbidden
	}

	return s.quizRepo.DeleteQuestion(ctx, questionID)
}

func (s *service) GetQuiz(ctx context.Context, quizID, userID int) (*models.QuizWithFields, error) {
	quiz, err := s.quizRepo.GetQuiz(ctx, quizID)
	if err != nil {
		return nil, err
	}

	if quiz.UserID != userID {
		return nil, errors.ErrForbidden
	}

	user, err := s.userRepo.GetUserByID(ctx, quiz.UserID)
	if err != nil {
		return nil, err
	}

	book, err := s.bookRepo.GetBookByID(ctx, quiz.BookID)
	if err != nil {
		return nil, err
	}

	questions, err := s.quizRepo.GetQuestions(ctx, quizID)
	if err != nil {
		return nil, err
	}

	return &models.QuizWithFields{
		Quiz:      quiz,
		User:      user.ToUserWithoutPassword(),
		Book:      book,
		Questions: questions,
	}, nil
}

func (s *service) GetQuizWithoutAnswers(ctx context.Context, quizID int) (*models.QuizWithFields, error) {
	quiz, err := s.quizRepo.GetQuiz(ctx, quizID)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetUserByID(ctx, quiz.UserID)
	if err != nil {
		return nil, err
	}

	book, err := s.bookRepo.GetBookByID(ctx, quiz.BookID)
	if err != nil {
		return nil, err
	}

	questions, err := s.quizRepo.GetQuestions(ctx, quizID)
	if err != nil {
		return nil, err
	}

	for _, question := range questions {
		question.Answer = ""
	}

	return &models.QuizWithFields{
		Quiz:      quiz,
		User:      user.ToUserWithoutPassword(),
		Book:      book,
		Questions: questions,
	}, nil
}

func (s *service) AddComment(ctx context.Context, comment *models.QuizComment) error {
	return s.quizRepo.InsertComment(ctx, comment.QuizID, comment.UserID, comment.Comment)
}

func (s *service) UpdateComment(ctx context.Context, comment *models.QuizComment) error {
	com, err := s.quizRepo.GetComment(ctx, comment.ID)
	if err != nil {
		return err
	}

	if com.UserID != comment.UserID {
		return errors.ErrForbidden
	}

	return s.quizRepo.UpdateComment(ctx, comment.ID, comment.Comment)
}

func (s *service) SetRating(ctx context.Context, rating *models.QuizRating) error {
	return s.quizRepo.SetRating(ctx, rating.QuizID, rating.UserID, rating.Rating)
}

func (s *service) DeleteComment(ctx context.Context, commentID, userID int) error {
	comment, err := s.quizRepo.GetComment(ctx, commentID)
	if err != nil {
		return err
	}

	if comment.UserID != userID {
		return errors.ErrForbidden
	}

	return s.quizRepo.DeleteComment(ctx, commentID)
}

func (s *service) ListComments(ctx context.Context, quizID int) ([]*models.QuizComment, error) {
	return s.quizRepo.ListComments(ctx, quizID)
}
