package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getRelatedLabelsHandler handles getting related labels
func (h *Handler) getRelatedLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetRelatedLabelsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	labels, err := h.client.GetRelatedLabels(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get related labels failed: %w", err)
	}

	return nil, labels, nil
}

// getLabelsHandler handles getting labels
func (h *Handler) getLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetLabelsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	labels, err := h.client.GetLabels(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get labels failed: %w", err)
	}

	return nil, labels, nil
}

// AddLabelTools registers the label-related tools with the MCP server
func AddLabelTools(server *mcp.Server, client *confluence.ConfluenceClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[confluence.GetRelatedLabelsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_related_labels",
		Description: "Get labels related to a specific label. This tool allows you to find labels that are commonly used together with a given label.",
	}, handler.getRelatedLabelsHandler)

	mcp.AddTool[confluence.GetLabelsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_labels",
		Description: "Get labels with various filter options. This tool allows you to retrieve labels based on name, owner, namespace, or space.",
	}, handler.getLabelsHandler)
}