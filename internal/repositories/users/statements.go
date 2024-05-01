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

	createFriendRequestStmt = `INSERT INTO friends (user_id, friend_id, status)
					VALUES ($1, $2, $3)`

	acceptFriendRequestStmt = `UPDATE friends SET status = $3
					WHERE user_id = $1 AND friend_id = $2`

	getFriendsStmt = `SELECT u.id, u.username, u.email, u.password, u.role, u.city
	FROM users u
	JOIN friends f ON 
		(f.user_id = $1 AND u.id = f.friend_id)
		OR (f.friend_id = $1 AND u.id = f.user_id)
	WHERE f.status = 'accepted'
	`

	getFriendRequestStmt = `SELECT user_id, friend_id, status
	FROM friends WHERE (user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1)
	`

	getSentRequestFriendsStmt = `SELECT u.id, u.username, u.email, u.city
	FROM users u JOIN friends f ON u.id = f.friend_id WHERE f.user_id = $1 AND f.status = 'sent'`

	getReceivedRequestFriendsStmt = `SELECT u.id, u.username, u.email, u.city
	FROM users u JOIN friends f ON u.id = f.user_id WHERE f.friend_id = $1 AND f.status = 'sent'`

	updateUsernameStmt = `UPDATE users SET username = $2 WHERE id = $1`

	updatePasswordStmt = `UPDATE users SET password = $2 WHERE id = $1`
)
