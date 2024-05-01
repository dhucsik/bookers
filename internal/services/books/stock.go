package books

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/nickalie/go-webpbin"
	"github.com/samber/lo"
	"github.com/yusukebe/go-pngquant"
)

func (s *service) UploadStockBook(ctx context.Context, book *models.UploadStockBook) (int, string, error) {
	id, err := s.bookRepo.UploadStockBook(ctx, &models.StockBook{
		UserID: book.UserID,
		BookID: book.BookID,
	})
	if err != nil {
		return 0, "", err
	}

	src, err := book.Image.Open()
	if err != nil {
		return 0, "", err
	}
	defer src.Close()

	var file bytes.Buffer
	_, err = io.Copy(&file, src)
	if err != nil {
		return 0, "", err
	}

	compressed, err := pngquant.CompressBytes(file.Bytes(), "5")
	if err != nil {
		return 0, "", err
	}

	imageURL, err := s.UploadImage(ctx, compressed, fmt.Sprintf("stock/books/%d.png", id))
	if err != nil {
		return 0, "", err
	}

	return id, imageURL, nil
}

func (s *service) UpdateImage(ctx context.Context, userID, stockID int, image *multipart.FileHeader) (string, error) {
	stock, err := s.bookRepo.GetStockBook(ctx, stockID)
	if err != nil {
		return "", err
	}

	if stock.UserID != userID {
		return "", errors.ErrForbiddenForUser
	}

	src, err := image.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	var file bytes.Buffer
	_, err = io.Copy(&file, src)
	if err != nil {
		return "", err
	}

	compressed, err := pngquant.CompressBytes(file.Bytes(), "5")
	if err != nil {
		return "", err
	}

	imageURL, err := s.UploadImage(ctx, compressed, fmt.Sprintf("stock/books/%d.png", stockID))
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func (s *service) DeleteStockBook(ctx context.Context, userID, stockID int) error {
	stock, err := s.bookRepo.GetStockBook(ctx, stockID)
	if err != nil {
		return err
	}

	if stock.UserID != userID {
		return errors.ErrForbiddenForUser
	}

	err = s.bookRepo.DeleteStockBook(ctx, stockID)
	if err != nil {
		return err
	}

	return s.DeleteImage(ctx, fmt.Sprintf("stock/books/%d.png", stockID))
}

func (s *service) UploadImage(ctx context.Context, body []byte, filename string) (string, error) {
	object := s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(body),
		ContentType: aws.String("image/webp"),
		ACL:         aws.String("public-read"),
	}

	_, err := s.s3Client.PutObjectWithContext(ctx, &object)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://bookers-images.hb.kz-ast.vkcs.cloud/%s", filename), nil
}

func (s *service) DeleteImage(ctx context.Context, filename string) error {
	object := s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
	}

	_, err := s.s3Client.DeleteObjectWithContext(ctx, &object)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) ConvertToWebp(img image.Image, imageType string) ([]byte, error) {
	webpBuffer := new(bytes.Buffer)

	if err := webpbin.Encode(webpBuffer, img); err != nil {
		return nil, err
	}

	webpData := webpBuffer.Bytes()

	return webpData, nil
}

func (s *service) GetStockBooks(ctx context.Context, userID int) ([]*models.StockBookWithFields, error) {
	stockBooks, err := s.bookRepo.GetStockBooksByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	bookIDs := lo.Map(stockBooks, func(item *models.StockBook, _ int) int {
		return item.BookID
	})

	books, err := s.bookRepo.GetBooksByIDs(ctx, bookIDs)
	if err != nil {
		return nil, err
	}

	booksMap := lo.SliceToMap(books, func(item *models.Book) (int, *models.Book) {
		return item.ID, item
	})

	return lo.Map(stockBooks, func(item *models.StockBook, _ int) *models.StockBookWithFields {
		return &models.StockBookWithFields{
			StockBook: item,
			Book:      booksMap[item.BookID],
			User:      user.ToUserWithoutPassword(),
		}
	}), nil
}

func (s *service) GetStockByBook(ctx context.Context, bookID int) ([]*models.StockBookWithFields, error) {
	stockBooks, err := s.bookRepo.GetStockByBook(ctx, bookID)
	if err != nil {
		return nil, err
	}

	book, err := s.bookRepo.GetBookByID(ctx, bookID)
	if err != nil {
		return nil, err
	}

	userIDs := lo.Map(stockBooks, func(item *models.StockBook, _ int) int {
		return item.UserID
	})

	users, err := s.userRepo.GetUsersByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	usersMap := lo.SliceToMap(users, func(item *models.User) (int, *models.User) {
		return item.ID, item
	})

	return lo.Map(stockBooks, func(item *models.StockBook, _ int) *models.StockBookWithFields {
		return &models.StockBookWithFields{
			StockBook: item,
			User:      usersMap[item.UserID].ToUserWithoutPassword(),
			Book:      book,
		}
	}), nil
}
