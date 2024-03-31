package personal

import (
	"github.com/labstack/echo/v4"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (r *Controller) Init(router *echo.Group) {
	router.GET("/dhucsik", func(c echo.Context) error {
		return c.HTML(200, `
			<a href="/api/v1/dhucs/love"><h1>from dhucs to ayau</h1></a>		
		`)
	})

	router.GET("/alii", func(c echo.Context) error {
		return c.HTML(200, `
			<h1>ali chert</h1>		
		`)
	})

	router.GET("/anu", func(c echo.Context) error {
		return c.HTML(200, `
			<h1>anu chert</h1>
		`)
	})

	router.GET("/erema", func(c echo.Context) error {
		return c.HTML(200, `
			<h1>erema chert</h1>
		`)
	})

	router.File("/dhucs/love", "files/index.html")
}
