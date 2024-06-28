package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jckli/gitcloser/algorithm"
	_ "github.com/joho/godotenv/autoload"
	"github.com/valyala/fasthttp"
)

func main() {
	client := &fasthttp.Client{}
	app := fiber.New()
	startTime := time.Now()
	path, err := algorithm.FindShortestPath("jckli", "Phineas", client)
	endTime := time.Now()

	fmt.Println(path, err)
	fmt.Println("Time taken: ", endTime.Sub(startTime))
	app.Listen(":" + os.Getenv("PORT"))
}
