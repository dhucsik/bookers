package books

import (
	"time"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/samber/lo"
)

type bookModel struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	PubDate     time.Time `db:"pub_date"`
	Edition     string    `db:"edition"`
	Language    string    `db:"language"`
	Rating      float64   `db:"rating"`
	Image       string    `db:"image"`
	Description string    `db:"description"`
}

type bookModels []*bookModel

func (b *bookModel) convert() *models.Book {
	return &models.Book{
		ID:          b.ID,
		Title:       b.Title,
		PubDate:     b.PubDate,
		Edition:     b.Edition,
		Language:    b.Language,
		Rating:      b.Rating,
		Image:       b.Image,
		Description: b.Description,
	}
}

func convertBook(book *models.Book) *bookModel {
	return &bookModel{
		ID:          book.ID,
		Title:       book.Title,
		PubDate:     book.PubDate,
		Edition:     book.Edition,
		Language:    book.Language,
		Rating:      book.Rating,
		Image:       book.Image,
		Description: book.Description,
	}
}

func (b bookModels) convert() []*models.Book {
	return lo.Map(b, func(model *bookModel, _ int) *models.Book {
		return model.convert()
	})
}

type bookComment struct {
	ID        int       `db:"id"`
	BookID    int       `db:"book_id"`
	UserID    int       `db:"user_id"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
}

func (b *bookComment) convert() *models.BookComment {
	return &models.BookComment{
		ID:        b.ID,
		BookID:    b.BookID,
		UserID:    b.UserID,
		Comment:   b.Comment,
		CreatedAt: b.CreatedAt,
	}
}
