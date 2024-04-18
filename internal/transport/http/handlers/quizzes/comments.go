package quizzes

import (
	"net/http"
	"strconv"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// addCommentHandler godoc
// @Summary Add comment
// @Description Add comment
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "quiz ID"
// @Param request body addCommentRequest true "request"
// @Success 200 {object} createResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/{id}/comments [post]
func (c *Controller) addCommentHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	quizIDStr := ctx.Param("id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req addCommentRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	id, err := c.quizService.AddComment(ctx.Request().Context(), req.convert(session.UserID, quizID))
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, createResponse{
		Response: response.NewResponse(),
		Result:   createResp{ID: id},
	})
}

// updateCommentHandler godoc
// @Summary Update comment request
// @Description Update comment request
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "comment ID"
// @Param request body updateCommentRequest true "request"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/comments/{id} [put]
func (c *Controller) updateCommentHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	commentIDStr := ctx.Param("id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req updateCommentRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.quizService.UpdateComment(ctx.Request().Context(), req.convert(commentID, session.UserID))
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// deleteCommentHandler godoc
// @Summary Delete comment
// @Description Delete comment
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "comment ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/comments/{id} [delete]
func (c *Controller) deleteCommentHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	commentIDStr := ctx.Param("id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.quizService.DeleteComment(ctx.Request().Context(), commentID, session.UserID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}

// listCommentsHandler godoc
// @Summary List comments
// @Description List comments
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "quiz ID"
// @Success 200 {object} listCommentsResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/{id}/comments [get]
func (c *Controller) listCommentsHandler(ctx echo.Context) error {
	quizIDStr := ctx.Param("id")
	quizID, err := strconv.Atoi(quizIDStr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	comments, err := c.quizService.ListComments(ctx.Request().Context(), quizID)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, listCommentsResponse{
		Response: response.NewResponse(),
		Result:   comments,
	})
}

// setRatingHandler godoc
// @Summary Set rating
// @Description Set rating
// @Tags quizzes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param id path int true "quiz ID"
// @Param request body setRatingRequest true "request"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /quizzes/{id}/rating [post]
func (c *Controller) setRatingHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	quizIDstr := ctx.Param("id")
	quizID, err := strconv.Atoi(quizIDstr)
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	var req setRatingRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	err = c.quizService.SetRating(ctx.Request().Context(), req.convert(quizID, session.UserID))
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, response.NewResponse())
}
