package users

import (
	"github.com/dhucsik/bookers/internal/services/books"
	"github.com/dhucsik/bookers/internal/services/users"
	"github.com/dhucsik/bookers/internal/transport/http/middlewares"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	auth         *middlewares.AuthMiddleware
	usersService users.Service
	bookService  books.Service
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

	router.POST("/users/friends/:id/request", r.auth.Handler(r.sendFriendRequestHandler))
	router.PUT("/users/friends/:id/accept", r.auth.Handler(r.acceptFriendRequestHandler))
	router.GET("/users/friends", r.auth.Handler(r.getFriendsHandler))
	router.GET("/users/friends/sent", r.auth.Handler(r.getSentRequests))
	router.GET("/users/friends/received", r.auth.Handler(r.getReceivedRequests))
}
