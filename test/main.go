package main

import (
	"github.com/fasthttp/websocket"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func main() {
	url := "ws://localhost:3000/v1/github/pathway/jckli/Phineas/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatalf("Error connecting to websocket: %v", err)
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}
		log.Printf("Received: %s\n", msg)
	}
}
