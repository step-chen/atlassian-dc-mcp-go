// Package jira provides MCP tools for interacting with Jira.
package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/utils"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getIssueHandler retrieves a Jira issue by its key with default fields.
func (h *Handler) getIssueHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetIssueInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issue, err := h.client.GetIssue(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get issue failed: %w", err)
	}

	return nil, issue, nil
}

// createIssueHandler creates a new Jira issue.
func (h *Handler) createIssueHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.CreateIssueInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issue, err := h.client.CreateIssue(input)
	if err != nil {
		return nil, nil, fmt.Errorf("create issue failed: %w", err)
	}

	return nil, issue, nil
}

// updateIssueHandler updates an existing Jira issue.
func (h *Handler) updateIssueHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.UpdateIssueInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issue, err := h.client.UpdateIssue(input)
	if err != nil {
		return nil, nil, fmt.Errorf("update issue failed: %w", err)
	}

	return nil, issue, nil
}

// searchIssuesHandler searches for Jira issues using a JQL query.
func (h *Handler) searchIssuesHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.SearchIssuesInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issues, err := h.client.SearchIssues(input)
	if err != nil {
		return nil, nil, fmt.Errorf("search issues failed: %w", err)
	}

	return nil, issues, nil
}

// AddIssueTools registers the issue-related tools with the MCP server
func AddIssueTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[jira.SearchIssuesInput, map[string]interface{}](server, "jira_search_issues", "Search for Jira issues using JQL", handler.searchIssuesHandler)
	utils.RegisterTool[jira.GetIssueInput, map[string]interface{}](server, "jira_get_issue", "Get a specific Jira issue by key or ID", handler.getIssueHandler)

	if permissions["jira_create_issue"] {
		utils.RegisterTool[jira.CreateIssueInput, map[string]interface{}](server, "jira_create_issue", "Create a new Jira issue", handler.createIssueHandler)
	}

	if permissions["jira_update_issue"] {
		utils.RegisterTool[jira.UpdateIssueInput, map[string]interface{}](server, "jira_update_issue", "Update an existing Jira issue", handler.updateIssueHandler)
	}
}
