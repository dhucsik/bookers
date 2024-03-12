package users

import (
	"github.com/dhucsik/bookers/internal/services/users"
	"github.com/dhucsik/bookers/internal/transport/http/middlewares"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	auth         *middlewares.AuthMiddleware
	usersService users.Service
}

func NewController(
	auth *middlewares.AuthMiddleware,
	usersService users.Service,
) *Controller {
	return &Controller{
		auth:         auth,
		usersService: usersService,
	}
}

func (r *Controller) Init(router *echo.Group) {
	router.POST("/users", r.createUser)
	router.PATCH("/users/:id/city", r.auth.Handler(r.setCity))
	router.DELETE("/users/:id", r.auth.Handler(r.deleteUser))

	router.GET("/users/:id", r.auth.Handler(r.getByID))
}
