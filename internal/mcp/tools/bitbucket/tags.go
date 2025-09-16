package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetTagsInput represents the input parameters for getting tags
type GetTagsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	FilterText string `json:"filterText,omitempty" jsonschema:"Filter tags by text"`
	OrderBy    string `json:"orderBy,omitempty" jsonschema:"Order tags by"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned tags"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of tags to return"`
}

// GetTagsOutput represents the output for getting tags
type GetTagsOutput = map[string]interface{}

// getTagsHandler handles getting tags
func (h *Handler) getTagsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetTagsInput) (*mcp.CallToolResult, GetTagsOutput, error) {
	tags, err := h.client.GetTags(input.ProjectKey, input.RepoSlug, input.FilterText, input.OrderBy, input.Start, input.Limit)
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

// GetTagInput represents the input parameters for getting a specific tag
type GetTagInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	TagName    string `json:"tagName" jsonschema:"required,The name of the tag to retrieve"`
}

// GetTagOutput represents the output for getting a specific tag
type GetTagOutput = map[string]interface{}

// getTagHandler handles getting a specific tag
func (h *Handler) getTagHandler(ctx context.Context, req *mcp.CallToolRequest, input GetTagInput) (*mcp.CallToolResult, GetTagOutput, error) {
	tag, err := h.client.GetTag(input.ProjectKey, input.RepoSlug, input.TagName)
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

	mcp.AddTool[GetTagsInput, GetTagsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_tags",
		Description: "Get tags in a repository",
	}, handler.getTagsHandler)

	mcp.AddTool[GetTagInput, GetTagOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_tag",
		Description: "Get a specific tag in a repository",
	}, handler.getTagHandler)
}
