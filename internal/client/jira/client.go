// Package jira provides a client for interacting with Jira Data Center APIs.
package jira

import (
	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/config"
)

// JiraClient represents a client for interacting with Jira Data Center APIs
type JiraClient struct {
	*client.BaseClient
}

// NewJiraClient creates a new Jira client with the provided configuration.
func NewJiraClient(config *config.ClientConfig) *JiraClient {
	return &JiraClient{
		BaseClient: client.NewBaseClient(config, "jira"),
	}
}
