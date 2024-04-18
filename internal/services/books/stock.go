package books

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/png"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/nickalie/go-webpbin"
	"github.com/samber/lo"
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

	img, err := png.Decode(bytes.NewReader(file.Bytes()))
	if err != nil {
		return 0, "", err
	}

	webpImg, err := s.ConvertToWebp(img, book.Image.Header.Get("Content-Type"))
	if err != nil {
		return 0, "", err
	}

	imageURL, err := s.UploadImage(ctx, webpImg, fmt.Sprintf("stock/books/%d.webp", id))
	if err != nil {
		return 0, "", err
	}

	return id, imageURL, nil
}

func (s *service) UploadImage(ctx context.Context, body []byte, filename string) (string, error) {
	object := s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(body),
		ContentType: aws.String("image/webp"),
	}

	out, err := s.s3Client.PutObjectWithContext(ctx, &object)
	if err != nil {
		return "", err
	}

	return out.String(), nil
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
		}
	}), nil
}
