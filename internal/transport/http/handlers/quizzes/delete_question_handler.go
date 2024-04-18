package quizzes

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// deleteQuestionHandler godoc
// @Summary Delete question
// @Description Delete question
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "question ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/questions/{id} [delete]
func (c *Controller) deleteQuestionHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	questionIDStr := ctx.Param("id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	if err := c.quizService.DeleteQuestion(ctx.Request().Context(), session.UserID, questionID); err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}
