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
	SetRating(ctx context.Context, bookID, userID, rating int) error
	InsertComment(ctx context.Context, bookID, userID int, comment string) error
	UpdateComment(ctx context.Context, id int, comment string) error
	ListComments(ctx context.Context, bookID int) ([]*models.BookComment, error)
	DeleteComment(ctx context.Context, id int) error
	GetComment(ctx context.Context, id int) (*models.BookComment, error)
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

func (r *repository) InsertComment(ctx context.Context, bookID, userID int, comment string) error {
	_, err := r.db.Exec(ctx, insertCommentStmt, bookID, userID, comment)
	return err
}

func (r *repository) UpdateComment(ctx context.Context, id int, comment string) error {
	_, err := r.db.Exec(ctx, updateCommentStmt, id, comment)
	return err
}

func (r *repository) SetRating(ctx context.Context, bookID, userID, rating int) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, insertBookRatingStmt, bookID, userID, rating)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	var avgRating float64
	err = tx.QueryRow(ctx, getAvgRatingStmt, bookID).Scan(&avgRating)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, updateBookRatingStmt, avgRating, bookID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}

func (r *repository) ListComments(ctx context.Context, bookID int) ([]*models.BookComment, error) {
	rows, err := r.db.Query(ctx, listCommentsStmt, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models.BookComment
	for rows.Next() {
		comment := &bookComment{}
		if err := rows.Scan(&comment.ID, &comment.BookID, &comment.UserID, &comment.Comment, &comment.CreatedAt); err != nil {
			return nil, err
		}

		out = append(out, comment.convert())
	}

	return out, nil
}

func (r *repository) DeleteComment(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, deleteCommentStmt, id)
	return err
}

func (r *repository) GetComment(ctx context.Context, id int) (*models.BookComment, error) {
	comment := &bookComment{}
	err := r.db.QueryRow(ctx, getCommentStmt, id).Scan(&comment.ID, &comment.BookID, &comment.UserID, &comment.Comment, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return comment.convert(), nil
}
