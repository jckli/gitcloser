package github

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func SearchHandler(c *fiber.Ctx, client *fasthttp.Client) error {
	query := c.Params("query")

	c.Response().Header.Set("Access-Control-Allow-Origin", "*")
	c.Response().Header.Set("Content-Type", "application/json")

	users, _, err := getSearchQuery(query, client)
	if err != nil {
		return c.Status(500).JSON(&ErrorResponse{
			Status: 500,
			Error:  err.Error(),
		})
	}

	response := &SearchUsersResponse{
		Status: 200,
	}
	response.Data.Results = users.Search.Nodes

	return c.JSON(response)
}
