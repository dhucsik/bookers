package models

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Role     string
	City     *string
}

func (u *User) ToUserWithoutPassword() *UserWithoutPassword {
	return &UserWithoutPassword{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

type UserWithoutPassword struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

const (
	FriendRequestSent     = "sent"
	FriendRequestAccepted = "accepted"
)

type FriendRequest struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	FriendID int    `json:"friend_id"`
	Status   string `json:"status"`
}
