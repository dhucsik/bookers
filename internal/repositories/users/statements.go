package users

const (
	createUserStmt = `INSERT INTO users (username, email, password) 
					VALUES ($1, $2, $3) 
					RETURNING id`

	setCityStmt = `UPDATE users SET city = $2 WHERE id = $1`

	getUserByIDStmt = `SELECT id, username, email, password, role, city 
					FROM users WHERE id = $1`

	getUserByUsernameStmt = `SELECT id, username, email, password, role, city
					FROM users WHERE username = $1`

	deleteUserStmt = `DELETE FROM users WHERE id = $1`

	getUsersByIDsStmt = `SELECT id, username, email, role, city
					FROM users WHERE id = ANY($1)`
)
