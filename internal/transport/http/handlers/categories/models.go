package categories

import (
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
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

type listCategoriesResponse struct {
	response.Response
	Result []*categoryItemResponse `json:"categories"`
}

type categoryItemResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newListCategoriesResponse(categories []*models.Category) listCategoriesResponse {
	out := lo.Map(categories, func(item *models.Category, _ int) *categoryItemResponse {
		return &categoryItemResponse{
			ID:   item.ID,
			Name: item.Name,
		}
	})

	return listCategoriesResponse{
		Response: response.NewResponse(),
		Result:   out,
	}
}
