package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getBoardsHandler handles getting Jira boards
func (h *Handler) getBoardsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get boards", func() (interface{}, error) {
		startAt := tools.GetIntArg(args, "startAt", 0)
		maxResults := tools.GetIntArg(args, "maxResults", 50)

		name, _ := tools.GetStringArg(args, "name")
		projectKeyOrId, _ := tools.GetStringArg(args, "projectKeyOrId")
		boardType, _ := tools.GetStringArg(args, "boardType")

		return h.client.GetBoards(startAt, maxResults, name, projectKeyOrId, boardType)
	})
}

// getBoardHandler handles getting a Jira board
func (h *Handler) getBoardHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get board", func() (interface{}, error) {
		id, ok := args["id"].(float64)
		if !ok {
			return nil, fmt.Errorf("missing or invalid id parameter")
		}

		return h.client.GetBoard(int(id))
	})
}

// getBoardBacklogHandler handles getting backlog for a Jira board
func (h *Handler) getBoardBacklogHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get board backlog", func() (interface{}, error) {
		boardId, ok := args["boardId"].(float64)
		if !ok {
			return nil, fmt.Errorf("missing or invalid boardId parameter")
		}

		startAt := tools.GetIntArg(args, "startAt", 0)
		maxResults := tools.GetIntArg(args, "maxResults", 50)

		jql, _ := tools.GetStringArg(args, "jql")
		validateQuery := tools.GetBoolArg(args, "validateQuery", false)
		fields := tools.GetStringSliceArg(args, "fields")
		expand, _ := tools.GetStringArg(args, "expand")

		return h.client.GetBoardBacklog(int(boardId), startAt, maxResults, jql, validateQuery, fields, expand)
	})
}

// getBoardEpicsHandler handles getting epics for a Jira board
func (h *Handler) getBoardEpicsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get board epics", func() (interface{}, error) {
		boardId, ok := args["boardId"].(float64)
		if !ok {
			return nil, fmt.Errorf("missing or invalid boardId parameter")
		}

		startAt := tools.GetIntArg(args, "startAt", 0)
		maxResults := tools.GetIntArg(args, "maxResults", 50)

		done, _ := tools.GetStringArg(args, "done")

		return h.client.GetBoardEpics(int(boardId), startAt, maxResults, done)
	})
}

// getBoardSprintsHandler handles getting sprints for a Jira board
func (h *Handler) getBoardSprintsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get board sprints", func() (interface{}, error) {
		boardId, ok := args["boardId"].(float64)
		if !ok {
			return nil, fmt.Errorf("missing or invalid boardId parameter")
		}

		startAt := tools.GetIntArg(args, "startAt", 0)
		maxResults := tools.GetIntArg(args, "maxResults", 50)

		state, _ := tools.GetStringArg(args, "state")

		return h.client.GetBoardSprints(int(boardId), startAt, maxResults, state)
	})
}

// getSprintHandler handles getting a Jira sprint
func (h *Handler) getSprintHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get sprint", func() (interface{}, error) {
		sprintId, ok := args["sprintId"].(float64)
		if !ok {
			return nil, fmt.Errorf("missing or invalid sprintId parameter")
		}

		return h.client.GetSprint(int(sprintId))
	})
}

// getSprintIssuesHandler handles getting issues for a Jira sprint
func (h *Handler) getSprintIssuesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get sprint issues", func() (interface{}, error) {
		sprintId, ok := args["sprintId"].(float64)
		if !ok {
			return nil, fmt.Errorf("missing or invalid sprintId parameter")
		}

		startAt := tools.GetIntArg(args, "startAt", 0)
		maxResults := tools.GetIntArg(args, "maxResults", 50)

		jql, _ := tools.GetStringArg(args, "jql")
		validateQuery := tools.GetBoolArg(args, "validateQuery", false)
		fields := tools.GetStringSliceArg(args, "fields")
		expand, _ := tools.GetStringArg(args, "expand")

		return h.client.GetSprintIssues(int(sprintId), startAt, maxResults, jql, validateQuery, fields, expand)
	})
}

// AddBoardTools registers the board-related tools with the MCP server
func AddBoardTools(server *mcp.Server, client *jira.JiraClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_boards",
		Description: "Get Jira boards with optional filters",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"startAt": {
					Type:        "integer",
					Description: "The starting index of the returned boards. Default: 0",
				},
				"maxResults": {
					Type:        "integer",
					Description: "The maximum number of boards to return per page. Default: 50",
				},
				"name": {
					Type:        "string",
					Description: "Filters results to boards that match the specified name",
				},
				"projectKeyOrId": {
					Type:        "string",
					Description: "Filters results to boards that match the specified project key or ID",
				},
				"boardType": {
					Type:        "string",
					Description: "Filters results to boards of the specified type",
				},
			},
		},
	}, handler.getBoardsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_board",
		Description: "Get a specific Jira board by its ID",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"id": {
					Type:        "integer",
					Description: "The ID of the board to retrieve",
				},
			},
			Required: []string{"id"},
		},
	}, handler.getBoardHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_board_backlog",
		Description: "Get backlog issues for a Jira board",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"boardId": {
					Type:        "integer",
					Description: "The ID of the board",
				},
				"startAt": {
					Type:        "integer",
					Description: "The starting index of the returned issues. Default: 0",
				},
				"maxResults": {
					Type:        "integer",
					Description: "The maximum number of issues to return per page. Default: 50",
				},
				"jql": {
					Type:        "string",
					Description: "JQL filter for board issues",
				},
				"validateQuery": {
					Type:        "boolean",
					Description: "Specifies whether to validate the JQL query. Default: false",
				},
				"fields": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "The list of fields to return for each issue",
				},
				"expand": {
					Type:        "string",
					Description: "A comma-separated list of parameters to expand",
				},
			},
			Required: []string{"boardId"},
		},
	}, handler.getBoardBacklogHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_board_epics",
		Description: "Get epics associated with a Jira board",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"boardId": {
					Type:        "integer",
					Description: "The ID of the board",
				},
				"startAt": {
					Type:        "integer",
					Description: "The starting index of the returned epics. Default: 0",
				},
				"maxResults": {
					Type:        "integer",
					Description: "The maximum number of epics to return per page. Default: 50",
				},
				"done": {
					Type:        "string",
					Description: "Filters results to epics that are either done or not done",
				},
			},
			Required: []string{"boardId"},
		},
	}, handler.getBoardEpicsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_board_sprints",
		Description: "Get sprints associated with a Jira board",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"boardId": {
					Type:        "integer",
					Description: "The ID of the board",
				},
				"startAt": {
					Type:        "integer",
					Description: "The starting index of the returned sprints. Default: 0",
				},
				"maxResults": {
					Type:        "integer",
					Description: "The maximum number of sprints to return per page. Default: 50",
				},
				"state": {
					Type:        "string",
					Description: "Filters results to sprints in the specified states",
				},
			},
			Required: []string{"boardId"},
		},
	}, handler.getBoardSprintsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_sprint",
		Description: "Get a specific Jira sprint by its ID",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"sprintId": {
					Type:        "integer",
					Description: "The ID of the sprint to retrieve",
				},
			},
			Required: []string{"sprintId"},
		},
	}, handler.getSprintHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_sprint_issues",
		Description: "Get issues in a specific Jira sprint",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"sprintId": {
					Type:        "integer",
					Description: "The ID of the sprint",
				},
				"startAt": {
					Type:        "integer",
					Description: "The starting index of the returned issues. Default: 0",
				},
				"maxResults": {
					Type:        "integer",
					Description: "The maximum number of issues to return per page. Default: 50",
				},
				"jql": {
					Type:        "string",
					Description: "JQL filter for sprint issues",
				},
				"validateQuery": {
					Type:        "boolean",
					Description: "Specifies whether to validate the JQL query. Default: false",
				},
				"fields": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "The list of fields to return for each issue",
				},
				"expand": {
					Type:        "string",
					Description: "A comma-separated list of parameters to expand",
				},
			},
			Required: []string{"sprintId"},
		},
	}, handler.getSprintIssuesHandler)
}