package admin

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// createCategory godoc
// @Summary Create category
// @Description Create category
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param request body createCategoryRequest true "request"
// @Success 201 {object} createAuthorResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 403 {object} response.Response "Forbidden"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/categories [post]
func (c *Controller) createCategory(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	if session.Role != "admin" {
		return response.NewErrorResponse(ctx, errors.ErrForbiddenForRole)
	}

	var req createCategoryRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	id, err := c.categoriesService.CreateCategory(ctx.Request().Context(), req.convert())
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, createAuthorResponse{
		Response: response.NewResponse(),
		Result:   newCreateAuthorResp(id),
	})
}
