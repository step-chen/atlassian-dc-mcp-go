// Package confluence provides MCP tools for interacting with Confluence.
package confluence

import (
	"atlassian-dc-mcp-go/internal/client/confluence"
)

// Handler encapsulates all Confluence tool handlers
type Handler struct {
	client *confluence.ConfluenceClient
}

// NewHandler creates a new Handler with the provided client
func NewHandler(client *confluence.ConfluenceClient) *Handler {
	return &Handler{
		client: client,
	}
}
