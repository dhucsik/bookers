package quizzes

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/samber/lo"
)

func (s *service) CheckQuiz(ctx context.Context, userID, quizID int, userAnswers []*models.UserAnswer) (*models.QuizWithQuestionResults, error) {
	questions, err := s.quizRepo.GetQuestions(ctx, quizID)
	if err != nil {
		return nil, err
	}

	correctAnswers := lo.SliceToMap(questions, func(item *models.Question) (int, string) {
		return item.ID, item.Answer
	})

	quizResult := &models.QuizResult{
		QuizID: quizID,
		UserID: userID,
	}

	var correct, total int
	questionResults := make([]*models.QuestionResult, 0, len(userAnswers))
	for _, userAnswer := range userAnswers {
		var isCorrect bool

		if correctAnswers[userAnswer.QuestionID] == userAnswer.UserAnswer {
			correct++
			isCorrect = true
		}
		total++

		questionResults = append(questionResults, &models.QuestionResult{
			QuestionID: userAnswer.QuestionID,
			UserAnswer: userAnswer.UserAnswer,
			IsCorrect:  isCorrect,
		})
	}

	quizResult.Correct = correct
	quizResult.Total = total

	results := &models.QuizWithQuestionResults{
		QuizResult: quizResult,
		Results:    questionResults,
	}
	err = s.quizRepo.SaveResults(ctx, results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *service) GetQuizResults(ctx context.Context, userID int) ([]*models.QuizResultWithFields, error) {
	quizResults, err := s.quizRepo.GetQuizResults(ctx, userID)
	if err != nil {
		return nil, err
	}

	quizIDs := lo.Map(quizResults, func(item *models.QuizResult, _ int) int {
		return item.QuizID
	})

	quizzes, err := s.quizRepo.GetQuizzes(ctx, quizIDs)
	if err != nil {
		return nil, err
	}

	bookIDs := lo.Map(quizzes, func(item *models.Quiz, _ int) int {
		return item.BookID
	})

	books, err := s.bookRepo.GetBooksByIDs(ctx, bookIDs)
	if err != nil {
		return nil, err
	}

	quizMap := lo.SliceToMap(quizzes, func(item *models.Quiz) (int, *models.Quiz) {
		return item.ID, item
	})

	bookMap := lo.SliceToMap(books, func(item *models.Book) (int, *models.Book) {
		return item.ID, item
	})

	out := lo.Map(quizResults, func(item *models.QuizResult, _ int) *models.QuizResultWithFields {
		return &models.QuizResultWithFields{
			QuizResult: item,
			Quiz:       quizMap[item.QuizID],
			Book:       bookMap[quizMap[item.QuizID].BookID],
		}
	})

	return out, nil
}

func (s *service) GetQuizResultWithAnswers(ctx context.Context, resultID int) (*models.QuizQuestionWithFields, error) {
	quizResult, err := s.quizRepo.GetQuizResult(ctx, resultID)
	if err != nil {
		return nil, err
	}

	quiz, err := s.quizRepo.GetQuiz(ctx, quizResult.QuizID)
	if err != nil {
		return nil, err
	}

	book, err := s.bookRepo.GetBookByID(ctx, quiz.BookID)
	if err != nil {
		return nil, err
	}

	out := &models.QuizQuestionWithFields{
		QuizWithQuestionResults: quizResult,
		Quiz:                    quiz,
		Book:                    book,
	}

	return out, nil
}
