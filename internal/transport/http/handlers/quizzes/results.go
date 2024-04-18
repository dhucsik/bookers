package quizzes

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// checkQuizHandler godoc
// @Summary Check quiz
// @Description Check quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "quiz ID"
// @Param request body checkQuizRequest true "request"
// @Success 200 {object} checkQuizResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/{id}/check [post]
func (c *Controller) checkQuizHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	quizIDStr := ctx.Param("id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req checkQuizRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	results, err := c.quizService.CheckQuiz(ctx.Request().Context(), session.UserID, quizID, req.convert())
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, checkQuizResponse{
		Response: response.NewResponse(),
		Result:   results,
	})
}

// getQuizResultsHandler godoc
// @Summary Get quiz results
// @Description Get quiz results
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} getQuizResultsResponse "Success"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/results [get]
func (c *Controller) getQuizResultsHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	results, err := c.quizService.GetQuizResults(ctx.Request().Context(), session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getQuizResultsResponse{
		Response: response.NewResponse(),
		Result:   results,
	})
}

// getQuizResultHandler godoc
// @Summary Get quiz result
// @Description Get quiz result
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "result ID"
// @Success 200 {object} getQuizResultResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/results/{id} [get]
func (c *Controller) getQuizResultHandler(ctx echo.Context) error {
	resultIDStr := ctx.Param("id")
	resultID, err := strconv.Atoi(resultIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	results, err := c.quizService.GetQuizResultWithAnswers(ctx.Request().Context(), resultID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getQuizResultResponse{
		Response: response.NewResponse(),
		Result:   results,
	})
}
