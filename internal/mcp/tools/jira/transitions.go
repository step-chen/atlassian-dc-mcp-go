package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getTransitionsHandler handles getting transitions for a Jira issue
func (h *Handler) getTransitionsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetTransitionsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	transitions, err := h.client.GetTransitions(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get transitions")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(transitions)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create transitions result")
		return result, nil, err
	}

	return result, transitions, nil
}

// transitionIssueHandler handles transitioning a Jira issue
func (h *Handler) transitionIssueHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.TransitionIssueInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	err := h.client.TransitionIssue(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "transition issue")
		return result, nil, err
	}

	resultMap := map[string]interface{}{"success": true}
	result, err := tools.CreateToolResult(resultMap)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create transition result")
		return result, nil, err
	}

	return result, resultMap, nil
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
