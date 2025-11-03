// Package jira provides a client for interacting with Jira Data Center APIs.
package jira

import (
	"time"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/config"
)

// JiraClient represents a client for interacting with Jira Data Center APIs
type JiraClient struct {
	*client.BaseClient
}

// NewJiraClient creates a new Jira client with the provided configuration.
func NewJiraClient(config *config.ClientConfig) *JiraClient {
	clientConfig := client.DefaultHTTPClientConfig()
	if config.Timeout > 0 {
		clientConfig.Timeout = time.Duration(config.Timeout) * time.Second
	}

	return &JiraClient{
		BaseClient: &client.BaseClient{
			Config:     config,
			HTTPClient: client.NewRetryableHTTPClient(clientConfig),
			Name:       "jira",
		},
	}
}
