package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getIssueTypesHandler handles getting Jira issue types
func (h *Handler) getIssueTypesHandler(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	issueTypes, err := h.client.GetIssueTypes()
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get issue types")
		return result, nil, err
	}

	resultMap := map[string]interface{}{
		"issueTypes": issueTypes,
	}

	result, err := tools.CreateToolResult(resultMap)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create issue types result")
		return result, nil, err
	}

	return result, resultMap, nil
}

// AddIssueTypeTools registers the issue type-related tools with the MCP server
func AddIssueTypeTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[struct{}, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_issue_types",
		Description: "Get Jira issue types",
	}, handler.getIssueTypesHandler)
}
