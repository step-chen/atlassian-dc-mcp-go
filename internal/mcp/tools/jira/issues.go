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

// createIssueWithPayloadHandler creates a new Jira issue with a custom payload.
func (h *Handler) createIssueWithPayloadHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.CreateIssueWithPayloadInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issue, err := h.client.CreateIssueWithPayload(input)
	if err != nil {
		return nil, nil, fmt.Errorf("create issue with payload failed: %w", err)
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

// updateIssueWithOptionsHandler updates an existing Jira issue with additional options.
func (h *Handler) updateIssueWithOptionsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.UpdateIssueWithOptionsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issue, err := h.client.UpdateIssueWithOptions(input)
	if err != nil {
		return nil, nil, fmt.Errorf("update issue with options failed: %w", err)
	}

	return nil, issue, nil
}

// getAgileIssueHandler retrieves an agile Jira issue by its key.
func (h *Handler) getAgileIssueHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetAgileIssueInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issue, err := h.client.GetAgileIssue(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get agile issue failed: %w", err)
	}

	return nil, issue, nil
}

// getIssueEstimationForBoardHandler gets issue estimation for a board.
func (h *Handler) getIssueEstimationForBoardHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetIssueEstimationForBoardInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	estimation, err := h.client.GetIssueEstimationForBoard(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get issue estimation for board failed: %w", err)
	}

	return nil, estimation, nil
}

// setIssueEstimationForBoardHandler sets issue estimation for a board.
func (h *Handler) setIssueEstimationForBoardHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.SetIssueEstimationForBoardInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	estimation, err := h.client.SetIssueEstimationForBoard(input)
	if err != nil {
		return nil, nil, fmt.Errorf("set issue estimation for board failed: %w", err)
	}

	return nil, estimation, nil
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
	utils.RegisterTool[jira.GetAgileIssueInput, map[string]interface{}](server, "jira_get_agile_issue", "Get an agile Jira issue by key or ID", handler.getAgileIssueHandler)
	utils.RegisterTool[jira.GetIssueEstimationForBoardInput, map[string]interface{}](server, "jira_get_issue_estimation_for_board", "Get issue estimation for a board", handler.getIssueEstimationForBoardHandler)
	utils.RegisterTool[jira.SetIssueEstimationForBoardInput, map[string]interface{}](server, "jira_set_issue_estimation_for_board", "Set issue estimation for a board", handler.setIssueEstimationForBoardHandler)

	if permissions["jira_create_issue"] {
		utils.RegisterTool[jira.CreateIssueInput, map[string]interface{}](server, "jira_create_issue", "Create a new Jira issue", handler.createIssueHandler)
		utils.RegisterTool[jira.CreateIssueWithPayloadInput, map[string]interface{}](server, "jira_create_issue_with_payload", "Create a new Jira issue with a custom payload", handler.createIssueWithPayloadHandler)
	}

	if permissions["jira_update_issue"] {
		utils.RegisterTool[jira.UpdateIssueInput, map[string]interface{}](server, "jira_update_issue", "Update an existing Jira issue", handler.updateIssueHandler)
		utils.RegisterTool[jira.UpdateIssueWithOptionsInput, map[string]interface{}](server, "jira_update_issue_with_options", "Update an existing Jira issue with additional options", handler.updateIssueWithOptionsHandler)
	}
}
