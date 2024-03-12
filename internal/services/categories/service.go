package categories

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/repositories/categories"
)

type Service interface {
	ListCategories(ctx context.Context) ([]*models.Category, error)
}

type service struct {
	categoriesRepo categories.Repository
}

func NewService(categoriesRepo categories.Repository) Service {
	return &service{
		categoriesRepo: categoriesRepo,
	}
}

func (s *service) ListCategories(ctx context.Context) ([]*models.Category, error) {
	return s.categoriesRepo.ListCategories(ctx)
}
