package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getContentChildrenHandler handles getting content children
func (h *Handler) getContentChildrenHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentChildrenInput) (*mcp.CallToolResult, types.MapOutput, error) {
	children, err := h.client.GetContentChildren(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get content children failed: %w", err)
	}

	return nil, children, nil
}

// getContentChildrenByTypeHandler handles getting content children by type
func (h *Handler) getContentChildrenByTypeHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentChildrenByTypeInput) (*mcp.CallToolResult, types.MapOutput, error) {
	children, err := h.client.GetContentChildrenByType(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get content children by type failed: %w", err)
	}

	return nil, children, nil
}

// getContentCommentsHandler handles getting content comments
func (h *Handler) getContentCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentCommentsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	comments, err := h.client.GetContentComments(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get content comments failed: %w", err)
	}

	return nil, comments, nil
}

// AddChildrenTools registers the children-related tools with the MCP server
func AddChildrenTools(server *mcp.Server, client *confluence.ConfluenceClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[confluence.GetContentChildrenInput, types.MapOutput](server, "confluence_get_content_children", "Get content children", handler.getContentChildrenHandler)
	utils.RegisterTool[confluence.GetContentChildrenByTypeInput, types.MapOutput](server, "confluence_get_content_children_by_type", "Get content children by type", handler.getContentChildrenByTypeHandler)
	utils.RegisterTool[confluence.GetContentCommentsInput, types.MapOutput](server, "confluence_get_content_comments", "Get content comments", handler.getContentCommentsHandler)
}
