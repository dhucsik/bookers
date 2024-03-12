package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// authHandler godoc
// @Summary Authenticate user
// @Description Authenticate user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body authRequest true "request"
// @Success 200 {object} authResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /auth [post]
func (c *Controller) authHandler(ctx echo.Context) error {
	var req authRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	accessToken, refreshToken, err := c.authService.GetAuth(ctx.Request().Context(), req.Username, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	resp := authResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return ctx.JSON(http.StatusOK, resp)
}
