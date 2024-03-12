package middlewares

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/util/jwt"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		session, isRefresh, err := jwt.ParseJWT(token)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, err.Error())
		}

		if isRefresh {
			return ctx.JSON(http.StatusUnauthorized, "unexpected refresh token")
		}

		if session == nil {
			return ctx.JSON(http.StatusUnauthorized, "empty session")
		}

		sessionCtx := session.SetInCtx(ctx.Request().Context())
		ctx.SetRequest(ctx.Request().WithContext(sessionCtx))

		return next(ctx)
	}
}
