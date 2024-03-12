package categories

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	ListCategories(ctx context.Context) ([]*models.Category, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ListCategories(ctx context.Context) ([]*models.Category, error) {
	rows, err := r.db.Query(ctx, listCategoriesStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out categoryModels
	for rows.Next() {
		category := &categoryModel{}
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}

		out = append(out, category)
	}

	return out.convert(), nil
}
