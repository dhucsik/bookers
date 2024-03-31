package books

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// getBookByIDHandler godoc
// @Summary Get book by ID
// @Description Get book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} bookResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/{id} [get]
func (c *Controller) getBookByIDHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse("invalid id"))
	}

	book, err := c.bookService.GetBookByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, newBookResp(book))
}
