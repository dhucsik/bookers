package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// createUser godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param request body createUserRequest true "request"
// @Success 201 {object} createUserResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 500 {object} errorResponse "Internal server error"
func (c *Controller) createUser(ctx echo.Context) error {
	var req createUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	user, err := c.usersService.CreateUser(ctx.Request().Context(), req.convert())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, newCreateUserResponse(user))
}
