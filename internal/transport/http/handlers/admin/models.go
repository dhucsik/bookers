package admin

import "github.com/dhucsik/bookers/internal/models"

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
