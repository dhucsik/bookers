package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/util/response"
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
// @Failure 400 {object} response.Response "Invalid limit or offset"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books [get]
func (c *Controller) listBooksHandler(ctx echo.Context) error {
	limitStr := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	search := ctx.QueryParam("search")
	offsetStr := ctx.QueryParam("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	books, totalCount, err := c.bookService.ListBooks(ctx.Request().Context(), search, limit, offset)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, newListBooksResponse(books, totalCount))
}

// searchBooksHandler godoc
// @Summary Search books
// @Description Search books
// @Tags books
// @Accept json
// @Produce json
// @Param search body searchBooksRequest true "request"
// @Success 200 {object} listBooksResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/search [post]
func (c *Controller) searchBooksHandler(ctx echo.Context) error {
	var req searchBooksRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	books, totalCount, err := c.bookService.SearchBooks(ctx.Request().Context(), req.convert())
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, newListBooksResponse(books, totalCount))
}
