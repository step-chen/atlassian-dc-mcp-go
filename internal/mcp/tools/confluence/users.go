package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCurrentUserHandler handles getting the current Confluence user
func (h *Handler) getCurrentUserHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.EmptyInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	user, err := h.client.GetCurrentUser()
	if err != nil {
		return nil, nil, fmt.Errorf("get current user failed: %w", err)
	}

	return nil, user, nil
}

// AddUserTools registers the user-related tools with the MCP server
func AddUserTools(server *mcp.Server, client *confluence.ConfluenceClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[confluence.EmptyInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_current_user",
		Description: "Get current Confluence user. This tool retrieves information about the currently authenticated user.",
		InputSchema: &jsonschema.Schema{
			Type:       "object",
			Properties: map[string]*jsonschema.Schema{},
		},
	}, handler.getCurrentUserHandler)
}
