package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func main() {
	// WebSocket server URL
	//serverURL := "ws://server.example.com/chat"
	serverURL := "wss://echo.websocket.org"

	// Additional WebSocket headers
	headers := http.Header{
		"Upgrade":                  {"websocket"},
		"Connection":               {"Upgrade"},
		"Sec-WebSocket-Key":        {"dGhlIHNhbXBsZSBub25jZQ=="},
		"Origin":                   {"http://example.com"},
		"Sec-WebSocket-Protocol":   {"chat, superchat"},
		"Sec-WebSocket-Version":    {"13"},
	}

	// Establish WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, headers)
	if err != nil {
		fmt.Println("WebSocket connection error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket handshake successful!")

	// Your WebSocket communication logic goes here
	// For example, you can send and receive messages using conn.WriteMessage and conn.ReadMessage

	// Here, we'll just wait for user input to keep the program running
	fmt.Println("Press Enter to exit.")
	fmt.Scanln()
}
