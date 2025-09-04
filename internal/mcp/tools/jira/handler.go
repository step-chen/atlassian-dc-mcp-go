// Package jira provides MCP tools for interacting with Jira.
package jira

import (
	"atlassian-dc-mcp-go/internal/client/jira"
)

// Handler encapsulates all Jira tool handlers
type Handler struct {
	client *jira.JiraClient
}

// NewHandler creates a new Handler with the provided client
func NewHandler(client *jira.JiraClient) *Handler {
	return &Handler{
		client: client,
	}
}
