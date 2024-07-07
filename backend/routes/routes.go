package backend

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jckli/gitcloser/v1/github"
	"github.com/jckli/gitcloser/v1/index"
	"github.com/valyala/fasthttp"
)

func InitRoutes(app *fiber.App, client *fasthttp.Client) {
	// v1 routes
	app.Get("/", index.IndexHandler)

	app.Get("/v1/github/pathway/:user1/:user2", func(c *fiber.Ctx) error {
		return github.PathwayHandler(c, client)
	})

	app.Get("/v1/github/pathway/:user1/:user2/ws", websocket.New(func(c *websocket.Conn) {
		github.PathwayHandlerWS(c, client)

	}))

	app.Get("/v1/github/search/:query", func(c *fiber.Ctx) error {
		return github.SearchHandler(c, client)
	})

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Page not found")
	})

	/*
		app.Use(cors.New(cors.Config{
			AllowOrigins: "https://*.hayasaka.moe",
		}))
	*/
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
}
