package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// Packet defines the structure for the WebSocket request
type Packet struct {
	Cmd    string `json:"Cmd"`
	Symbol string `json:"Symbol"`
}

func main() {
	// WebSocket server URL
	url := "wss://ccws.ggex.io/market"

	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to WebSocket")

	// Create the request packet
	packet := Packet{
		Cmd:    "TRI",
		Symbol: "RPGUSDT",
	}

	// Serialize the packet to JSON
	jsonData, err := json.Marshal(packet)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Send the JSON data over the WebSocket
	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	log.Printf("Sent request: %s", string(jsonData))

	// Start receiving messages from the server
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// Keep the connection alive with ping messages
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Send a ping message to keep the connection alive
			err := conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				log.Printf("Ping error: %v", err)
				return
			}
		}
	}
}