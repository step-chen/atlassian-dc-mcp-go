package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getAttachmentHandler handles getting an attachment
func (h *Handler) getAttachmentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get attachment", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		attachmentId, ok := tools.GetStringArg(args, "attachmentId")
		if !ok {
			return nil, fmt.Errorf("missing or invalid attachmentId parameter")
		}

		attachment, err := h.client.GetAttachment(projectKey, repoSlug, attachmentId)
		if err != nil {
			return nil, err
		}

		return attachment, nil
	})
}

// getAttachmentMetadataHandler handles getting attachment metadata
func (h *Handler) getAttachmentMetadataHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get attachment metadata", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		attachmentId, ok := tools.GetStringArg(args, "attachmentId")
		if !ok {
			return nil, fmt.Errorf("missing or invalid attachmentId parameter")
		}

		metadata, err := h.client.GetAttachmentMetadata(projectKey, repoSlug, attachmentId)
		if err != nil {
			return nil, err
		}

		return metadata, nil
	})
}

// AddAttachmentTools registers the attachment-related tools with the MCP server
func AddAttachmentTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_attachment",
		Description: "Get a specific attachment",
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
				"attachmentId": {
					Type:        "string",
					Description: "The attachment ID",
				},
			},
			Required: []string{"projectKey", "repoSlug", "attachmentId"},
		},
	}, handler.getAttachmentHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_attachment_metadata",
		Description: "Get metadata for a specific attachment",
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
				"attachmentId": {
					Type:        "string",
					Description: "The attachment ID",
				},
			},
			Required: []string{"projectKey", "repoSlug", "attachmentId"},
		},
	}, handler.getAttachmentMetadataHandler)
}
