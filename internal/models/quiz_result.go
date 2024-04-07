package models

type QuizResult struct {
	ID      int `json:"id"`
	QuizID  int `json:"quiz_id"`
	UserID  int `json:"user_id"`
	Correct int `json:"correct"`
	Total   int `json:"total"`
}

type QuestionResult struct {
	ID           int    `json:"id"`
	QuizResultID int    `json:"quiz_result_id,omitempty"`
	QuestionID   int    `json:"question_id"`
	UserAnswer   string `json:"user_answer"`
	IsCorrect    bool   `json:"is_correct"`
}

type QuizWithQuestionResults struct {
	*QuizResult
	Results []*QuestionResult `json:"results"`
}

type UserAnswer struct {
	QuestionID int    `json:"question_id"`
	UserAnswer string `json:"user_answer"`
}

type QuizResultWithFields struct {
	*QuizResult
	Quiz *Quiz `json:"quiz"`
	Book *Book `json:"book"`
}

type QuizQuestionWithFields struct {
	*QuizWithQuestionResults
	Quiz *Quiz `json:"quiz"`
	Book *Book `json:"book"`
}
