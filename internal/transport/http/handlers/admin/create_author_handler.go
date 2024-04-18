package admin

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// createAuthor godoc
// @Summary Create author
// @Description Create author
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param request body createAuthorRequest true "request"
// @Success 200 {object} createAuthorResponse "Success"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/authors [post]
func (c *Controller) createAuthor(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	if session.Role != "admin" {
		return response.NewErrorResponse(ctx, errors.ErrForbiddenForRole)
	}

	var req createAuthorRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	id, err := c.authorService.CreateAuthor(ctx.Request().Context(), req.convert())
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, createAuthorResponse{
		Response: response.NewResponse(),
		Result:   newCreateAuthorResp(id),
	})
}
