package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCommentsHandler handles getting comments for a Jira issue
func (h *Handler) getCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetCommentsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	comments, err := h.client.GetComments(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get comments failed: %w", err)
	}

	return nil, comments, nil
}

// addCommentHandler handles adding a comment to a Jira issue
func (h *Handler) addCommentHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.AddCommentInput) (*mcp.CallToolResult, types.MapOutput, error) {
	comment, err := h.client.AddComment(input)
	if err != nil {
		return nil, nil, fmt.Errorf("add comment failed: %w", err)
	}

	return nil, comment, nil
}

// AddCommentTools registers the comment-related tools with the MCP server
func AddCommentTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[jira.GetCommentsInput, types.MapOutput](server, "jira_get_comments", "Get comments for a Jira issue", handler.getCommentsHandler)

	// Only register write tools if write permission is enabled
	if permissions["jira_add_comment"] {
		utils.RegisterTool[jira.AddCommentInput, types.MapOutput](server, "jira_add_comment", "Add a comment to a Jira issue", handler.addCommentHandler)
	}
}
