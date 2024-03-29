package books

import (
	"github.com/dhucsik/bookers/internal/services/books"
	"github.com/dhucsik/bookers/internal/transport/http/middlewares"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	auth        *middlewares.AuthMiddleware
	bookService books.Service
}

func NewController(
	auth *middlewares.AuthMiddleware,
	bookService books.Service,
) *Controller {
	return &Controller{
		auth:        auth,
		bookService: bookService,
	}
}

func (r *Controller) Init(router *echo.Group) {

}
