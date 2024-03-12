package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// refreshHandler godoc
// @Summary Refresh token
// @Description Refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body refreshRequest true "request"
// @Success 200 {object} authResponse "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /auth/refresh [post]
func (c *Controller) refreshHandler(ctx echo.Context) error {
	var req refreshRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	accessToken, refreshToken, err := c.authService.Refresh(ctx.Request().Context(), req.Token)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	resp := authResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return ctx.JSON(http.StatusOK, resp)
}
