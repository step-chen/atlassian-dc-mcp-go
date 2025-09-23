package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getWorklogsHandler handles getting worklogs for a Jira issue
func (h *Handler) getWorklogsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetWorklogsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	worklogs, err := h.client.GetWorklogs(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get worklogs failed: %w", err)
	}

	return nil, worklogs, nil
}

// AddWorklogTools registers the worklog-related tools with the MCP server
func AddWorklogTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[jira.GetWorklogsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_worklogs",
		Description: "Get worklogs for a Jira issue or a specific worklog by ID",
	}, handler.getWorklogsHandler)
}
