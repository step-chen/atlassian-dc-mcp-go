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
func AddChildrenTools(server *mcp.Server, client *confluence.ConfluenceClient, hasWritePermission bool) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content_children",
		Description: "Get content children",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "The content ID",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Fields to expand in the results",
				},
				"parentVersion": {
					Type:        "string",
					Description: "The parent version",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.getContentChildrenHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content_children_by_type",
		Description: "Get content children by type",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "The content ID",
				},
				"childType": {
					Type:        "string",
					Description: "The child type",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Fields to expand in the results",
				},
				"start": {
					Type:        "integer",
					Description: "The start index for pagination",
				},
				"limit": {
					Type:        "integer",
					Description: "The maximum number of results to return",
				},
				"orderBy": {
					Type:        "string",
					Description: "The order by field",
				},
			},
			Required: []string{"contentID", "childType"},
		},
	}, handler.getContentChildrenByTypeHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content_comments",
		Description: "Get content comments",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "The content ID",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Fields to expand in the results",
				},
				"parentVersion": {
					Type:        "string",
					Description: "The parent version",
				},
				"start": {
					Type:        "integer",
					Description: "The start index for pagination",
				},
				"limit": {
					Type:        "integer",
					Description: "The maximum number of results to return",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.getContentCommentsHandler)
}