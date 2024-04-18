package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// setRatingHandler godoc
// @Summary Set rating
// @Description Set rating
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "book ID"
// @Param request body setRatingRequest true "request"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/{id}/rating [post]
func (c *Controller) setRatingHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	bookIDstr := ctx.Param("id")
	bookID, err := strconv.Atoi(bookIDstr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req setRatingRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.bookService.SetRating(ctx.Request().Context(), req.convert(bookID, session.UserID))
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}
