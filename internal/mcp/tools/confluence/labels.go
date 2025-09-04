package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getRelatedLabelsHandler handles getting related labels
func (h *Handler) getRelatedLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get related labels", func() (interface{}, error) {
		labelName, ok := tools.GetStringArg(args, "labelName")
		if !ok {
			return nil, fmt.Errorf("missing or invalid labelName parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		return h.client.GetRelatedLabels(labelName, start, limit)
	})
}

// getLabelsHandler handles getting labels
func (h *Handler) getLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get labels", func() (interface{}, error) {
		labelName, _ := tools.GetStringArg(args, "labelName")
		owner, _ := tools.GetStringArg(args, "owner")
		namespace, _ := tools.GetStringArg(args, "namespace")
		spaceKey, _ := tools.GetStringArg(args, "spaceKey")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		return h.client.GetLabels(labelName, owner, namespace, spaceKey, start, limit)
	})
}

// AddLabelTools registers the label-related tools with the MCP server
func AddLabelTools(server *mcp.Server, client *confluence.ConfluenceClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_related_labels",
		Description: "Get labels related to a specific label. This tool allows you to find labels that are commonly used together with a given label.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"labelName": {
					Type:        "string",
					Description: "Name of the label to find related labels for",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned labels",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of labels to return",
				},
			},
			Required: []string{"labelName"},
		},
	}, handler.getRelatedLabelsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_labels",
		Description: "Get labels with various filter options. This tool allows you to retrieve labels based on name, owner, namespace, or space.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"labelName": {
					Type:        "string",
					Description: "Name of the label to find",
				},
				"owner": {
					Type:        "string",
					Description: "Owner of the labels",
				},
				"namespace": {
					Type:        "string",
					Description: "Namespace of the labels",
				},
				"spaceKey": {
					Type:        "string",
					Description: "Space key to filter labels by",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned labels",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of labels to return",
				},
			},
		},
	}, handler.getLabelsHandler)
}