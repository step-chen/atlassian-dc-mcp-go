package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getTagsHandler handles getting tags
func (h *Handler) getTagsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetTagsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	tags, err := h.client.GetTags(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get tags failed: %w", err)
	}

	return nil, tags, nil
}

// getTagHandler handles getting a tag
func (h *Handler) getTagHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetTagInput) (*mcp.CallToolResult, types.MapOutput, error) {
	tag, err := h.client.GetTag(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get tag failed: %w", err)
	}

	return nil, tag, nil
}

// AddTagTools registers the tag-related tools with the MCP server
func AddTagTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[bitbucket.GetTagsInput, types.MapOutput](server, "bitbucket_get_tags", "Get tags in a repository", handler.getTagsHandler)
	utils.RegisterTool[bitbucket.GetTagInput, types.MapOutput](server, "bitbucket_get_tag", "Get a specific tag in a repository", handler.getTagHandler)
}
