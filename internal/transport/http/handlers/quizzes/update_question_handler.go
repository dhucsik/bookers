package quizzes

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// updateQuestionHandler godoc
// @Summary Update question
// @Description Update question
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "question ID"
// @Param body body updateQuestionRequest true "Request body"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/questions/{id} [put]
func (c *Controller) updateQuestionHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	questionIDStr := ctx.Param("id")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req updateQuestionRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	if err := c.quizService.UpdateQuestion(ctx.Request().Context(), session.UserID, req.convert(questionID)); err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}
