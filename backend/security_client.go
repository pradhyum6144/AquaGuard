package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SecurityClient handles communication with the Rust security service
type SecurityClient struct {
	baseURL string
}

// NewSecurityClient creates a new client for the security service
func NewSecurityClient() *SecurityClient {
	// Fixed: Security service runs on port 8081
	return &SecurityClient{
		baseURL: "http://localhost:8081",
	}
}

// EncryptCommand encrypts a bot command before sending it
func (c *SecurityClient) EncryptCommand(command string) (string, string, error) {
	payload := map[string]string{"data": command}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := http.Post(c.baseURL+"/encrypt", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", "", fmt.Errorf("failed to call security service: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("security service error: %s", string(body))
	}

	var result struct {
		Encrypted string `json:"encrypted"`
		Nonce     string `json:"nonce"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", fmt.Errorf("failed to parse response: %w", err)
	}

	return result.Encrypted, result.Nonce, nil
}
