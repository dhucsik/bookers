package books

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// searchStockBooksRequest godoc
// @Summary Search stock books
// @Description Search stock books
// @Tags books
// @Accept json
// @Produce json
// @Param search body searchStockBooksRequest true "request"
// @Success 200 {object} searchStockBooksResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/stock/search [post]
func (c *Controller) searchStockBooks(ctx echo.Context) error {
	var req searchStockBooksRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	books, err := c.bookService.SearchStockByParams(ctx.Request().Context(), req.convert())
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, searchStockBooksResponse{
		Response: response.NewResponse(),
		Result:   books,
	})
}
