package models

import (
	"fmt"
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
	ID       int    `json:"id"`
	BookID   int    `json:"book_id"`
	UserID   int    `json:"user_id"`
	ImageURL string `json:"image_url"`
}

func (sb *StockBook) SetImage() {
	sb.ImageURL = fmt.Sprintf("https://bookers-images.fra1.digitaloceanspaces.com/stock/books/%d.png", sb.ID)

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

type SearchParams struct {
	Search     string
	Limit      int
	Offset     int
	Categories []int
	Authors    []int
	City       string
}
