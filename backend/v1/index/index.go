package index

import (
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	response := &IndexResponse{
		Status: 200,
		Data: &IndexData{
			Message: "GitCloser API v1",
		},
	}

	c.Status(response.Status)
	c.Response().Header.Set("Content-Type", "application/json")

	return c.JSON(response)
}
