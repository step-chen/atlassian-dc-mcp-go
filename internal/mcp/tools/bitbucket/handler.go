// Package bitbucket provides MCP tools for interacting with Bitbucket.
package bitbucket

import (
	"atlassian-dc-mcp-go/internal/client/bitbucket"
)

// Handler encapsulates all Bitbucket tool handlers
type Handler struct {
	client *bitbucket.BitbucketClient
}

// NewHandler creates a new Handler with the provided client
func NewHandler(client *bitbucket.BitbucketClient) *Handler {
	return &Handler{
		client: client,
	}
}