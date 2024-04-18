package quizzes

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// viewQuizHandler godoc
// @Summary View quiz
// @Description View quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "Quiz ID"
// @Success 200 {object} getQuizResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/{id}/view [get]
func (c *Controller) viewQuizHandler(ctx echo.Context) error {
	quizIDStr := ctx.Param("id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	quiz, err := c.quizService.GetQuizWithoutAnswers(ctx.Request().Context(), quizID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getQuizResponse{
		Response: response.NewResponse(),
		Result:   quiz,
	})
}
