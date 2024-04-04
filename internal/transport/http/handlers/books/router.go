package books

import (
	"github.com/dhucsik/bookers/internal/services/books"
	"github.com/dhucsik/bookers/internal/services/quizzes"
	"github.com/dhucsik/bookers/internal/transport/http/middlewares"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	auth        *middlewares.AuthMiddleware
	bookService books.Service
	quizService quizzes.Service
}

func NewController(
	auth *middlewares.AuthMiddleware,
	bookService books.Service,
	quizService quizzes.Service,
) *Controller {
	return &Controller{
		auth:        auth,
		bookService: bookService,
		quizService: quizService,
	}
}

func (r *Controller) Init(router *echo.Group) {
	router.GET("/books", r.listBooksHandler)
	router.GET("/books/:id", r.getBookByIDHandler)

	router.POST("/books/:id/quizzes", r.auth.Handler(r.createQuizHandler))

	router.GET("/books/:id/comments", r.listCommentsHandler)
	router.POST("/books/:id/comments", r.auth.Handler(r.addCommentHandler))
	router.PUT("/books/comments/:id", r.auth.Handler(r.updateCommentHandler))
	router.DELETE("/books/comments/:id", r.auth.Handler(r.deleteCommentHandler))

	router.POST("/books/:id/rating", r.auth.Handler(r.setRatingHandler))
}
