package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

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
		result, _, err := tools.HandleToolError(err, "get attachment")
		return result, GetAttachmentOutput{}, err
	}

	result, err := tools.CreateToolResult(attachment)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create attachment result")
		return result, GetAttachmentOutput{}, err
	}

	return result, GetAttachmentOutput{Content: attachment}, nil
}

// getAttachmentMetadataHandler handles getting attachment metadata
func (h *Handler) getAttachmentMetadataHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetAttachmentMetadataInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	metadata, err := h.client.GetAttachmentMetadata(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get attachment metadata")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(metadata)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create attachment metadata result")
		return result, nil, err
	}

	return result, metadata, nil
}

// AddAttachmentTools registers the attachment-related tools with the MCP server
func AddAttachmentTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[bitbucket.GetAttachmentInput, GetAttachmentOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_attachment",
		Description: "Get a specific attachment",
	}, handler.getAttachmentHandler)

	mcp.AddTool[bitbucket.GetAttachmentMetadataInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "bitbucket_get_attachment_metadata",
		Description: "Get metadata for a specific attachment",
	}, handler.getAttachmentMetadataHandler)
}
