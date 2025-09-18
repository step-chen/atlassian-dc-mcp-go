// Package bitbucket provides type definitions for Bitbucket MCP tools.
package bitbucket

// MapOutput represents a generic output map for Bitbucket MCP tools
type MapOutput = map[string]interface{}

// ContentOutput represents the output for getting content
type ContentOutput struct {
	Content []byte `json:"content" jsonschema:"the content"`
}

// GetUserOutput represents the output for getting a user
type GetUserOutput struct {
	User map[string]interface{} `json:"user" jsonschema:"the user details"`
}

type DiffOutput struct {
	Diff string `json:"diff"`
}
