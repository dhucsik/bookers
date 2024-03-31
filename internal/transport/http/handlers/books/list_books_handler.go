package books

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// listBooksHandler godoc
// @Summary List books
// @Description List books
// @Tags books
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param offset query int true "Offset"
// @Param search query string false "Search"
// @Success 200 {object} listBooksResponse "Success"
// @Failure 400 {object} errorResponse "Invalid limit or offset"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books [get]
func (c *Controller) listBooksHandler(ctx echo.Context) error {
	limitStr := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return ctx.JSON(400, newErrorResponse("invalid limit"))
	}

	search := ctx.QueryParam("search")
	offsetStr := ctx.QueryParam("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return ctx.JSON(400, newErrorResponse("invalid offset"))
	}

	books, err := c.bookService.ListBooks(ctx.Request().Context(), search, limit, offset)
	if err != nil {
		return ctx.JSON(500, newErrorResponse(err.Error()))
	}

	return ctx.JSON(200, newListBooksResponse(books))
}
