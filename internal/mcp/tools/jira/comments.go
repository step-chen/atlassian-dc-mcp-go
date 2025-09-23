package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCommentsHandler handles getting comments for a Jira issue
func (h *Handler) getCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetCommentsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	comments, err := h.client.GetComments(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get comments")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(comments)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create comments result")
		return result, nil, err
	}

	return result, comments, nil
}

// addCommentHandler handles adding a comment to a Jira issue
func (h *Handler) addCommentHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.AddCommentInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	comment, err := h.client.AddComment(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "add comment")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(comment)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create comment result")
		return result, nil, err
	}

	return result, comment, nil
}

// AddCommentTools registers the comment-related tools with the MCP server
func AddCommentTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[jira.GetCommentsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_comments",
		Description: "Get comments for a Jira issue",
	}, handler.getCommentsHandler)

	// Only register write tools if write permission is enabled
	if permissions["jira_add_comment"] {
		mcp.AddTool[jira.AddCommentInput, map[string]interface{}](server, &mcp.Tool{
			Name:        "jira_add_comment",
			Description: "Add a comment to a Jira issue",
		}, handler.addCommentHandler)
	}
}