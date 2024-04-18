package admin

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// deleteAuthor godoc
// @Summary Delete author
// @Description Delete author
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "author ID"
// @Success 200 {object} response.Response "Success"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /admin/authors/{id} [delete]
func (c *Controller) deleteAuthor(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	if session.Role != "admin" {
		return response.NewErrorResponse(ctx, errors.ErrForbiddenForRole)
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.authorService.DeleteAuthor(ctx.Request().Context(), id)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}
