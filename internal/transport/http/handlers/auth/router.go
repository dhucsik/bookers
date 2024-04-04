package auth

import (
	"github.com/dhucsik/bookers/internal/services/auth"
	"github.com/dhucsik/bookers/internal/services/users"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	authService  auth.Service
	usersService users.Service
}

func NewController(authService auth.Service, usersService users.Service) *Controller {
	return &Controller{
		authService:  authService,
		usersService: usersService,
	}
}

func (r *Controller) Init(router *echo.Group) {
	router.POST("/auth", r.authHandler)
	router.POST("/auth/refresh", r.refreshHandler)
}
