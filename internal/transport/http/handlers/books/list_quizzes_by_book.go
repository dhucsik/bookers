package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// listQuizzesByBookHandler godoc
// @Summary List quizzes by book
// @Description List quizzes by book
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "Book ID"
// @Success 200 {object} listQuizzesResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/{id}/quizzes [get]
func (c *Controller) listQuizzesByBookHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	quizzes, err := c.quizService.ListQuizzesByBookID(ctx.Request().Context(), id)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, listQuizzesResponse{
		Response: response.NewResponse(),
		Result:   quizzes,
	})
}
