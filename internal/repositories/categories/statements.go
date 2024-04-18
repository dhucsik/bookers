package categories

const (
	listCategoriesStmt = `SELECT id, name FROM categories`

	createCategoryStmt = `INSERT INTO categories (name) VALUES ($1) RETURNING id`

	deleteCategoryStmt = `DELETE FROM categories WHERE id = $1`

	listCategoriesByBookIDStmt = `SELECT c.id, c.name FROM categories c
	JOIN books_categories bc ON c.id = bc.category_id
	WHERE bc.book_id = $1`

	listCategoriesByBookIDsStmt = `SELECT bc.book_id, c.id, c.name FROM categories c
	JOIN books_categories bc ON c.id = bc.category_id
	WHERE bc.book_id = ANY($1)`
)
