package categories

import (
	"github.com/dhucsik/bookers/internal/services/categories"
	"github.com/dhucsik/bookers/internal/transport/http/middlewares"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	auth              *middlewares.AuthMiddleware
	categoriesService categories.Service
}

func NewController(
	auth *middlewares.AuthMiddleware,
	categoriesService categories.Service,
) *Controller {
	return &Controller{
		auth:              auth,
		categoriesService: categoriesService,
	}
}

func (r *Controller) Init(router *echo.Group) {
	router.GET("/categories", r.listCategories)
}
