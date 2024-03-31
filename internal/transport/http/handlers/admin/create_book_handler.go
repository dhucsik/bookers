package admin

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/labstack/echo/v4"
)

// createBookHandler godoc
// @Summary Create book
// @Description Create book
// @Tags admin
// @Accept json
// @Produce json
// @Param request body createBookRequest true "request"
// @Success 201 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 403 {object} errorResponse "Forbidden"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /admin/books [post]
func (c *Controller) createBookHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	if session.Role != "admin" {
		return ctx.JSON(http.StatusForbidden, newErrorResponse("forbidden"))
	}

	var req createBookRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	book, err := req.convert()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	err = c.bookService.CreateBook(ctx.Request().Context(), book, req.AuthorIDs, req.CategoryIDs)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, nil)
}
