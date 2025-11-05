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
func NewBitbucketClient(config *config.ClientConfig) (*BitbucketClient, error) {
	baseClient, err := client.NewBaseClient(config, "bitbucket", client.BitbucketTokenKey)
	if err != nil {
		return nil, err
	}

	return &BitbucketClient{
		BaseClient: baseClient,
	}, nil
}
