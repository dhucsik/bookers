package models

type BookRating struct {
	ID     int `json:"id"`
	BookID int `json:"book_id"`
	UserID int `json:"user_id"`
	Rating int `json:"rating"`
}
