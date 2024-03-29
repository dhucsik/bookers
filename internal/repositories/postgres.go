package repositories

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/avast/retry-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
)

func NewPostgres(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var inErr error

	err := retry.Do(func() error {
		log.Println("Trying to connect to the database...")
		pool, inErr = pgxpool.New(ctx, dsn)
		if inErr != nil {
			return inErr
		}

		return nil
	}, retry.Attempts(3), retry.Delay(1*time.Second), retry.Context(ctx))
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		return nil, err
	}

	return pool, nil
}
