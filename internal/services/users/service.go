package users

import (
	"context"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/repositories/users"
)

type Service interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	SetCity(ctx context.Context, id int, city string) error
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	DeleteUser(ctx context.Context, userID int) error
}

type service struct {
	userRepo users.Repository
}

func NewService(userRepo users.Repository) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.ErrUsernameExists
	}

	return s.userRepo.CreateUser(ctx, user)
}

func (s *service) SetCity(ctx context.Context, userID int, city string) error {
	session, ok := models.GetSession(ctx)
	if !ok {
		return errors.ErrInvalidJWTToken
	}

	if session.UserID != userID && session.Role != "admin" {
		return errors.ErrForbiddenForUser
	}

	return s.userRepo.SetCity(ctx, session.UserID, city)
}

func (s *service) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	return s.userRepo.GetUserByID(ctx, userID)
}

func (s *service) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.userRepo.GetUserByUsername(ctx, username)
}

func (s *service) DeleteUser(ctx context.Context, userID int) error {
	session, ok := models.GetSession(ctx)
	if !ok {
		return errors.ErrInvalidJWTToken
	}

	if session.UserID != userID && session.Role != "admin" {
		return errors.ErrForbiddenForRole
	}

	return s.userRepo.DeleteUser(ctx, userID)
}
