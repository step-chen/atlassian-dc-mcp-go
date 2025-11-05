// Package confluence provides a client for interacting with Confluence Data Center APIs.
package confluence

import (
	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/config"
)

// ConfluenceClient represents a client for interacting with Confluence Data Center APIs
type ConfluenceClient struct {
	*client.BaseClient
}

// NewConfluenceClient creates a new Confluence client with the provided configuration.
func NewConfluenceClient(config *config.ClientConfig) (*ConfluenceClient, error) {
	baseClient, err := client.NewBaseClient(config, "confluence", client.ConfluenceTokenKey)
	if err != nil {
		return nil, err
	}

	return &ConfluenceClient{
		BaseClient: baseClient,
	}, nil
}
