package models

import "fmt"

type User struct {
	ID         int     `json:"id"`
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	Password   string  `json:"password,omitempty"`
	Role       string  `json:"role"`
	City       *string `json:"city,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	ProfilePic string  `json:"profile_pic,omitempty"`
}

func (u *User) SetProfilePic() {
	u.ProfilePic = fmt.Sprintf("https://bookers-images.hb.kz-ast.vkcs.cloud/users/%d.png", u.ID)
}

func (u *User) ToUserWithoutPassword() *UserWithoutPassword {
	return &UserWithoutPassword{
		ID:         u.ID,
		Username:   u.Username,
		Email:      u.Email,
		City:       u.City,
		ProfilePic: u.ProfilePic,
	}
}

type UserWithoutPassword struct {
	ID         int     `json:"id"`
	Username   string  `json:"username"`
	Email      string  `json:"email,omitempty"`
	City       *string `json:"city,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	ProfilePic string  `json:"profile_pic,omitempty"`
}

func (u *UserWithoutPassword) SetProfilePic() {
	u.ProfilePic = fmt.Sprintf("https://bookers-images.hb.kz-ast.vkcs.cloud/users/%d.png", u.ID)
}

type UserWithCounts struct {
	*User
	BooksCount int `json:"books_count"`
	ShareCount int `json:"share_count"`
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
