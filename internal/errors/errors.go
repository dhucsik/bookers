package errors

import "errors"

var (
	ErrInvalidJWTToken   = errors.New("invalid jwt token")
	ErrEmptyAuthHeader   = errors.New("empty auth header")
	ErrTokenExpired      = errors.New("token is expired")
	ErrUnexpectedRefresh = errors.New("unexpected refresh token")
	ErrNotRefreshToken   = errors.New("not a refresh token")
	ErrForbidden         = errors.New("forbidden")
)
