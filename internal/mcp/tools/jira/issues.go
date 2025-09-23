// Package jira provides MCP tools for interacting with Jira.
package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"

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

	mcp.AddTool[jira.GetIssueInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_issue",
		Description: "Get a specific Jira issue by its key with default fields",
	}, handler.getIssueHandler)

	mcp.AddTool[jira.CreateIssueInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_create_issue",
		Description: "Create a new Jira issue",
	}, handler.createIssueHandler)

	mcp.AddTool[jira.SearchIssuesInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_search_issues",
		Description: "Search for Jira issues using JQL",
	}, handler.searchIssuesHandler)
}
