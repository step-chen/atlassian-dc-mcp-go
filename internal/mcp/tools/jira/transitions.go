package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getTransitionsHandler handles getting transitions for a Jira issue
func (h *Handler) getTransitionsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetTransitionsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	transitions, err := h.client.GetTransitions(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get transitions failed: %w", err)
	}

	return nil, transitions, nil
}

// transitionIssueHandler handles transitioning a Jira issue
func (h *Handler) transitionIssueHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.TransitionIssueInput) (*mcp.CallToolResult, types.MapOutput, error) {
	err := h.client.TransitionIssue(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("transition issue failed: %w", err)
	}

	resultMap := types.MapOutput{"success": true}
	return nil, resultMap, nil
}

// AddTransitionTools registers the transition-related tools with the MCP server
func AddTransitionTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[jira.GetTransitionsInput, types.MapOutput](server, "jira_get_transitions", "Get transitions for a Jira issue", handler.getTransitionsHandler)

	if permissions["jira_transition_issue"] {
		utils.RegisterTool[jira.TransitionIssueInput, types.MapOutput](server, "jira_transition_issue", "Transition a Jira issue", handler.transitionIssueHandler)
	}
}
