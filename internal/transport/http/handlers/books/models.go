package books

import (
	"time"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/samber/lo"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(message string) errorResponse {
	return errorResponse{
		Message: message,
	}
}

type categoryResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newCategoryResp(category *models.Category) categoryResp {
	return categoryResp{
		ID:   category.ID,
		Name: category.Name,
	}
}

type authorResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newAuthorResp(author *models.Author) authorResp {
	return authorResp{
		ID:   author.ID,
		Name: author.Name,
	}
}

type bookResponse struct {
	ID         int            `json:"id"`
	Title      string         `json:"title"`
	PubDate    string         `json:"pub_date"`
	Edition    string         `json:"edition"`
	Language   string         `json:"language"`
	Rating     float64        `json:"rating"`
	Authors    []authorResp   `json:"authors"`
	Categories []categoryResp `json:"categories"`
}

func newBookResp(book *models.BookWithFields) bookResponse {
	authors := lo.Map(book.Authors, func(author *models.Author, _ int) authorResp {
		return newAuthorResp(author)
	})

	categories := lo.Map(book.Categories, func(category *models.Category, _ int) categoryResp {
		return newCategoryResp(category)
	})

	return bookResponse{
		ID:         book.ID,
		Title:      book.Title,
		PubDate:    book.PubDate.Format(time.DateOnly),
		Edition:    book.Edition,
		Language:   book.Language,
		Rating:     book.Rating,
		Authors:    authors,
		Categories: categories,
	}
}

type listBooksResponse struct {
	Books []bookResponse `json:"books"`
}

func newListBooksResponse(books []*models.BookWithFields) listBooksResponse {
	out := lo.Map(books, func(book *models.BookWithFields, _ int) bookResponse {
		return newBookResp(book)
	})

	return listBooksResponse{
		Books: out,
	}
}
