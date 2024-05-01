package books

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
)

func (r *repository) AddBookToLiked(ctx context.Context, userID, bookID int) error {
	_, err := r.db.Exec(ctx, addBookToLikedStmt, userID, bookID)
	return err
}

func (r *repository) RemoveBookFromLiked(ctx context.Context, userID, bookID int) error {
	_, err := r.db.Exec(ctx, deleteFromLikedStmt, userID, bookID)
	return err
}

func (r *repository) GetLikedBooks(ctx context.Context, userID int) ([]*models.Book, error) {
	rows, err := r.db.Query(ctx, getLikedBooksStmt, userID)
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
