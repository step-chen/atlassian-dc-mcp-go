package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetAttachmentInput represents the input parameters for getting an attachment
type GetAttachmentInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required, the project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required, the repository slug"`
	AttachmentId string `json:"attachmentId" jsonschema:"required, the attachment ID"`
}

// GetAttachmentMetadataInput represents the input parameters for getting attachment metadata
type GetAttachmentMetadataInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required, the project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required, the repository slug"`
	AttachmentId string `json:"attachmentId" jsonschema:"required, the attachment ID"`
}

// AttachmentOutput represents the output for an attachment
type AttachmentOutput struct {
	Content []byte `json:"content" jsonschema:"the attachment content"`
}

// getAttachmentHandler handles getting an attachment
func (h *Handler) getAttachmentHandler(ctx context.Context, req *mcp.CallToolRequest, input GetAttachmentInput) (*mcp.CallToolResult, AttachmentOutput, error) {
	attachment, err := h.client.GetAttachment(input.ProjectKey, input.RepoSlug, input.AttachmentId)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get attachment")
		return result, AttachmentOutput{}, err
	}

	result, err := tools.CreateToolResult(attachment)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create attachment result")
		return result, AttachmentOutput{}, err
	}

	return result, AttachmentOutput{Content: attachment}, nil
}

// getAttachmentMetadataHandler handles getting attachment metadata
func (h *Handler) getAttachmentMetadataHandler(ctx context.Context, req *mcp.CallToolRequest, input GetAttachmentMetadataInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	metadata, err := h.client.GetAttachmentMetadata(input.ProjectKey, input.RepoSlug, input.AttachmentId)
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

	mcp.AddTool[GetAttachmentInput, AttachmentOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_attachment",
		Description: "Get a specific attachment",
	}, handler.getAttachmentHandler)

	mcp.AddTool[GetAttachmentMetadataInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "bitbucket_get_attachment_metadata",
		Description: "Get metadata for a specific attachment",
	}, handler.getAttachmentMetadataHandler)
}
