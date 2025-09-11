package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getTagsHandler handles getting tags
func (h *Handler) getTagsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get tags", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		filterText, _ := tools.GetStringArg(args, "filterText")
		orderBy, _ := tools.GetStringArg(args, "orderBy")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetTags(projectKey, repoSlug, filterText, orderBy, start, limit)
	})
}

// getTagHandler handles getting a specific tag
func (h *Handler) getTagHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get tag", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		tagName, ok := tools.GetStringArg(args, "tagName")
		if !ok {
			return nil, fmt.Errorf("missing or invalid tagName parameter")
		}

		return h.client.GetTag(projectKey, repoSlug, tagName)
	})
}

// AddTagTools registers the tag-related tools with the MCP server
func AddTagTools(server *mcp.Server, client *bitbucket.BitbucketClient, hasWritePermission bool) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_tags",
		Description: "Get tags for a repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"filterText": {
					Type:        "string",
					Description: "Filter text to apply to the tag names",
				},
				"orderBy": {
					Type:        "string",
					Description: "Field to order tags by",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned tags",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of tags to return",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getTagsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_tag",
		Description: "Get a specific tag",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"tagName": {
					Type:        "string",
					Description: "The name of the tag to retrieve",
				},
			},
			Required: []string{"projectKey", "repoSlug", "tagName"},
		},
	}, handler.getTagHandler)
}
