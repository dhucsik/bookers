package quizzes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// viewQuizHandler godoc
// @Summary View quiz
// @Description View quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Param id path int true "Quiz ID"
// @Success 200 {object} models.QuizWithFields "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /quizzes/{id}/view [get]
func (c *Controller) viewQuizHandler(ctx echo.Context) error {
	quizIDStr := ctx.Param("id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	quiz, err := c.quizService.GetQuizWithoutAnswers(ctx.Request().Context(), quizID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, quiz)
}
