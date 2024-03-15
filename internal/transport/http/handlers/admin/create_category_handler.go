package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// createCategory godoc
// @Summary Create category
// @Description Create category
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} nil "Success"
// @Failure 500 {object} nil "Internal server error"
// @Router /admin/categories [post]
func (c *Controller) createCategory(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
