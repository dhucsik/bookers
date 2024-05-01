package errors

import (
	"net/http"
)

var (
	ErrNotRefreshToken    = NewApiError("not a refresh token", http.StatusBadRequest, "400001")
	ErrWrongStatus        = NewApiError("wrong status", http.StatusBadRequest, "400002")
	ErrInvalidImageFormat = NewApiError("invalid image format, only png", http.StatusBadRequest, "400003")

	ErrInvalidJWTToken     = NewApiError("invalid jwt token", http.StatusUnauthorized, "401001")
	ErrEmptyAuthHeader     = NewApiError("empty auth header", http.StatusUnauthorized, "401002")
	ErrTokenExpired        = NewApiError("token is expired", http.StatusUnauthorized, "401003")
	ErrUnexpectedRefresh   = NewApiError("unexpected refresh token", http.StatusUnauthorized, "401004")
	ErrInvalidPassword     = NewApiError("invalid password", http.StatusUnauthorized, "401005")
	ErrUsernameExists      = NewApiError("username exists", http.StatusUnauthorized, "401006")
	ErrFriendRequestExists = NewApiError("friend request exists", http.StatusUnauthorized, "401007")

	ErrForbiddenForRole = NewApiError("forbidden for role", http.StatusForbidden, "403001")
	ErrForbiddenForUser = NewApiError("forbidden for user", http.StatusForbidden, "403002")

	ErrUserNotFound    = NewApiError("user not found", http.StatusNotFound, "404001")
	ErrQuizNotFound    = NewApiError("quiz not found", http.StatusNotFound, "404002")
	ErrBookNotFound    = NewApiError("book not found", http.StatusNotFound, "404003")
	ErrResultNotFound  = NewApiError("result not found", http.StatusNotFound, "404004")
	ErrRequestNotFound = NewApiError("request not found", http.StatusNotFound, "404005")
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
