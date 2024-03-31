package books

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateBook(ctx context.Context, book *models.Book, authorIDs, categoryIDs []int) error
	ListBooks(ctx context.Context, search string, limit, offset int) ([]*models.Book, error)
	GetBookByID(ctx context.Context, id int) (*models.Book, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateBook(ctx context.Context, book *models.Book, authorIDs, categoryIDs []int) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	err = tx.QueryRow(ctx, createBookStmt, book.Title, book.PubDate, book.Edition, book.Language, book.Rating).Scan(&book.ID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	for _, authorID := range authorIDs {
		_, err = tx.Exec(ctx, createBookAuthorStmt, book.ID, authorID)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	for _, categoryID := range categoryIDs {
		_, err = tx.Exec(ctx, createBookCategoryStmt, book.ID, categoryID)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *repository) ListBooks(ctx context.Context, search string, limit, offset int) ([]*models.Book, error) {
	rows, err := r.db.Query(ctx, listBooksStmt, search, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out bookModels
	for rows.Next() {
		book := &bookModel{}
		if err := rows.Scan(&book.ID, &book.Title, &book.PubDate, &book.Edition, &book.Language, &book.Rating); err != nil {
			return nil, err
		}

		out = append(out, book)
	}

	return out.convert(), nil
}

func (r *repository) GetBookByID(ctx context.Context, id int) (*models.Book, error) {
	book := &bookModel{}
	err := r.db.QueryRow(ctx, getBookStmt, id).Scan(&book.ID, &book.Title, &book.PubDate, &book.Edition, &book.Language, &book.Rating)
	if err != nil {
		return nil, err
	}

	return book.convert(), nil
}
