package main

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

var startTime = time.Now()

// HealthStatus represents the overall system health
type HealthStatus struct {
	Status       string            `json:"status"`
	Service      string            `json:"service"`
	Version      string            `json:"version"`
	Uptime       string            `json:"uptime"`
	GoVersion    string            `json:"go_version"`
	NumGoroutine int               `json:"num_goroutine"`
	Dependencies map[string]string `json:"dependencies"`
}

// detailedHealthHandler provides extended health information
func detailedHealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	uptime := time.Since(startTime)

	health := HealthStatus{
		Status:       "healthy",
		Service:      "aquaguard-api",
		Version:      "0.1.0",
		Uptime:       uptime.String(),
		GoVersion:    runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
		Dependencies: map[string]string{
			"security-service": checkServiceHealth("http://localhost:8081/health"),
			"redis":            "not_configured",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

// checkServiceHealth checks if a service is healthy
func checkServiceHealth(url string) string {
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "unhealthy"
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return "healthy"
	}
	return "unhealthy"
}
