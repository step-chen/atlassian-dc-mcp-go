// Package utils provides utility functions for HTTP clients and other common operations.
package utils

import (
	"fmt"
	"net"
	"net/http"
	"time"
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

// NewHTTPClient creates a new HTTP client with the provided configuration
func NewHTTPClient(config *HTTPClientConfig) *http.Client {
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

	return &http.Client{
		Transport: transport,
		Timeout:   config.Timeout,
	}
}

// ExecuteHTTPRequestWithRetry executes an HTTP request with retry logic
func ExecuteHTTPRequestWithRetry(client *http.Client, req *http.Request, service string, result interface{}, config *HTTPClientConfig) error {
	var lastErr error

	attempts := 1
	if config != nil {
		attempts += config.RetryAttempts
	}

	for i := 0; i < attempts; i++ {
		// Clone the request for retries since the body might be consumed
		clonedReq := req.Clone(req.Context())
		if req.Body != nil {
			// For simplicity, we assume the body can be cloned/recreated
			// In a more robust implementation, we would need to handle this more carefully
			clonedReq.Body = req.Body
		}

		err := ExecuteHTTPRequest(client, clonedReq, service, result)
		if err == nil {
			return nil
		}

		lastErr = err

		// If this is not the last attempt, wait before retrying
		if i < attempts-1 && config != nil && config.RetryDelay > 0 {
			time.Sleep(config.RetryDelay)
		}
	}

	return fmt.Errorf("request failed after %d attempts: %w", attempts, lastErr)
}
