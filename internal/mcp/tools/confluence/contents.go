// Package confluence provides MCP tools for interacting with Confluence.
package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getContentHandler handles getting Confluence content
func (h *Handler) getContentHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	content, err := h.client.GetContent(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get content failed: %w", err)
	}

	return nil, content, nil
}

// searchContentHandler handles searching Confluence content
func (h *Handler) searchContentHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.SearchContentInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	content, err := h.client.SearchContent(input)
	if err != nil {
		return nil, nil, fmt.Errorf("search content failed: %w", err)
	}

	return nil, content, nil
}

// getContentByIDHandler handles getting Confluence content by ID
func (h *Handler) getContentByIDHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentByIDInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	content, err := h.client.GetContentByID(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get content by ID failed: %w", err)
	}

	return nil, content, nil
}

// createContentHandler handles creating Confluence content
func (h *Handler) createContentHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.CreateContentInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	content, err := h.client.CreateContent(input)
	if err != nil {
		return nil, nil, fmt.Errorf("create content failed: %w", err)
	}

	return nil, content, nil
}

// updateContentHandler handles updating Confluence content
func (h *Handler) updateContentHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.UpdateContentInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	content, err := h.client.UpdateContent(input)
	if err != nil {
		return nil, nil, fmt.Errorf("update content failed: %w", err)
	}

	return nil, content, nil
}

// deleteContentHandler handles deleting Confluence content
func (h *Handler) deleteContentHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.DeleteContentInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	err := h.client.DeleteContent(input)
	if err != nil {
		return nil, nil, fmt.Errorf("delete content failed: %w", err)
	}

	response := map[string]interface{}{
		"message": fmt.Sprintf("Successfully deleted content with ID: %s", input.ContentID),
	}

	return nil, response, nil
}

// getContentHistoryHandler handles getting Confluence content history
func (h *Handler) getContentHistoryHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentHistoryInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	history, err := h.client.GetContentHistory(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get content history failed: %w", err)
	}

	return nil, history, nil
}

// getAttachmentsHandler handles getting attachments for Confluence content
func (h *Handler) getAttachmentsHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetAttachmentsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	attachments, err := h.client.GetAttachments(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get attachments failed: %w", err)
	}

	return nil, attachments, nil
}

// getExtractedTextHandler handles getting extracted text from Confluence attachment
func (h *Handler) getExtractedTextHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetExtractedTextInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	extractedText, err := h.client.GetExtractedText(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get extracted text failed: %w", err)
	}

	return nil, extractedText, nil
}

// getContentLabelsHandler handles getting labels for Confluence content
func (h *Handler) getContentLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.GetContentLabelsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	labels, err := h.client.GetContentLabels(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get content labels failed: %w", err)
	}

	return nil, labels, nil
}

// scanContentBySpaceKeyHandler handles scanning Confluence content by space key
func (h *Handler) scanContentBySpaceKeyHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.ScanContentBySpaceKeyInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	content, err := h.client.ScanContentBySpaceKey(input)
	if err != nil {
		return nil, nil, fmt.Errorf("scan content by space key failed: %w", err)
	}

	return nil, content, nil
}

// searchHandler handles searching Confluence content using the Search API
func (h *Handler) searchHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.SearchInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	content, err := h.client.Search(input)
	if err != nil {
		return nil, nil, fmt.Errorf("search failed: %w", err)
	}

	return nil, content, nil
}

// addCommentHandler handles adding a comment to Confluence content
func (h *Handler) addCommentHandler(ctx context.Context, req *mcp.CallToolRequest, input confluence.AddCommentInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	comment, err := h.client.AddComment(input)
	if err != nil {
		return nil, nil, fmt.Errorf("add comment failed: %w", err)
	}

	return nil, comment, nil
}

// AddContentTools registers the content-related tools with the MCP server
func AddContentTools(server *mcp.Server, client *confluence.ConfluenceClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[confluence.GetContentInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_content",
		Description: "Get a list of Confluence content. This tool allows you to retrieve multiple content items with various filter options.",
	}, handler.getContentHandler)

	mcp.AddTool[confluence.SearchContentInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_search_content",
		Description: "Search for Confluence content using CQL (Confluence Query Language). This tool allows you to find content based on various criteria such as text, space, labels, and more.",
	}, handler.searchContentHandler)

	mcp.AddTool[confluence.GetContentByIDInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_content_by_id",
		Description: "Get a specific Confluence content item by its ID. This tool allows you to retrieve detailed information about a content item including its body, metadata, and version history.",
	}, handler.getContentByIDHandler)

	mcp.AddTool[confluence.GetContentHistoryInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_content_history",
		Description: "Retrieve the history of a Confluence content item. This tool provides detailed information about all versions of a content item.",
	}, handler.getContentHistoryHandler)

	mcp.AddTool[confluence.GetContentCommentsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_content_comments",
		Description: "Get comments for a specific Confluence content item. This tool allows you to retrieve all comments associated with a content item.",
	}, handler.getContentCommentsHandler)

	mcp.AddTool[confluence.GetContentLabelsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_content_labels",
		Description: "Get labels for a specific Confluence content item. This tool allows you to retrieve all labels associated with a content item.",
	}, handler.getContentLabelsHandler)

	mcp.AddTool[confluence.GetAttachmentsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_attachments",
		Description: "Get attachments for a specific Confluence content item.",
	}, handler.getAttachmentsHandler)

	mcp.AddTool[confluence.GetExtractedTextInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_get_extracted_text",
		Description: "Get extracted text from a Confluence attachment.",
	}, handler.getExtractedTextHandler)

	mcp.AddTool[confluence.ScanContentBySpaceKeyInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_scan_content_by_space_key",
		Description: "Scan Confluence content by space key.",
	}, handler.scanContentBySpaceKeyHandler)

	mcp.AddTool[confluence.SearchInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "confluence_search",
		Description: "Search Confluence using the Search API.",
	}, handler.searchHandler)

	if permissions["confluence_create_content"] {
		mcp.AddTool[confluence.CreateContentInput, map[string]interface{}](server, &mcp.Tool{
			Name:        "confluence_create_content",
			Description: "Create new Confluence content. This tool allows you to create pages, blog posts, and other content types.",
		}, handler.createContentHandler)
	}

	if permissions["confluence_update_content"] {
		mcp.AddTool[confluence.UpdateContentInput, map[string]interface{}](server, &mcp.Tool{
			Name:        "confluence_update_content",
			Description: "Update existing Confluence content. This tool allows you to modify various aspects of existing content such as title, body, and other properties.",
		}, handler.updateContentHandler)
	}

	if permissions["confluence_delete_content"] {
		mcp.AddTool[confluence.DeleteContentInput, map[string]interface{}](server, &mcp.Tool{
			Name:        "confluence_delete_content",
			Description: "Delete Confluence content by ID. This tool allows you to permanently remove content from Confluence.",
		}, handler.deleteContentHandler)
	}

	if permissions["confluence_add_comment"] {
		mcp.AddTool[confluence.AddCommentInput, map[string]interface{}](server, &mcp.Tool{
			Name:        "confluence_add_comment",
			Description: "Add a comment to Confluence content. This tool allows you to attach comments to specific content items.",
		}, handler.addCommentHandler)
	}
}
