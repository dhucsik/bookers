package errors

import (
	"net/http"
)

var (
	ErrNotRefreshToken    = NewApiError("not a refresh token", http.StatusBadRequest, "400001")
	ErrWrongStatus        = NewApiError("wrong status", http.StatusBadRequest, "400002")
	ErrInvalidImageFormat = NewApiError("invalid image format, only png", http.StatusBadRequest, "400003")

	ErrInvalidJWTToken   = NewApiError("invalid jwt token", http.StatusUnauthorized, "401001")
	ErrEmptyAuthHeader   = NewApiError("empty auth header", http.StatusUnauthorized, "401002")
	ErrTokenExpired      = NewApiError("token is expired", http.StatusUnauthorized, "401003")
	ErrUnexpectedRefresh = NewApiError("unexpected refresh token", http.StatusUnauthorized, "401004")

	ErrForbiddenForRole = NewApiError("forbidden for role", http.StatusForbidden, "403001")
	ErrForbiddenForUser = NewApiError("forbidden for user", http.StatusForbidden, "403002")
)

type ApiError struct {
	Message   string `json:"message"`
	Code      int    `json:"code"`
	InnerCode string `json:"inner_code"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(message string, code int, innerCode string) *ApiError {
	return &ApiError{
		Message:   message,
		Code:      code,
		InnerCode: innerCode,
	}
}
