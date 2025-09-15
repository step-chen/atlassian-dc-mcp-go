package confluence

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCurrentUserHandler handles getting the current Confluence user
func (h *Handler) getCurrentUserHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get current user", func() (interface{}, error) {
		return h.client.GetCurrentUser()
	})
}

// AddUserTools registers the user-related tools with the MCP server
func AddUserTools(server *mcp.Server, client *confluence.ConfluenceClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_current_user",
		Description: "Get current Confluence user. This tool retrieves information about the currently authenticated user.",
		InputSchema: &jsonschema.Schema{
			Type:       "object",
			Properties: map[string]*jsonschema.Schema{},
		},
	}, handler.getCurrentUserHandler)
}
