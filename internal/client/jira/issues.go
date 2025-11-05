// Package jira provides a client for interacting with Jira Data Center APIs.
package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// GetIssue retrieves a specific issue by its key.
//
// Parameters:
//   - input: GetIssueInput containing issueKey and fields
//
// Returns:
//   - types.MapOutput: The issue data
//   - error: An error if the request fails
func (c *JiraClient) GetIssue(ctx context.Context, input GetIssueInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	// Pass nil as the invalid value for fields since we want to include fields when the slice is empty
	client.SetQueryParam(queryParams, "fields", input.Fields, nil)

	var output types.MapOutput
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "2", "issue", input.IssueKey},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// CreateIssue creates a new issue.
//
// Parameters:
//   - input: CreateIssueInput containing projectKey, summary, issueType, description, and priority
//
// Returns:
//   - types.MapOutput: The created issue data
//   - error: An error if the request fails
func (c *JiraClient) CreateIssue(ctx context.Context, input CreateIssueInput) (types.MapOutput, error) {
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

	return c.CreateIssueWithPayload(ctx, createPayloadInput)
}

// CreateIssueWithPayload creates a new issue with a custom payload.
//
// Parameters:
//   - input: CreateIssueWithPayloadInput containing payload and updateHistory
//
// Returns:
//   - types.MapOutput: The created issue data
//   - error: An error if the request fails
func (c *JiraClient) CreateIssueWithPayload(ctx context.Context, input CreateIssueWithPayloadInput) (types.MapOutput, error) {
	jsonPayload, err := json.Marshal(input.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	queryParams := url.Values{}
	if input.UpdateHistory {
		queryParams.Set("updateHistory", "true")
	}

	var output types.MapOutput
	err = client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodPost,
		[]any{"rest", "api", "2", "issue"},
		queryParams,
		jsonPayload,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// CreateSubTask creates a new sub-task for an issue.
//
// Parameters:
//   - input: CreateSubTaskInput containing parentKeyOrID, projectKey, summary, issueType, description, and priority
//
// Returns:
//   - types.MapOutput: The created sub-task data
//   - error: An error if the request fails
func (c *JiraClient) CreateSubTask(ctx context.Context, input CreateSubTaskInput) (types.MapOutput, error) {
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

	return c.CreateIssueWithPayload(ctx, createPayloadInput)
}

// UpdateIssue updates an existing issue.
//
// Parameters:
//   - input: UpdateIssueInput containing issueKey and updates
//
// Returns:
//   - types.MapOutput: The updated issue data
//   - error: An error if the request fails
func (c *JiraClient) UpdateIssue(ctx context.Context, input UpdateIssueInput) (types.MapOutput, error) {
	return c.UpdateIssueWithOptions(ctx, UpdateIssueWithOptionsInput{
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
func (c *JiraClient) UpdateIssueWithOptions(ctx context.Context, input UpdateIssueWithOptionsInput) (types.MapOutput, error) {

	payload := types.MapOutput{}

	if input.Updates != nil {
		payload["fields"] = input.Updates
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var queryParams url.Values
	if input.Options != nil {
		queryParams = url.Values{}
		for k, v := range input.Options {
			queryParams.Set(k, v)
		}
	}

	var output types.MapOutput
	err = client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodPut,
		[]any{"rest", "api", "2", "issue", input.IssueKey},
		queryParams,
		jsonPayload,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// GetTransitions retrieves transitions available for an issue.
//
// Parameters:
//   - input: GetTransitionsInput containing issueKey
//
// Returns:
//   - types.MapOutput: The transitions data
//   - error: An error if the request fails
func (c *JiraClient) GetTransitions(ctx context.Context, input GetTransitionsInput) (types.MapOutput, error) {
	var output types.MapOutput
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "2", "issue", input.IssueKey, "transitions"},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// TransitionIssue transitions an issue to a new status.
//
// Parameters:
//   - input: TransitionIssueInput containing issueKey and transitionID
//
// Returns:
//   - error: An error if the request fails
func (c *JiraClient) TransitionIssue(ctx context.Context, input TransitionIssueInput) error {
	transition := types.MapOutput{}
	client.SetRequestBodyParam(transition, "id", input.TransitionID)

	payload := types.MapOutput{}
	client.SetRequestBodyParam(payload, "transition", transition)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	err = client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodPost,
		[]any{"rest", "api", "2", "issue", input.IssueKey, "transitions"},
		nil,
		jsonPayload,
		client.AcceptJSON,
		nil,
	)
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
func (c *JiraClient) GetSubtasks(ctx context.Context, input GetSubtasksInput) ([]types.MapOutput, error) {
	var outputs []types.MapOutput
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "2", "issue", input.IssueKey, "subtask"},
		nil,
		nil,
		client.AcceptJSON,
		&outputs,
	)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}

// GetAgileIssue retrieves a specific agile issue by its key or ID.
//
// Parameters:
//   - input: GetAgileIssueInput containing issueIdOrKey, expand, fields, and updateHistory
//
// Returns:
//   - types.MapOutput: The agile issue data
//   - error: An error if the request fails
func (c *JiraClient) GetAgileIssue(ctx context.Context, input GetAgileIssueInput) (types.MapOutput, error) {

	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "expand", input.Expand, "")

	if len(input.Fields) > 0 {
		for _, field := range input.Fields {
			queryParams.Add("fields", field)
		}
	}

	client.SetQueryParam(queryParams, "updateHistory", input.UpdateHistory, false)

	var output types.MapOutput
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "agile", "1.0", "issue", input.IssueIdOrKey},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// GetIssueEstimationForBoard retrieves the estimation for an issue on a board.
//
// Parameters:
//   - input: GetIssueEstimationForBoardInput containing issueIdOrKey and boardId
//
// Returns:
//   - types.MapOutput: The estimation data
//   - error: An error if the request fails
func (c *JiraClient) GetIssueEstimationForBoard(ctx context.Context, input GetIssueEstimationForBoardInput) (types.MapOutput, error) {

	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "boardId", input.BoardId, int64(0))

	var output types.MapOutput
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "agile", "1.0", "issue", input.IssueIdOrKey, "estimation"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// SetIssueEstimationForBoard sets the estimation for an issue on a board.
//
// Parameters:
//   - input: SetIssueEstimationForBoardInput containing issueIdOrKey, boardId, and value
//
// Returns:
//   - types.MapOutput: The estimation data
//   - error: An error if the request fails
func (c *JiraClient) SetIssueEstimationForBoard(ctx context.Context, input SetIssueEstimationForBoardInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "boardId", input.BoardId, int64(0))

	payload := types.MapOutput{}
	client.SetRequestBodyParam(payload, "value", input.Value)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var output types.MapOutput
	err = client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodPut,
		[]any{"rest", "agile", "1.0", "issue", input.IssueIdOrKey, "estimation"},
		queryParams,
		jsonPayload,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}
