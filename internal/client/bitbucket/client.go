// Package bitbucket provides a client for interacting with Bitbucket Data Center APIs.
package bitbucket

import (
	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/config"
)

// BitbucketClient represents a client for interacting with Bitbucket Data Center APIs
type BitbucketClient struct {
	*client.BaseClient
}

// NewBitbucketClient creates a new Bitbucket client with the provided configuration.
func NewBitbucketClient(config *config.ClientConfig) *BitbucketClient {
	return &BitbucketClient{
		BaseClient: client.NewBaseClient(config, "bitbucket"),
	}
}