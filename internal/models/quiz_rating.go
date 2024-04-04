package models

type QuizRating struct {
	ID     int `json:"id"`
	QuizID int `json:"quiz_id"`
	UserID int `json:"user_id"`
	Rating int `json:"rating"`
}
