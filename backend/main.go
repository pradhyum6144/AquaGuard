package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Bot represents an autonomous water-cleaning robot in the system.
// This is our core data model that tracks the state of each bot.
type Bot struct {
	ID           string  `json:"id"`
	Status       string  `json:"status"`        // "active", "idle", "charging", "maintenance"
	BatteryLevel float64 `json:"battery_level"` // 0-100 percentage
	TrashCount   int     `json:"trash_count"`   // total items collected
}

func main() {
	// Set up our HTTP routes
	http.HandleFunc("/health", healthHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("AquaGuard API server starting on port %s\n", port)
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
