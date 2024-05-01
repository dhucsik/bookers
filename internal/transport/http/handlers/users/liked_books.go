package users

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// addLikedBook godoc
// @Summary Add liked book
// @Description Add liked book
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param request body addLikedBookRequest true "request"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/liked-books [post]
func (c *Controller) addLikedBook(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	var req addLikedBookRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err := c.usersService.AddBookToLiked(ctx.Request().Context(), session.UserID, req.BookID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// removeLikedBook godoc
// @Summary Remove liked book
// @Description Remove liked book
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param request body removeLikedBookRequest true "request"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/liked-books [delete]
func (c *Controller) removeLikedBook(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	var req removeLikedBookRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err := c.usersService.RemoveBookFromLiked(ctx.Request().Context(), session.UserID, req.BookID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// getLikedBooks godoc
// @Summary Get liked books
// @Description Get liked books
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} listLikedBooksResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/liked-books [get]
func (c *Controller) getLikedBooks(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	books, err := c.usersService.GetLikedBooks(ctx.Request().Context(), session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, listLikedBooksResponse{
		Response: response.NewResponse(),
		Result:   books,
	})
}
