// Package jira provides MCP tools for interacting with Jira.
package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getIssueHandler retrieves a Jira issue by its key with default fields.
func (h *Handler) getIssueHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get issue", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		// Define default fields to retrieve
		fields := []string{"summary", "description", "status", "assignee", "reporter", "priority", "issuetype", "project", "created", "updated", "comment"}

		result, err := h.client.GetIssue(issueKey, fields)
		return result, err
	})
}

// getIssueWithFieldsHandler retrieves a Jira issue by its key.
func (h *Handler) getIssueWithFieldsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get issue", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		fields := tools.GetStringSliceArg(args, "fields")

		result, err := h.client.GetIssue(issueKey, fields)
		return result, err
	})
}

// searchIssuesHandler searches for Jira issues using a JQL query.
func (h *Handler) searchIssuesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("search issues", func() (interface{}, error) {
		jql, ok := tools.GetStringArg(args, "jql")
		if !ok {
			return nil, fmt.Errorf("missing or invalid jql parameter")
		}

		projectKeyOrId, _ := tools.GetStringArg(args, "projectKeyOrId")
		orderBy, _ := tools.GetStringArg(args, "orderBy")

		statuses := tools.GetStringSliceArg(args, "statuses")
		fields := tools.GetStringSliceArg(args, "fields")

		startAt := tools.GetIntArg(args, "startAt", 0)
		maxResults := tools.GetIntArg(args, "maxResults", 50)

		return h.client.SearchIssues(jql, projectKeyOrId, orderBy, statuses, maxResults, startAt, fields)
	})
}

// createIssueHandler creates a new Jira issue.
func (h *Handler) createIssueHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("create issue", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		summary, ok := tools.GetStringArg(args, "summary")
		if !ok {
			return nil, fmt.Errorf("missing or invalid summary parameter")
		}

		issueType, ok := tools.GetStringArg(args, "issueType")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueType parameter")
		}

		description, _ := tools.GetStringArg(args, "description")
		priority, _ := tools.GetStringArg(args, "priority")

		return h.client.CreateIssue(projectKey, summary, issueType, description, priority)
	})
}

// updateIssueHandler handles updating an existing Jira issue
func (h *Handler) updateIssueHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("update issue", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		updates, ok := args["updates"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("missing or invalid updates parameter")
		}

		return h.client.UpdateIssue(issueKey, updates)
	})
}

// getSubtasksHandler handles getting subtasks for a Jira issue
func (h *Handler) getSubtasksHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get subtasks", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		return h.client.GetSubtasks(issueKey)
	})
}

// createSubTaskHandler handles creating a subtask for a Jira issue
func (h *Handler) createSubTaskHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("create subtask", func() (interface{}, error) {
		parentKeyOrID, ok := tools.GetStringArg(args, "parentKeyOrID")
		if !ok {
			return nil, fmt.Errorf("missing or invalid parentKeyOrID parameter")
		}

		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		summary, ok := tools.GetStringArg(args, "summary")
		if !ok {
			return nil, fmt.Errorf("missing or invalid summary parameter")
		}

		issueType, ok := tools.GetStringArg(args, "issueType")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueType parameter")
		}

		description, _ := tools.GetStringArg(args, "description")
		priority, _ := tools.GetStringArg(args, "priority")

		return h.client.CreateSubTask(parentKeyOrID, projectKey, summary, issueType, description, priority)
	})
}

// updateIssueWithOptionsHandler handles updating an existing Jira issue with options
func (h *Handler) updateIssueWithOptionsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("update issue with options", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		updates, ok := args["updates"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("missing or invalid updates parameter")
		}

		options, ok := args["options"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("missing or invalid options parameter")
		}

		// Convert options map to map[string]string
		optionsStr := make(map[string]string)
		for k, v := range options {
			if str, ok := v.(string); ok {
				optionsStr[k] = str
			}
		}

		return h.client.UpdateIssueWithOptions(issueKey, updates, optionsStr)
	})
}

// createIssueWithPayloadHandler handles creating a new Jira issue with a custom payload
func (h *Handler) createIssueWithPayloadHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("create issue with payload", func() (interface{}, error) {
		payload, ok := args["payload"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("missing or invalid payload parameter")
		}

		updateHistory := tools.GetBoolArg(args, "updateHistory", false)

		return h.client.CreateIssueWithPayload(payload, updateHistory)
	})
}

// getAgileIssueHandler handles getting an agile issue
func (h *Handler) getAgileIssueHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get agile issue", func() (interface{}, error) {
		issueIdOrKey, ok := tools.GetStringArg(args, "issueIdOrKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueIdOrKey parameter")
		}

		expand, _ := tools.GetStringArg(args, "expand")
		fields := tools.GetStringSliceArg(args, "fields")
		updateHistory := tools.GetBoolArg(args, "updateHistory", false)

		return h.client.GetAgileIssue(issueIdOrKey, expand, fields, updateHistory)
	})
}

// getIssueEstimationForBoardHandler handles getting issue estimation for board
func (h *Handler) getIssueEstimationForBoardHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get issue estimation for board", func() (interface{}, error) {
		issueIdOrKey, ok := tools.GetStringArg(args, "issueIdOrKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueIdOrKey parameter")
		}

		boardId, ok := args["boardId"].(float64)
		if !ok {
			return nil, fmt.Errorf("missing or invalid boardId parameter")
		}

		return h.client.GetIssueEstimationForBoard(issueIdOrKey, int64(boardId))
	})
}

// setIssueEstimationForBoardHandler handles setting issue estimation for board
func (h *Handler) setIssueEstimationForBoardHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("set issue estimation for board", func() (interface{}, error) {
		issueIdOrKey, ok := tools.GetStringArg(args, "issueIdOrKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueIdOrKey parameter")
		}

		boardId, ok := args["boardId"].(float64)
		if !ok {
			return nil, fmt.Errorf("missing or invalid boardId parameter")
		}

		value, ok := tools.GetStringArg(args, "value")
		if !ok {
			return nil, fmt.Errorf("missing or invalid value parameter")
		}

		return h.client.SetIssueEstimationForBoard(issueIdOrKey, int64(boardId), value)
	})
}

// AddIssueTools registers the issue-related tools with the MCP server
func AddIssueTools(server *mcp.Server, client *jira.JiraClient, hasWritePermission bool) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_issue",
		Description: "Get a specific Jira issue with default fields (summary, description, status, assignee, reporter, priority, issuetype, project, created, updated, comment).",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"issueKey": {
					Type:        "string",
					Description: "The key of the issue to retrieve",
				},
			},
			Required: []string{"issueKey"},
		},
	}, handler.getIssueHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_issue_with_fields",
		Description: "Get a specific Jira issue. You can specify which fields to retrieve using the 'fields' parameter. If no fields are specified, all fields will be returned. Common fields include: summary, description, status, assignee, reporter, priority, issuetype, project, created, updated, and resolution.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"issueKey": {
					Type:        "string",
					Description: "The key of the issue to retrieve",
				},
				"fields": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Fields to include in the response. Common fields: summary, description, status, assignee, reporter, priority, issuetype, project, created, updated, resolution. If empty or not provided, all fields are returned.",
				},
			},
			Required: []string{"issueKey"},
		},
	}, handler.getIssueWithFieldsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_search_issues",
		Description: "Search for Jira issues using JQL (Jira Query Language). This tool allows you to find issues based on various criteria such as project, status, assignee, and more using the powerful JQL syntax.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"jql": {
					Type:        "string",
					Description: "JQL query string to filter issues. Example: 'project = PROJ AND status = Open'",
				},
				"projectKeyOrId": {
					Type:        "string",
					Description: "Project key or ID to filter by",
				},
				"orderBy": {
					Type:        "string",
					Description: "Field to order results by",
				},
				"statuses": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Statuses to filter by",
				},
				"fields": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Fields to include in the response. If not provided, all fields are returned.",
				},
				"startAt": {
					Type:        "integer",
					Description: "The starting index of the returned issues (for pagination). Default: 0",
				},
				"maxResults": {
					Type:        "integer",
					Description: "The maximum number of issues to return (for pagination). Default: 50, Max: 100",
				},
			},
			Required: []string{"jql"},
		},
	}, handler.searchIssuesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_subtasks",
		Description: "Get all subtasks for a specific Jira issue.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"issueKey": {
					Type:        "string",
					Description: "The key of the issue to get subtasks for",
				},
			},
			Required: []string{"issueKey"},
		},
	}, handler.getSubtasksHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_issue_estimation_for_board",
		Description: "Get issue estimation for board to retrieve estimation data for a specific issue on a board.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"issueIdOrKey": {
					Type:        "string",
					Description: "Issue ID or key",
				},
				"boardId": {
					Type:        "integer",
					Description: "Board ID",
				},
			},
			Required: []string{"issueIdOrKey", "boardId"},
		},
	}, handler.getIssueEstimationForBoardHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_agile_issue",
		Description: "Get an agile issue with additional parameters for expand, fields, and history tracking.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"issueIdOrKey": {
					Type:        "string",
					Description: "Issue ID or key",
				},
				"expand": {
					Type:        "string",
					Description: "Fields to expand",
				},
				"fields": {
					Type:        "array",
					Items:       &jsonschema.Schema{Type: "string"},
					Description: "Fields to include",
				},
				"updateHistory": {
					Type:        "boolean",
					Description: "Whether to update history",
				},
			},
			Required: []string{"issueIdOrKey"},
		},
	}, handler.getAgileIssueHandler)

	// Only register write tools if write permission is enabled
	if hasWritePermission {
		mcp.AddTool(server, &mcp.Tool{
			Name:        "jira_create_issue",
			Description: "Create a new Jira issue in the specified project with the provided details.",
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"projectKey": {
						Type:        "string",
						Description: "Project key for the new issue",
					},
					"summary": {
						Type:        "string",
						Description: "Summary of the new issue",
					},
					"issueType": {
						Type:        "string",
						Description: "Type of the new issue",
					},
					"description": {
						Type:        "string",
						Description: "Description of the new issue",
					},
					"priority": {
						Type:        "string",
						Description: "Priority of the new issue",
					},
				},
				Required: []string{"projectKey", "summary", "issueType"},
			},
		}, handler.createIssueHandler)

		mcp.AddTool(server, &mcp.Tool{
			Name:        "jira_update_issue",
			Description: "Update an existing Jira issue with the specified fields.",
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"issueKey": {
						Type:        "string",
						Description: "The key of the issue to update",
					},
					"updates": {
						Type:        "object",
						Description: "Fields to update",
					},
				},
				Required: []string{"issueKey", "updates"},
			},
		}, handler.updateIssueHandler)

		mcp.AddTool(server, &mcp.Tool{
			Name:        "jira_create_subtask",
			Description: "Create a subtask for a specific Jira issue.",
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"parentKeyOrID": {
						Type:        "string",
						Description: "The key or ID of the parent issue",
					},
					"projectKey": {
						Type:        "string",
						Description: "Project key for the new subtask",
					},
					"summary": {
						Type:        "string",
						Description: "Summary of the new subtask",
					},
					"issueType": {
						Type:        "string",
						Description: "Type of the new subtask",
					},
					"description": {
						Type:        "string",
						Description: "Description of the new subtask",
					},
					"priority": {
						Type:        "string",
						Description: "Priority of the new subtask",
					},
				},
				Required: []string{"parentKeyOrID", "projectKey", "summary", "issueType"},
			},
		}, handler.createSubTaskHandler)

		mcp.AddTool(server, &mcp.Tool{
			Name:        "jira_set_issue_estimation_for_board",
			Description: "Set issue estimation for board to update the estimation value for a specific issue on a board.",
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"issueIdOrKey": {
						Type:        "string",
						Description: "Issue ID or key",
					},
					"boardId": {
						Type:        "integer",
						Description: "Board ID",
					},
					"value": {
						Type:        "string",
						Description: "Estimation value",
					},
				},
				Required: []string{"issueIdOrKey", "boardId", "value"},
			},
		}, handler.setIssueEstimationForBoardHandler)

		mcp.AddTool(server, &mcp.Tool{
			Name:        "jira_update_issue_with_options",
			Description: "Update an existing Jira issue with additional options for controlling the update behavior.",
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"issueKey": {
						Type:        "string",
						Description: "The key of the issue to update",
					},
					"updates": {
						Type:        "object",
						Description: "Fields to update",
					},
					"options": {
						Type:        "object",
						Description: "Options for the update",
					},
				},
				Required: []string{"issueKey", "updates", "options"},
			},
		}, handler.updateIssueWithOptionsHandler)

		mcp.AddTool(server, &mcp.Tool{
			Name:        "jira_create_issue_with_payload",
			Description: "Create a new Jira issue with a custom payload for advanced issue creation scenarios.",
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"payload": {
						Type:        "object",
						Description: "Custom payload for the issue creation",
					},
					"updateHistory": {
						Type:        "boolean",
						Description: "Whether to update the issue view history",
					},
				},
				Required: []string{"payload"},
			},
		}, handler.createIssueWithPayloadHandler)

	}
}
