package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// createAuthor godoc
// @Summary Create author
// @Description Create author
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} nil "Success"
// @Failure 500 {object} nil "Internal server error"
// @Router /admin/authors [post]
func (c *Controller) createAuthor(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
