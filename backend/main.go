package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

// Bot represents an autonomous water-cleaning robot in the system.
// This is our core data model that tracks the state of each bot.
type Bot struct {
	ID           string  `json:"id"`
	Status       string  `json:"status"`        // "active", "idle", "charging", "maintenance"
	BatteryLevel float64 `json:"battery_level"` // 0-100 percentage
	TrashCount   int     `json:"trash_count"`   // total items collected
}

// TelemetryData represents the data sent by bots
type TelemetryData struct {
	BotID        string    `json:"bot_id"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	BatteryLevel float64   `json:"battery_level"`
	Timestamp    time.Time `json:"timestamp"`
}

// botsMutex protects the bots map from concurrent access.
// We need this because multiple bots could send telemetry at the same time,
// and without synchronization, we could have race conditions where two
// goroutines try to read/write the map simultaneously, causing data corruption.
var botsMutex sync.Mutex

// bots stores the current state of all known bots in the system.
// The key is the bot ID, value is the latest telemetry data.
var bots = make(map[string]TelemetryData)

func main() {
	// Load persisted stats on startup
	if err := LoadStats(); err != nil {
		log.Printf("Warning: could not load stats: %v", err)
	}

	// Set up our HTTP routes
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/telemetry", telemetryHandler)
	http.HandleFunc("/detection", detectionHandler)
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/command", commandHandler)
	http.HandleFunc("/bots", botsHandler)

	// WebSocket endpoint for real-time updates
	http.Handle("/ws", websocket.Handler(wsHandler))

	// Start the server
	port := ":8080"
	fmt.Printf("🤖 AquaGuard API server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// healthHandler returns a simple health check response
// This is useful for load balancers and monitoring systems
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{
		"status":  "healthy",
		"service": "aquaguard-api",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// telemetryHandler receives location and battery data from bots
// This is called frequently by each bot to report its current state
func telemetryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data TelemetryData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Set timestamp if not provided
	if data.Timestamp.IsZero() {
		data.Timestamp = time.Now()
	}

	// Lock the mutex before accessing the shared map
	// This ensures only one goroutine can modify the map at a time
	botsMutex.Lock()
	bots[data.BotID] = data
	botsMutex.Unlock()

	log.Printf("Received telemetry from bot %s: battery=%.1f%%, location=(%.4f, %.4f)",
		data.BotID, data.BatteryLevel, data.Latitude, data.Longitude)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}
