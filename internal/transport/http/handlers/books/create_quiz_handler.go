package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// createQuizHandler godoc
// @Summary Create quiz
// @Description Create quiz
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "book ID"
// @Param request body createQuizRequest true "request"
// @Success 201 {object} createResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/{id}/quizzes [post]
func (c *Controller) createQuizHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	bookIDStr := ctx.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req createQuizRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	id, err := c.quizService.CreateQuiz(ctx.Request().Context(), req.convert(bookID, session.UserID))
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, createResponse{
		Response: response.NewResponse(),
		Result:   newCreateResp(id),
	})
}
