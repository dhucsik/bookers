package categories

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type Repository interface {
	ListCategories(ctx context.Context) ([]*models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, id int) error
	GetByBookID(ctx context.Context, bookID int) ([]*models.Category, error)
	GetByBookIDs(ctx context.Context, bookIDs []int) (map[int][]*models.Category, error)
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

func (r *repository) CreateCategory(ctx context.Context, category *models.Category) error {
	_, err := r.db.Exec(ctx, createCategoryStmt, category.Name)
	return err
}

func (r *repository) DeleteCategory(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, deleteCategoryStmt, id)
	return err
}

func (r *repository) GetByBookID(ctx context.Context, bookID int) ([]*models.Category, error) {
	rows, err := r.db.Query(ctx, listCategoriesByBookIDStmt, bookID)
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

func (r *repository) GetByBookIDs(ctx context.Context, bookIDs []int) (map[int][]*models.Category, error) {
	rows, err := r.db.Query(ctx, listCategoriesByBookIDsStmt, pq.Array(bookIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make(map[int][]*models.Category)
	for rows.Next() {
		category := &categoryModel{}
		var bookID int
		if err := rows.Scan(&bookID, &category.ID, &category.Name); err != nil {
			return nil, err
		}

		out[bookID] = append(out[bookID], category.convert())
	}

	return out, nil
}
