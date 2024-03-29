package models

import "time"

type Book struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	PubDate  time.Time `json:"pub_date"`
	Edition  int       `json:"edition"`
	Language string    `json:"language"`
	Rating   float64   `json:"rating"`
}

type BookWithFields struct {
	*Book
	Authors    []*Author   `json:"authors"`
	Categories []*Category `json:"categories"`
}
