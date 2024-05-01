package books

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type Repository interface {
	CreateBook(ctx context.Context, book *models.Book, authorIDs, categoryIDs []int) (int, error)
	ListBooks(ctx context.Context, search string, limit, offset int) ([]*models.Book, int, error)
	SearchBooks(ctx context.Context, params *models.SearchParams) ([]*models.Book, int, error)
	GetBookByID(ctx context.Context, id int) (*models.Book, error)
	GetBooksByIDs(ctx context.Context, ids []int) ([]*models.Book, error)
	SetRating(ctx context.Context, bookID, userID, rating int) error
	InsertComment(ctx context.Context, bookID, userID int, comment string) (int, error)
	UpdateComment(ctx context.Context, id int, comment string) error
	ListComments(ctx context.Context, bookID int) ([]*models.BookComment, error)
	DeleteComment(ctx context.Context, id int) error
	GetComment(ctx context.Context, id int) (*models.BookComment, error)
	UploadStockBook(ctx context.Context, book *models.StockBook) (int, error)
	CreateRequest(ctx context.Context, req *models.ShareRequest) error
	UpdateRequest(ctx context.Context, req *models.ShareRequest) error
	GetRequest(ctx context.Context, id int) (*models.ShareRequest, error)
	GetStockBook(ctx context.Context, bookID int) (*models.StockBook, error)
	GetBooksByStockIDs(ctx context.Context, ids []int) (map[int]*models.Book, error)
	ListRequests(ctx context.Context, userID int) ([]*models.ShareRequest, error)
	GetStockBooksByUser(ctx context.Context, userID int) ([]*models.StockBook, error)
	GetStockByBook(ctx context.Context, bookID int) ([]*models.StockBook, error)
	GetUserStockCount(ctx context.Context, userID int) (int, error)
	GetSuccessRequestCount(ctx context.Context, userID int) (int, error)
	DeleteStockBook(ctx context.Context, bookID int) error
	AddBookToLiked(ctx context.Context, userID, bookID int) error
	RemoveBookFromLiked(ctx context.Context, userID, bookID int) error
	GetLikedBooks(ctx context.Context, userID int) ([]*models.Book, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateBook(ctx context.Context, book *models.Book, authorIDs, categoryIDs []int) (int, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}

	err = tx.QueryRow(ctx, createBookStmt, book.Title, book.PubDate, book.Edition, book.Language, book.Rating, book.Image, book.Description).Scan(&book.ID)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	for _, authorID := range authorIDs {
		_, err = tx.Exec(ctx, createBookAuthorStmt, book.ID, authorID)
		if err != nil {
			tx.Rollback(ctx)
			return 0, err
		}
	}

	for _, categoryID := range categoryIDs {
		_, err = tx.Exec(ctx, createBookCategoryStmt, book.ID, categoryID)
		if err != nil {
			tx.Rollback(ctx)
			return 0, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return book.ID, nil
}

func (r *repository) ListBooks(ctx context.Context, search string, limit, offset int) ([]*models.Book, int, error) {
	var totalCount int

	rows, err := r.db.Query(ctx, listBooksStmt, search, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var out bookModels
	for rows.Next() {
		book := &bookModel{}
		if err := rows.Scan(&book.ID, &book.Title, &book.PubDate, &book.Edition, &book.Language, &book.Rating, &book.Image, &book.Description, &totalCount); err != nil {
			return nil, 0, err
		}

		out = append(out, book)
	}

	return out.convert(), totalCount, nil
}

func (r *repository) SearchBooks(ctx context.Context, params *models.SearchParams) ([]*models.Book, int, error) {
	q, args, err := searchQueryBuild(params)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(ctx, q, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var totalCount int
	var out bookModels
	for rows.Next() {
		book := &bookModel{}
		if err := rows.Scan(&book.ID, &book.Title, &book.PubDate, &book.Edition, &book.Language, &book.Rating, &book.Image, &book.Description, &totalCount); err != nil {
			return nil, 0, err
		}

		out = append(out, book)
	}

	return out.convert(), totalCount, nil
}

func (r *repository) GetBookByID(ctx context.Context, id int) (*models.Book, error) {
	book := &bookModel{}
	err := r.db.QueryRow(ctx, getBookStmt, id).Scan(&book.ID, &book.Title, &book.PubDate, &book.Edition, &book.Language, &book.Rating, &book.Image, &book.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return book.convert(), nil
}

func (r *repository) GetBooksByIDs(ctx context.Context, ids []int) ([]*models.Book, error) {
	rows, err := r.db.Query(ctx, getBooksByIDsStmt, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out bookModels
	for rows.Next() {
		book := &bookModel{}
		if err := rows.Scan(&book.ID, &book.Title, &book.PubDate, &book.Edition, &book.Language, &book.Rating, &book.Image, &book.Description); err != nil {
			return nil, err
		}

		out = append(out, book)
	}

	return out.convert(), nil
}

func (r *repository) InsertComment(ctx context.Context, bookID, userID int, comment string) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, insertCommentStmt, bookID, userID, comment).Scan(&id)
	return id, err
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
