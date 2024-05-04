package users

import (
	"database/sql"

	"github.com/dhucsik/bookers/internal/models"
)

type userModel struct {
	ID       int            `db:"id"`
	Username string         `db:"username"`
	Email    string         `db:"email"`
	Password string         `db:"password"`
	Role     string         `db:"role"`
	City     sql.NullString `db:"city"`
}

func (m *userModel) convert() *models.User {
	var city string
	if m.City.Valid {
		city = m.City.String
	}

	out := &models.User{
		ID:       m.ID,
		Username: m.Username,
		Email:    m.Email,
		Password: m.Password,
		Role:     m.Role,
		City:     &city,
	}

	out.SetProfilePic()
	return out
}

func convertUser(user *models.User) *userModel {
	city := sql.NullString{}
	if user.City != nil {
		city = sql.NullString{
			String: *user.City,
			Valid:  true,
		}
	}

	return &userModel{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		City:     city,
	}
}
