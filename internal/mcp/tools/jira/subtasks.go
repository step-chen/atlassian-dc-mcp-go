package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getSubtasksHandler handles getting subtasks for a Jira issue
func (h *Handler) getSubtasksHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetSubtasksInput) (*mcp.CallToolResult, []map[string]interface{}, error) {
	subtasks, err := h.client.GetSubtasks(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get subtasks")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(subtasks)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create subtasks result")
		return result, nil, err
	}

	return result, subtasks, nil
}

// createSubTaskHandler handles creating a subtask for a Jira issue
func (h *Handler) createSubTaskHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.CreateSubTaskInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	subtask, err := h.client.CreateSubTask(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create subtask")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(subtask)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create subtask result")
		return result, nil, err
	}

	return result, subtask, nil
}

// AddSubtaskTools registers the subtask-related tools with the MCP server
func AddSubtaskTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[jira.GetSubtasksInput, []map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_subtasks",
		Description: "Get subtasks for a Jira issue",
	}, handler.getSubtasksHandler)

	if permissions["jira_create_subtask"] {
		mcp.AddTool[jira.CreateSubTaskInput, map[string]interface{}](server, &mcp.Tool{
			Name:        "jira_create_subtask",
			Description: "Create a subtask for a Jira issue",
		}, handler.createSubTaskHandler)
	}
}