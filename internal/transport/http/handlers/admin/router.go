package admin

import (
	"github.com/labstack/echo/v4"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (r *Controller) Init(router *echo.Group) {
	router.POST("/admin/categories", r.createCategory)
	router.DELETE("/admin/categories/:id", r.deleteCategory)
	router.POST("/admin/authors", r.createAuthor)
	router.DELETE("/admin/authors/:id", r.deleteAuthor)
}
