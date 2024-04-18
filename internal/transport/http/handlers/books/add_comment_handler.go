package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// addCommentHandler godoc
// @Summary Add comment
// @Description Add comment
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "book ID"
// @Param request body addCommentRequest true "request"
// @Success 200 {object} createResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/{id}/comments [post]
func (c *Controller) addCommentHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	bookIDStr := ctx.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req addCommentRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	id, err := c.bookService.AddComment(ctx.Request().Context(), req.convert(session.UserID, bookID))
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, createResponse{
		Response: response.NewResponse(),
		Result:   newCreateResp(id),
	})
}
