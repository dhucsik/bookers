package quizzes

import (
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

type listQuizzesResponse struct {
	response.Response
	Result listQuizzesResp `json:"result"`
}

type listQuizzesResp struct {
	Quizzes    []*models.QuizWithBase `json:"quizzes"`
	TotalCount int                    `json:"total_count"`
}

type listCommentsResponse struct {
	response.Response
	Result []*models.QuizComment `json:"result"`
}

type getQuizResponse struct {
	response.Response
	Result *models.QuizWithFields `json:"result"`
}

type checkQuizResponse struct {
	response.Response
	Result *models.QuizWithQuestionResults
}

type getQuizResultsResponse struct {
	response.Response
	Result []*models.QuizResultWithFields
}

type getQuizResultResponse struct {
	response.Response
	Result *models.QuizQuestionWithFields
}

type updateQuizRequest struct {
	Title string `json:"title"`
}

type addQuestionRequest struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

func (r *addQuestionRequest) convert(quizID int) *models.Question {
	return &models.Question{
		QuizID:   quizID,
		Question: r.Question,
		Options:  r.Options,
		Answer:   r.Answer,
	}
}

type updateQuestionRequest struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

func (r *updateQuestionRequest) convert(questionID int) *models.Question {
	return &models.Question{
		ID:       questionID,
		Question: r.Question,
		Options:  r.Options,
		Answer:   r.Answer,
	}
}

type addCommentRequest struct {
	Comment string `json:"comment"`
}

func (r *addCommentRequest) convert(userID, quizID int) *models.QuizComment {
	return &models.QuizComment{
		UserID:  userID,
		QuizID:  quizID,
		Comment: r.Comment,
	}
}

type updateCommentRequest struct {
	Comment string `json:"comment"`
}

func (r *updateCommentRequest) convert(commentID, userID int) *models.QuizComment {
	return &models.QuizComment{
		ID:      commentID,
		UserID:  userID,
		Comment: r.Comment,
	}
}

type setRatingRequest struct {
	Rating int `json:"rating"`
}

func (r *setRatingRequest) convert(quizID, userID int) *models.QuizRating {
	return &models.QuizRating{
		UserID: userID,
		QuizID: quizID,
		Rating: r.Rating,
	}
}

type checkQuizRequest struct {
	UserAnswers []checkQuestionRequest `json:"user_answers"`
}

type checkQuestionRequest struct {
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
}

func (r *checkQuestionRequest) convert() *models.UserAnswer {
	return &models.UserAnswer{
		QuestionID: r.QuestionID,
		UserAnswer: r.Answer,
	}
}

func (r *checkQuizRequest) convert() []*models.UserAnswer {
	return lo.Map(r.UserAnswers, func(item checkQuestionRequest, _ int) *models.UserAnswer {
		return item.convert()
	})
}
