package authors

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/repositories/authors"
)

type Service interface {
	ListAuthors(ctx context.Context) ([]*models.Author, error)
}

type service struct {
	authorsRepository authors.Repository
}

func NewService(authorsRepository authors.Repository) Service {
	return &service{
		authorsRepository: authorsRepository,
	}
}

func (s *service) ListAuthors(ctx context.Context) ([]*models.Author, error) {
	return s.authorsRepository.ListAuthors(ctx)
}
