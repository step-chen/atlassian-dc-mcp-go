package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCurrentUserHandler handles getting the current Confluence user
func (h *Handler) getCurrentUserHandler(ctx context.Context, req *mcp.CallToolRequest, input types.EmptyInput) (*mcp.CallToolResult, types.MapOutput, error) {
	user, err := h.client.GetCurrentUser()
	if err != nil {
		return nil, nil, fmt.Errorf("get current user failed: %w", err)
	}

	return nil, user, nil
}

// AddUserTools registers the user-related tools with the MCP server
func AddUserTools(server *mcp.Server, client *confluence.ConfluenceClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[types.EmptyInput, types.MapOutput](server, "confluence_get_current_user", "Get current Confluence user. This tool retrieves information about the currently authenticated user.", handler.getCurrentUserHandler)
}
