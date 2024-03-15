package authors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// listAuthors godoc
// @Summary List authors
// @Description List authors
// @Tags authors
// @Accept json
// @Produce json
// @Success 200 {object} listAuthorsResponse "Success"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /authors [get]
func (c *Controller) listAuthors(ctx echo.Context) error {
	authors, err := c.authorsService.ListAuthors(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, newListAuthorsResponse(authors))
}
