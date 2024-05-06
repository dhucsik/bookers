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

// updateStockImageHandler godoc
// @Summary Update stock image
// @Description Update stock image
// @Tags books
// @Accept mpfd
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "Stock ID"
// @Param image formData file true "image"
// @Success 200 {object} updateStockResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/stock/{id}/image [put]
func (c *Controller) updateStockImageHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	stockIDStr := ctx.Param("id")
	stockID, err := strconv.Atoi(stockIDStr)
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

	imageURL, err := c.bookService.UpdateImage(ctx.Request().Context(), session.UserID, stockID, img)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, updateStockResponse{
		Response: response.NewResponse(),
		Result:   imageURL,
	})
}

// deleteStockBookHandler godoc
// @Summary Delete stock book
// @Description Delete stock book
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "Stock ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/stock/{id} [delete]
func (c *Controller) deleteStockBookHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	stockIDStr := ctx.Param("id")
	stockID, err := strconv.Atoi(stockIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.bookService.DeleteStockBook(ctx.Request().Context(), session.UserID, stockID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
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

// getStockByUserHanlder godoc
// @Summary Get stock books by user
// @Description Get stock books by user
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "User ID"
// @Success 200 {object} getStockBookResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/{id}/stock [get]
func (c *Controller) getStockByUserHanlder(ctx echo.Context) error {
	_, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	userIDStr := ctx.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	books, err := c.bookService.GetStockBooks(ctx.Request().Context(), userID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getStockBookResponse{
		Response: response.NewResponse(),
		Result:   books,
	})
}

// getStockByBookHandler godoc
// @Summary Get stock by book
// @Description Get stock by book
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "book ID"
// @Success 200 {object} getStockBookResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/{id}/stock [get]
func (c *Controller) getStockByBookHandler(ctx echo.Context) error {
	_, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	bookIDStr := ctx.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	books, err := c.bookService.GetStockByBook(ctx.Request().Context(), bookID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getStockBookResponse{
		Response: response.NewResponse(),
		Result:   books,
	})
}

// getStockBookByIDHandler godoc
// @Summary Get stock book by ID
// @Description Get stock book by ID
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "Stock ID"
// @Success 200 {object} getStockBookByIDResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/stock/{id} [get]
func (c *Controller) getStockBookByIDHandler(ctx echo.Context) error {
	_, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	stockIDStr := ctx.Param("id")
	stockID, err := strconv.Atoi(stockIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	book, err := c.bookService.GetStockBook(ctx.Request().Context(), stockID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getStockBookByIDResponse{
		Response: response.NewResponse(),
		Result:   book,
	})
}
