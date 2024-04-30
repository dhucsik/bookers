package personal

import (
	"github.com/labstack/echo/v4"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (r *Controller) Init(router *echo.Group) {
	router.File("/main", "dist/index.html")
}
