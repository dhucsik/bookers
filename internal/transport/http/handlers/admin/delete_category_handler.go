package admin

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/labstack/echo/v4"
)

// deleteCategory godoc
// @Summary Delete category
// @Description Delete category
// @Tags admin
// @Accept json
// @Produce json
// @Param id path int true "category ID"
// @Success 200 {object} nil "Success"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /admin/categories/{id} [delete]
func (c *Controller) deleteCategory(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return errors.New("session not found")
	}

	if session.Role != "admin" {
		return errors.New("forbidden")
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	err = c.categoriesService.DeleteCategory(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, nil)
}
