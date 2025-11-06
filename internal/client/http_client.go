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

	"atlassian-dc-mcp-go/internal/types"
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

// HTTPErrorHandlingOptions defines optional configuration for HTTP error handling
type HTTPErrorHandlingOptions struct {
	SkipLogging     bool
	MaxBodySize     int64
	SkipBodyReading bool
	CustomErrorHandler func(*http.Response) error
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
func NewRetryableHTTPClient(config *HTTPClientConfig, transport http.RoundTripper) *retryablehttp.Client {
	// Create base transport with the provided configuration
	baseTransport := &http.Transport{
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

	// Set the custom transport to wrap the base transport.
	if rt, ok := transport.(*TokenAuthTransport); ok {
		rt.Transport = baseTransport
	} else {
		// If it's not a TokenAuthTransport, use the base transport directly.
		transport = baseTransport
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
func ExecuteRequest(ctx context.Context, client *BaseClient, method string, pathSegments []any, queryParams map[string][]string, body []byte, accept Accept, result any) error {
	// Build the HTTP request
	req, err := buildHttpRequest(method, client.Config.URL, pathSegments, queryParams, body, accept)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	// Convert http.Request to retryablehttp.Request
	retryReq, err := retryablehttp.FromRequest(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("[%s] failed to convert request: %w", client.Name, err)
	}

	// Execute the request with retry mechanism
	resp, err := client.HTTPClient.Do(retryReq)
	if err != nil {
		return fmt.Errorf("[%s] request failed: %w", client.Name, err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if err := HandleHTTPError(resp, client.Name); err != nil {
		return err
	}

	// Decode response if needed
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("[%s] failed to decode response: %w", client.Name, err)
		}
	}

	return nil
}

// ExecuteStream executes an HTTP request and returns a stream of the response body.
// It builds the request and executes it with retry logic and timeout.
func ExecuteStream(ctx context.Context, client *BaseClient, method string, pathSegments []any, queryParams map[string][]string, body []byte, accept Accept, timeout time.Duration) (io.ReadCloser, error) {
	// Build the HTTP request
	req, err := buildHttpRequest(method, client.Config.URL, pathSegments, queryParams, body, accept)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	// Create a context with timeout
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	req = req.WithContext(ctx)

	// Convert the request to a retryable request
	retryableReq, err := retryablehttp.FromRequest(req)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to convert request: %w", client.Name, err)
	}

	// Execute the request
	resp, err := client.HTTPClient.Do(retryableReq)
	if err != nil {
		return nil, fmt.Errorf("[%s] request failed: %w", client.Name, err)
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		resp.Body.Close()
		return nil, fmt.Errorf("[%s] request failed with status %d", client.Name, resp.StatusCode)
	}

	return resp.Body, nil
}

// HandleHTTPError handles HTTP errors based on status codes and logs them
func HandleHTTPError(resp *http.Response, service string, opts ...HTTPErrorHandlingOptions) error {
	options := HTTPErrorHandlingOptions{
		MaxBodySize: 1024 * 1024, // 1MB default
	}
	
	if len(opts) > 0 {
		options = opts[0]
	}
	
	// Use custom error handler if provided
	if options.CustomErrorHandler != nil {
		if err := options.CustomErrorHandler(resp); err != nil {
			return err
		}
		// Continue with default handling if custom handler doesn't return an error
	}
	
	// Check if status code indicates success
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	
	var bodyString string
	if !options.SkipBodyReading {
		// Limit the size of response body read
		limitReader := io.LimitReader(resp.Body, options.MaxBodySize)
		bodyBytes, _ := io.ReadAll(limitReader)
		bodyString = string(bodyBytes)
	}
	
	// Log error (unless skipped)
	if !options.SkipLogging {
		logging.GetLogger().Warn("HTTP request failed",
			zap.String("service", service),
			zap.Int("status_code", resp.StatusCode),
			zap.String("response_body", bodyString))
	}
	
	// Return appropriate error based on status code
	switch resp.StatusCode {
	case http.StatusBadRequest:
		return &types.Error{
			Code:    "BAD_REQUEST",
			Message: fmt.Sprintf("[%s] bad request: %s", service, bodyString),
		}
	case http.StatusUnauthorized:
		return &types.Error{
			Code:    "UNAUTHORIZED",
			Message: fmt.Sprintf("[%s] unauthorized: %s", service, bodyString),
		}
	case http.StatusForbidden:
		return &types.Error{
			Code:    "FORBIDDEN",
			Message: fmt.Sprintf("[%s] forbidden: %s", service, bodyString),
		}
	case http.StatusNotFound:
		return &types.Error{
			Code:    "NOT_FOUND",
			Message: fmt.Sprintf("[%s] not found: %s", service, bodyString),
		}
	case http.StatusTooManyRequests:
		return &types.Error{
			Code:    "TOO_MANY_REQUESTS",
			Message: fmt.Sprintf("[%s] too many requests: %s", service, bodyString),
		}
	case http.StatusInternalServerError:
		return &types.Error{
			Code:    "INTERNAL_SERVER_ERROR",
			Message: fmt.Sprintf("[%s] internal server error: %s", service, bodyString),
		}
	case http.StatusBadGateway:
		return &types.Error{
			Code:    "BAD_GATEWAY",
			Message: fmt.Sprintf("[%s] bad gateway: %s", service, bodyString),
		}
	case http.StatusServiceUnavailable:
		return &types.Error{
			Code:    "SERVICE_UNAVAILABLE",
			Message: fmt.Sprintf("[%s] service unavailable: %s", service, bodyString),
		}
	case http.StatusGatewayTimeout:
		return &types.Error{
			Code:    "GATEWAY_TIMEOUT",
			Message: fmt.Sprintf("[%s] gateway timeout: %s", service, bodyString),
		}
	default:
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return &types.Error{
				Code:    "CLIENT_ERROR",
				Message: fmt.Sprintf("[%s] client error %d: %s", service, resp.StatusCode, bodyString),
			}
		} else if resp.StatusCode >= 500 {
			return &types.Error{
				Code:    "SERVER_ERROR",
				Message: fmt.Sprintf("[%s] server error %d: %s", service, resp.StatusCode, bodyString),
			}
		}
		return &types.Error{
			Code:    "UNKNOWN_ERROR",
			Message: fmt.Sprintf("[%s] unexpected error with status %d: %s", service, resp.StatusCode, bodyString),
		}
	}
}