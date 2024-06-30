package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jckli/gitcloser/backend"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	app := fiber.New()

	/*
		client := &fasthttp.Client{}
		startTime := time.Now()
		path, err := algorithm.FindShortestPath("jckli", "Phineas", client)
		endTime := time.Now()

		fmt.Println(path, err)
		fmt.Println("Time taken: ", endTime.Sub(startTime))
	*/

	backend.InitRoutes(app)
	app.Listen(":" + os.Getenv("PORT"))
}
