// Package jira provides a client for interacting with Jira Data Center APIs.
package jira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetIssue retrieves a specific issue by its key.
//
// Parameters:
//   - issueKey: The key of the issue to retrieve
//   - fields: The list of fields to return for the issue
//
// Returns:
//   - map[string]any: The issue data
//   - error: An error if the request fails
func (c *JiraClient) GetIssue(issueKey string, fields []string) (map[string]any, error) {
	queryParams := make(url.Values)
	// Pass nil as the invalid value for fields since we want to include fields when the slice is empty
	utils.SetQueryParam(queryParams, "fields", fields, nil)

	var issue map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "issue", issueKey}, queryParams, nil, &issue)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

// CreateIssue creates a new issue.
//
// Parameters:
//   - projectKey: The key of the project to create the issue in
//   - summary: The summary of the issue
//   - issueType: The type of the issue
//   - description: The description of the issue
//   - priority: The priority of the issue
//
// Returns:
//   - map[string]any: The created issue data
//   - error: An error if the request fails
func (c *JiraClient) CreateIssue(projectKey, summary, issueType, description, priority string) (map[string]any, error) {
	payload := map[string]any{
		"fields": map[string]any{
			"project": map[string]string{
				"key": projectKey,
			},
			"summary": summary,
			"issuetype": map[string]string{
				"name": issueType,
			},
			"description": description,
			"priority": map[string]string{
				"name": priority,
			},
		},
	}

	return c.CreateIssueWithPayload(payload, false)
}

// CreateIssueWithPayload creates a new issue with a custom payload.
//
// Parameters:
//   - payload: The payload containing issue data
//   - updateHistory: Whether to update the user's history
//
// Returns:
//   - map[string]any: The created issue data
//   - error: An error if the request fails
func (c *JiraClient) CreateIssueWithPayload(payload map[string]any, updateHistory bool) (map[string]any, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	queryParams := make(url.Values)
	if updateHistory {
		queryParams.Set("updateHistory", "true")
	}

	var issue map[string]interface{}
	err = c.executeRequest(http.MethodPost, []string{"rest", "api", "2", "issue"}, queryParams, jsonPayload, &issue)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

// CreateSubTask creates a new sub-task for an issue.
//
// Parameters:
//   - parentKeyOrID: The key or ID of the parent issue
//   - projectKey: The key of the project to create the sub-task in
//   - summary: The summary of the sub-task
//   - issueType: The type of the sub-task
//   - description: The description of the sub-task
//   - priority: The priority of the sub-task
//
// Returns:
//   - map[string]any: The created sub-task data
//   - error: An error if the request fails
func (c *JiraClient) CreateSubTask(parentKeyOrID, projectKey, summary, issueType, description, priority string) (map[string]any, error) {
	payload := map[string]any{
		"fields": map[string]any{
			"project": map[string]string{
				"key": projectKey,
			},
			"summary": summary,
			"issuetype": map[string]string{
				"name": issueType,
			},
			"description": description,
			"priority": map[string]string{
				"name": priority,
			},
			"parent": map[string]string{
				"key": parentKeyOrID,
			},
		},
	}

	return c.CreateIssueWithPayload(payload, false)
}

// UpdateIssue updates an existing issue.
//
// Parameters:
//   - issueKey: The key of the issue to update
//   - updates: The fields to update
//
// Returns:
//   - map[string]any: The updated issue data
//   - error: An error if the request fails
func (c *JiraClient) UpdateIssue(issueKey string, updates map[string]any) (map[string]any, error) {
	return c.UpdateIssueWithOptions(issueKey, updates, nil)
}

// UpdateIssueWithOptions updates an existing issue with additional options.
//
// Parameters:
//   - issueKey: The key of the issue to update
//   - updates: The fields to update
//   - options: Additional options for the update
//
// Returns:
//   - map[string]any: The updated issue data
//   - error: An error if the request fails
func (c *JiraClient) UpdateIssueWithOptions(issueKey string, updates map[string]any, options map[string]string) (map[string]any, error) {

	payload := make(map[string]any)

	if updates != nil {
		payload["fields"] = updates
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var queryParams url.Values
	if options != nil {
		queryParams = make(url.Values)
		for k, v := range options {
			queryParams.Set(k, v)
		}
	}

	var result map[string]interface{}
	err = c.executeRequest(http.MethodPut, []string{"rest", "api", "2", "issue", issueKey}, queryParams, jsonPayload, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTransitions retrieves transitions available for an issue.
//
// Parameters:
//   - issueKey: The key of the issue
//
// Returns:
//   - map[string]any: The transitions data
//   - error: An error if the request fails
func (c *JiraClient) GetTransitions(issueKey string) (map[string]any, error) {
	var transitions map[string]interface{}
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "issue", issueKey, "transitions"}, nil, nil, &transitions)
	if err != nil {
		return nil, err
	}

	return transitions, nil
}

// TransitionIssue transitions an issue to a new status.
//
// Parameters:
//   - issueKey: The key of the issue to transition
//   - transitionID: The ID of the transition to apply
//
// Returns:
//   - error: An error if the request fails
func (c *JiraClient) TransitionIssue(issueKey, transitionID string) error {
	payload := map[string]any{
		"transition": map[string]string{
			"id": transitionID,
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	err = c.executeRequest(http.MethodPost, []string{"rest", "api", "2", "issue", issueKey, "transitions"}, nil, jsonPayload, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetSubtasks retrieves sub-tasks of an issue.
//
// Parameters:
//   - issueKey: The key of the issue
//
// Returns:
//   - []map[string]any: The sub-tasks data
//   - error: An error if the request fails
func (c *JiraClient) GetSubtasks(issueKey string) ([]map[string]any, error) {
	var subtasks []map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "issue", issueKey, "subtask"}, nil, nil, &subtasks)
	if err != nil {
		return nil, err
	}

	return subtasks, nil
}

// GetAgileIssue retrieves a specific agile issue by its key or ID.
//
// Parameters:
//   - issueIdOrKey: The ID or key of the issue
//   - expand: Parameters to expand in the response
//   - fields: The list of fields to return for the issue
//   - updateHistory: Whether to update the user's history
//
// Returns:
//   - map[string]any: The agile issue data
//   - error: An error if the request fails
func (c *JiraClient) GetAgileIssue(issueIdOrKey string, expand string, fields []string, updateHistory bool) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "expand", expand, "")

	if len(fields) > 0 {
		for _, field := range fields {
			queryParams.Add("fields", field)
		}
	}

	utils.SetQueryParam(queryParams, "updateHistory", updateHistory, false)

	var issue map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "issue", issueIdOrKey}, queryParams, nil, &issue)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

// GetIssueEstimationForBoard retrieves the estimation for an issue on a board.
//
// Parameters:
//   - issueIdOrKey: The ID or key of the issue
//   - boardId: The ID of the board
//
// Returns:
//   - map[string]any: The estimation data
//   - error: An error if the request fails
func (c *JiraClient) GetIssueEstimationForBoard(issueIdOrKey string, boardId int64) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "boardId", boardId, int64(0))

	var estimation map[string]interface{}
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "issue", issueIdOrKey, "estimation"}, queryParams, nil, &estimation)
	if err != nil {
		return nil, err
	}

	return estimation, nil
}

// SetIssueEstimationForBoard sets the estimation for an issue on a board.
//
// Parameters:
//   - issueIdOrKey: The ID or key of the issue
//   - boardId: The ID of the board
//   - value: The estimation value
//
// Returns:
//   - map[string]any: The estimation data
//   - error: An error if the request fails
func (c *JiraClient) SetIssueEstimationForBoard(issueIdOrKey string, boardId int64, value string) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "boardId", boardId, int64(0))

	payload := map[string]any{
		"value": value,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var estimation map[string]interface{}
	err = c.executeRequest(http.MethodPut, []string{"rest", "agile", "1.0", "issue", issueIdOrKey, "estimation"}, queryParams, jsonPayload, &estimation)
	if err != nil {
		return nil, err
	}

	return estimation, nil
}
