package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getWorklogsHandler handles getting worklogs for a Jira issue
func (h *Handler) getWorklogsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetWorklogsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	worklogs, err := h.client.GetWorklogs(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get worklogs failed: %w", err)
	}

	return nil, worklogs, nil
}

// AddWorklogTools registers the worklog-related tools with the MCP server
func AddWorklogTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[jira.GetWorklogsInput, types.MapOutput](server, "jira_get_worklogs", "Get worklogs for a Jira issue or a specific worklog by ID", handler.getWorklogsHandler)
}
