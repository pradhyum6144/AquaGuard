package main

import (
	"os"
)

// Config holds application configuration
type Config struct {
	Port           string
	RedisURL       string
	SecurityURL    string
	LogLevel       string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		RedisURL:    getEnv("REDIS_URL", "localhost:6379"),
		SecurityURL: getEnv("SECURITY_URL", "http://localhost:8081"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}
}

// getEnv returns environment variable or default
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
