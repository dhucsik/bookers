package admin

import (
	"errors"
	"net/http"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/labstack/echo/v4"
)

// createCategory godoc
// @Summary Create category
// @Description Create category
// @Tags admin
// @Accept json
// @Produce json
// @Param request body createCategoryRequest true "request"
// @Success 200 {object} nil "Success"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /admin/categories [post]
func (c *Controller) createCategory(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return errors.New("session not found")
	}

	if session.Role != "admin" {
		return errors.New("forbidden")
	}

	var req createCategoryRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	if err := c.categoriesService.CreateCategory(ctx.Request().Context(), req.convert()); err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, nil)
}
