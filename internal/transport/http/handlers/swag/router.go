package swag

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Init(r *echo.Group) {
	r.GET("/swagger/*", echoSwagger.WrapHandler)
}
