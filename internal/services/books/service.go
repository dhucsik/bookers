package books

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/repositories/authors"
	"github.com/dhucsik/bookers/internal/repositories/books"
	"github.com/dhucsik/bookers/internal/repositories/categories"
	"github.com/dhucsik/bookers/internal/repositories/users"
	"github.com/samber/lo"
)

type Service interface {
	CreateBook(ctx context.Context, book *models.Book, authorIDs, categoryIDs []int) error
	GetBookByID(ctx context.Context, id int) (*models.BookWithFields, error)
	ListBooks(ctx context.Context, search string, limit, offset int) ([]*models.BookWithFields, error)
	SetRating(ctx context.Context, rating *models.BookRating) error
	ListComments(ctx context.Context, bookID int) ([]*models.BookComment, error)
	AddComment(ctx context.Context, comment *models.BookComment) error
	UpdateComment(ctx context.Context, comment *models.BookComment) error
	DeleteComment(ctx context.Context, commentID, userID int) error
	UploadStockBook(ctx context.Context, book *models.UploadStockBook) (string, error)
	ApproveRequest(ctx context.Context, userID, id int) error
	SenderAccepted(ctx context.Context, userID, id int) error
	ReceiverRequested(ctx context.Context, stockBookID, userID, id int) error
	CancelRequest(ctx context.Context, userID, id int) error
	CreateRequest(ctx context.Context, userID, stockBookID int) error
	GetRequest(ctx context.Context, id int) (*models.RequestWithFields, error)
	GetRequests(ctx context.Context, userID int) ([]*models.RequestWithFields, error)
	GetStockBooks(ctx context.Context, userID int) ([]*models.StockBookWithFields, error)
}

type service struct {
	s3Client     *s3.S3
	bucket       string
	endpoint     string
	userRepo     users.Repository
	bookRepo     books.Repository
	authorRepo   authors.Repository
	categoryRepo categories.Repository
}

func NewService(
	bookRepo books.Repository,
	authorRepo authors.Repository,
	categoryRepo categories.Repository,
	endpoint string,
	bucket string,
	accessKey string,
	secretKey string,
) (Service, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String(endpoint),
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return &service{
		s3Client:     s3.New(sess),
		bucket:       bucket,
		endpoint:     endpoint,
		bookRepo:     bookRepo,
		authorRepo:   authorRepo,
		categoryRepo: categoryRepo,
	}, nil
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

func (s *service) ListBooks(ctx context.Context, search string, limit, offset int) ([]*models.BookWithFields, error) {
	books, err := s.bookRepo.ListBooks(ctx, search, limit, offset)
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

func (s *service) AddComment(ctx context.Context, comment *models.BookComment) error {
	return s.bookRepo.InsertComment(ctx, comment.BookID, comment.UserID, comment.Comment)
}

func (s *service) UpdateComment(ctx context.Context, comment *models.BookComment) error {
	com, err := s.bookRepo.GetComment(ctx, comment.ID)
	if err != nil {
		return err
	}

	if com.UserID != comment.UserID {
		return errors.ErrForbidden
	}

	return s.bookRepo.UpdateComment(ctx, comment.ID, comment.Comment)
}

func (s *service) SetRating(ctx context.Context, rating *models.BookRating) error {
	return s.bookRepo.SetRating(ctx, rating.BookID, rating.UserID, rating.Rating)
}

func (s *service) DeleteComment(ctx context.Context, commentID, userID int) error {
	// TODO: check if the user is the owner of the comment
	comment, err := s.bookRepo.GetComment(ctx, commentID)
	if err != nil {
		return err
	}

	if comment.UserID != userID {
		return errors.ErrForbidden
	}

	return s.bookRepo.DeleteComment(ctx, commentID)
}

func (s *service) ListComments(ctx context.Context, bookID int) ([]*models.BookComment, error) {
	return s.bookRepo.ListComments(ctx, bookID)
}
