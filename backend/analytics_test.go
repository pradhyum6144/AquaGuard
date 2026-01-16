package main

import (
	"testing"
)

func TestProcessDetection_Plastic(t *testing.T) {
	// Reset stats for test
	impactStats = ImpactStats{}

	event := DetectionEvent{
		BotID:     "test-bot",
		TrashType: TrashTypePlastic,
		Latitude:  19.0760,
		Longitude: 72.8777,
	}

	ProcessDetection(event)

	if impactStats.TotalPlastic != 1 {
		t.Errorf("Expected TotalPlastic=1, got %d", impactStats.TotalPlastic)
	}
}

func TestProcessDetection_Organic(t *testing.T) {
	impactStats = ImpactStats{}

	event := DetectionEvent{
		BotID:     "test-bot",
		TrashType: TrashTypeOrganic,
	}

	ProcessDetection(event)

	if impactStats.TotalOrganic != 1 {
		t.Errorf("Expected TotalOrganic=1, got %d", impactStats.TotalOrganic)
	}
	if impactStats.TotalPlastic != 0 {
		t.Errorf("Expected TotalPlastic=0, got %d", impactStats.TotalPlastic)
	}
}

func TestGetImpactStats(t *testing.T) {
	impactStats = ImpactStats{
		TotalPlastic: 10,
		TotalMetal:   5,
	}

	stats := GetImpactStats()

	if stats.TotalPlastic != 10 {
		t.Errorf("Expected TotalPlastic=10, got %d", stats.TotalPlastic)
	}
}
