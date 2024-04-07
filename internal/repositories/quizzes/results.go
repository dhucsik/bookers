package quizzes

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
)

func (r *repository) SaveResults(ctx context.Context, result *models.QuizWithQuestionResults) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	var quizResultID int
	err = tx.QueryRow(ctx, insertQuizResultStmt, result.QuizID, result.UserID, result.Correct, result.Total).Scan(&quizResultID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	result.ID = quizResultID
	for _, questionResult := range result.Results {
		_, err = tx.Exec(ctx, insertQuestionResultStmt, quizResultID, questionResult.QuestionID, questionResult.UserAnswer, questionResult.IsCorrect)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *repository) GetQuizResults(ctx context.Context, userID int) ([]*models.QuizResult, error) {
	rows, err := r.db.Query(ctx, getQuizResultsStmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models.QuizResult
	for rows.Next() {
		result := &models.QuizResult{}
		if err := rows.Scan(&result.ID, &result.QuizID, &result.UserID, &result.Correct, &result.Total); err != nil {
			return nil, err
		}

		out = append(out, result)
	}

	return out, nil
}

func (r *repository) GetQuizResult(ctx context.Context, id int) (*models.QuizWithQuestionResults, error) {
	result := &models.QuizWithQuestionResults{}
	err := r.db.QueryRow(ctx, getQuizResultStmt, id).Scan(&result.ID, &result.QuizID, &result.UserID, &result.Correct, &result.Total)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, getQuestionResultsStmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questionResults []*models.QuestionResult
	for rows.Next() {
		questionResult := &models.QuestionResult{}
		if err := rows.Scan(&questionResult.ID, &questionResult.QuizResultID, &questionResult.QuestionID, &questionResult.UserAnswer, &questionResult.IsCorrect); err != nil {
			return nil, err
		}

		questionResults = append(questionResults, questionResult)
	}

	result.Results = questionResults
	return result, nil
}
