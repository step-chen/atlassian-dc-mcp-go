package jira

import (
	"net/http"
	"net/url"
	"strconv"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// GetBoards retrieves boards based on filters.
//
// Parameters:
//   - input: GetBoardsInput containing startAt, maxResults, name, projectKeyOrId, and boardType
//
// Returns:
//   - types.MapOutput: The boards data
//   - error: An error if the request fails
func (c *JiraClient) GetBoards(input GetBoardsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	client.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	client.SetQueryParam(queryParams, "name", input.Name, "")
	client.SetQueryParam(queryParams, "projectKeyOrId", input.ProjectKeyOrId, "")
	client.SetQueryParam(queryParams, "type", input.BoardType, "")

	var output types.MapOutput
	err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "agile", "1.0", "board"},
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

// GetBoard retrieves a specific board by its ID.
//
// Parameters:
//   - input: GetBoardInput containing id
//
// Returns:
//   - types.MapOutput: The board data
//   - error: An error if the request fails
func (c *JiraClient) GetBoard(input GetBoardInput) (types.MapOutput, error) {
	var output types.MapOutput
	err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "agile", "1.0", "board", strconv.Itoa(input.Id)},
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

// GetBoardBacklog retrieves the backlog of a specific board.
//
// Parameters:
//   - input: GetBoardBacklogInput containing boardId, startAt, maxResults, jql, validateQuery, fields, and expand
//
// Returns:
//   - types.MapOutput: The backlog data
//   - error: An error if the request fails
func (c *JiraClient) GetBoardBacklog(input GetBoardBacklogInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	client.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	client.SetQueryParam(queryParams, "jql", input.JQL, "")
	client.SetQueryParam(queryParams, "validateQuery", input.ValidateQuery, false)
	client.SetQueryParam(queryParams, "expand", input.Expand, "")

	if len(input.Fields) > 0 {
		for _, field := range input.Fields {
			queryParams.Add("fields", field)
		}
	}

	var output types.MapOutput
	err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "agile", "1.0", "board", strconv.Itoa(input.BoardId), "backlog"},
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

// GetBoardEpics retrieves the epics associated with a specific board.
//
// Parameters:
//   - input: GetBoardEpicsInput containing boardId, startAt, maxResults, and done
//
// Returns:
//   - types.MapOutput: The epics data
//   - error: An error if the request fails
func (c *JiraClient) GetBoardEpics(input GetBoardEpicsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	client.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	client.SetQueryParam(queryParams, "done", input.Done, "")

	var output types.MapOutput
	err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "agile", "1.0", "board", strconv.Itoa(input.BoardId), "epic"},
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

// GetBoardSprints retrieves the sprints associated with a specific board.
//
// Parameters:
//   - input: GetBoardSprintsInput containing boardId, startAt, maxResults, and state
//
// Returns:
//   - types.MapOutput: The sprints data
//   - error: An error if the request fails
func (c *JiraClient) GetBoardSprints(input GetBoardSprintsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	client.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	client.SetQueryParam(queryParams, "state", input.State, "")

	var output types.MapOutput
	err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "agile", "1.0", "board", strconv.Itoa(input.BoardId), "sprint"},
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

// GetSprint retrieves a specific sprint by its ID.
//
// Parameters:
//   - input: GetSprintInput containing sprintId
//
// Returns:
//   - types.MapOutput: The sprint data
//   - error: An error if the request fails
func (c *JiraClient) GetSprint(input GetSprintInput) (types.MapOutput, error) {
	var output types.MapOutput
	err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "agile", "1.0", "sprint", strconv.Itoa(input.SprintId)},
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

// GetSprintIssues retrieves issues in a specific sprint.
//
// Parameters:
//   - input: GetSprintIssuesInput containing sprintId, startAt, maxResults, jql, validateQuery, fields, and expand
//
// Returns:
//   - types.MapOutput: The issues data
//   - error: An error if the request fails
func (c *JiraClient) GetSprintIssues(input GetSprintIssuesInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	client.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	client.SetQueryParam(queryParams, "jql", input.JQL, "")
	client.SetQueryParam(queryParams, "validateQuery", input.ValidateQuery, false)
	client.SetQueryParam(queryParams, "expand", input.Expand, "")

	if len(input.Fields) > 0 {
		for _, field := range input.Fields {
			queryParams.Add("fields", field)
		}
	}

	var output types.MapOutput
	err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "agile", "1.0", "sprint", strconv.Itoa(input.SprintId), "issue"},
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
