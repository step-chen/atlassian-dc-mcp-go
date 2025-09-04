package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getPrioritiesHandler handles getting Jira priorities
func (h *Handler) getPrioritiesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get priorities", func() (interface{}, error) {
		priorities, err := h.client.GetPriorities()
		if err != nil {
			return nil, err
		}

		resultMap := map[string]interface{}{
			"priorities": priorities,
		}
		return resultMap, nil
	})
}

// AddPriorityTools registers the priority-related tools with the MCP server
func AddPriorityTools(server *mcp.Server, client *jira.JiraClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_priorities",
		Description: "Get all Jira priorities",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{},
		},
	}, handler.getPrioritiesHandler)
}