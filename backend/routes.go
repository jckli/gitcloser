package backend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jckli/gitcloser/backend/handlers/index"
)

func InitRoutes(app *fiber.App) {
	app.Get("/", index.IndexHandler)
}
