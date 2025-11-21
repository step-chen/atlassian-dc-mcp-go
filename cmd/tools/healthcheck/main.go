package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

// HealthResponse represents the response from the health check endpoint
type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version,omitempty"`
}

// ReadinessResponse represents the response from the readiness check endpoint
type ReadinessResponse struct {
	Status string   `json:"status"`
	Issues []string `json:"issues,omitempty"`
}

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("MCP_PORT")
	if port == "" {
		port = "8090"
	}

	// Validate port is a number
	_, err := strconv.Atoi(port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid port: %s\n", port)
		os.Exit(1)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// First check the health endpoint for basic liveness
	healthURL := fmt.Sprintf("http://localhost:%s/health", port)
	healthResp, err := client.Get(healthURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to health endpoint: %v\n", err)
		os.Exit(1)
	}
	healthResp.Body.Close()

	// Check health status code
	if healthResp.StatusCode < 200 || healthResp.StatusCode >= 300 {
		fmt.Fprintf(os.Stderr, "Health check failed with status code: %d\n", healthResp.StatusCode)
		os.Exit(1)
	}

	// Then check the readiness endpoint for service dependencies
	readyURL := fmt.Sprintf("http://localhost:%s/ready", port)
	resp, err := client.Get(readyURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to readiness endpoint: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check readiness status code
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// Try to decode the response
		var readyResp ReadinessResponse
		if err := json.NewDecoder(resp.Body).Decode(&readyResp); err == nil {
			fmt.Println("Service is ready")
		} else {
			fmt.Println("Service is ready")
		}
		os.Exit(0) // Ready
	} else {
		// Try to decode the error response
		var readyResp ReadinessResponse
		if err := json.NewDecoder(resp.Body).Decode(&readyResp); err == nil && len(readyResp.Issues) > 0 {
			fmt.Fprintf(os.Stderr, "Service is not ready. Issues:\n")
			for _, issue := range readyResp.Issues {
				fmt.Fprintf(os.Stderr, "  - %s\n", issue)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Readiness check failed with status code: %d\n", resp.StatusCode)
		}
		os.Exit(1) // Not ready
	}
}
