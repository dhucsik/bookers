package authors

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/repositories/authors"
)

type Service interface {
	ListAuthors(ctx context.Context, search string, limit, offset int) ([]*models.Author, int, error)
	CreateAuthor(ctx context.Context, author *models.Author) (int, error)
	DeleteAuthor(ctx context.Context, id int) error
}

type service struct {
	authorsRepository authors.Repository
}

func NewService(authorsRepository authors.Repository) Service {
	return &service{
		authorsRepository: authorsRepository,
	}
}

func (s *service) ListAuthors(ctx context.Context, search string, limit, offset int) ([]*models.Author, int, error) {
	return s.authorsRepository.ListAuthors(ctx, search, limit, offset)
}

func (s *service) CreateAuthor(ctx context.Context, author *models.Author) (int, error) {
	return s.authorsRepository.CreateAuthor(ctx, author)
}

func (s *service) DeleteAuthor(ctx context.Context, id int) error {
	return s.authorsRepository.DeleteAuthor(ctx, id)
}
