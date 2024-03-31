package authors

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// listAuthors godoc
// @Summary List authors
// @Description List authors
// @Tags authors
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int true "Offset"
// @Param search query string false "Search"
// @Success 200 {object} listAuthorsResponse "Success"
// @Failure 400 {object} errorResponse "Invalid limit or offset"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /authors [get]
func (c *Controller) listAuthors(ctx echo.Context) error {
	limitStr := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse("invalid limit"))
	}

	search := ctx.QueryParam("search")
	offsetStr := ctx.QueryParam("offset")
	offset, err := strconv.Atoi(offsetStr)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse("invalid offset"))
	}

	authors, err := c.authorsService.ListAuthors(ctx.Request().Context(), search, limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, newListAuthorsResponse(authors))
}
