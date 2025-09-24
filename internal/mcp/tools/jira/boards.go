package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getBoardsHandler handles getting Jira boards
func (h *Handler) getBoardsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	boards, err := h.client.GetBoards(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get boards failed: %w", err)
	}

	return nil, boards, nil
}

// getBoardHandler handles getting a Jira board
func (h *Handler) getBoardHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardInput) (*mcp.CallToolResult, types.MapOutput, error) {
	board, err := h.client.GetBoard(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get board failed: %w", err)
	}

	return nil, board, nil
}

// getBoardBacklogHandler handles getting backlog for a Jira board
func (h *Handler) getBoardBacklogHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardBacklogInput) (*mcp.CallToolResult, types.MapOutput, error) {
	backlog, err := h.client.GetBoardBacklog(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get board backlog failed: %w", err)
	}

	return nil, backlog, nil
}

// getBoardEpicsHandler handles getting epics for a Jira board
func (h *Handler) getBoardEpicsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardEpicsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	epics, err := h.client.GetBoardEpics(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get board epics failed: %w", err)
	}

	return nil, epics, nil
}

// getBoardSprintsHandler handles getting sprints for a Jira board
func (h *Handler) getBoardSprintsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardSprintsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	sprints, err := h.client.GetBoardSprints(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get board sprints failed: %w", err)
	}

	return nil, sprints, nil
}

// getSprintHandler handles getting a Jira sprint
func (h *Handler) getSprintHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetSprintInput) (*mcp.CallToolResult, types.MapOutput, error) {
	sprint, err := h.client.GetSprint(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get sprint failed: %w", err)
	}

	return nil, sprint, nil
}

// getSprintIssuesHandler handles getting issues for a Jira sprint
func (h *Handler) getSprintIssuesHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetSprintIssuesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	issues, err := h.client.GetSprintIssues(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get sprint issues failed: %w", err)
	}

	return nil, issues, nil
}

// AddBoardTools registers the board-related tools with the MCP server
func AddBoardTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[jira.GetBoardsInput, types.MapOutput](server, "jira_get_boards", "Get Jira boards with optional filters", handler.getBoardsHandler)
	utils.RegisterTool[jira.GetBoardInput, types.MapOutput](server, "jira_get_board", "Get a specific Jira board by its ID", handler.getBoardHandler)
	utils.RegisterTool[jira.GetBoardBacklogInput, types.MapOutput](server, "jira_get_board_backlog", "Get backlog issues for a Jira board", handler.getBoardBacklogHandler)
	utils.RegisterTool[jira.GetBoardEpicsInput, types.MapOutput](server, "jira_get_board_epics", "Get epics associated with a Jira board", handler.getBoardEpicsHandler)
	utils.RegisterTool[jira.GetBoardSprintsInput, types.MapOutput](server, "jira_get_board_sprints", "Get sprints associated with a Jira board", handler.getBoardSprintsHandler)
	utils.RegisterTool[jira.GetSprintInput, types.MapOutput](server, "jira_get_sprint", "Get a specific Jira sprint by its ID", handler.getSprintHandler)
	utils.RegisterTool[jira.GetSprintIssuesInput, types.MapOutput](server, "jira_get_sprint_issues", "Get issues in a specific Jira sprint", handler.getSprintIssuesHandler)
}
