package quizzes

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// updateQuizHandler godoc
// @Summary Update quiz
// @Description Update quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "quiz ID"
// @Param request body updateQuizRequest true "Title"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/{id} [put]
func (c *Controller) updateQuizHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	quizIDStr := ctx.Param("id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req updateQuizRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	if err := c.quizService.UpdateQuizTitle(ctx.Request().Context(), session.UserID, quizID, req.Title); err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}
