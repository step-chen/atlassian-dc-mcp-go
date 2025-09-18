package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getTagsHandler handles getting tags
func (h *Handler) getTagsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetTagsInput) (*mcp.CallToolResult, MapOutput, error) {
	tags, err := h.client.GetTags(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get tags")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(tags)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create tags result")
		return result, nil, err
	}

	return result, tags, nil
}

// getTagHandler handles getting a tag
func (h *Handler) getTagHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetTagInput) (*mcp.CallToolResult, MapOutput, error) {
	tag, err := h.client.GetTag(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get tag")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(tag)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create tag result")
		return result, nil, err
	}

	return result, tag, nil
}

// AddTagTools registers the tag-related tools with the MCP server
func AddTagTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[bitbucket.GetTagsInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_tags",
		Description: "Get tags in a repository",
	}, handler.getTagsHandler)

	mcp.AddTool[bitbucket.GetTagInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_tag",
		Description: "Get a specific tag in a repository",
	}, handler.getTagHandler)
}
