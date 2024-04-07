package books

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/models"
	"github.com/labstack/echo/v4"
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
// @Success 201 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/{id}/request [post]
func (c *Controller) createRequestHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	stockBookIDStr := ctx.Param("id")
	stockBookID, err := strconv.Atoi(stockBookIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	err = c.bookService.CreateRequest(ctx.Request().Context(), session.UserID, stockBookID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, nil)
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
// @Success 200 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/request/{id}/cancel [put]
func (c *Controller) cancelRequestHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	err = c.bookService.CancelRequest(ctx.Request().Context(), session.UserID, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, nil)
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
// @Success 200 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/request/{id}/received [put]
func (c *Controller) requestReceivedHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	var req requestReceived
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	err = c.bookService.ReceiverRequested(ctx.Request().Context(), req.BookID, session.UserID, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, nil)
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
// @Success 200 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/request/{id}/sender_accepted [put]
func (c *Controller) senderAcceptedHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	err = c.bookService.SenderAccepted(ctx.Request().Context(), session.UserID, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, nil)
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
// @Success 200 {object} nil "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/request/{id}/approve [put]
func (c *Controller) approveRequest(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	err = c.bookService.ApproveRequest(ctx.Request().Context(), session.UserID, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, nil)
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
// @Success 200 {object} models.RequestWithFields "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/request/{id} [get]
func (c *Controller) getRequestHandler(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	}

	request, err := c.bookService.GetRequest(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, request)
}

// getRequestsHandler godoc
// @Summary Get requests
// @Description Get requests
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {array} models.RequestWithFields "Success"
// @Failure 400 {object} errorResponse "Bad request"
// @Failure 401 {object} errorResponse "Unauthorized"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /books/requests [get]
func (c *Controller) getRequestsHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, newErrorResponse("session not found"))
	}

	requests, err := c.bookService.GetRequests(ctx.Request().Context(), session.UserID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, requests)
}
