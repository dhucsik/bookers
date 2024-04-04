package auth

import "github.com/dhucsik/bookers/internal/models"

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponse struct {
	AccessToken  string                      `json:"access_token"`
	RefreshToken string                      `json:"refresh_token"`
	User         *models.UserWithoutPassword `json:"user"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(message string) errorResponse {
	return errorResponse{
		Message: message,
	}
}

type refreshRequest struct {
	Token string `json:"token"`
}
