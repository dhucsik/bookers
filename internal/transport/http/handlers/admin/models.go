package admin

import (
	"time"

	"github.com/dhucsik/bookers/internal/models"
)

type createCategoryRequest struct {
	Name string `json:"name"`
}

func (r createCategoryRequest) convert() *models.Category {
	return &models.Category{
		Name: r.Name,
	}
}

type createAuthorRequest struct {
	Name string `json:"name"`
}

func (r createAuthorRequest) convert() *models.Author {
	return &models.Author{
		Name: r.Name,
	}
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(message string) errorResponse {
	return errorResponse{Message: message}
}

type createBookRequest struct {
	Title       string `json:"title"`
	PubDate     string `json:"pub_date"`
	Edition     string `json:"edition"`
	Language    string `json:"language"`
	Image       string `json:"image"`
	Description string `json:"description"`
	AuthorIDs   []int  `json:"author_ids"`
	CategoryIDs []int  `json:"category_ids"`
}

func (r createBookRequest) convert() (*models.Book, error) {
	pubDate, err := time.Parse(time.DateOnly, r.PubDate)
	if err != nil {
		return nil, err
	}

	return &models.Book{
		Title:       r.Title,
		PubDate:     pubDate,
		Edition:     r.Edition,
		Language:    r.Language,
		Image:       r.Image,
		Description: r.Description,
	}, nil
}
