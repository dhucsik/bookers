package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// getBookByIDHandler godoc
// @Summary Get book by ID
// @Description Get book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} getBookResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/{id} [get]
func (c *Controller) getBookByIDHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	book, err := c.bookService.GetBookByID(ctx.Request().Context(), id)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getBookResponse{
		Response: response.NewResponse(),
		Result:   newBookResp(book),
	})
}
