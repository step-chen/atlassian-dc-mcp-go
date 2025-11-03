// Package bitbucket provides a client for interacting with Bitbucket Data Center APIs.
package bitbucket

import (
	"time"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/config"
)

// BitbucketClient represents a client for interacting with Bitbucket Data Center APIs
type BitbucketClient struct {
	*client.BaseClient
}

// NewBitbucketClient creates a new Bitbucket client with the provided configuration.
func NewBitbucketClient(config *config.ClientConfig) *BitbucketClient {
	clientConfig := client.DefaultHTTPClientConfig()
	if config.Timeout > 0 {
		clientConfig.Timeout = time.Duration(config.Timeout) * time.Second
	}

	return &BitbucketClient{
		BaseClient: &client.BaseClient{
			Config:     config,
			HTTPClient: client.NewRetryableHTTPClient(clientConfig),
			Name:       "bitbucket",
		},
	}
}
