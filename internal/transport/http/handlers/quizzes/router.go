package quizzes

import (
	"github.com/dhucsik/bookers/internal/services/quizzes"
	"github.com/dhucsik/bookers/internal/transport/http/middlewares"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	quizService quizzes.Service
	auth        *middlewares.AuthMiddleware
}

func NewController(auth *middlewares.AuthMiddleware, quizService quizzes.Service) *Controller {
	return &Controller{
		quizService: quizService,
		auth:        auth,
	}
}

func (c *Controller) Init(router *echo.Group) {
	router.GET("/quizzes/:id", c.auth.Handler(c.getQuizHandler))
	router.GET("/quizzes/:id/view", c.auth.Handler(c.viewQuizHandler))
	router.PUT("/quizzes/:id", c.auth.Handler(c.updateQuizHandler))
	router.DELETE("/quizzes/:id", c.auth.Handler(c.deleteQuizHandler))

	router.GET("/quizzes/results/:id", c.auth.Handler(c.getQuizResultHandler))
	router.POST("/quizzes/:id/check", c.auth.Handler(c.checkQuizHandler))
	router.GET("/quizzes/results", c.auth.Handler(c.getQuizResultsHandler))

	router.POST("/quizzes/:id/questions", c.auth.Handler(c.addQuestionHandler))
	router.PUT("/quizzes/questions/:id", c.auth.Handler(c.updateQuestionHandler))
	router.DELETE("/quizzes/questions/:id", c.auth.Handler(c.deleteQuestionHandler))

	router.GET("/quizzes/:id/comments", c.auth.Handler(c.listCommentsHandler))
	router.POST("/quizzes/:id/comments", c.auth.Handler(c.addCommentHandler))
	router.PUT("/quizzes/comments/:id", c.auth.Handler(c.updateCommentHandler))
	router.DELETE("/quizzes/comments/:id", c.auth.Handler(c.deleteCommentHandler))

	router.POST("/quizzes/:id/rating", c.auth.Handler(c.setRatingHandler))
}
