package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// deleteCategory godoc
// @Summary Delete category
// @Description Delete category
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} nil "Success"
// @Failure 500 {object} nil "Internal server error"
// @Router /admin/categories/{id} [delete]
func (c *Controller) deleteCategory(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
