package authors

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	ListAuthors(ctx context.Context) ([]*models.Author, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ListAuthors(ctx context.Context) ([]*models.Author, error) {
	rows, err := r.db.Query(ctx, listAuthorsStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out authorModels
	for rows.Next() {
		author := &authorModel{}
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, err
		}

		out = append(out, author)
	}

	return out.convert(), nil
}
