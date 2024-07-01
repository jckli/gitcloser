package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jckli/gitcloser/routes"
	_ "github.com/joho/godotenv/autoload"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	client := &fasthttp.Client{}

	/*
		client := &fasthttp.Client{}
		startTime := time.Now()
		path, err := algorithm.FindShortestPath("jckli", "Phineas", client)
		endTime := time.Now()

		fmt.Println(path, err)
		fmt.Println("Time taken: ", endTime.Sub(startTime))
	*/

	backend.InitRoutes(app, client)

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
