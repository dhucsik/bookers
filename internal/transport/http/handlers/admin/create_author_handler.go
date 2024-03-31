package admin

import (
	"errors"
	"net/http"

	"github.com/dhucsik/bookers/internal/models"
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
// @Success 200 {object} nil "Success"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /admin/authors [post]
func (c *Controller) createAuthor(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return errors.New("session not found")
	}

	if session.Role != "admin" {
		return errors.New("forbidden")
	}

	var req createAuthorRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	if err := c.authorService.CreateAuthor(ctx.Request().Context(), req.convert()); err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, nil)
}
