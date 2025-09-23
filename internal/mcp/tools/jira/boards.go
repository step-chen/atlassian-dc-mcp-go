package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getBoardsHandler handles getting Jira boards
func (h *Handler) getBoardsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	boards, err := h.client.GetBoards(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get boards failed: %w", err)
	}

	return nil, boards, nil
}

// getBoardHandler handles getting a Jira board
func (h *Handler) getBoardHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	board, err := h.client.GetBoard(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get board failed: %w", err)
	}

	return nil, board, nil
}

// getBoardBacklogHandler handles getting backlog for a Jira board
func (h *Handler) getBoardBacklogHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardBacklogInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	backlog, err := h.client.GetBoardBacklog(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get board backlog failed: %w", err)
	}

	return nil, backlog, nil
}

// getBoardEpicsHandler handles getting epics for a Jira board
func (h *Handler) getBoardEpicsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardEpicsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	epics, err := h.client.GetBoardEpics(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get board epics failed: %w", err)
	}

	return nil, epics, nil
}

// getBoardSprintsHandler handles getting sprints for a Jira board
func (h *Handler) getBoardSprintsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardSprintsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	sprints, err := h.client.GetBoardSprints(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get board sprints failed: %w", err)
	}

	return nil, sprints, nil
}

// getSprintHandler handles getting a Jira sprint
func (h *Handler) getSprintHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetSprintInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	sprint, err := h.client.GetSprint(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get sprint failed: %w", err)
	}

	return nil, sprint, nil
}

// getSprintIssuesHandler handles getting issues for a Jira sprint
func (h *Handler) getSprintIssuesHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetSprintIssuesInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issues, err := h.client.GetSprintIssues(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get sprint issues failed: %w", err)
	}

	return nil, issues, nil
}

// AddBoardTools registers the board-related tools with the MCP server
func AddBoardTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[jira.GetBoardsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_boards",
		Description: "Get Jira boards with optional filters",
	}, handler.getBoardsHandler)

	mcp.AddTool[jira.GetBoardInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_board",
		Description: "Get a specific Jira board by its ID",
	}, handler.getBoardHandler)

	mcp.AddTool[jira.GetBoardBacklogInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_board_backlog",
		Description: "Get backlog issues for a Jira board",
	}, handler.getBoardBacklogHandler)

	mcp.AddTool[jira.GetBoardEpicsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_board_epics",
		Description: "Get epics associated with a Jira board",
	}, handler.getBoardEpicsHandler)

	mcp.AddTool[jira.GetBoardSprintsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_board_sprints",
		Description: "Get sprints associated with a Jira board",
	}, handler.getBoardSprintsHandler)

	mcp.AddTool[jira.GetSprintInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_sprint",
		Description: "Get a specific Jira sprint by its ID",
	}, handler.getSprintHandler)

	mcp.AddTool[jira.GetSprintIssuesInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_sprint_issues",
		Description: "Get issues in a specific Jira sprint",
	}, handler.getSprintIssuesHandler)
}
