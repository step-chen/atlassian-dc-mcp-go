package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getTransitionsHandler handles getting transitions for a Jira issue
func (h *Handler) getTransitionsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetTransitionsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	transitions, err := h.client.GetTransitions(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get transitions failed: %w", err)
	}

	return nil, transitions, nil
}

// transitionIssueHandler handles transitioning a Jira issue
func (h *Handler) transitionIssueHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.TransitionIssueInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	err := h.client.TransitionIssue(input)
	if err != nil {
		return nil, nil, fmt.Errorf("transition issue failed: %w", err)
	}

	resultMap := map[string]interface{}{"success": true}
	return nil, resultMap, nil
}

// AddTransitionTools registers the transition-related tools with the MCP server
func AddTransitionTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[jira.GetTransitionsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_transitions",
		Description: "Get transitions for a Jira issue",
	}, handler.getTransitionsHandler)

	if permissions["jira_transition_issue"] {
		mcp.AddTool[jira.TransitionIssueInput, map[string]interface{}](server, &mcp.Tool{
			Name:        "jira_transition_issue",
			Description: "Transition a Jira issue",
		}, handler.transitionIssueHandler)
	}
}
