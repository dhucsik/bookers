package authors

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type Repository interface {
	ListAuthors(ctx context.Context, search string, limit, offset int) ([]*models.Author, int, error)
	CreateAuthor(ctx context.Context, author *models.Author) (int, error)
	DeleteAuthor(ctx context.Context, id int) error
	GetByBookID(ctx context.Context, bookID int) ([]*models.Author, error)
	GetByBookIDs(ctx context.Context, bookIDs []int) (map[int][]*models.Author, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ListAuthors(ctx context.Context, search string, limit, offset int) ([]*models.Author, int, error) {
	var totalCount int
	rows, err := r.db.Query(ctx, listAuthorsStmt, search, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var out authorModels
	for rows.Next() {
		author := &authorModel{}
		if err := rows.Scan(&author.ID, &author.Name, &totalCount); err != nil {
			return nil, 0, err
		}

		out = append(out, author)
	}

	return out.convert(), totalCount, nil
}

func (r *repository) CreateAuthor(ctx context.Context, author *models.Author) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, createAuthorStmt, author.Name).Scan(&id)
	return id, err
}

func (r *repository) DeleteAuthor(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, deleteAuthorStmt, id)
	return err
}

func (r *repository) GetByBookID(ctx context.Context, bookID int) ([]*models.Author, error) {
	rows, err := r.db.Query(ctx, listAuthorsByBookIDStmt, bookID)
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

func (r *repository) GetByBookIDs(ctx context.Context, bookIDs []int) (map[int][]*models.Author, error) {
	rows, err := r.db.Query(ctx, listAuthorsByBookIDsStmt, pq.Array(bookIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make(map[int][]*models.Author)
	for rows.Next() {
		var bookID int
		author := &authorModel{}
		if err := rows.Scan(&bookID, &author.ID, &author.Name); err != nil {
			return nil, err
		}

		out[bookID] = append(out[bookID], author.convert())
	}

	return out, nil
}
