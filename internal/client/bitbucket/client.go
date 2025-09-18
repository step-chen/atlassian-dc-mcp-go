// Package bitbucket provides a client for interacting with Bitbucket Data Center APIs.
package bitbucket

import (
	"fmt"
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

// executeRequest executes an HTTP request to the Bitbucket API.
func (c *BitbucketClient) executeRequest(method string, pathSegments []string, queryParams url.Values, body []byte, result interface{}) error {
	req, err := utils.BuildHttpRequest(method, c.Config.URL, pathSegments, queryParams, body, c.Config.Token)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	if err := utils.ExecuteHTTPRequestWithRetry(c.HTTPClient, req, "bitbucket", result, c.ClientConfig); err != nil {
		return err
	}

	return nil
}

// ExecuteRequest is a public wrapper for executeRequest method.
func (c *BitbucketClient) ExecuteRequest(method string, pathSegments []string, queryParams url.Values, body []byte, result interface{}) error {
	return c.executeRequest(method, pathSegments, queryParams, body, result)
}
