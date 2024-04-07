package books

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
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
		return nil, err
	}

	return book, nil
}

func (r *repository) GetBooksByStockIDs(ctx context.Context, ids []int) ([]*models.Book, error) {
	rows, err := r.db.Query(ctx, getBooksByStockIDsStmt, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		book := &models.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.PubDate, &book.Edition, &book.Language, &book.Rating, &book.Image, &book.Description)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
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
		stockBooks = append(stockBooks, book)
	}

	return stockBooks, nil
}
