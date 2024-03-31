package books

import (
	"context"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/repositories/authors"
	"github.com/dhucsik/bookers/internal/repositories/books"
	"github.com/dhucsik/bookers/internal/repositories/categories"
	"github.com/samber/lo"
)

type Service interface {
	CreateBook(ctx context.Context, book *models.Book, authorIDs, categoryIDs []int) error
	GetBookByID(ctx context.Context, id int) (*models.BookWithFields, error)
	ListBooks(ctx context.Context, search string, limit, offset int) ([]*models.BookWithFields, error)
}

type service struct {
	bookRepo     books.Repository
	authorRepo   authors.Repository
	categoryRepo categories.Repository
}

func NewService(
	bookRepo books.Repository,
	authorRepo authors.Repository,
	categoryRepo categories.Repository,
) Service {
	return &service{
		bookRepo:     bookRepo,
		authorRepo:   authorRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *service) CreateBook(ctx context.Context, book *models.Book, authorIDs, categoryIDs []int) error {
	return s.bookRepo.CreateBook(ctx, book, authorIDs, categoryIDs)
}

func (s *service) GetBookByID(ctx context.Context, id int) (*models.BookWithFields, error) {
	book, err := s.bookRepo.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}

	authors, err := s.authorRepo.GetByBookID(ctx, id)
	if err != nil {
		return nil, err
	}

	categories, err := s.categoryRepo.GetByBookID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.BookWithFields{
		Book:       book,
		Authors:    authors,
		Categories: categories,
	}, nil
}

func (s *service) ListBooks(ctx context.Context, search string, offset, limit int) ([]*models.BookWithFields, error) {
	books, err := s.bookRepo.ListBooks(ctx, search, offset, limit)
	if err != nil {
		return nil, err
	}

	ids := lo.Map(books, func(book *models.Book, _ int) int {
		return book.ID
	})

	authors, err := s.authorRepo.GetByBookIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	categories, err := s.categoryRepo.GetByBookIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	out := lo.Map(books, func(book *models.Book, _ int) *models.BookWithFields {
		return &models.BookWithFields{
			Book:       book,
			Authors:    authors[book.ID],
			Categories: categories[book.ID],
		}
	})

	return out, nil
}
