package authors

import (
	"github.com/dhucsik/bookers/internal/services/authors"
	"github.com/dhucsik/bookers/internal/transport/http/middlewares"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	auth           *middlewares.AuthMiddleware
	authorsService authors.Service
}

func NewController(
	auth *middlewares.AuthMiddleware,
	authorsService authors.Service,
) *Controller {
	return &Controller{
		auth:           auth,
		authorsService: authorsService,
	}
}

func (r *Controller) Init(router *echo.Group) {
	router.GET("/authors", r.listAuthors)
}
