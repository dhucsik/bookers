package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/models"
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
// @Success 201 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/{id}/quizzes [post]
func (c *Controller) createQuizHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	bookIDStr := ctx.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	var req createQuizRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := c.quizService.CreateQuiz(ctx.Request().Context(), req.convert(bookID, session.UserID)); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, nil)
}
