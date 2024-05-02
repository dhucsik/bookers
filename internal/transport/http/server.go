package http

import (
	"context"
	"log"
	"net/http"

	"github.com/dhucsik/bookers/internal/services/auth"
	"github.com/dhucsik/bookers/internal/services/authors"
	"github.com/dhucsik/bookers/internal/services/books"
	"github.com/dhucsik/bookers/internal/services/categories"
	"github.com/dhucsik/bookers/internal/services/quizzes"
	"github.com/dhucsik/bookers/internal/services/users"
	"github.com/dhucsik/bookers/internal/transport/http/handlers/admin"
	authC "github.com/dhucsik/bookers/internal/transport/http/handlers/auth"
	authorsC "github.com/dhucsik/bookers/internal/transport/http/handlers/authors"
	booksC "github.com/dhucsik/bookers/internal/transport/http/handlers/books"
	categoriesC "github.com/dhucsik/bookers/internal/transport/http/handlers/categories"
	"github.com/dhucsik/bookers/internal/transport/http/handlers/personal"
	quizzesC "github.com/dhucsik/bookers/internal/transport/http/handlers/quizzes"
	"github.com/dhucsik/bookers/internal/transport/http/handlers/swag"
	usersC "github.com/dhucsik/bookers/internal/transport/http/handlers/users"
	"github.com/dhucsik/bookers/internal/transport/http/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Swagger Bookers API
// @version 1.0
// @description bookers server.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.htm

// @host bookers.kz
// @BasePath /api/v1

type IController interface {
	Init(group *echo.Group)
}

type Server struct {
	server *echo.Echo
	router *echo.Group
}

func NewServer(
	authService auth.Service,
	usersService users.Service,
	authorsService authors.Service,
	categoriesService categories.Service,
	booksService books.Service,
	quizzesService quizzes.Service,
	chat chan string,
) *Server {
	srv := echo.New()
	router := srv.Group("/api/v1")

	srv.Use(middleware.Logger())
	srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderXRequestedWith,
		},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	},
	))

	server := &Server{
		server: srv,
		router: router,
	}

	authMiddleware := middlewares.NewAuthMiddleware()

	server.WithControllers(
		authC.NewController(authService, usersService),
		usersC.NewController(authMiddleware, usersService),
		authorsC.NewController(authMiddleware, authorsService),
		categoriesC.NewController(authMiddleware, categoriesService),
		admin.NewController(authMiddleware, authorsService, categoriesService, booksService),
		booksC.NewController(authMiddleware, booksService, quizzesService),
		quizzesC.NewController(authMiddleware, quizzesService),
		personal.NewController(chat),
		swag.NewController(),
	)

	return server
}

func (s *Server) WithControllers(controllers ...IController) {
	for _, ctrl := range controllers {
		ctrl.Init(s.router)
	}
}

func (s *Server) Start(ctx context.Context, addr string) error {
	go func() {
		if err := s.server.Start(addr); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()

	return nil
}
