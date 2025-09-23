// Package jira provides MCP tools for interacting with Jira.
package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getIssueHandler retrieves a Jira issue by its key with default fields.
func (h *Handler) getIssueHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetIssueInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issue, err := h.client.GetIssue(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get issue")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(issue)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create issue result")
		return result, nil, err
	}

	return result, issue, nil
}

// createIssueHandler creates a new Jira issue.
func (h *Handler) createIssueHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.CreateIssueInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issue, err := h.client.CreateIssue(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create issue")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(issue)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create issue result")
		return result, nil, err
	}

	return result, issue, nil
}

// searchIssuesHandler searches for Jira issues using a JQL query.
func (h *Handler) searchIssuesHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.SearchIssuesInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issues, err := h.client.SearchIssues(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "search issues")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(issues)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create search issues result")
		return result, nil, err
	}

	return result, issues, nil
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
