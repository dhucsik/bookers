package quizzes

import (
	"encoding/json"
	"time"

	"github.com/dhucsik/bookers/internal/models"
)

type questionModel struct {
	ID       int    `db:"id"`
	QuizID   int    `db:"quiz_id"`
	Question string `db:"question"`
	Options  []byte `db:"options"`
	Answer   string `db:"answer"`
}

func newQuestionModel(q *models.Question) (*questionModel, error) {
	options, err := json.Marshal(q.Options)
	if err != nil {
		return nil, err
	}

	return &questionModel{
		ID:       q.ID,
		QuizID:   q.QuizID,
		Question: q.Question,
		Options:  options,
		Answer:   q.Answer,
	}, nil
}

func (q *questionModel) convert() (*models.Question, error) {
	var options []string
	if err := json.Unmarshal(q.Options, &options); err != nil {
		return nil, err
	}

	return &models.Question{
		ID:       q.ID,
		QuizID:   q.QuizID,
		Question: q.Question,
		Options:  options,
		Answer:   q.Answer,
	}, nil
}

type quizComment struct {
	ID        int       `db:"id"`
	QuizID    int       `db:"quiz_id"`
	UserID    int       `db:"user_id"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
}

func (b *quizComment) convert() *models.QuizComment {
	return &models.QuizComment{
		ID:        b.ID,
		QuizID:    b.QuizID,
		UserID:    b.UserID,
		Comment:   b.Comment,
		CreatedAt: b.CreatedAt,
	}
}
