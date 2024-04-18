package auth

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/util/response"
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
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /auth [post]
func (c *Controller) authHandler(ctx echo.Context) error {
	var req authRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	accessToken, refreshToken, err := c.authService.GetAuth(ctx.Request().Context(), req.Username, req.Password)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	user, err := c.usersService.GetUserByUsername(ctx.Request().Context(), req.Username)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, authResponse{
		Response: response.NewResponse(),
		Result: authResp{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			User:         user.ToUserWithoutPassword(),
		},
	})
}
