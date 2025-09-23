package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getBoardsHandler handles getting Jira boards
func (h *Handler) getBoardsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	boards, err := h.client.GetBoards(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get boards")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(boards)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create boards result")
		return result, nil, err
	}

	return result, boards, nil
}

// getBoardHandler handles getting a Jira board
func (h *Handler) getBoardHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	board, err := h.client.GetBoard(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get board")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(board)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create board result")
		return result, nil, err
	}

	return result, board, nil
}

// getBoardBacklogHandler handles getting backlog for a Jira board
func (h *Handler) getBoardBacklogHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardBacklogInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	backlog, err := h.client.GetBoardBacklog(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get board backlog")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(backlog)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create board backlog result")
		return result, nil, err
	}

	return result, backlog, nil
}

// getBoardEpicsHandler handles getting epics for a Jira board
func (h *Handler) getBoardEpicsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardEpicsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	epics, err := h.client.GetBoardEpics(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get board epics")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(epics)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create board epics result")
		return result, nil, err
	}

	return result, epics, nil
}

// getBoardSprintsHandler handles getting sprints for a Jira board
func (h *Handler) getBoardSprintsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetBoardSprintsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	sprints, err := h.client.GetBoardSprints(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get board sprints")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(sprints)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create board sprints result")
		return result, nil, err
	}

	return result, sprints, nil
}

// getSprintHandler handles getting a Jira sprint
func (h *Handler) getSprintHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetSprintInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	sprint, err := h.client.GetSprint(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get sprint")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(sprint)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create sprint result")
		return result, nil, err
	}

	return result, sprint, nil
}

// getSprintIssuesHandler handles getting issues for a Jira sprint
func (h *Handler) getSprintIssuesHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetSprintIssuesInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	issues, err := h.client.GetSprintIssues(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get sprint issues")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(issues)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create sprint issues result")
		return result, nil, err
	}

	return result, issues, nil
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
