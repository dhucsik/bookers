package users

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// sendFriendRequestHandler godoc
// @Summary Send friend request
// @Description Send friend request
// @Tags friends
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "Friend ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/friends/{id}/request [post]
func (c *Controller) sendFriendRequestHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	friendIDStr := ctx.Param("id")
	friendID, err := strconv.Atoi(friendIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.usersService.CreateRequest(ctx.Request().Context(), session.UserID, friendID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// acceptFriendRequestHandler godoc
// @Summary Accept friend request
// @Description Accept friend request
// @Tags friends
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "Friend ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/friends/{id}/accept [put]
func (c *Controller) acceptFriendRequestHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	friendIDStr := ctx.Param("id")
	friendID, err := strconv.Atoi(friendIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.usersService.AcceptRequest(ctx.Request().Context(), session.UserID, friendID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// getFriendsHandler godoc
// @Summary Get friends
// @Description Get friends
// @Tags friends
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} listFriendsResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/friends [get]
func (c *Controller) getFriendsHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	friends, err := c.usersService.GetFriends(ctx.Request().Context(), session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, listFriendsResponse{
		Response: response.NewResponse(),
		Result:   newListFriendsResponse(friends),
	})
}

// getSentRequests godoc
// @Summary Get sent friend requests
// @Description Get sent friend requests
// @Tags friends
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} listFriendsResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/friends/sent [get]
func (c *Controller) getSentRequests(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	requests, err := c.usersService.GetSentRequestFriends(ctx.Request().Context(), session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, listFriendsResponse{
		Response: response.NewResponse(),
		Result:   newListFriendsResponse(requests),
	})
}

// getReceivedRequests godoc
// @Summary Get received friend requests
// @Description Get received friend requests
// @Tags friends
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Success 200 {object} listFriendsResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/friends/received [get]
func (c *Controller) getReceivedRequests(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	requests, err := c.usersService.GetReceivedRequestFriends(ctx.Request().Context(), session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, listFriendsResponse{
		Response: response.NewResponse(),
		Result:   newListFriendsResponse(requests),
	})
}
