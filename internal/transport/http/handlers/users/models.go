package users

import "github.com/dhucsik/bookers/internal/models"

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(message string) errorResponse {
	return errorResponse{
		Message: message,
	}
}

type createUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r createUserRequest) convert() *models.User {
	return &models.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

type createUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func newCreateUserResponse(user *models.User) createUserResponse {
	return createUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

type setCityRequest struct {
	City string `json:"city"`
}

type getUserByIDResponse struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	City     *string `json:"city,omitempty"`
}

func newGetUserByIDResponse(user *models.User) getUserByIDResponse {
	return getUserByIDResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		City:     user.City,
	}
}
