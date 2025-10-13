// Package confluence provides a client for interacting with Confluence Data Center APIs.
package confluence

import (
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/utils"
)

// ConfluenceClient represents a client for interacting with Confluence Data Center APIs
type ConfluenceClient struct {
	Config       *config.ConfluenceConfig
	HTTPClient   *http.Client
	ClientConfig *utils.HTTPClientConfig
}

// NewConfluenceClient creates a new Confluence client with the provided configuration.
func NewConfluenceClient(config *config.ConfluenceConfig) *ConfluenceClient {
	clientConfig := utils.DefaultHTTPClientConfig()
	httpClient := utils.NewHTTPClient(clientConfig)

	return &ConfluenceClient{
		Config:       config,
		HTTPClient:   httpClient,
		ClientConfig: clientConfig,
	}
}

// executeRequest executes an HTTP request to the Confluence API.
func (c *ConfluenceClient) executeRequest(method string, pathSegments []string, queryParams url.Values, body []byte, result any, accept utils.Accept) error {
	req, err := utils.BuildHttpRequest(method, c.Config.URL, pathSegments, queryParams, body, c.Config.Token, accept)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	if err := utils.ExecuteHTTPRequestWithRetry(c.HTTPClient, req, "confluence", result, c.ClientConfig); err != nil {
		return err
	}

	return nil
}
