package personal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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

	reqStruct := gptRequest{
		Model: "gpt-3.5-turbo",
		Messages: []gptMessage{
			{
				Role:    "system",
				Content: "You are an experienced reader assistant who helps other novice readers by answering their questions.",
			},
			{
				Role:    "user",
				Content: req.Message,
			},
		},
	}

	reqBody, err := json.Marshal(reqStruct)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	gptreq, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/davinci-codex/completions", bytes.NewReader(reqBody))
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	gptreq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apikey))
	gptreq.Header.Set("Content-Type", "application/json")

	gptresp, err := c.client.Do(gptreq)
	if err != nil {
		return response.NewErrorResponse(ctx, err)
	}
	defer gptresp.Body.Close()

	var gptres gptResponse
	if err := json.NewDecoder(gptresp.Body).Decode(&gptres); err != nil {
		return response.NewErrorResponse(ctx, err)
	}

	if len(gptres.Choises) == 0 {
		return ctx.JSON(200, sendMessageResponse{
			Response: response.NewResponse(),
			Result:   "no response",
		})
	}

	return ctx.JSON(200, sendMessageResponse{
		Response: response.NewResponse(),
		Result:   gptres.Choises[0].Message.Content,
	})
}

/*
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
}*/

type sendMessageRequest struct {
	Message string `json:"message"`
}

type sendMessageResponse struct {
	response.Response
	Result string `json:"result"`
}

type gptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type gptRequest struct {
	Model    string       `json:"model"`
	Messages []gptMessage `json:"messages"`
}

type gptResponse struct {
	Choises []choice `json:"choices"`
}

type choice struct {
	Index   int        `json:"index"`
	Message gptMessage `json:"message"`
}
