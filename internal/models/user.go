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
