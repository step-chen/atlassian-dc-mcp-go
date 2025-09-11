package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getSpacesHandler handles getting Confluence spaces
func (h *Handler) getSpacesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get spaces", func() (interface{}, error) {
		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		return h.client.GetSpaces(limit, start)
	})
}

// getSpaceHandler handles getting a specific Confluence space
func (h *Handler) getSpaceHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get space", func() (interface{}, error) {
		spaceKey, ok := tools.GetStringArg(args, "spaceKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid spaceKey parameter")
		}

		expand := tools.GetStringSliceArg(args, "expand")

		return h.client.GetSpace(spaceKey, expand)
	})
}

// getContentsInSpaceHandler handles getting contents in a specific Confluence space
func (h *Handler) getContentsInSpaceHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get contents in space", func() (interface{}, error) {
		spaceKey, ok := tools.GetStringArg(args, "spaceKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid spaceKey parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		expand := tools.GetStringSliceArg(args, "expand")

		return h.client.GetContentsInSpace(spaceKey, start, limit, expand)
	})
}

// getContentsByTypeHandler handles getting contents by type in a specific Confluence space
func (h *Handler) getContentsByTypeHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get contents by type", func() (interface{}, error) {
		spaceKey, ok := tools.GetStringArg(args, "spaceKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid spaceKey parameter")
		}

		contentType, ok := tools.GetStringArg(args, "contentType")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentType parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		expand := tools.GetStringSliceArg(args, "expand")

		return h.client.GetContentsByType(spaceKey, contentType, start, limit, expand)
	})
}

// getSpacesByKeyHandler handles getting spaces by key
func (h *Handler) getSpacesByKeyHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get spaces by key", func() (interface{}, error) {
		keys := tools.GetStringSliceArg(args, "keys")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		expand := tools.GetStringSliceArg(args, "expand")

		spaceIds := tools.GetStringSliceArg(args, "spaceIds")

		spaceKeys, _ := tools.GetStringArg(args, "spaceKeys")

		spaceId := tools.GetStringSliceArg(args, "spaceId")

		spaceKeySingle, _ := tools.GetStringArg(args, "spaceKeySingle")

		typ, _ := tools.GetStringArg(args, "type")

		status, _ := tools.GetStringArg(args, "status")

		label := tools.GetStringSliceArg(args, "label")

		contentLabel := tools.GetStringSliceArg(args, "contentLabel")

		var favourite *bool
		if val, ok := args["favourite"].(bool); ok {
			favourite = &val
		}

		var hasRetentionPolicy *bool
		if val, ok := args["hasRetentionPolicy"].(bool); ok {
			hasRetentionPolicy = &val
		}

		return h.client.GetSpacesByKey(keys, start, limit, expand, spaceIds, spaceKeys, spaceId, spaceKeySingle, typ, status, label, contentLabel, favourite, hasRetentionPolicy)
	})
}

// AddSpaceTools registers the space-related tools with the MCP server
func AddSpaceTools(server *mcp.Server, client *confluence.ConfluenceClient, hasWritePermission bool) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_spaces",
		Description: "Get a list of Confluence spaces. This tool allows you to retrieve multiple spaces with pagination support.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned spaces. Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of spaces to return. Default: 25, Max: 100",
				},
			},
		},
	}, handler.getSpacesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_space",
		Description: "Get a specific Confluence space by its key. This tool allows you to retrieve detailed information about a space including its name, description, and metadata.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"spaceKey": {
					Type:        "string",
					Description: "The key of the space to retrieve (e.g., 'DEV', 'TEAM').",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "An array of properties to expand in the result (e.g., ['description.plain', 'homepage']).",
				},
			},
			Required: []string{"spaceKey"},
		},
	}, handler.getSpaceHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_contents_in_space",
		Description: "Get contents in a specific Confluence space. This tool allows you to retrieve all content items within a space with pagination support.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"spaceKey": {
					Type:        "string",
					Description: "The key of the space to retrieve contents from (e.g., 'DEV', 'TEAM').",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned contents. Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of contents to return. Default: 25, Max: 100",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "An array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
			},
			Required: []string{"spaceKey"},
		},
	}, handler.getContentsInSpaceHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_contents_by_type",
		Description: "Get contents by type in a specific Confluence space. This tool allows you to retrieve content items of a specific type (e.g., page, blogpost) within a space.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"spaceKey": {
					Type:        "string",
					Description: "The key of the space to retrieve contents from (e.g., 'DEV', 'TEAM').",
				},
				"contentType": {
					Type:        "string",
					Description: "The type of content to retrieve (e.g., 'page', 'blogpost', 'comment').",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned contents. Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of contents to return. Default: 25, Max: 100",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "An array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
			},
			Required: []string{"spaceKey", "contentType"},
		},
	}, handler.getContentsByTypeHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_spaces_by_key",
		Description: "Get spaces by key with various filter options. This tool allows you to retrieve spaces using multiple filter criteria including keys, IDs, types, status, and labels.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"keys": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of space keys to retrieve (e.g., ['DEV', 'TEAM']).",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned spaces. Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of spaces to return. Default: 25, Max: 100",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "An array of properties to expand in the result (e.g., ['description.plain', 'homepage']).",
				},
				"spaceIds": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of space IDs to retrieve.",
				},
				"spaceKeys": {
					Type:        "string",
					Description: "Comma-separated list of space keys.",
				},
				"spaceId": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of space IDs.",
				},
				"spaceKeySingle": {
					Type:        "string",
					Description: "Single space key.",
				},
				"type": {
					Type:        "string",
					Description: "Type of spaces to retrieve (e.g., 'global', 'personal').",
				},
				"status": {
					Type:        "string",
					Description: "Status of spaces to retrieve (e.g., 'current', 'archived').",
				},
				"label": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of labels to filter by.",
				},
				"contentLabel": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of content labels to filter by.",
				},
				"favourite": {
					Type:        "boolean",
					Description: "Whether to retrieve only favourite spaces.",
				},
				"hasRetentionPolicy": {
					Type:        "boolean",
					Description: "Whether to retrieve only spaces with retention policy.",
				},
			},
		},
	}, handler.getSpacesByKeyHandler)
}