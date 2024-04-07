package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/models"
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
// @Success 201 {object} map[string]string "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/stock/upload [post]
func (c *Controller) uploadStockBookHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	bookIDStr := ctx.FormValue("book_id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	img, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	if img.Header.Get("Content-Type") != "image/png" {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse("only png are allowed"))
	}

	imageURL, err := c.bookService.UploadStockBook(ctx.Request().Context(), &models.UploadStockBook{
		UserID: session.UserID,
		BookID: bookID,
		Image:  img,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"image_url": imageURL})
}

// getStockBooksHandler godoc
// @Summary Get stock books
// @Description Get stock books
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} []models.StockBookWithFields "Success"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/stock [get]
func (c *Controller) getStockBooksHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	books, err := c.bookService.GetStockBooks(ctx.Request().Context(), session.UserID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, books)
}
