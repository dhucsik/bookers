package admin

import (
	"time"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
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

type createAuthorResponse struct {
	response.Response
	Result createAuthorResp `json:"result"`
}

type createAuthorResp struct {
	ID int `json:"id"`
}

func newCreateAuthorResp(id int) createAuthorResp {
	return createAuthorResp{ID: id}
}
