package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getIssueTypesHandler handles getting Jira issue types
func (h *Handler) getIssueTypesHandler(ctx context.Context, req *mcp.CallToolRequest, input types.EmptyInput) (*mcp.CallToolResult, types.MapOutput, error) {
	issueTypes, err := h.client.GetIssueTypes()
	if err != nil {
		return nil, nil, fmt.Errorf("get issue types failed: %w", err)
	}

	resultMap := types.MapOutput{
		"issueTypes": issueTypes,
	}

	return nil, resultMap, nil
}

// AddIssueTypeTools registers the issue type-related tools with the MCP server
func AddIssueTypeTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[types.EmptyInput, types.MapOutput](server, "jira_get_issue_types", "Get Jira issue types", handler.getIssueTypesHandler)
}
