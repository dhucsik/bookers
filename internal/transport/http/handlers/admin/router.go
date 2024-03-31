package admin

import (
	"github.com/dhucsik/bookers/internal/services/authors"
	"github.com/dhucsik/bookers/internal/services/books"
	"github.com/dhucsik/bookers/internal/services/categories"
	"github.com/dhucsik/bookers/internal/transport/http/middlewares"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	auth              *middlewares.AuthMiddleware
	authorService     authors.Service
	categoriesService categories.Service
	bookService       books.Service
}

func NewController(
	auth *middlewares.AuthMiddleware,
	authorService authors.Service,
	categoriesService categories.Service,
	booksService books.Service,
) *Controller {
	return &Controller{
		auth:              auth,
		authorService:     authorService,
		categoriesService: categoriesService,
		bookService:       booksService,
	}
}

func (r *Controller) Init(router *echo.Group) {
	router.POST("/admin/categories", r.auth.Handler(r.createCategory))
	router.DELETE("/admin/categories/:id", r.auth.Handler(r.deleteCategory))
	router.POST("/admin/authors", r.auth.Handler(r.createAuthor))
	router.DELETE("/admin/authors/:id", r.auth.Handler(r.deleteAuthor))
	router.POST("/admin/books", r.auth.Handler(r.createBookHandler))
}
