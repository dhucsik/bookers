package users

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// getByID godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "User ID"
// @Success 200 {object} getUserByIDResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/{id} [get]
func (c *Controller) getByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	user, err := c.usersService.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getUserByIDResponse{
		Response: response.NewResponse(),
		Result:   newGetUserByIDResponse(user),
	})
}
