// Package bitbucket provides type definitions for Bitbucket MCP tools.
package bitbucket

import "atlassian-dc-mcp-go/internal/types"

// ContentOutput represents the output for getting content
type ContentOutput struct {
	Content []byte `json:"content" jsonschema:"the content"`
}

// GetUserOutput represents the output for getting a user
type GetUserOutput struct {
	User types.MapOutput `json:"user" jsonschema:"the user details"`
}

type DiffOutput struct {
	Diff string `json:"diff"`
}
