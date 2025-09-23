package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/utils"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getSpaceHandler handles getting a specific Confluence space
func (h *Handler) getSpaceHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetSpaceInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	space, err := h.client.GetSpace(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get space failed: %w", err)
	}

	return nil, space, nil
}

// getContentsInSpaceHandler handles getting contents in a specific Confluence space
func (h *Handler) getContentsInSpaceHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentsInSpaceInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	contents, err := h.client.GetContentsInSpace(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get contents in space failed: %w", err)
	}

	return nil, contents, nil
}

// getContentsByTypeHandler handles getting contents by type in a specific Confluence space
func (h *Handler) getContentsByTypeHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentsByTypeInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	contents, err := h.client.GetContentsByType(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get contents by type failed: %w", err)
	}

	return nil, contents, nil
}

// getSpacesByKeyHandler handles getting spaces by key
func (h *Handler) getSpacesByKeyHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetSpacesByKeyInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	spaces, err := h.client.GetSpacesByKey(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get spaces by key failed: %w", err)
	}

	return nil, spaces, nil
}

// AddSpaceTools registers the space-related tools with the MCP server
func AddSpaceTools(server *mcp.Server, client *confluence.ConfluenceClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[confluence.GetSpaceInput, map[string]interface{}](server, "confluence_get_space", "Get a specific Confluence space by its key. This tool allows you to retrieve detailed information about a space including its name, description, and metadata.", handler.getSpaceHandler)
	utils.RegisterTool[confluence.GetContentsInSpaceInput, map[string]interface{}](server, "confluence_get_contents_in_space", "Get contents in a specific Confluence space. This tool allows you to retrieve all content items within a space.", handler.getContentsInSpaceHandler)
	utils.RegisterTool[confluence.GetContentsByTypeInput, map[string]interface{}](server, "confluence_get_contents_by_type", "Get contents by type in a specific Confluence space. This tool allows you to retrieve content items of a specific type (e.g., page, blogpost) within a space.", handler.getContentsByTypeHandler)
	utils.RegisterTool[confluence.GetSpacesByKeyInput, map[string]interface{}](server, "confluence_get_spaces_by_key", "Get spaces by key with various filter options. This tool allows you to retrieve spaces using multiple filter criteria including keys, IDs, types, status, and labels.", handler.getSpacesByKeyHandler)
}
