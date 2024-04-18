package quizzes

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// getQuizHandler godoc
// @Summary Get quiz
// @Description Get quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "quiz ID"
// @Success 200 {object} getQuizResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/{id} [get]
func (c *Controller) getQuizHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	quizIDStr := ctx.Param("id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	quiz, err := c.quizService.GetQuiz(ctx.Request().Context(), quizID, session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getQuizResponse{
		Response: response.NewResponse(),
		Result:   quiz,
	})
}

// listQuizzesHandler godoc
// @Summary List quizzes
// @Description List quizzes
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Success 200 {object} listQuizzesResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes [get]
func (c *Controller) listQuizzesHandler(ctx echo.Context) error {
	limitStr := ctx.QueryParam("limit")
	offsetStr := ctx.QueryParam("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	quizzes, totalCount, err := c.quizService.ListQuizzes(ctx.Request().Context(), limit, offset)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, listQuizzesResponse{
		Response: response.NewResponse(),
		Result:   listQuizzesResp{Quizzes: quizzes, TotalCount: totalCount},
	})
}

// listQuizzesByUserID godoc
// @Summary List quizzes by user ID
// @Description List quizzes by user ID
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} listUserQuizzesResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/user [get]
func (c *Controller) listQuizzesByUserID(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	quizzes, err := c.quizService.ListQuizzesByUserID(ctx.Request().Context(), session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, listUserQuizzesResponse{
		Response: response.NewResponse(),
		Result:   quizzes,
	})
}
