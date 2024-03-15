package authors

import (
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

type listAuthorsResponse struct {
	Authors []*authorItemResponse `json:"authors"`
}

type authorItemResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newListAuthorsResponse(authors []*models.Author) listAuthorsResponse {
	out := lo.Map(authors, func(item *models.Author, _ int) *authorItemResponse {
		return &authorItemResponse{
			ID:   item.ID,
			Name: item.Name,
		}
	})

	return listAuthorsResponse{
		Authors: out,
	}
}
