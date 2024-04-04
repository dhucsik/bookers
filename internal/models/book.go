package models

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	PubDate     time.Time `json:"pub_date"`
	Edition     string    `json:"edition"`
	Language    string    `json:"language"`
	Rating      float64   `json:"rating"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
}

type BookWithFields struct {
	*Book
	Authors    []*Author   `json:"authors"`
	Categories []*Category `json:"categories"`
}
