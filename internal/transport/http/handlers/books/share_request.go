package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

// createRequestHandler godoc
// @Summary Create request
// @Description Create request
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "stock book ID"
// @Success 201 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/{id}/request [post]
func (c *Controller) createRequestHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	stockBookIDStr := ctx.Param("id")
	stockBookID, err := strconv.Atoi(stockBookIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.bookService.CreateRequest(ctx.Request().Context(), session.UserID, stockBookID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, response.NewResponse())
}

// cancelRequestHandler godoc
// @Summary Cancel request
// @Description Cancel request
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "request ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/request/{id}/cancel [put]
func (c *Controller) cancelRequestHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.bookService.CancelRequest(ctx.Request().Context(), session.UserID, id)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// requestReceivedHandler godoc
// @Summary Request received
// @Description Request received
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "request ID"
// @Param request body requestReceived true "request"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/request/{id}/received [put]
func (c *Controller) requestReceivedHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req requestReceived
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.bookService.ReceiverRequested(ctx.Request().Context(), req.BookID, session.UserID, id)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// senderAcceptedHandler godoc
// @Summary Sender accepted
// @Description Sender accepted
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "request ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/request/{id}/sender_accepted [put]
func (c *Controller) senderAcceptedHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.bookService.SenderAccepted(ctx.Request().Context(), session.UserID, id)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// approveRequest godoc
// @Summary Approve request
// @Description Approve request
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "request ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/request/{id}/approve [put]
func (c *Controller) approveRequest(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.bookService.ApproveRequest(ctx.Request().Context(), session.UserID, id)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// getRequestHandler godoc
// @Summary Get request
// @Description Get request
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "request ID"
// @Success 200 {object} getRequestResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/request/{id} [get]
func (c *Controller) getRequestHandler(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	request, err := c.bookService.GetRequest(ctx.Request().Context(), id)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, getRequestResponse{
		Response: response.NewResponse(),
		Result:   request,
	})
}

// getRequestsHandler godoc
// @Summary Get requests
// @Description Get requests
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} listRequestsResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/request [get]
func (c *Controller) getRequestsHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	requests, err := c.bookService.GetRequests(ctx.Request().Context(), session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, listRequestsResponse{
		Response: response.NewResponse(),
		Result:   requests,
	})
}

// getApprovedRequestsHandler godoc
// @Summary Get approved requests
// @Description Get approved requests
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} listRequestsResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /books/exchanges [get]
func (c *Controller) getApprovedRequestsHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	requests, err := c.bookService.GetRequests(ctx.Request().Context(), session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	out := lo.Filter(requests, func(req *models.RequestWithFields, _ int) bool {
		return req.SenderStatus == models.StatusSenderProved && req.ReceiverStatus == models.StatusReceiverProved
	})

	return ctx.JSON(http.StatusOK, listRequestsResponse{
		Response: response.NewResponse(),
		Result:   out,
	})
}
