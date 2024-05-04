package books

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/dhucsik/bookers/internal/models"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func searchQueryBuild(params *models.SearchParams) (string, []interface{}, error) {
	cols := []string{"id", "title", "pub_date", "edition", "language", "rating", "image", "description", "COUNT(*) OVER()"}

	q := psql.Select(cols...).From("books")
	q = q.Where(`title ILIKE '%' || ? || '%'`, params.Search)

	if len(params.Categories) > 0 {
		q = q.Join("books_categories bc ON bc.book_id = books.id")
		q = q.Where("bc.category_id = ANY(?)", params.Categories)
	}

	if len(params.Authors) > 0 {
		q = q.Join("books_authors ba ON ba.book_id = books.id")
		q = q.Where("ba.author_id = ANY(?)", params.Authors)
	}

	q = q.OrderBy("books.id").Limit(uint64(params.Limit)).Offset(uint64(params.Offset))

	return q.ToSql()
}

func getStockBooks(params *models.SearchParams) (string, []interface{}, error) {
	cols := []string{"sb.id", "sb.user_id", "sb.book_id", "b.id", "b.title", "b.pub_date", "b.edition", "b.language", "b.rating", "b.image", "b.description", "u.id", "u.username", "u.email", "u.city"}

	q := psql.Select(cols...).From("stock_books sb")
	q = q.Join("books b ON b.id = sb.book_id")
	q = q.Join("users u ON u.id = sb.user_id")
	q = q.Where(`b.title ILIKE '%' || ? || '%'`, params.Search)

	if len(params.City) > 0 {
		q = q.Where("u.city = ?", params.City)
	}

	if len(params.Categories) > 0 {
		q = q.Join("books_categories bc ON bc.book_id = b.id")
		q = q.Where("bc.category_id = ANY(?)", params.Categories)
	}

	if len(params.Authors) > 0 {
		q = q.Join("books_authors ba ON ba.book_id = b.id")
		q = q.Where("ba.author_id = ANY(?)", params.Authors)
	}

	q = q.OrderBy("sb.id").Limit(uint64(params.Limit)).Offset(uint64(params.Offset))

	return q.ToSql()
}
