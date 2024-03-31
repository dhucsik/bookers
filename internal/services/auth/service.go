package auth

import (
	"context"
	"time"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/services/users"
	"github.com/dhucsik/bookers/internal/util/jwt"
)

type Service interface {
	CheckToken(ctx context.Context, token string) (*models.Session, error)
	GetAuth(ctx context.Context, username, password string) (string, string, error)
	Refresh(ctx context.Context, token string) (string, string, error)
}

type service struct {
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	userService     users.Service
}

func NewService(
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
	usersService users.Service,
) Service {
	return &service{
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		userService:     usersService,
	}
}

func (s *service) CheckToken(ctx context.Context, token string) (*models.Session, error) {
	if token == "" {
		return nil, errors.ErrEmptyAuthHeader
	}

	session, isRefresh, err := jwt.ParseJWT(token)
	if err != nil {
		return nil, err
	}

	if isRefresh {
		return nil, errors.ErrUnexpectedRefresh
	}

	if session.Exp() < time.Now().Unix() {
		return nil, errors.ErrTokenExpired
	}

	return session, nil
}

func (s *service) GetAuth(ctx context.Context, username, password string) (string, string, error) {
	user, err := s.userService.GetUserByUsername(ctx, username)
	if err != nil {
		return "", "", err
	}

	session := &models.Session{
		UserID: user.ID,
		Role:   user.Role,
	}

	accessToken, err := jwt.GenerateJWT(session, s.accessTokenTTL, false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.GenerateJWT(session, s.refreshTokenTTL, true)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *service) Refresh(ctx context.Context, token string) (string, string, error) {
	ses, isRefresh, err := jwt.ParseJWT(token)
	if err != nil {
		return "", "", err
	}

	if !isRefresh {
		return "", "", errors.ErrNotRefreshToken
	}

	session := &models.Session{
		UserID: ses.UserID,
		Role:   ses.Role,
	}

	accessToken, err := jwt.GenerateJWT(session, s.accessTokenTTL, false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.GenerateJWT(session, s.refreshTokenTTL, true)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
