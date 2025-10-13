// Package jira provides a client for interacting with Jira Data Center APIs.
package jira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
)

// GetIssue retrieves a specific issue by its key.
//
// Parameters:
//   - input: GetIssueInput containing issueKey and fields
//
// Returns:
//   - types.MapOutput: The issue data
//   - error: An error if the request fails
func (c *JiraClient) GetIssue(input GetIssueInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	// Pass nil as the invalid value for fields since we want to include fields when the slice is empty
	utils.SetQueryParam(queryParams, "fields", input.Fields, nil)

	var issue types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "issue", input.IssueKey}, queryParams, nil, &issue, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

// CreateIssue creates a new issue.
//
// Parameters:
//   - input: CreateIssueInput containing projectKey, summary, issueType, description, and priority
//
// Returns:
//   - types.MapOutput: The created issue data
//   - error: An error if the request fails
func (c *JiraClient) CreateIssue(input CreateIssueInput) (types.MapOutput, error) {
	payload := types.MapOutput{
		"fields": types.MapOutput{
			"project": map[string]string{
				"key": input.ProjectKey,
			},
			"summary": input.Summary,
			"issuetype": map[string]string{
				"name": input.IssueType,
			},
			"description": input.Description,
			"priority": map[string]string{
				"name": input.Priority,
			},
		},
	}

	createPayloadInput := CreateIssueWithPayloadInput{
		Payload:       payload,
		UpdateHistory: false,
	}

	return c.CreateIssueWithPayload(createPayloadInput)
}

// CreateIssueWithPayload creates a new issue with a custom payload.
//
// Parameters:
//   - input: CreateIssueWithPayloadInput containing payload and updateHistory
//
// Returns:
//   - types.MapOutput: The created issue data
//   - error: An error if the request fails
func (c *JiraClient) CreateIssueWithPayload(input CreateIssueWithPayloadInput) (types.MapOutput, error) {
	jsonPayload, err := json.Marshal(input.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	queryParams := make(url.Values)
	if input.UpdateHistory {
		queryParams.Set("updateHistory", "true")
	}

	var issue types.MapOutput
	err = c.executeRequest(http.MethodPost, []string{"rest", "api", "2", "issue"}, queryParams, jsonPayload, &issue, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

// CreateSubTask creates a new sub-task for an issue.
//
// Parameters:
//   - input: CreateSubTaskInput containing parentKeyOrID, projectKey, summary, issueType, description, and priority
//
// Returns:
//   - types.MapOutput: The created sub-task data
//   - error: An error if the request fails
func (c *JiraClient) CreateSubTask(input CreateSubTaskInput) (types.MapOutput, error) {
	payload := types.MapOutput{
		"fields": types.MapOutput{
			"project": map[string]string{
				"key": input.ProjectKey,
			},
			"summary": input.Summary,
			"issuetype": map[string]string{
				"name": input.IssueType,
			},
			"description": input.Description,
			"priority": map[string]string{
				"name": input.Priority,
			},
			"parent": map[string]string{
				"key": input.ParentKeyOrID,
			},
		},
	}

	createPayloadInput := CreateIssueWithPayloadInput{
		Payload:       payload,
		UpdateHistory: false,
	}

	return c.CreateIssueWithPayload(createPayloadInput)
}

// UpdateIssue updates an existing issue.
//
// Parameters:
//   - input: UpdateIssueInput containing issueKey and updates
//
// Returns:
//   - types.MapOutput: The updated issue data
//   - error: An error if the request fails
func (c *JiraClient) UpdateIssue(input UpdateIssueInput) (types.MapOutput, error) {
	return c.UpdateIssueWithOptions(UpdateIssueWithOptionsInput{
		IssueKey: input.IssueKey,
		Updates:  input.Updates,
		Options:  nil,
	})
}

// UpdateIssueWithOptions updates an existing issue with additional options.
//
// Parameters:
//   - input: UpdateIssueWithOptionsInput containing issueKey, updates, and options
//
// Returns:
//   - types.MapOutput: The updated issue data
//   - error: An error if the request fails
func (c *JiraClient) UpdateIssueWithOptions(input UpdateIssueWithOptionsInput) (types.MapOutput, error) {

	payload := make(types.MapOutput)

	if input.Updates != nil {
		payload["fields"] = input.Updates
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var queryParams url.Values
	if input.Options != nil {
		queryParams = make(url.Values)
		for k, v := range input.Options {
			queryParams.Set(k, v)
		}
	}

	var result types.MapOutput
	err = c.executeRequest(http.MethodPut, []string{"rest", "api", "2", "issue", input.IssueKey}, queryParams, jsonPayload, &result, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTransitions retrieves transitions available for an issue.
//
// Parameters:
//   - input: GetTransitionsInput containing issueKey
//
// Returns:
//   - types.MapOutput: The transitions data
//   - error: An error if the request fails
func (c *JiraClient) GetTransitions(input GetTransitionsInput) (types.MapOutput, error) {
	var transitions types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "issue", input.IssueKey, "transitions"}, nil, nil, &transitions, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return transitions, nil
}

// TransitionIssue transitions an issue to a new status.
//
// Parameters:
//   - input: TransitionIssueInput containing issueKey and transitionID
//
// Returns:
//   - error: An error if the request fails
func (c *JiraClient) TransitionIssue(input TransitionIssueInput) error {
	transition := make(types.MapOutput)
	utils.SetRequestBodyParam(transition, "id", input.TransitionID)

	payload := make(types.MapOutput)
	utils.SetRequestBodyParam(payload, "transition", transition)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	err = c.executeRequest(http.MethodPost, []string{"rest", "api", "2", "issue", input.IssueKey, "transitions"}, nil, jsonPayload, nil, utils.AcceptJSON)
	if err != nil {
		return err
	}

	return nil
}

// GetSubtasks retrieves sub-tasks of an issue.
//
// Parameters:
//   - input: GetSubtasksInput containing issueKey
//
// Returns:
//   - []types.MapOutput: The sub-tasks data
//   - error: An error if the request fails
func (c *JiraClient) GetSubtasks(input GetSubtasksInput) ([]types.MapOutput, error) {
	var subtasks []types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "issue", input.IssueKey, "subtask"}, nil, nil, &subtasks, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return subtasks, nil
}

// GetAgileIssue retrieves a specific agile issue by its key or ID.
//
// Parameters:
//   - input: GetAgileIssueInput containing issueIdOrKey, expand, fields, and updateHistory
//
// Returns:
//   - types.MapOutput: The agile issue data
//   - error: An error if the request fails
func (c *JiraClient) GetAgileIssue(input GetAgileIssueInput) (types.MapOutput, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "expand", input.Expand, "")

	if len(input.Fields) > 0 {
		for _, field := range input.Fields {
			queryParams.Add("fields", field)
		}
	}

	utils.SetQueryParam(queryParams, "updateHistory", input.UpdateHistory, false)

	var issue types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "issue", input.IssueIdOrKey}, queryParams, nil, &issue, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

// GetIssueEstimationForBoard retrieves the estimation for an issue on a board.
//
// Parameters:
//   - input: GetIssueEstimationForBoardInput containing issueIdOrKey and boardId
//
// Returns:
//   - types.MapOutput: The estimation data
//   - error: An error if the request fails
func (c *JiraClient) GetIssueEstimationForBoard(input GetIssueEstimationForBoardInput) (types.MapOutput, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "boardId", input.BoardId, int64(0))

	var estimation types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "issue", input.IssueIdOrKey, "estimation"}, queryParams, nil, &estimation, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return estimation, nil
}

// SetIssueEstimationForBoard sets the estimation for an issue on a board.
//
// Parameters:
//   - input: SetIssueEstimationForBoardInput containing issueIdOrKey, boardId, and value
//
// Returns:
//   - types.MapOutput: The estimation data
//   - error: An error if the request fails
func (c *JiraClient) SetIssueEstimationForBoard(input SetIssueEstimationForBoardInput) (types.MapOutput, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "boardId", input.BoardId, int64(0))

	payload := make(types.MapOutput)
	utils.SetRequestBodyParam(payload, "value", input.Value)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var estimation types.MapOutput
	err = c.executeRequest(http.MethodPut, []string{"rest", "agile", "1.0", "issue", input.IssueIdOrKey, "estimation"}, queryParams, jsonPayload, &estimation, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return estimation, nil
}
