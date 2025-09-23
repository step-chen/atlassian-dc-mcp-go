package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/utils"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getPrioritiesHandler handles getting Jira priorities
func (h *Handler) getPrioritiesHandler(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	priorities, err := h.client.GetPriorities()
	if err != nil {
		return nil, nil, fmt.Errorf("get priorities failed: %w", err)
	}

	resultMap := map[string]interface{}{
		"priorities": priorities,
	}

	return nil, resultMap, nil
}

// AddPriorityTools registers the priority-related tools with the MCP server
func AddPriorityTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[struct{}, map[string]interface{}](server, "jira_get_priorities", "Get all Jira priorities", handler.getPrioritiesHandler)
}
