package authors

import (
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/samber/lo"
)

type listAuthorsResponse struct {
	response.Response
	Result listAuthorsResp `json:"result"`
}

type listAuthorsResp struct {
	Authors    []*authorItemResponse `json:"authors"`
	TotalCount int                   `json:"total_count"`
}

type authorItemResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newListAuthorsResponse(authors []*models.Author, count int) listAuthorsResponse {
	out := lo.Map(authors, func(item *models.Author, _ int) *authorItemResponse {
		return &authorItemResponse{
			ID:   item.ID,
			Name: item.Name,
		}
	})

	return listAuthorsResponse{
		Response: response.NewResponse(),
		Result:   listAuthorsResp{Authors: out, TotalCount: count},
	}
}
