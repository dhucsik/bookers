package books

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
)

func (r *repository) UploadStockBook(ctx context.Context, book *models.StockBook) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, uploadStockBookStmt, book.UserID, book.BookID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) GetStockBook(ctx context.Context, bookID int) (*models.StockBook, error) {
	book := &models.StockBook{}
	err := r.db.QueryRow(ctx, getStockBookStmt, bookID).Scan(&book.ID, &book.UserID, &book.BookID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	book.SetImage()
	return book, nil
}

func (r *repository) GetBooksByStockIDs(ctx context.Context, ids []int) (map[int]*models.Book, error) {
	rows, err := r.db.Query(ctx, getBooksByStockIDsStmt, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make(map[int]*models.Book)
	for rows.Next() {
		var id int
		book := &models.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.PubDate, &book.Edition, &book.Language, &book.Rating, &book.Image, &book.Description, &id)
		if err != nil {
			return nil, err
		}
		out[id] = book
	}

	return out, nil
}

func (r *repository) GetStockBooksByUser(ctx context.Context, userID int) ([]*models.StockBook, error) {
	rows, err := r.db.Query(ctx, getStockBooksByUserStmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stockBooks []*models.StockBook
	for rows.Next() {
		book := &models.StockBook{}
		err := rows.Scan(&book.ID, &book.UserID, &book.BookID)
		if err != nil {
			return nil, err
		}

		book.SetImage()
		stockBooks = append(stockBooks, book)
	}

	return stockBooks, nil
}

func (r *repository) GetStockByBook(ctx context.Context, bookID int) ([]*models.StockBook, error) {
	rows, err := r.db.Query(ctx, getStockByBookStmt, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stockBooks []*models.StockBook
	for rows.Next() {
		book := &models.StockBook{}
		err := rows.Scan(&book.ID, &book.UserID, &book.BookID)
		if err != nil {
			return nil, err
		}

		book.SetImage()
		stockBooks = append(stockBooks, book)
	}

	return stockBooks, nil
}

func (r *repository) DeleteStockBook(ctx context.Context, bookID int) error {
	_, err := r.db.Exec(ctx, deleteStockBookStmt, bookID)
	return err
}

func (r *repository) GetUserStockCount(ctx context.Context, userID int) (int, error) {
	var count int
	err := r.db.QueryRow(ctx, getUserStockCountStmt, userID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *repository) SearchStockBooks(ctx context.Context, params *models.SearchParams) ([]*models.StockBookWithFields, error) {
	q, args, err := getStockBooks(params)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*models.StockBookWithFields
	for rows.Next() {
		stockBook := &models.StockBookWithFields{}
		book := &models.Book{}
		user := &models.UserWithoutPassword{}

		err := rows.Scan(&stockBook.ID,
			&stockBook.UserID, &stockBook.BookID,
			&book.ID, &book.Title, &book.PubDate,
			&book.Edition, &book.Language,
			&book.Rating, &book.Image,
			&book.Description, &user.ID,
			&user.Username, &user.Email, &user.City)
		if err != nil {
			return nil, err
		}

		user.SetProfilePic()
		stockBook.Book = book
		stockBook.User = user
		stockBook.SetImage()

		out = append(out, stockBook)
	}

	return out, nil
}
