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
	booksService books.Service,
) *Controller {
	return &Controller{
		auth:         auth,
		usersService: usersService,
		bookService:  booksService,
	}
}

func (r *Controller) Init(router *echo.Group) {
	router.POST("/users", r.createUser)

	router.POST("/users/liked-books", r.auth.Handler(r.addLikedBook))
	router.DELETE("/users/liked-books", r.auth.Handler(r.removeLikedBook))
	router.GET("/users/liked-books", r.auth.Handler(r.getLikedBooks))

	router.PUT("/users/city", r.auth.Handler(r.setCity))
	router.PUT("/users/username", r.auth.Handler(r.updateUsername))
	router.PUT("/users/password", r.auth.Handler(r.updatePassword))
	router.PUT("/users/phone", r.auth.Handler(r.updatePhone))
	router.DELETE("/users/:id", r.auth.Handler(r.deleteUser))
	router.POST("/users/profile/image", r.auth.Handler(r.uploadProfilePicHandler))

	router.GET("/users/:id", r.auth.Handler(r.getByID))

	router.POST("/users/friends/:id/request", r.auth.Handler(r.sendFriendRequestHandler))
	router.PUT("/users/friends/:id/accept", r.auth.Handler(r.acceptFriendRequestHandler))
	router.GET("/users/friends", r.auth.Handler(r.getFriendsHandler))
	router.GET("/users/friends/sent", r.auth.Handler(r.getSentRequests))
	router.GET("/users/friends/received", r.auth.Handler(r.getReceivedRequests))
}
