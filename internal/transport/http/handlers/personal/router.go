package personal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	chat   chan string
	client *http.Client
	apikey string
}

func NewController(chat chan string, apikey string) *Controller {
	return &Controller{
		apikey: apikey,
		chat:   chat,
		client: http.DefaultClient,
	}
}

func (r *Controller) Init(router *echo.Group) {
	router.POST("/webhook", r.receiveMessage)
	router.POST("/send", r.sendMessage)
}
