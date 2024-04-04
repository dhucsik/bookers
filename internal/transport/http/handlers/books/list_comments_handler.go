package books

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// listCommentsHandler godoc
// @Summary List comments
// @Description List comments
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} []models.BookComment "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/{id}/comments [get]
func (c *Controller) listCommentsHandler(ctx echo.Context) error {
	bookIDStr := ctx.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	comments, err := c.bookService.ListComments(ctx.Request().Context(), bookID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, comments)
}
