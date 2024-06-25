package main

import (
	"fmt"

	"github.com/jckli/gitcloser/algorithm"
	_ "github.com/joho/godotenv/autoload"
	"github.com/valyala/fasthttp"
)

func main() {
	client := &fasthttp.Client{}
	path, err := algorithm.FindShortestPath("jckli", "Vendicated", client)
	fmt.Println(path, err)
}
