// Package utils provides utility functions for the MCP server.
package utils

import (
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterTool is a helper function that simplifies the registration of MCP tools.
// It reduces boilerplate code by automatically creating the tool definition with
// the provided name and description.
//
// Example usage:
//
//	registerTool(server, "jira_get_issue", "Get a specific Jira issue by its key", handler.getIssueHandler)
func RegisterTool[In, Out any](server *mcp.Server, name, description string, handler mcp.ToolHandlerFor[In, Out]) {
	mcp.AddTool[In, Out](server, &mcp.Tool{
		Name:        name,
		Description: description,
	}, handler)
}