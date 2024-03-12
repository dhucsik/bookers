package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dhucsik/bookers/configs"
	"github.com/dhucsik/bookers/internal/repositories/users"
	"github.com/dhucsik/bookers/internal/services/auth"
	usersS "github.com/dhucsik/bookers/internal/services/users"
	"github.com/dhucsik/bookers/internal/transport/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	cfg *configs.Config

	usersRepository users.Repository

	usersService usersS.Service
	authService  auth.Service

	httpServer *http.Server
	db         *pgxpool.Pool
}

func InitApp(ctx context.Context) *App {
	app := NewApp(ctx)

	for _, init := range []func(ctx context.Context) error{
		app.InitRepositories,
		app.InitServices,
		app.InitHTTPServer,
	} {
		err := init(ctx)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}

	return app
}

func (a *App) Config() *configs.Config {
	return a.cfg
}

func NewApp(ctx context.Context) *App {
	config, err := configs.Parse()
	if err != nil {
		log.Fatal(err)
	}

	return NewAppWithConfig(ctx, config)
}

func NewAppWithConfig(ctx context.Context, cfg *configs.Config) *App {
	app := &App{
		cfg: cfg,
	}

	err := app.InitDatabase(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return app
}

func (a *App) InitRepositories(_ context.Context) error {
	a.usersRepository = users.NewRepository(a.db)

	return nil
}

func (a *App) InitServices(_ context.Context) error {
	a.usersService = usersS.NewService(a.usersRepository)
	a.authService = auth.NewService(time.Hour, time.Hour, a.usersService)

	return nil
}

func (a *App) InitHTTPServer(_ context.Context) error {
	a.httpServer = http.NewServer(a.authService, a.usersService)

	return nil
}

func (a *App) Start(ctx context.Context) error {
	return a.httpServer.Start(ctx, fmt.Sprintf(":%s", a.cfg.HTTP.Port))
}
