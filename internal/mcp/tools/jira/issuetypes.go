package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getIssueTypesHandler handles getting Jira issue types
func (h *Handler) getIssueTypesHandler(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	issueTypes, err := h.client.GetIssueTypes()
	if err != nil {
		return nil, nil, fmt.Errorf("get issue types failed: %w", err)
	}

	resultMap := map[string]interface{}{
		"issueTypes": issueTypes,
	}

	return nil, resultMap, nil
}

// AddIssueTypeTools registers the issue type-related tools with the MCP server
func AddIssueTypeTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[struct{}, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_issue_types",
		Description: "Get Jira issue types",
	}, handler.getIssueTypesHandler)
}
