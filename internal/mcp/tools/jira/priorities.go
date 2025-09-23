package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getPrioritiesHandler handles getting Jira priorities
func (h *Handler) getPrioritiesHandler(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	priorities, err := h.client.GetPriorities()
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get priorities")
		return result, nil, err
	}

	resultMap := map[string]interface{}{
		"priorities": priorities,
	}

	result, err := tools.CreateToolResult(resultMap)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create priorities result")
		return result, nil, err
	}

	return result, resultMap, nil
}

// AddPriorityTools registers the priority-related tools with the MCP server
func AddPriorityTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[struct{}, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_priorities",
		Description: "Get all Jira priorities",
	}, handler.getPrioritiesHandler)
}
