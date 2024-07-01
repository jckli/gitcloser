package backend

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/jckli/gitcloser/backend/v1/github/pathway"
	"github.com/jckli/gitcloser/backend/v1/index"
	"github.com/valyala/fasthttp"
)

func InitRoutes(app *fiber.App, client *fasthttp.Client) {
	// v1 routes
	app.Get("/", index.IndexHandler)

	app.Get("/v1/github/pathway/:user1/:user2", func(c *fiber.Ctx) error {
		return pathway.PathwayHandler(c, client)
	})

	app.Get("/v1/github/pathway/:user1/:user2/ws", websocket.New(func(c *websocket.Conn) {
		pathway.PathwayHandlerWS(c, client)

	}))

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Page not found")
	})
}
