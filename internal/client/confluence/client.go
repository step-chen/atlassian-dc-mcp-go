// Package confluence provides a client for interacting with Confluence Data Center APIs.
package confluence

import (
	"time"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/config"
)

// ConfluenceClient represents a client for interacting with Confluence Data Center APIs
type ConfluenceClient struct {
	*client.BaseClient
}

// NewConfluenceClient creates a new Confluence client with the provided configuration.
func NewConfluenceClient(config *config.ClientConfig) *ConfluenceClient {
	clientConfig := client.DefaultHTTPClientConfig()
	if config.Timeout > 0 {
		clientConfig.Timeout = time.Duration(config.Timeout) * time.Second
	}

	return &ConfluenceClient{
		BaseClient: &client.BaseClient{
			Config:     config,
			HTTPClient: client.NewRetryableHTTPClient(clientConfig),
			Name:       "confluence",
		},
	}
}
