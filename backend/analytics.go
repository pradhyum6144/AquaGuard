package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// TrashType represents the classification of detected items
type TrashType int

const (
	TrashTypeOrganic TrashType = 0 // Leaves, twigs, etc - ignore
	TrashTypePlastic TrashType = 1 // Bottles, bags, etc - collect
	TrashTypeMetal   TrashType = 2 // Cans, etc - collect
	TrashTypeOther   TrashType = 3 // Unknown items
)

// DetectionEvent represents an item detected by the bot's camera/sensors
type DetectionEvent struct {
	BotID     string    `json:"bot_id"`
	TrashType TrashType `json:"trash_type"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}

// ImpactStats tracks the environmental impact of our bots
type ImpactStats struct {
	TotalPlastic int `json:"total_plastic"`
	TotalMetal   int `json:"total_metal"`
	TotalOther   int `json:"total_other"`
	TotalOrganic int `json:"total_organic"` // Ignored items
}

var (
	impactMutex sync.Mutex
	impactStats = ImpactStats{}
)

// ProcessDetection handles a detection event and updates counters
// Uses a switch statement to classify trash and decide what to count
func ProcessDetection(event DetectionEvent) {
	impactMutex.Lock()
	defer impactMutex.Unlock()

	switch event.TrashType {
	case TrashTypeOrganic:
		// Organic matter like leaves - we ignore these
		// Bot should let them decompose naturally
		impactStats.TotalOrganic++
		log.Printf("Bot %s: Detected organic matter, ignoring", event.BotID)

	case TrashTypePlastic:
		// Plastic items - high priority for collection
		impactStats.TotalPlastic++
		log.Printf("Bot %s: Collecting plastic item! Total: %d", event.BotID, impactStats.TotalPlastic)

	case TrashTypeMetal:
		// Metal items - collect for recycling
		impactStats.TotalMetal++
		log.Printf("Bot %s: Collecting metal item! Total: %d", event.BotID, impactStats.TotalMetal)

	case TrashTypeOther:
		// Unknown items - collect to be safe
		impactStats.TotalOther++
		log.Printf("Bot %s: Collecting unknown item", event.BotID)

	default:
		log.Printf("Bot %s: Unknown trash type: %d", event.BotID, event.TrashType)
	}
}

// GetImpactStats returns the current impact statistics
func GetImpactStats() ImpactStats {
	impactMutex.Lock()
	defer impactMutex.Unlock()
	return impactStats
}

// detectionHandler receives detection events from bots
func detectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event DetectionEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	ProcessDetection(event)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "processed"})
}

// statsHandler returns the current impact statistics
func statsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stats := GetImpactStats()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
