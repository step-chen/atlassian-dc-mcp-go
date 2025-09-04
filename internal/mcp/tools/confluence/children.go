package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getContentChildrenHandler handles getting content children
func (h *Handler) getContentChildrenHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get content children", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		expand := tools.GetStringSliceArg(args, "expand")

		parentVersion, _ := tools.GetStringArg(args, "parentVersion")

		return h.client.GetContentChildren(contentID, expand, parentVersion)
	})
}

// getContentChildrenByTypeHandler handles getting content children by type
func (h *Handler) getContentChildrenByTypeHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get content children by type", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		childType, ok := tools.GetStringArg(args, "childType")
		if !ok {
			return nil, fmt.Errorf("missing or invalid childType parameter")
		}

		expand := tools.GetStringSliceArg(args, "expand")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		orderBy, _ := tools.GetStringArg(args, "orderBy")

		return h.client.GetContentChildrenByType(contentID, childType, expand, start, limit, orderBy)
	})
}

// getContentCommentsHandler handles getting content comments
func (h *Handler) getContentCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get content comments", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		expand := tools.GetStringSliceArg(args, "expand")

		parentVersion, _ := tools.GetStringArg(args, "parentVersion")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		return h.client.GetContentComments(contentID, expand, parentVersion, start, limit)
	})
}

// AddChildrenTools registers the children-related tools with the MCP server
func AddChildrenTools(server *mcp.Server, client *confluence.ConfluenceClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content_children",
		Description: "Get children of a specific Confluence content item. This tool allows you to retrieve direct child pages or other content types of a parent content item.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "The ID of the parent content item.",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
				"parentVersion": {
					Type:        "string",
					Description: "The version of the parent content to retrieve children for. Default: \"\" (latest)",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.getContentChildrenHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content_children_by_type",
		Description: "Get content children filtered by a specific type (e.g., 'page', 'comment', 'attachment').",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "The ID of the parent content item.",
				},
				"childType": {
					Type:        "string",
					Description: "The type of child content to retrieve (e.g., 'page', 'comment', 'attachment').",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned children. Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of children to return. Default: 25, Max: 100",
				},
				"orderBy": {
					Type:        "string",
					Description: "Field to order results by.",
				},
			},
			Required: []string{"contentID", "childType"},
		},
	}, handler.getContentChildrenByTypeHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content_comments",
		Description: "Get comments for a specific Confluence content item.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "The ID of the content item to retrieve comments for.",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
				"parentVersion": {
					Type:        "string",
					Description: "The version of the parent content to retrieve comments for. Default: \"\" (latest)",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned comments. Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of comments to return. Default: 25, Max: 100",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.getContentCommentsHandler)
}