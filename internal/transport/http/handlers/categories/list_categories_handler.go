package categories

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// listCategories godoc
// @Summary List categories
// @Description List categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} listCategoriesResponse "Success"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /categories [get]
func (c *Controller) listCategories(ctx echo.Context) error {
	categories, err := c.categoriesService.ListCategories(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, newListCategoriesResponse(categories))
}
