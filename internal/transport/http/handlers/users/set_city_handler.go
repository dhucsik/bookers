package users

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// setCity godoc
// @Summary Set city
// @Description Set city
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "User ID"
// @Param request body setCityRequest true "request"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/{id}/city [patch]
func (c *Controller) setCity(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req setCityRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.usersService.SetCity(ctx.Request().Context(), id, req.City)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// updateUsername godoc
// @Summary Update username
// @Description Update username
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param request body updateUsernameRequest true "request"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/username [patch]
func (c *Controller) updateUsername(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	var req updateUsernameRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err := c.usersService.UpdateUsername(ctx.Request().Context(), session.UserID, req.Username)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// updatePassword godoc
// @Summary Update password
// @Description Update password
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param request body updatePasswordRequest true "request"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/password [patch]
func (c *Controller) updatePassword(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	var req updatePasswordRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err := c.usersService.UpdatePassword(ctx.Request().Context(), session.UserID, req.Password)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}
