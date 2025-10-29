// Package utils provides utility functions for HTTP clients and other common operations.
package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// HTTPClientConfig defines the configuration for HTTP clients
type HTTPClientConfig struct {
	Timeout             time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
	RetryAttempts       int
	RetryDelay          time.Duration
}

// DefaultHTTPClientConfig returns a default HTTP client configuration
func DefaultHTTPClientConfig() *HTTPClientConfig {
	return &HTTPClientConfig{
		Timeout:             30 * time.Second,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
		RetryAttempts:       3,
		RetryDelay:          1 * time.Second,
	}
}

// NewRetryableHTTPClient creates a new retryable HTTP client with the provided configuration
func NewRetryableHTTPClient(config *HTTPClientConfig) *retryablehttp.Client {
	// Create base transport with the provided configuration
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          config.MaxIdleConns,
		MaxIdleConnsPerHost:   config.MaxIdleConnsPerHost,
		IdleConnTimeout:       config.IdleConnTimeout,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// Create base HTTP client
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   config.Timeout,
	}

	// Create retryable client
	client := retryablehttp.NewClient()
	client.HTTPClient = httpClient
	client.RetryMax = config.RetryAttempts
	client.RetryWaitMin = config.RetryDelay
	client.RetryWaitMax = config.RetryDelay * 10
	client.Logger = nil // Disable logging, we'll handle it ourselves

	return client
}

// ExecuteHTTPRequestWithRetry executes an HTTP request with retry logic using hashicorp/go-retryablehttp
func ExecuteHTTPRequestWithRetry(client *retryablehttp.Client, req *http.Request, service string, result interface{}) error {
	// Convert http.Request to retryablehttp.Request
	retryReq, err := retryablehttp.FromRequest(req)
	if err != nil {
		return fmt.Errorf("[%s] failed to convert request: %w", service, err)
	}

	// Execute the request with retry mechanism
	resp, err := client.Do(retryReq)
	if err != nil {
		return fmt.Errorf("[%s] request failed: %w", service, err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if err := HandleHTTPError(resp, service); err != nil {
		return err
	}

	// Decode response if needed
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("[%s] failed to decode response: %w", service, err)
		}
	}

	return nil
}