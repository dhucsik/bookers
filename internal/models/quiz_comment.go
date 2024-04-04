package models

import "time"

type QuizComment struct {
	ID        int       `json:"id"`
	QuizID    int       `json:"quiz_id"`
	UserID    int       `json:"user_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
