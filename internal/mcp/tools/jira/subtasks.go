package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetSubtasksResult represents the result structure for getSubtasksHandler
type GetSubtasksResult struct {
	Subtasks []types.MapOutput `json:"subtasks"`
}

// getSubtasksHandler handles getting subtasks for a Jira issue
func (h *Handler) getSubtasksHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetSubtasksInput) (*mcp.CallToolResult, GetSubtasksResult, error) {
	subtasks, err := h.client.GetSubtasks(ctx, input)
	if err != nil {
		return nil, GetSubtasksResult{}, fmt.Errorf("get subtasks failed: %w", err)
	}

	result := GetSubtasksResult{
		Subtasks: subtasks,
	}

	return nil, result, nil
}

// createSubTaskHandler handles creating a subtask for a Jira issue
func (h *Handler) createSubTaskHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.CreateSubTaskInput) (*mcp.CallToolResult, types.MapOutput, error) {
	subtask, err := h.client.CreateSubTask(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("create subtask failed: %w", err)
	}

	return nil, subtask, nil
}

// AddSubtaskTools registers the subtask-related tools with the MCP server
func AddSubtaskTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[jira.GetSubtasksInput, GetSubtasksResult](server, "jira_get_subtasks", "Get subtasks for a Jira issue", handler.getSubtasksHandler)

	if permissions["jira_create_subtask"] {
		utils.RegisterTool[jira.CreateSubTaskInput, types.MapOutput](server, "jira_create_subtask", "Create a subtask for a Jira issue", handler.createSubTaskHandler)
	}
}
