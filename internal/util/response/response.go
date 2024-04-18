package response

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	InnerCode  string `json:"inner_code"`
}

func NewResponse() Response {
	return Response{
		Message:    "success",
		StatusCode: http.StatusOK,
		InnerCode:  "200000",
	}
}

func NewErrorResponse(ctx echo.Context, err error) error {
	code := http.StatusInternalServerError
	message := err.Error()
	innerCode := "500000"

	apiErr, ok := err.(*errors.ApiError)
	if ok {
		code = apiErr.Code
		message = apiErr.Message
		innerCode = apiErr.InnerCode
	}

	return ctx.JSON(code, &Response{
		Message:    message,
		StatusCode: code,
		InnerCode:  innerCode,
	})
}

func NewBadRequest(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusBadRequest, &Response{
		Message:    err.Error(),
		StatusCode: http.StatusBadRequest,
		InnerCode:  "400000",
	})
}
