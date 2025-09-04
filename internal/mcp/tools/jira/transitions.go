package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getTransitionsHandler handles getting transitions for a Jira issue
func (h *Handler) getTransitionsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get transitions", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		return h.client.GetTransitions(issueKey)
	})
}

// transitionIssueHandler handles transitioning a Jira issue
func (h *Handler) transitionIssueHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("transition issue", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		transitionID, ok := tools.GetStringArg(args, "transitionID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid transitionID parameter")
		}

		err := h.client.TransitionIssue(issueKey, transitionID)
		return map[string]interface{}{"success": true}, err
	})
}

// AddTransitionTools registers the transition-related tools with the MCP server
func AddTransitionTools(server *mcp.Server, client *jira.JiraClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_transitions",
		Description: "Get transitions for a Jira issue",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"issueKey": {
					Type:        "string",
					Description: "The key of the issue to get transitions for",
				},
			},
			Required: []string{"issueKey"},
		},
	}, handler.getTransitionsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_transition_issue",
		Description: "Transition a Jira issue",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"issueKey": {
					Type:        "string",
					Description: "The key of the issue to transition",
				},
				"transitionID": {
					Type:        "string",
					Description: "The ID of the transition to apply",
				},
			},
			Required: []string{"issueKey", "transitionID"},
		},
	}, handler.transitionIssueHandler)
}