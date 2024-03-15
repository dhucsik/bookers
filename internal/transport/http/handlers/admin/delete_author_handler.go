package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// deleteAuthor godoc
// @Summary Delete author
// @Description Delete author
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} nil "Success"
// @Failure 500 {object} nil "Internal server error"
// @Router /admin/authors/{id} [delete]
func (c *Controller) deleteAuthor(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, nil)
}
