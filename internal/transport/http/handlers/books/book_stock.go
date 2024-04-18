package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// uploadStockBookHandler godoc
// @Summary Upload stock book
// @Description Upload stock book
// @Tags books
// @Accept mpfd
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param book_id formData int true "book ID"
// @Param image formData file true "image"
// @Success 201 {object} uploadStockBookResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/stock/upload [post]
func (c *Controller) uploadStockBookHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	bookIDStr := ctx.FormValue("book_id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	img, err := ctx.FormFile("image")
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	if img.Header.Get("Content-Type") != "image/png" {
		return response.NewBadRequest(ctx, errors.ErrInvalidImageFormat)
	}

	id, imageURL, err := c.bookService.UploadStockBook(ctx.Request().Context(), &models.UploadStockBook{
		UserID: session.UserID,
		BookID: bookID,
		Image:  img,
	})
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, uploadStockBookResponse{
		Response: response.NewResponse(),
		Result:   newUploadStockBookResp(id, imageURL),
	})
}

// getStockBooksHandler godoc
// @Summary Get stock books
// @Description Get stock books
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} getStockBookResponse "Success"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/stock [get]
func (c *Controller) getStockBooksHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	books, err := c.bookService.GetStockBooks(ctx.Request().Context(), session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getStockBookResponse{
		Response: response.NewResponse(),
		Result:   books,
	})
}
