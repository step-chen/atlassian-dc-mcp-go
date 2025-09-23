package jira

import (
	"net/http"
	"net/url"
	"strconv"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetBoards retrieves boards based on filters.
//
// Parameters:
//   - input: GetBoardsInput containing startAt, maxResults, name, projectKeyOrId, and boardType
//
// Returns:
//   - map[string]any: The boards data
//   - error: An error if the request fails
func (c *JiraClient) GetBoards(input GetBoardsInput) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	utils.SetQueryParam(queryParams, "name", input.Name, "")
	utils.SetQueryParam(queryParams, "projectKeyOrId", input.ProjectKeyOrId, "")
	utils.SetQueryParam(queryParams, "type", input.BoardType, "")

	var boardsResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "board"}, queryParams, nil, &boardsResponse)
	if err != nil {
		return nil, err
	}

	return boardsResponse, nil
}

// GetBoard retrieves a specific board by its ID.
//
// Parameters:
//   - input: GetBoardInput containing id
//
// Returns:
//   - map[string]any: The board data
//   - error: An error if the request fails
func (c *JiraClient) GetBoard(input GetBoardInput) (map[string]any, error) {
	var boardResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "board", strconv.Itoa(input.Id)}, nil, nil, &boardResponse)
	if err != nil {
		return nil, err
	}

	return boardResponse, nil
}

// GetBoardBacklog retrieves the backlog of a specific board.
//
// Parameters:
//   - input: GetBoardBacklogInput containing boardId, startAt, maxResults, jql, validateQuery, fields, and expand
//
// Returns:
//   - map[string]any: The backlog data
//   - error: An error if the request fails
func (c *JiraClient) GetBoardBacklog(input GetBoardBacklogInput) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	utils.SetQueryParam(queryParams, "jql", input.JQL, "")
	utils.SetQueryParam(queryParams, "validateQuery", input.ValidateQuery, false)
	utils.SetQueryParam(queryParams, "expand", input.Expand, "")

	if len(input.Fields) > 0 {
		for _, field := range input.Fields {
			queryParams.Add("fields", field)
		}
	}

	var backlogResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "board", strconv.Itoa(input.BoardId), "backlog"}, queryParams, nil, &backlogResponse)
	if err != nil {
		return nil, err
	}

	return backlogResponse, nil
}

// GetBoardEpics retrieves the epics associated with a specific board.
//
// Parameters:
//   - input: GetBoardEpicsInput containing boardId, startAt, maxResults, and done
//
// Returns:
//   - map[string]any: The epics data
//   - error: An error if the request fails
func (c *JiraClient) GetBoardEpics(input GetBoardEpicsInput) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	utils.SetQueryParam(queryParams, "done", input.Done, "")

	var epicsResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "board", strconv.Itoa(input.BoardId), "epic"}, queryParams, nil, &epicsResponse)
	if err != nil {
		return nil, err
	}

	return epicsResponse, nil
}

// GetBoardSprints retrieves the sprints associated with a specific board.
//
// Parameters:
//   - input: GetBoardSprintsInput containing boardId, startAt, maxResults, and state
//
// Returns:
//   - map[string]any: The sprints data
//   - error: An error if the request fails
func (c *JiraClient) GetBoardSprints(input GetBoardSprintsInput) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	utils.SetQueryParam(queryParams, "state", input.State, "")

	var sprintsResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "board", strconv.Itoa(input.BoardId), "sprint"}, queryParams, nil, &sprintsResponse)
	if err != nil {
		return nil, err
	}

	return sprintsResponse, nil
}

// GetSprint retrieves a specific sprint by its ID.
//
// Parameters:
//   - input: GetSprintInput containing sprintId
//
// Returns:
//   - map[string]any: The sprint data
//   - error: An error if the request fails
func (c *JiraClient) GetSprint(input GetSprintInput) (map[string]any, error) {
	var sprintResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "sprint", strconv.Itoa(input.SprintId)}, nil, nil, &sprintResponse)
	if err != nil {
		return nil, err
	}

	return sprintResponse, nil
}

// GetSprintIssues retrieves issues in a specific sprint.
//
// Parameters:
//   - input: GetSprintIssuesInput containing sprintId, startAt, maxResults, jql, validateQuery, fields, and expand
//
// Returns:
//   - map[string]any: The issues data
//   - error: An error if the request fails
func (c *JiraClient) GetSprintIssues(input GetSprintIssuesInput) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	utils.SetQueryParam(queryParams, "jql", input.JQL, "")
	utils.SetQueryParam(queryParams, "validateQuery", input.ValidateQuery, false)
	utils.SetQueryParam(queryParams, "expand", input.Expand, "")

	if len(input.Fields) > 0 {
		for _, field := range input.Fields {
			queryParams.Add("fields", field)
		}
	}

	var issuesResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "sprint", strconv.Itoa(input.SprintId), "issue"}, queryParams, nil, &issuesResponse)
	if err != nil {
		return nil, err
	}

	return issuesResponse, nil
}