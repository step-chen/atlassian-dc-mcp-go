// Package bitbucket provides a client for interacting with Bitbucket Data Center APIs.
package bitbucket

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"

	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/utils"
)

// BitbucketClient represents a client for interacting with Bitbucket Data Center APIs
type BitbucketClient struct {
	Config       *config.BitbucketConfig
	HTTPClient   *retryablehttp.Client
	ClientConfig *utils.HTTPClientConfig
}

// NewBitbucketClient creates a new Bitbucket client with the provided configuration.
func NewBitbucketClient(cfg *config.BitbucketConfig) *BitbucketClient {
	clientConfig := utils.DefaultHTTPClientConfig()
	// 使用配置中的超时时间
	if cfg.Timeout > 0 {
		clientConfig.Timeout = time.Duration(cfg.Timeout) * time.Second
	}
	httpClient := utils.NewRetryableHTTPClient(clientConfig)

	return &BitbucketClient{
		Config:       cfg,
		HTTPClient:   httpClient,
		ClientConfig: clientConfig,
	}
}

func (c *BitbucketClient) executeRequest(method string, pathParams []string, queryParams url.Values, body []byte, result any, accept utils.Accept) error {
	req, err := utils.BuildHttpRequest(method, c.Config.URL, pathParams, queryParams, body, c.Config.Token, accept)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	// Execute the request with retry mechanism
	err = utils.ExecuteHTTPRequestWithRetry(c.HTTPClient, req, "Bitbucket", result)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	return nil
}

func (c *BitbucketClient) executeStreamRequest(method string, pathParams []string, queryParams url.Values, body []byte, accept utils.Accept) (io.ReadCloser, error) {
	// Create the HTTP request using utils
	req, err := utils.BuildHttpRequest(method, c.Config.URL, pathParams, queryParams, body, c.Config.Token, accept)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	// Create a context with timeout
	ctx := context.Background()
	if c.Config.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(c.Config.Timeout)*time.Second)
		defer cancel()
	}
	req = req.WithContext(ctx)

	// Convert the request to a retryable request
	retryableReq, err := retryablehttp.FromRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convert request: %w", err)
	}

	// Execute the request
	resp, err := c.HTTPClient.Do(retryableReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		resp.Body.Close()
		return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	return resp.Body, nil
}
