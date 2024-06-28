package main

import (
	"fmt"
	"time"

	"github.com/jckli/gitcloser/algorithm"
	_ "github.com/joho/godotenv/autoload"
	"github.com/valyala/fasthttp"
)

func main() {
	client := &fasthttp.Client{}
	startTime := time.Now()
	path, err := algorithm.FindShortestPath("alii", "char", client)
	endTime := time.Now()

	fmt.Println(path, err)
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}
