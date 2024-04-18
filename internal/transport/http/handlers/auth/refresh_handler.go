package auth

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/util/response"
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
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /auth/refresh [post]
func (c *Controller) refreshHandler(ctx echo.Context) error {
	var req refreshRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	accessToken, refreshToken, err := c.authService.Refresh(ctx.Request().Context(), req.Token)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, authResponse{
		Response: response.NewResponse(),
		Result: authResp{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}
