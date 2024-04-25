package models

import (
	"mime/multipart"
	"time"
)

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

type StockBook struct {
	ID     int `json:"id"`
	BookID int `json:"book_id"`
	UserID int `json:"user_id"`
}

type StockBookWithFields struct {
	*StockBook
	Book *Book                `json:"book"`
	User *UserWithoutPassword `json:"user"`
}

type UploadStockBook struct {
	BookID int
	UserID int
	Image  *multipart.FileHeader
}
