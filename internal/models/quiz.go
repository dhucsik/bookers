package models

import "time"

type Quiz struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	BookID         int       `json:"book_id"`
	Title          string    `json:"title"`
	Rating         float64   `json:"rating"`
	CreatedAt      time.Time `json:"created_at"`
	QuestionsCount int       `json:"questions_count"`
}

type QuizWithBase struct {
	*Quiz
	User *UserWithoutPassword `json:"user"`
	Book *Book                `json:"book"`
}

type QuizWithFields struct {
	*Quiz
	User      *UserWithoutPassword `json:"user"`
	Book      *Book                `json:"book"`
	Questions []*Question          `json:"questions"`
}

type Question struct {
	ID       int      `json:"id"`
	QuizID   int      `json:"quiz_id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}
