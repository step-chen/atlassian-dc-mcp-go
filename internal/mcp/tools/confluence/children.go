package confluence

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getContentChildrenHandler handles getting content children
func (h *Handler) getContentChildrenHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentChildrenInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	children, err := h.client.GetContentChildren(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get content children")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(children)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create content children result")
		return result, nil, err
	}

	return result, children, nil
}

// getContentChildrenByTypeHandler handles getting content children by type
func (h *Handler) getContentChildrenByTypeHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentChildrenByTypeInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	children, err := h.client.GetContentChildrenByType(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get content children by type")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(children)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create content children by type result")
		return result, nil, err
	}

	return result, children, nil
}

// getContentCommentsHandler handles getting content comments
func (h *Handler) getContentCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentCommentsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	comments, err := h.client.GetContentComments(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get content comments")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(comments)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create content comments result")
		return result, nil, err
	}

	return result, comments, nil
}

// AddChildrenTools registers the children-related tools with the MCP server
func AddChildrenTools(server *mcp.Server, client *confluence.ConfluenceClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[confluence.GetContentChildrenInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_content_children",
		Description: "Get content children",
	}, handler.getContentChildrenHandler)

	mcp.AddTool[confluence.GetContentChildrenByTypeInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_content_children_by_type",
		Description: "Get content children by type",
	}, handler.getContentChildrenByTypeHandler)

	mcp.AddTool[confluence.GetContentCommentsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_content_comments",
		Description: "Get content comments",
	}, handler.getContentCommentsHandler)
}
