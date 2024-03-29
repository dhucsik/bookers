package books

import (
	"time"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/samber/lo"
)

type bookModel struct {
	ID       int       `db:"id"`
	Title    string    `db:"title"`
	PubDate  time.Time `db:"pub_date"`
	Edition  int       `db:"edition"`
	Language string    `db:"language"`
	Rating   float64   `db:"rating"`
}

type bookModels []*bookModel

func (b *bookModel) convert() *models.Book {
	return &models.Book{
		ID:       b.ID,
		Title:    b.Title,
		PubDate:  b.PubDate,
		Edition:  b.Edition,
		Language: b.Language,
		Rating:   b.Rating,
	}
}

func convertBook(book *models.Book) *bookModel {
	return &bookModel{
		ID:       book.ID,
		Title:    book.Title,
		PubDate:  book.PubDate,
		Edition:  book.Edition,
		Language: book.Language,
		Rating:   book.Rating,
	}
}

func (b bookModels) convert() []*models.Book {
	return lo.Map(b, func(model *bookModel, _ int) *models.Book {
		return model.convert()
	})
}
