package personal

import (
	"github.com/labstack/echo/v4"
)

func (c *Controller) receiveMessage(ctx echo.Context) error {
	var update Update
	if err := ctx.Bind(&update); err != nil {
		return err
	}

	if update.Message == nil {
		return ctx.JSON(200, "ok")
	}

	if update.Message.Chat == nil {
		return ctx.JSON(200, "ok")
	}

	return ctx.JSON(200, "ok")
}

type Update struct {
	Message *Message `json:"message,omitempty"`
}

type Message struct {
	ID   int    `json:"message_id,omitempty"`
	Chat *Chat  `json:"chat,omitempty"`
	Text string `json:"text,omitempty"`
}

type Chat struct {
	ID int64 `json:"id,omitempty"`
}
