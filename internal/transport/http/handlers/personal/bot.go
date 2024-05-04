package personal

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/dhucsik/bookers/internal/util/response"
	"github.com/labstack/echo/v4"
)

// sendMessage godoc
// @Summary Send message
// @Description Send message
// @Tags bot
// @Accept json
// @Produce json
// @Param message body sendMessageRequest true "request"
// @Success 200 {object} sendMessageResponse "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /send [post]
func (c *Controller) sendMessage(ctx echo.Context) error {
	var req sendMessageRequest
	if err := ctx.Bind(&req); err != nil {
		return response.NewBadRequest(ctx, err)
	}

	sendURL := "https://api.telegram.org/bot7136868845:AAGARWDVHtWLJQpctFjbEUQo3ISjKtGpu8M/sendMessage"

	_, err := http.PostForm(sendURL, url.Values{
		"chat_id": {strconv.FormatInt(-1002123116824, 10)},
		"text":    {req.Message},
	})
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	var resp string
	select {
	case resp = <-c.chat:
		return ctx.JSON(200, sendMessageResponse{
			Response: response.NewResponse(),
			Result:   resp,
		})
	case <-time.After(30 * time.Second):
		return ctx.JSON(200, sendMessageResponse{
			Response: response.NewResponse(),
			Result:   "timeout",
		})
	}
}

type sendMessageRequest struct {
	Message string `json:"message"`
}

type sendMessageResponse struct {
	response.Response
	Result string `json:"result"`
}
