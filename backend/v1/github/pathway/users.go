package pathway

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/jckli/gitcloser/algorithm"
	"github.com/valyala/fasthttp"
)

func PathwayHandler(c *fiber.Ctx, client *fasthttp.Client) error {
	// get parameter of user1 and user2
	user1 := c.Params("user1")
	user2 := c.Params("user2")

	pathway, err := algorithm.FindShortestPath(user1, user2, client)

	c.Response().Header.Set("Content-Type", "application/json")

	if err != nil {
		c.Status(500)
		return c.JSON(&ErrorResponse{
			Status: 500,
			Error:  err.Error(),
		})
	}

	response := &PathwayResponse{
		Status: 200,
	}
	response.Data.Pathway = ParsePathwayUser(pathway)

	c.Status(response.Status)

	return c.JSON(response)
}

func PathwayHandlerWS(c *websocket.Conn, client *fasthttp.Client) {
	// get parameter of user1 and user2
	user1 := c.Params("user1")
	user2 := c.Params("user2")

	pathway, err := algorithm.FindShortestPathWS(user1, user2, c, client)
	if err != nil {
		c.WriteJSON(&ErrorResponse{
			Status: 500,
			Error:  err.Error(),
		})
		return
	}

	response := &PathwayResponse{
		Status: 200,
	}
	response.Data.Pathway = ParsePathwayUser(pathway)

	c.WriteJSON(response)

	c.Close()
	return
}

func ParsePathwayUser(pathway []algorithm.UserNode) []PathwayUser {
	pathwayUsers := make([]PathwayUser, len(pathway))

	for i, user := range pathway {
		pathwayUsers[i] = PathwayUser{
			Login:     user.Login,
			AvatarUrl: user.AvatarUrl,
			Url:       user.Url,
			Followers: struct {
				TotalCount int `json:"totalCount"`
			}{
				TotalCount: user.Followers.TotalCount,
			},
			Following: struct {
				TotalCount int `json:"totalCount"`
			}{
				TotalCount: user.Following.TotalCount,
			},
		}
	}

	return pathwayUsers
}
