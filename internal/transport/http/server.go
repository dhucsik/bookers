package http

import (
	"context"
	"log"
	"net/http"

	"github.com/dhucsik/bookers/internal/services/auth"
	"github.com/dhucsik/bookers/internal/services/users"
	authC "github.com/dhucsik/bookers/internal/transport/http/handlers/auth"
	usersC "github.com/dhucsik/bookers/internal/transport/http/handlers/users"
	"github.com/dhucsik/bookers/internal/transport/http/middlewares"
	"github.com/labstack/echo/v4"
)

// @title Swagger Bookers API
// @version 1.0
// @description bookers server.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.htm

// @host localhost:8080
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
) *Server {
	srv := echo.New()
	router := srv.Group("/api/v1")

	server := &Server{
		server: srv,
		router: router,
	}

	authMiddleware := middlewares.NewAuthMiddleware()

	server.WithControllers(
		authC.NewController(authService),
		usersC.NewController(authMiddleware, usersService),
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
