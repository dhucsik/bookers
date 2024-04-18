package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// deleteCommentHandler godoc
// @Summary Delete comment
// @Description Delete comment
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "comment ID"
// @Success 200 {object} response.Response "Success"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/comments/{id} [delete]
func (c *Controller) deleteCommentHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	commentIDStr := ctx.Param("id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.bookService.DeleteComment(ctx.Request().Context(), commentID, session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}
