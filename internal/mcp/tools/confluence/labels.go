package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getRelatedLabelsHandler handles getting related labels
func (h *Handler) getRelatedLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetRelatedLabelsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	labels, err := h.client.GetRelatedLabels(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get related labels failed: %w", err)
	}

	return nil, labels, nil
}

// getLabelsHandler handles getting labels
func (h *Handler) getLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetLabelsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	labels, err := h.client.GetLabels(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get labels failed: %w", err)
	}

	return nil, labels, nil
}

// AddLabelTools registers the label-related tools with the MCP server
func AddLabelTools(server *mcp.Server, client *confluence.ConfluenceClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[confluence.GetRelatedLabelsInput, types.MapOutput](server, "confluence_get_related_labels", "Get labels related to a specific label. This tool allows you to find labels that are commonly used together with a given label.", handler.getRelatedLabelsHandler)
	utils.RegisterTool[confluence.GetLabelsInput, types.MapOutput](server, "confluence_get_labels", "Get labels with various filter options. This tool allows you to retrieve labels based on name, owner, namespace, or space.", handler.getLabelsHandler)
}
