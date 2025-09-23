package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getWorklogsHandler handles getting worklogs for a Jira issue
func (h *Handler) getWorklogsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetWorklogsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	worklogs, err := h.client.GetWorklogs(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get worklogs")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(worklogs)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create worklogs result")
		return result, nil, err
	}

	return result, worklogs, nil
}

// AddWorklogTools registers the worklog-related tools with the MCP server
func AddWorklogTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[jira.GetWorklogsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_worklogs",
		Description: "Get worklogs for a Jira issue or a specific worklog by ID",
	}, handler.getWorklogsHandler)
}