package auth

import (
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
)

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponse struct {
	response.Response
	Result authResp `json:"result"`
}

type authResp struct {
	AccessToken  string                      `json:"access_token"`
	RefreshToken string                      `json:"refresh_token"`
	User         *models.UserWithoutPassword `json:"user"`
}

type refreshRequest struct {
	Token string `json:"token"`
}
