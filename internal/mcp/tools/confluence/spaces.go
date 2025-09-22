package confluence

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getSpaceHandler handles getting a specific Confluence space
func (h *Handler) getSpaceHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetSpaceInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	space, err := h.client.GetSpace(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get space")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(space)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create space result")
		return result, nil, err
	}

	return result, space, nil
}

// getContentsInSpaceHandler handles getting contents in a specific Confluence space
func (h *Handler) getContentsInSpaceHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentsInSpaceInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	contents, err := h.client.GetContentsInSpace(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get contents in space")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(contents)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create contents in space result")
		return result, nil, err
	}

	return result, contents, nil
}

// getContentsByTypeHandler handles getting contents by type in a specific Confluence space
func (h *Handler) getContentsByTypeHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentsByTypeInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	contents, err := h.client.GetContentsByType(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get contents by type")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(contents)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create contents by type result")
		return result, nil, err
	}

	return result, contents, nil
}

// getSpacesByKeyHandler handles getting spaces by key
func (h *Handler) getSpacesByKeyHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetSpacesByKeyInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	spaces, err := h.client.GetSpacesByKey(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get spaces by key")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(spaces)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create spaces by key result")
		return result, nil, err
	}

	return result, spaces, nil
}

// AddSpaceTools registers the space-related tools with the MCP server
func AddSpaceTools(server *mcp.Server, client *confluence.ConfluenceClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[confluence.GetSpaceInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_space",
		Description: "Get a specific Confluence space by its key. This tool allows you to retrieve detailed information about a space including its name, description, and metadata.",
	}, handler.getSpaceHandler)

	mcp.AddTool[confluence.GetContentsInSpaceInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_contents_in_space",
		Description: "Get contents in a specific Confluence space. This tool allows you to retrieve all content items within a space.",
	}, handler.getContentsInSpaceHandler)

	mcp.AddTool[confluence.GetContentsByTypeInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_contents_by_type",
		Description: "Get contents by type in a specific Confluence space. This tool allows you to retrieve content items of a specific type (e.g., page, blogpost) within a space.",
	}, handler.getContentsByTypeHandler)

	mcp.AddTool[confluence.GetSpacesByKeyInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_spaces_by_key",
		Description: "Get spaces by key with various filter options. This tool allows you to retrieve spaces using multiple filter criteria including keys, IDs, types, status, and labels.",
	}, handler.getSpacesByKeyHandler)
}
