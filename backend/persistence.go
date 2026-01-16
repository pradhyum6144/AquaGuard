package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

const dataDir = "./data"
const statsFile = "impact_stats.json"

// SaveStats persists the current impact statistics to a JSON file
// This ensures we don't lose data if the server restarts
func SaveStats() error {
	stats := GetImpactStats()

	// Ensure data directory exists
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}

	filePath := filepath.Join(dataDir, statsFile)
	data, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}

	log.Printf("Stats saved to %s", filePath)
	return nil
}

// LoadStats loads previously saved impact statistics from disk
// Called on server startup to restore state
func LoadStats() error {
	filePath := filepath.Join(dataDir, statsFile)

	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		// No saved data yet, start fresh
		log.Println("No saved stats found, starting fresh")
		return nil
	}
	if err != nil {
		return err
	}

	var stats ImpactStats
	if err := json.Unmarshal(data, &stats); err != nil {
		return err
	}

	// Restore the stats
	impactMutex.Lock()
	impactStats = stats
	impactMutex.Unlock()

	log.Printf("Loaded stats: %d plastic, %d metal recovered", stats.TotalPlastic, stats.TotalMetal)
	return nil
}
