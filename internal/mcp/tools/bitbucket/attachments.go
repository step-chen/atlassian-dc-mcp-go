package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/utils"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetAttachmentOutput represents the output for an attachment
type GetAttachmentOutput struct {
	Content []byte `json:"content" jsonschema:"the attachment content"`
}

// getAttachmentHandler handles getting an attachment
func (h *Handler) getAttachmentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetAttachmentInput) (*mcp.CallToolResult, GetAttachmentOutput, error) {
	attachment, err := h.client.GetAttachment(input)
	if err != nil {
		return nil, GetAttachmentOutput{}, fmt.Errorf("get attachment failed: %w", err)
	}

	return nil, GetAttachmentOutput{Content: attachment}, nil
}

// getAttachmentMetadataHandler handles getting attachment metadata
func (h *Handler) getAttachmentMetadataHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetAttachmentMetadataInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	metadata, err := h.client.GetAttachmentMetadata(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get attachment metadata failed: %w", err)
	}

	return nil, metadata, nil
}

// createAttachmentHandler handles creating an attachment
func (h *Handler) createAttachmentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.CreateAttachmentInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	attachment, err := h.client.CreateAttachment(input)
	if err != nil {
		return nil, nil, fmt.Errorf("create attachment failed: %w", err)
	}

	return nil, attachment, nil
}

// deleteAttachmentHandler handles deleting an attachment
func (h *Handler) deleteAttachmentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.DeleteAttachmentInput) (*mcp.CallToolResult, interface{}, error) {
	err := h.client.DeleteAttachment(input)
	if err != nil {
		return nil, nil, fmt.Errorf("delete attachment failed: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: "Successfully deleted attachment",
			},
		},
	}, nil, nil
}

// AddAttachmentTools registers the attachment-related tools with the MCP server
func AddAttachmentTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[bitbucket.GetAttachmentInput, GetAttachmentOutput](server, "bitbucket_get_attachment", "Get a specific attachment", handler.getAttachmentHandler)
	utils.RegisterTool[bitbucket.GetAttachmentMetadataInput, map[string]interface{}](server, "bitbucket_get_attachment_metadata", "Get metadata for a specific attachment", handler.getAttachmentMetadataHandler)

	if permissions["bitbucket_create_attachment"] {
		utils.RegisterTool[bitbucket.CreateAttachmentInput, map[string]interface{}](server, "bitbucket_create_attachment", "Create a new attachment", handler.createAttachmentHandler)
	}

	if permissions["bitbucket_delete_attachment"] {
		utils.RegisterTool[bitbucket.DeleteAttachmentInput, interface{}](server, "bitbucket_delete_attachment", "Delete an attachment", handler.deleteAttachmentHandler)
	}
}
