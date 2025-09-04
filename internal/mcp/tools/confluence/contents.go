// Package confluence provides MCP tools for interacting with Confluence.
package confluence

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getContentHandler handles getting Confluence content
func (h *Handler) getContentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get content", func() (interface{}, error) {
		typeParam, _ := tools.GetStringArg(args, "type")
		spaceKey, _ := tools.GetStringArg(args, "spaceKey")
		title, _ := tools.GetStringArg(args, "title")
		status := tools.GetStringSliceArg(args, "status")
		postingDay, _ := tools.GetStringArg(args, "postingDay")
		expand := tools.GetStringSliceArg(args, "expand")
		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 0)

		return h.client.GetContent(typeParam, spaceKey, title, status, postingDay, expand, start, limit)
	})
}

// searchContentHandler handles searching Confluence content
func (h *Handler) searchContentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("search content", func() (interface{}, error) {
		cql, ok := tools.GetStringArg(args, "cql")
		if !ok {
			return nil, fmt.Errorf("missing or invalid cql parameter")
		}

		cqlcontext, _ := tools.GetStringArg(args, "cqlcontext")
		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 0)
		expand := tools.GetStringSliceArg(args, "expand")

		return h.client.SearchContent(cql, cqlcontext, start, limit, expand)
	})
}

// getContentByIDHandler handles getting Confluence content by ID
func (h *Handler) getContentByIDHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get content by ID", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		expand := tools.GetStringSliceArg(args, "expand")

		return h.client.GetContentByID(contentID, expand)
	})
}

// createContentHandler handles creating Confluence content
func (h *Handler) createContentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("create content", func() (interface{}, error) {
		contentType, ok := tools.GetStringArg(args, "contentType")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentType parameter")
		}

		title, ok := tools.GetStringArg(args, "title")
		if !ok {
			return nil, fmt.Errorf("missing or invalid title parameter")
		}

		spaceData, ok := args["space"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("missing or invalid space parameter")
		}

		body, ok := args["body"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("missing or invalid body parameter")
		}

		// Optional parameters
		var ancestors []map[string]interface{}
		if anc, ok := args["ancestors"].([]interface{}); ok {
			ancestors = make([]map[string]interface{}, len(anc))
			for i, a := range anc {
				if ancestorMap, ok := a.(map[string]interface{}); ok {
					ancestors[i] = ancestorMap
				}
			}
		}

		var metadata map[string]interface{}
		if meta, ok := args["metadata"].(map[string]interface{}); ok {
			metadata = meta
		}

		return h.client.CreateContent(contentType, title, spaceData, body, ancestors, metadata)
	})
}

// updateContentHandler handles updating Confluence content
func (h *Handler) updateContentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("update content", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		contentData, ok := args["contentData"].(map[string]interface{})
		if !ok {
			// If contentData is not provided, construct it from individual parameters
			contentData = map[string]interface{}{}

			// Copy all provided arguments to contentData
			for k, v := range args {
				if k != "contentID" {
					contentData[k] = v
				}
			}
		}

		return h.client.UpdateContent(contentID, contentData)
	})
}

// deleteContentHandler handles deleting Confluence content
func (h *Handler) deleteContentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("delete content", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		err := h.client.DeleteContent(contentID)
		if err != nil {
			return nil, fmt.Errorf("failed to delete content: %w", err)
		}

		return map[string]interface{}{"message": fmt.Sprintf("Successfully deleted content with ID: %s", contentID)}, nil
	})
}

// getContentHistoryHandler handles getting Confluence content history
func (h *Handler) getContentHistoryHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get content history", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		expand := tools.GetStringSliceArg(args, "expand")

		return h.client.GetContentHistory(contentID, expand)
	})
}

// getCommentsHandler handles getting comments for Confluence content
func (h *Handler) getCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get comments", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		expand := tools.GetStringSliceArg(args, "expand")
		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		return h.client.GetComments(contentID, expand, start, limit)
	})
}

// addCommentHandler handles adding a comment to Confluence content
func (h *Handler) addCommentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("add comment", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		body, ok := args["body"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("missing or invalid body parameter")
		}

		// 构造符合客户端方法要求的评论正文
		commentBody := ""
		if storage, ok := body["storage"].(map[string]interface{}); ok {
			if value, ok := storage["value"].(string); ok {
				commentBody = value
			}
		}

		return h.client.AddComment(contentID, commentBody)
	})
}

// getAttachmentsHandler handles getting attachments for Confluence content
func (h *Handler) getAttachmentsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get attachments", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		expand := tools.GetStringSliceArg(args, "expand")
		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)
		filename, _ := tools.GetStringArg(args, "filename")
		mediaType, _ := tools.GetStringArg(args, "mediaType")

		return h.client.GetAttachments(contentID, expand, start, limit, filename, mediaType)
	})
}

// getExtractedTextHandler handles getting extracted text from Confluence attachment
func (h *Handler) getExtractedTextHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get extracted text", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		attachmentID, ok := tools.GetStringArg(args, "attachmentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid attachmentID parameter")
		}

		return h.client.GetExtractedText(contentID, attachmentID)
	})
}

// getContentLabelsHandler handles getting labels for Confluence content
func (h *Handler) getContentLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get content labels", func() (interface{}, error) {
		contentID, ok := tools.GetStringArg(args, "contentID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid contentID parameter")
		}

		return h.client.GetContentLabels(contentID)
	})
}

// scanContentBySpaceKeyHandler handles scanning Confluence content by space key
func (h *Handler) scanContentBySpaceKeyHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("scan content by space key", func() (interface{}, error) {
		typeParam, _ := tools.GetStringArg(args, "type")
		spaceKey, ok := tools.GetStringArg(args, "spaceKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid spaceKey parameter")
		}

		status := tools.GetStringSliceArg(args, "status")
		postingDay, _ := tools.GetStringArg(args, "postingDay")
		expand := tools.GetStringSliceArg(args, "expand")
		cursor, _ := tools.GetStringArg(args, "cursor")
		limit := tools.GetIntArg(args, "limit", 0)

		return h.client.ScanContentBySpaceKey(typeParam, spaceKey, status, postingDay, expand, cursor, limit)
	})
}

// searchHandler handles searching Confluence content using the Search API
func (h *Handler) searchHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("search", func() (interface{}, error) {
		cql, ok := tools.GetStringArg(args, "cql")
		if !ok {
			return nil, fmt.Errorf("missing or invalid cql parameter")
		}

		cqlcontext, _ := tools.GetStringArg(args, "cqlcontext")
		excerpt, _ := tools.GetStringArg(args, "excerpt")
		expand := tools.GetStringSliceArg(args, "expand")
		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 0)
		includeArchivedSpaces := tools.GetBoolArg(args, "includeArchivedSpaces", false)

		return h.client.Search(cql, cqlcontext, excerpt, expand, start, limit, includeArchivedSpaces)
	})
}

// AddContentTools registers the content-related tools with the MCP server
func AddContentTools(server *mcp.Server, client *confluence.ConfluenceClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content",
		Description: "Get Confluence content with various filter options. This tool allows you to retrieve content items with support for filtering by type, space, title, status, and more.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"type": {
					Type:        "string",
					Description: "Type of content to retrieve (e.g., 'page', 'blogpost', 'comment').",
				},
				"spaceKey": {
					Type:        "string",
					Description: "Key of the space to retrieve content from (e.g., 'DEV', 'TEAM').",
				},
				"title": {
					Type:        "string",
					Description: "Title of the content to retrieve.",
				},
				"status": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of status values to filter by (e.g., ['current', 'draft', 'trashed']).",
				},
				"postingDay": {
					Type:        "string",
					Description: "Posting day to filter by (format: 'YYYY-MM-DD').",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned contents. Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of contents to return. Default: 25, Max: 100",
				},
			},
		},
	}, handler.getContentHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_search_content",
		Description: "Search for Confluence content using CQL (Confluence Query Language). This tool allows you to find content based on various criteria such as text, space, labels, and more.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"cql": {
					Type:        "string",
					Description: "CQL query string to filter content. Example: 'type=page AND space=DEV'",
				},
				"cqlcontext": {
					Type:        "string",
					Description: "Context for the CQL query.",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned content (for pagination). Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of content items to return (for pagination). Default: 25, Max: 100",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
			},
			Required: []string{"cql"},
		},
	}, handler.searchContentHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content_by_id",
		Description: "Get a specific Confluence content item by its ID. This tool allows you to retrieve detailed information about a content item including its body, metadata, and version history.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "The ID of the content item to retrieve.",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version', 'history']).",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.getContentByIDHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_create_content",
		Description: "Create new Confluence content such as pages or blog posts. This tool allows you to create content with a title, body, space, and optional metadata or parent page relationships.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentType": {
					Type:        "string",
					Description: "Type of content to create (e.g., 'page', 'blogpost').",
				},
				"title": {
					Type:        "string",
					Description: "Title of the content to create.",
				},
				"space": {
					Type:        "object",
					Description: "Space object containing the key of the space to create content in.",
					Properties: map[string]*jsonschema.Schema{
						"key": {
							Type:        "string",
							Description: "Key of the space to create content in (e.g., 'DEV').",
						},
					},
					Required: []string{"key"},
				},
				"body": {
					Type:        "object",
					Description: "Body content of the page/blog post.",
				},
				"ancestors": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "object"},
					Description: "Array of ancestor content (for creating child pages).",
				},
				"metadata": {
					Type:        "object",
					Description: "Metadata for the content.",
				},
			},
			Required: []string{"contentType", "title", "space", "body"},
		},
	}, handler.createContentHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_update_content",
		Description: "Update existing Confluence content. This tool allows you to modify various aspects of existing content such as title, body, and other properties.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "ID of the content to update.",
				},
				"contentData": {
					Type:        "object",
					Description: "Complete content data to update. Can include type, title, body, status, ancestors, etc.",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.updateContentHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_delete_content",
		Description: "Delete Confluence content by ID. This tool allows you to permanently remove content from Confluence.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "ID of the content to delete.",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.deleteContentHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content_history",
		Description: "Retrieve the history of a Confluence content item. This tool provides detailed information about all versions of a content item.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "ID of the content to get history for.",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.getContentHistoryHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content_comments",
		Description: "Retrieve comments on Confluence content. This tool allows you to get detailed information about all comments made on a specific content item.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "ID of the content to get comments for.",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned comments (for pagination). Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of comments to return (for pagination). Default: 25, Max: 100",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.getCommentsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_create_content_comment",
		Description: "Create a new comment on Confluence content. This tool allows you to add comments to existing content items.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "ID of the content to add a comment to.",
				},
				"body": {
					Type:        "object",
					Description: "Body content of the comment with storage representation.",
					Properties: map[string]*jsonschema.Schema{
						"storage": {
							Type:        "object",
							Description: "Storage representation of the comment body.",
							Properties: map[string]*jsonschema.Schema{
								"value": {
									Type:        "string",
									Description: "The actual comment text content.",
								},
								"representation": {
									Type:        "string",
									Description: "Representation format, typically 'storage'.",
								},
							},
							Required: []string{"value", "representation"},
						},
					},
					Required: []string{"storage"},
				},
			},
			Required: []string{"contentID", "body"},
		},
	}, handler.addCommentHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_attachments",
		Description: "Get attachments for Confluence content. This tool allows you to retrieve all attachments associated with a specific piece of content.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "ID of the content to get attachments for.",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned attachments (for pagination). Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of attachments to return (for pagination). Default: 25, Max: 100",
				},
				"filename": {
					Type:        "string",
					Description: "Filename to filter attachments by.",
				},
				"mediaType": {
					Type:        "string",
					Description: "Media type to filter attachments by.",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.getAttachmentsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_extracted_text",
		Description: "Get extracted text from a Confluence attachment. This tool allows you to retrieve the text content of an attachment.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "ID of the content containing the attachment.",
				},
				"attachmentID": {
					Type:        "string",
					Description: "ID of the attachment to extract text from.",
				},
			},
			Required: []string{"contentID", "attachmentID"},
		},
	}, handler.getExtractedTextHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_get_content_labels",
		Description: "Get labels for Confluence content. This tool allows you to retrieve all labels associated with a specific piece of content.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"contentID": {
					Type:        "string",
					Description: "ID of the content to get labels for.",
				},
			},
			Required: []string{"contentID"},
		},
	}, handler.getContentLabelsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_scan_content_by_space_key",
		Description: "Scan Confluence content by space key. This tool allows you to retrieve all content within a specific space.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"type": {
					Type:        "string",
					Description: "Type of content to scan for (e.g., 'page', 'blogpost').",
				},
				"spaceKey": {
					Type:        "string",
					Description: "Key of the space to scan content in (e.g., 'DEV').",
				},
				"status": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of status values to filter by (e.g., ['current', 'draft']).",
				},
				"postingDay": {
					Type:        "string",
					Description: "Posting day to filter by (format: 'YYYY-MM-DD').",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
				"cursor": {
					Type:        "string",
					Description: "Cursor for pagination.",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of content items to return (for pagination). Default: 25, Max: 100",
				},
			},
			Required: []string{"spaceKey"},
		},
	}, handler.scanContentBySpaceKeyHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "confluence_search",
		Description: "Search for Confluence content using CQL (Confluence Query Language) with additional parameters. This tool provides more search options compared to confluence_search_content.",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"cql": {
					Type:        "string",
					Description: "CQL query string to filter content. Example: 'type=page AND space=DEV'",
				},
				"cqlcontext": {
					Type:        "string",
					Description: "Context for the CQL query.",
				},
				"excerpt": {
					Type:        "string",
					Description: "The excerpt strategy to apply to the result. Valid values: 'indexed', 'highlight', 'none'.",
				},
				"expand": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Array of properties to expand in the result (e.g., ['body.storage', 'version']).",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned content (for pagination). Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of content items to return (for pagination). Default: 25, Max: 100",
				},
				"includeArchivedSpaces": {
					Type:        "boolean",
					Description: "Whether to include content from archived spaces in the results. Default: false",
				},
			},
			Required: []string{"cql"},
		},
	}, handler.searchHandler)
}
