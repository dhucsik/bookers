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
		q = q.Join("book_categories bc ON bc.book_id = books.id")
		q = q.Where("bc.category_id = ANY(?)", params.Categories)
	}

	if len(params.Authors) > 0 {
		q = q.Join("book_authors ba ON ba.book_id = books.id")
		q = q.Where("ba.author_id = ANY(?)", params.Authors)
	}

	q = q.OrderBy("books.id").Limit(uint64(params.Limit)).Offset(uint64(params.Offset))

	return q.ToSql()
}
