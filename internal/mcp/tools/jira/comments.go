package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCommentsHandler handles getting comments for a Jira issue
func (h *Handler) getCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get comments", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		startAt := tools.GetIntArg(args, "startAt", 0)
		maxResults := tools.GetIntArg(args, "maxResults", 50)

		expand, _ := tools.GetStringArg(args, "expand")
		orderBy, _ := tools.GetStringArg(args, "orderBy")

		return h.client.GetComments(issueKey, startAt, maxResults, expand, orderBy)
	})
}

// addCommentHandler handles adding a comment to a Jira issue
func (h *Handler) addCommentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("add comment", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		comment, ok := tools.GetStringArg(args, "comment")
		if !ok {
			return nil, fmt.Errorf("missing or invalid comment parameter")
		}

		return h.client.AddComment(issueKey, comment)
	})
}

// AddCommentTools registers the comment-related tools with the MCP server
func AddCommentTools(server *mcp.Server, client *jira.JiraClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_comments",
		Description: "Get comments for a Jira issue",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"issueKey": {
					Type:        "string",
					Description: "The key of the issue to get comments for",
				},
				"startAt": {
					Type:        "integer",
					Description: "The starting index of the returned comments",
				},
				"maxResults": {
					Type:        "integer",
					Description: "The maximum number of comments to return",
				},
				"expand": {
					Type:        "string",
					Description: "Fields to expand in the response",
				},
				"orderBy": {
					Type:        "string",
					Description: "Field to order comments by",
				},
			},
			Required: []string{"issueKey"},
		},
	}, handler.getCommentsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_add_comment",
		Description: "Add a comment to a Jira issue",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"issueKey": {
					Type:        "string",
					Description: "The key of the issue to add a comment to",
				},
				"comment": {
					Type:        "string",
					Description: "The text content of the comment to add",
				},
			},
			Required: []string{"issueKey", "comment"},
		},
	}, handler.addCommentHandler)
}