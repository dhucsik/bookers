package quizzes

import "github.com/dhucsik/bookers/internal/models"

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(message string) errorResponse {
	return errorResponse{
		Message: message,
	}
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
