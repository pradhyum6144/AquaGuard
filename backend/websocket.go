package main

import (
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

var (
	clientsMutex sync.Mutex
	clients      = make(map[*websocket.Conn]bool)
	botOverride  = false
)

// WebSocket handler for real-time dashboard updates
func wsHandler(ws *websocket.Conn) {
	clientsMutex.Lock()
	clients[ws] = true
	clientsMutex.Unlock()

	log.Printf("New WebSocket client connected")

	defer func() {
		clientsMutex.Lock()
		delete(clients, ws)
		clientsMutex.Unlock()
		ws.Close()
	}()

	// Listen for commands from the dashboard
	for {
		var msg string
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			break
		}

		switch msg {
		case "HALT":
			botOverride = true
			log.Println("🛑 HALT command received - all bots stopping!")
			broadcastMessage("OVERRIDE_ACTIVE")
		case "RESUME":
			botOverride = false
			log.Println("✅ RESUME command received - bots resuming operation")
			broadcastMessage("OVERRIDE_INACTIVE")
		}
	}
}

// broadcastMessage sends a message to all connected clients
func broadcastMessage(msg string) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	for client := range clients {
		websocket.Message.Send(client, msg)
	}
}

// IsOverrideActive returns whether manual override is active
func IsOverrideActive() bool {
	return botOverride
}
