// Package utils provides utility functions for HTTP clients and other common operations.
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/zap"

	"atlassian-dc-mcp-go/internal/utils/logging"
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
			Timeout:   config.Timeout,
			KeepAlive: config.Timeout,
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

// ExecuteRequest executes an HTTP request with the provided parameters.
// It builds the request and executes it with retry logic.
func ExecuteRequest(client *BaseClient, method string, pathSegments []string, queryParams map[string][]string, body []byte, accept Accept, result any) error {
	// Build the HTTP request
	req, err := BuildHttpRequest(method, client.Config.URL, pathSegments, queryParams, body, client.Config.Token, accept)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	// Execute the request with retry mechanism
	err = executeHTTPRequest(client.HTTPClient, req, client.Name, result)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	return nil
}

// ExecuteStream executes an HTTP request and returns a stream of the response body.
// It builds the request and executes it with retry logic and timeout.
func ExecuteStream(client *BaseClient, method string, pathSegments []string, queryParams map[string][]string, body []byte, accept Accept, timeout time.Duration) (io.ReadCloser, error) {
	// Build the HTTP request
	req, err := BuildHttpRequest(method, client.Config.URL, pathSegments, queryParams, body, client.Config.Token, accept)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	// Execute the stream request using the common utility function
	return executeStreamRequest(client.HTTPClient, req, client.Name, timeout)
}

// ExecuteHTTPRequest executes an HTTP request with retry logic using hashicorp/go-retryablehttp
func executeHTTPRequest(client *retryablehttp.Client, req *http.Request, service string, result interface{}) error {
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

// ExecuteStreamRequest executes an HTTP request with retry logic and returns a stream of the response body.
func executeStreamRequest(client *retryablehttp.Client, req *http.Request, service string, timeout time.Duration) (io.ReadCloser, error) {
	// Create a context with timeout
	ctx := context.Background()
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	req = req.WithContext(ctx)

	// Convert the request to a retryable request
	retryableReq, err := retryablehttp.FromRequest(req)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to convert request: %w", service, err)
	}

	// Execute the request
	resp, err := client.Do(retryableReq)
	if err != nil {
		return nil, fmt.Errorf("[%s] request failed: %w", service, err)
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		resp.Body.Close()
		return nil, fmt.Errorf("[%s] request failed with status %d", service, resp.StatusCode)
	}

	return resp.Body, nil
}

// HandleHTTPError handles HTTP errors based on status codes and logs them
func HandleHTTPError(resp *http.Response, service string) error {
	strErr := ""
	switch resp.StatusCode {
	case http.StatusBadRequest:
		strErr = "bad request"
	case http.StatusUnauthorized:
		strErr = "unauthorized"
	case http.StatusForbidden:
		strErr = "forbidden"
	case http.StatusNotFound:
		strErr = "not found"
	case http.StatusTooManyRequests:
		strErr = "too many requests"
	case http.StatusInternalServerError:
		strErr = "internal server error"
	case http.StatusBadGateway:
		strErr = "bad gateway"
	case http.StatusServiceUnavailable:
		strErr = "service unavailable"
	case http.StatusGatewayTimeout:
		strErr = "gateway timeout"
	default:
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		} else if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			strErr = "client error"
		} else if resp.StatusCode >= 500 {
			strErr = "server error"
		} else {
			strErr = "unknown error"
		}
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	logging.GetLogger().Warn("HTTP request failed",
		zap.String("service", service),
		zap.Int("status_code", resp.StatusCode),
		zap.String("error", strErr),
		zap.String("response_body", bodyString))

	return fmt.Errorf("[%s] %s : %d - %s", service, strErr, resp.StatusCode, bodyString)
}
