package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// listCommentsHandler godoc
// @Summary List comments
// @Description List comments
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} listCommentsResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/{id}/comments [get]
func (c *Controller) listCommentsHandler(ctx echo.Context) error {
	bookIDStr := ctx.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	comments, err := c.bookService.ListComments(ctx.Request().Context(), bookID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, listCommentsResponse{
		Response: response.NewResponse(),
		Result:   comments,
	})
}
