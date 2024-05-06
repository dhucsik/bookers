package app

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dhucsik/bookers/configs"
	"github.com/dhucsik/bookers/internal/repositories/authors"
	"github.com/dhucsik/bookers/internal/repositories/books"
	"github.com/dhucsik/bookers/internal/repositories/categories"
	"github.com/dhucsik/bookers/internal/repositories/quizzes"
	"github.com/dhucsik/bookers/internal/repositories/users"
	"github.com/dhucsik/bookers/internal/services/auth"
	authorsS "github.com/dhucsik/bookers/internal/services/authors"
	booksS "github.com/dhucsik/bookers/internal/services/books"
	categoriesS "github.com/dhucsik/bookers/internal/services/categories"
	quizzesS "github.com/dhucsik/bookers/internal/services/quizzes"
	usersS "github.com/dhucsik/bookers/internal/services/users"
	"github.com/dhucsik/bookers/internal/transport/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	cfg *configs.Config

	usersRepository      users.Repository
	authorsRepository    authors.Repository
	categoriesRepository categories.Repository
	booksRepository      books.Repository
	quizzesRepository    quizzes.Repository

	usersService      usersS.Service
	authService       auth.Service
	authorsService    authorsS.Service
	categoriesService categoriesS.Service
	booksService      booksS.Service
	quizzesService    quizzesS.Service
	messageChat       chan string

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
		cfg:         cfg,
		messageChat: make(chan string, 5),
	}

	err := app.InitDatabase(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return app
}

func (a *App) InitRepositories(_ context.Context) error {
	a.usersRepository = users.NewRepository(a.db)
	a.authorsRepository = authors.NewRepository(a.db)
	a.categoriesRepository = categories.NewRepository(a.db)
	a.booksRepository = books.NewRepository(a.db)
	a.quizzesRepository = quizzes.NewRepository(a.db)

	return nil
}

func (a *App) InitServices(_ context.Context) error {
	var err error

	a.usersService = usersS.NewService(a.usersRepository, a.booksRepository)
	a.authService = auth.NewService(time.Hour, time.Hour, a.usersService)
	a.authorsService = authorsS.NewService(a.authorsRepository)
	a.categoriesService = categoriesS.NewService(a.categoriesRepository)
	a.booksService, err = booksS.NewService(
		a.booksRepository,
		a.authorsRepository,
		a.categoriesRepository,
		a.usersRepository,
		a.cfg.Env.Get("s3_endpoint"),
		a.cfg.Env.Get("s3_bucket"),
		a.cfg.Env.Get("s3_access_key"),
		a.cfg.Env.Get("s3_secret_key"),
	)
	if err != nil {
		return err
	}
	a.quizzesService = quizzesS.NewService(a.quizzesRepository, a.booksRepository, a.usersRepository)

	return nil
}

func (a *App) InitHTTPServer(_ context.Context) error {
	a.httpServer = http.NewServer(
		a.authService,
		a.usersService,
		a.authorsService,
		a.categoriesService,
		a.booksService,
		a.quizzesService,
		a.messageChat,
		a.cfg.Env.Get("api_key"),
	)

	return nil
}

func (a *App) Start(ctx context.Context) error {
	return a.httpServer.Start(ctx, fmt.Sprintf(":%s", a.cfg.HTTP.Port))
}
