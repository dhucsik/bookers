package admin

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/labstack/echo/v4"
)

// createCategory godoc
// @Summary Create category
// @Description Create category
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param request body createCategoryRequest true "request"
// @Success 201 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 403 {object} errorResponse "Forbidden"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /admin/categories [post]
func (c *Controller) createCategory(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	if session.Role != "admin" {
		return ctx.JSON(http.StatusForbidden, newErrorResponse("forbidden"))
	}

	var req createCategoryRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	if err := c.categoriesService.CreateCategory(ctx.Request().Context(), req.convert()); err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, nil)
}
