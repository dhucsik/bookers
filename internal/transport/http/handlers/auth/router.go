package auth

import (
	"github.com/dhucsik/bookers/internal/services/auth"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	authService auth.Service
}

func NewController(authService auth.Service) *Controller {
	return &Controller{authService: authService}
}

func (r *Controller) Init(router *echo.Group) {
	router.POST("/auth", r.authHandler)
	router.POST("/auth/refresh", r.refreshHandler)
}
