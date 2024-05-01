package books

import (
	"time"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/samber/lo"
)

type createResponse struct {
	response.Response
	Result createResp `json:"result"`
}

type createResp struct {
	ID int `json:"id"`
}

func newCreateResp(id int) createResp {
	return createResp{
		ID: id,
	}
}

type uploadStockBookResponse struct {
	response.Response
	Result uploadStockBookResp `json:"result"`
}

type uploadStockBookResp struct {
	ID       int    `json:"id"`
	ImageURL string `json:"image_url"`
}

func newUploadStockBookResp(id int, imageURL string) uploadStockBookResp {
	return uploadStockBookResp{
		ID:       id,
		ImageURL: imageURL,
	}
}

type updateStockResponse struct {
	response.Response
	Result string `json:"result"`
}

type getStockBookResponse struct {
	response.Response
	Result []*models.StockBookWithFields `json:"result"`
}

type categoryResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newCategoryResp(category *models.Category) categoryResp {
	return categoryResp{
		ID:   category.ID,
		Name: category.Name,
	}
}

type authorResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newAuthorResp(author *models.Author) authorResp {
	return authorResp{
		ID:   author.ID,
		Name: author.Name,
	}
}

type getBookResponse struct {
	response.Response
	Result bookResponse `json:"result"`
}

type bookResponse struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	PubDate     string         `json:"pub_date"`
	Edition     string         `json:"edition"`
	Language    string         `json:"language"`
	Rating      float64        `json:"rating"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	Authors     []authorResp   `json:"authors"`
	Categories  []categoryResp `json:"categories"`
}

func newBookResp(book *models.BookWithFields) bookResponse {
	authors := lo.Map(book.Authors, func(author *models.Author, _ int) authorResp {
		return newAuthorResp(author)
	})

	categories := lo.Map(book.Categories, func(category *models.Category, _ int) categoryResp {
		return newCategoryResp(category)
	})

	return bookResponse{
		ID:          book.ID,
		Title:       book.Title,
		PubDate:     book.PubDate.Format(time.DateOnly),
		Edition:     book.Edition,
		Language:    book.Language,
		Rating:      book.Rating,
		Description: book.Description,
		Image:       book.Image,
		Authors:     authors,
		Categories:  categories,
	}
}

type listBooksResponse struct {
	response.Response
	Result listBooksResp `json:"result"`
}

type listBooksResp struct {
	Books []bookResponse `json:"books"`
	Total int            `json:"total"`
}

func newListBooksResponse(books []*models.BookWithFields, totalCount int) listBooksResponse {
	out := lo.Map(books, func(book *models.BookWithFields, _ int) bookResponse {
		return newBookResp(book)
	})

	return listBooksResponse{
		Response: response.NewResponse(),
		Result:   listBooksResp{Books: out, Total: totalCount},
	}
}

type listQuizzesResponse struct {
	response.Response
	Result []*models.QuizWithBase `json:"result"`
}

type listCommentsResponse struct {
	response.Response
	Result []*models.BookComment `json:"comments"`
}

type getRequestResponse struct {
	response.Response
	Result *models.RequestWithFields `json:"request"`
}

type listRequestsResponse struct {
	response.Response
	Result []*models.RequestWithFields `json:"requests"`
}

type addCommentRequest struct {
	Comment string `json:"comment"`
}

func (r addCommentRequest) convert(bookID, userID int) *models.BookComment {
	return &models.BookComment{
		BookID:  bookID,
		UserID:  userID,
		Comment: r.Comment,
	}
}

type setRatingRequest struct {
	Rating int `json:"rating"`
}

func (r setRatingRequest) convert(bookID, userID int) *models.BookRating {
	return &models.BookRating{
		BookID: bookID,
		UserID: userID,
		Rating: r.Rating,
	}
}

type updateCommentRequest struct {
	Comment string `json:"comment"`
}

func (r updateCommentRequest) convert(commentID, userID int) *models.BookComment {
	return &models.BookComment{
		ID:      commentID,
		UserID:  userID,
		Comment: r.Comment,
	}
}

type createQuizRequest struct {
	Title string `json:"title"`
}

func (r *createQuizRequest) convert(bookID, userID int) *models.Quiz {
	return &models.Quiz{
		Title:  r.Title,
		BookID: bookID,
		UserID: userID,
	}
}

type requestReceived struct {
	BookID int `json:"book_id"`
}

type searchBooksRequest struct {
	Query       string `json:"query"`
	CategoryIDs []int  `json:"category_ids"`
	AuthorIDs   []int  `json:"author_ids"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}

func (r searchBooksRequest) convert() *models.SearchParams {
	return &models.SearchParams{
		Search:     r.Query,
		Limit:      r.Limit,
		Offset:     r.Offset,
		Categories: r.CategoryIDs,
		Authors:    r.AuthorIDs,
	}
}
