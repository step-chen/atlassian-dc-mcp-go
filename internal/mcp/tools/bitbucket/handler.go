package bitbucket

import (
	"atlassian-dc-mcp-go/internal/client/bitbucket"
)

// Handler provides a struct for storing the Bitbucket client
type Handler struct {
	client *bitbucket.BitbucketClient
}

// NewHandler creates a new Handler with the provided client
func NewHandler(client *bitbucket.BitbucketClient) *Handler {
	return &Handler{
		client: client,
	}
}
