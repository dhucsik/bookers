package users

import (
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
)

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
	response.Response
	Result createUserResp `json:"result"`
}

type createUserResp struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newCreateUserResponse(user *models.User) createUserResp {
	return createUserResp{
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
	response.Response
	Result getUserByIDResp `json:"result"`
}

type getUserByIDResp struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Role     string  `json:"role"`
	City     *string `json:"city,omitempty"`
}

func newGetUserByIDResponse(user *models.User) getUserByIDResp {
	return getUserByIDResp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		City:     user.City,
	}
}
