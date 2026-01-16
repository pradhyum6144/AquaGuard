package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// BotCommand represents a command to send to a bot
type BotCommand struct {
	BotID   string `json:"bot_id"`
	Command string `json:"command"` // "halt", "resume", "return_home"
}

// commandHandler receives commands for specific bots
func commandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cmd BotCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Check if global override is active
	if IsOverrideActive() && cmd.Command != "halt" {
		http.Error(w, "Manual override active - only HALT commands accepted", http.StatusForbidden)
		return
	}

	log.Printf("Command for bot %s: %s", cmd.BotID, cmd.Command)

	// TODO: Send encrypted command to bot via security service
	// For now, just acknowledge
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "sent",
		"bot_id":  cmd.BotID,
		"command": cmd.Command,
	})
}

// botsHandler returns all known bots and their states
func botsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	botsMutex.Lock()
	botsData := make([]TelemetryData, 0, len(bots))
	for _, bot := range bots {
		botsData = append(botsData, bot)
	}
	botsMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"bots":             botsData,
		"override_active":  IsOverrideActive(),
		"total_bots":       len(botsData),
	})
}
