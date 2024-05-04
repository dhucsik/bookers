package users

import (
	"net/http"

	"github.com/dhucsik/bookers/internal/errors"
	"github.com/dhucsik/bookers/internal/models"
	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// uploadProfilePicHandler godoc
// @Summary Upload profile picture
// @Description Upload profile picture
// @Tags users
// @Accept mpfd
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Authorization"
// @Param image formData file true "image"
// @Success 200 {object} uploadProfilePicResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /users/profile/image [post]
func (c *Controller) uploadProfilePicHandler(ctx echo.Context) error {
	session, ok := models.GetSession(ctx.Request().Context())
	if !ok {
		return response.NewErrorResponse(ctx, errors.ErrInvalidJWTToken)
	}

	img, err := ctx.FormFile("image")
	if err != nil {
		return response.NewBadRequest(ctx, err)
	}

	if img.Header.Get("Content-Type") != "image/png" {
		return response.NewBadRequest(ctx, errors.ErrInvalidImageFormat)
	}

	imageURL, err := c.bookService.SetProfilePic(ctx.Request().Context(), session.UserID, img)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, uploadProfilePicResponse{
		Response: response.NewResponse(),
		Result:   imageURL,
	})
}
