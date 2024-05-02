package personal

import (
	"github.com/labstack/echo/v4"
)

type Controller struct {
	chat chan string
}

func NewController(chat chan string) *Controller {
	return &Controller{
		chat: chat,
	}
}

func (r *Controller) Init(router *echo.Group) {
	router.POST("/webhook", r.receiveMessage)
	router.POST("/send", r.sendMessage)
}
