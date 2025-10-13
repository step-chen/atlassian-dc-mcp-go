// Package bitbucket provides a client for interacting with Bitbucket Data Center APIs.
package bitbucket

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/utils"
)

// BitbucketClient represents a client for interacting with Bitbucket Data Center APIs
type BitbucketClient struct {
	Config       *config.BitbucketConfig
	HTTPClient   *http.Client
	ClientConfig *utils.HTTPClientConfig
}

// NewBitbucketClient creates a new Bitbucket client with the provided configuration.
func NewBitbucketClient(cfg *config.BitbucketConfig) *BitbucketClient {
	clientConfig := utils.DefaultHTTPClientConfig()
	httpClient := utils.NewHTTPClient(clientConfig)

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
	err = utils.ExecuteHTTPRequestWithRetry(c.HTTPClient, req, "Bitbucket", result, c.ClientConfig)
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

	// Execute the request
	resp, err := c.HTTPClient.Do(req)
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
