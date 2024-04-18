package quizzes

import (
	"context"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/repositories/books"
	"github.com/dhucsik/bookers/internal/repositories/quizzes"
	"github.com/dhucsik/bookers/internal/repositories/users"
	"github.com/samber/lo"
)

type Service interface {
	CreateQuiz(ctx context.Context, quiz *models.Quiz) (int, error)
	UpdateQuizTitle(ctx context.Context, userID, quizID int, title string) error
	DeleteQuiz(ctx context.Context, userID, quizID int) error
	AddQuestion(ctx context.Context, userID int, question *models.Question) (int, error)
	UpdateQuestion(ctx context.Context, userID int, question *models.Question) error
	DeleteQuestion(ctx context.Context, userID, questionID int) error
	GetQuiz(ctx context.Context, quizID, userID int) (*models.QuizWithFields, error)
	GetQuizWithoutAnswers(ctx context.Context, quizID int) (*models.QuizWithFields, error)
	ListComments(ctx context.Context, quizID int) ([]*models.QuizComment, error)
	DeleteComment(ctx context.Context, commentID, userID int) error
	UpdateComment(ctx context.Context, comment *models.QuizComment) error
	SetRating(ctx context.Context, rating *models.QuizRating) error
	AddComment(ctx context.Context, comment *models.QuizComment) (int, error)
	CheckQuiz(ctx context.Context, userID, quizID int, userAnswers []*models.UserAnswer) (*models.QuizWithQuestionResults, error)
	GetQuizResults(ctx context.Context, userID int) ([]*models.QuizResultWithFields, error)
	GetQuizResultWithAnswers(ctx context.Context, resultID int) (*models.QuizQuestionWithFields, error)
	ListQuizzes(ctx context.Context, limit, offset int) ([]*models.QuizWithBase, int, error)
	ListQuizzesByBookID(ctx context.Context, bookID int) ([]*models.QuizWithBase, error)
}

type service struct {
	quizRepo quizzes.Repository
	userRepo users.Repository
	bookRepo books.Repository
}

func NewService(
	quizRepo quizzes.Repository,
	bookRepo books.Repository,
	userRepo users.Repository,
) Service {
	return &service{
		quizRepo: quizRepo,
		userRepo: userRepo,
		bookRepo: bookRepo,
	}
}

func (s *service) CreateQuiz(ctx context.Context, quiz *models.Quiz) (int, error) {
	return s.quizRepo.CreateQuiz(ctx, quiz)
}

func (s *service) UpdateQuizTitle(ctx context.Context, userID, quizID int, title string) error {
	quiz, err := s.quizRepo.GetQuiz(ctx, quizID)
	if err != nil {
		return nil
	}

	if quiz.UserID != userID {
		return errors.ErrForbiddenForUser
	}

	return s.quizRepo.UpdateQuizTitle(ctx, quizID, title)
}

func (s *service) DeleteQuiz(ctx context.Context, userID, quizID int) error {
	quiz, err := s.quizRepo.GetQuiz(ctx, quizID)
	if err != nil {
		return nil
	}

	if quiz.UserID != userID {
		return errors.ErrForbiddenForUser
	}

	return s.quizRepo.DeleteQuiz(ctx, quizID)
}

func (s *service) AddQuestion(ctx context.Context, userID int, question *models.Question) (int, error) {
	quiz, err := s.quizRepo.GetQuiz(ctx, question.QuizID)
	if err != nil {
		return 0, err
	}

	if quiz.UserID != userID {
		return 0, errors.ErrForbiddenForUser
	}

	return s.quizRepo.InsertQuestion(ctx, question)
}

func (s *service) UpdateQuestion(ctx context.Context, userID int, question *models.Question) error {
	quiz, err := s.quizRepo.GetQuizByQuestion(ctx, question.ID)
	if err != nil {
		return nil
	}

	if quiz.UserID != userID {
		return errors.ErrForbiddenForUser
	}

	return s.quizRepo.UpdateQuestion(ctx, question)
}

func (s *service) DeleteQuestion(ctx context.Context, userID, questionID int) error {
	quiz, err := s.quizRepo.GetQuizByQuestion(ctx, questionID)
	if err != nil {
		return nil
	}

	if quiz.UserID != userID {
		return errors.ErrForbiddenForUser
	}

	return s.quizRepo.DeleteQuestion(ctx, questionID)
}

func (s *service) GetQuiz(ctx context.Context, quizID, userID int) (*models.QuizWithFields, error) {
	quiz, err := s.quizRepo.GetQuiz(ctx, quizID)
	if err != nil {
		return nil, err
	}

	if quiz.UserID != userID {
		return nil, errors.ErrForbiddenForUser
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

func (s *service) AddComment(ctx context.Context, comment *models.QuizComment) (int, error) {
	return s.quizRepo.InsertComment(ctx, comment.QuizID, comment.UserID, comment.Comment)
}

func (s *service) UpdateComment(ctx context.Context, comment *models.QuizComment) error {
	com, err := s.quizRepo.GetComment(ctx, comment.ID)
	if err != nil {
		return err
	}

	if com.UserID != comment.UserID {
		return errors.ErrForbiddenForUser
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
		return errors.ErrForbiddenForUser
	}

	return s.quizRepo.DeleteComment(ctx, commentID)
}

func (s *service) ListComments(ctx context.Context, quizID int) ([]*models.QuizComment, error) {
	return s.quizRepo.ListComments(ctx, quizID)
}

func (s *service) ListQuizzes(ctx context.Context, limit, offset int) ([]*models.QuizWithBase, int, error) {
	quizzes, totalCount, err := s.quizRepo.ListQuizzes(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	bookIDs := lo.Map(quizzes, func(quiz *models.Quiz, _ int) int {
		return quiz.BookID
	})

	books, err := s.bookRepo.GetBooksByIDs(ctx, bookIDs)
	if err != nil {
		return nil, 0, err
	}

	users, err := s.userRepo.GetUsersByIDs(ctx, lo.Map(quizzes, func(quiz *models.Quiz, _ int) int {
		return quiz.UserID
	}))
	if err != nil {
		return nil, 0, err
	}

	bookMap := lo.SliceToMap(books, func(book *models.Book) (int, *models.Book) {
		return book.ID, book
	})

	userMap := lo.SliceToMap(users, func(user *models.User) (int, *models.User) {
		return user.ID, user
	})

	out := lo.Map(quizzes, func(quiz *models.Quiz, _ int) *models.QuizWithBase {
		return &models.QuizWithBase{
			Quiz: quiz,
			Book: bookMap[quiz.BookID],
			User: userMap[quiz.UserID].ToUserWithoutPassword(),
		}
	})

	return out, totalCount, nil
}

func (s *service) ListQuizzesByBookID(ctx context.Context, bookID int) ([]*models.QuizWithBase, error) {
	quizzes, err := s.quizRepo.ListQuizzesByBookID(ctx, bookID)
	if err != nil {
		return nil, err
	}

	users, err := s.userRepo.GetUsersByIDs(ctx, lo.Map(quizzes, func(quiz *models.Quiz, _ int) int {
		return quiz.UserID
	}))
	if err != nil {
		return nil, err
	}

	book, err := s.bookRepo.GetBookByID(ctx, bookID)
	if err != nil {
		return nil, err
	}

	userMap := lo.SliceToMap(users, func(user *models.User) (int, *models.User) {
		return user.ID, user
	})

	out := lo.Map(quizzes, func(quiz *models.Quiz, _ int) *models.QuizWithBase {
		return &models.QuizWithBase{
			Quiz: quiz,
			Book: book,
			User: userMap[quiz.UserID].ToUserWithoutPassword(),
		}
	})

	return out, nil
}
