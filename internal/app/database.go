package app

import (
	"context"

	"github.com/dhucsik/bookers/internal/repositories"
)

func (a *App) InitDatabase(ctx context.Context) error {
	pool, err := repositories.NewPostgres(ctx, a.Config().Env["postgres_dsn"].Value)
	if err != nil {
		return err
	}

	a.db = pool
	return nil
}
