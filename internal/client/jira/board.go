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
//   - startAt: The starting index of the returned boards
//   - maxResults: The maximum number of boards to return per page
//   - name: Filters results to boards that match the specified name
//   - projectKeyOrId: Filters results to boards that match the specified project key or ID
//   - boardType: Filters results to boards of the specified type
//
// Returns:
//   - map[string]any: The boards data
//   - error: An error if the request fails
func (c *JiraClient) GetBoards(startAt, maxResults int, name, projectKeyOrId, boardType string) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", startAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", maxResults, 0)
	utils.SetQueryParam(queryParams, "name", name, "")
	utils.SetQueryParam(queryParams, "projectKeyOrId", projectKeyOrId, "")
	utils.SetQueryParam(queryParams, "type", boardType, "")

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
//   - id: The ID of the board to retrieve
//
// Returns:
//   - map[string]any: The board data
//   - error: An error if the request fails
func (c *JiraClient) GetBoard(id int) (map[string]any, error) {
	var boardResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "board", strconv.Itoa(id)}, nil, nil, &boardResponse)
	if err != nil {
		return nil, err
	}

	return boardResponse, nil
}

// GetBoardBacklog retrieves the backlog of a specific board.
//
// Parameters:
//   - boardId: The ID of the board
//   - startAt: The starting index of the returned issues
//   - maxResults: The maximum number of issues to return per page
//   - jql: Filters results using a JQL query
//   - validateQuery: Specifies whether to validate the JQL query
//   - fields: The list of fields to return for each issue
//   - expand: A comma-separated list of parameters to expand
//
// Returns:
//   - map[string]any: The backlog data
//   - error: An error if the request fails
func (c *JiraClient) GetBoardBacklog(boardId int, startAt int, maxResults int, jql string, validateQuery bool, fields []string, expand string) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", startAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", maxResults, 0)
	utils.SetQueryParam(queryParams, "jql", jql, "")
	utils.SetQueryParam(queryParams, "validateQuery", validateQuery, false)
	utils.SetQueryParam(queryParams, "expand", expand, "")

	if len(fields) > 0 {
		for _, field := range fields {
			queryParams.Add("fields", field)
		}
	}

	var backlogResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "board", strconv.Itoa(boardId), "backlog"}, queryParams, nil, &backlogResponse)
	if err != nil {
		return nil, err
	}

	return backlogResponse, nil
}

// GetBoardEpics retrieves the epics associated with a specific board.
//
// Parameters:
//   - boardId: The ID of the board
//   - startAt: The starting index of the returned epics
//   - maxResults: The maximum number of epics to return per page
//   - done: Filters results to epics that are either done or not done
//
// Returns:
//   - map[string]any: The epics data
//   - error: An error if the request fails
func (c *JiraClient) GetBoardEpics(boardId int, startAt int, maxResults int, done string) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", startAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", maxResults, 0)
	utils.SetQueryParam(queryParams, "done", done, "")

	var epicsResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "board", strconv.Itoa(boardId), "epic"}, queryParams, nil, &epicsResponse)
	if err != nil {
		return nil, err
	}

	return epicsResponse, nil
}

// GetBoardSprints retrieves the sprints associated with a specific board.
//
// Parameters:
//   - boardId: The ID of the board
//   - startAt: The starting index of the returned sprints
//   - maxResults: The maximum number of sprints to return per page
//   - state: Filters results to sprints in the specified states
//
// Returns:
//   - map[string]any: The sprints data
//   - error: An error if the request fails
func (c *JiraClient) GetBoardSprints(boardId int, startAt int, maxResults int, state string) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", startAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", maxResults, 0)
	utils.SetQueryParam(queryParams, "state", state, "")

	var sprintsResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "board", strconv.Itoa(boardId), "sprint"}, queryParams, nil, &sprintsResponse)
	if err != nil {
		return nil, err
	}

	return sprintsResponse, nil
}

// GetSprint retrieves a specific sprint by its ID.
//
// Parameters:
//   - sprintId: The ID of the sprint to retrieve
//
// Returns:
//   - map[string]any: The sprint data
//   - error: An error if the request fails
func (c *JiraClient) GetSprint(sprintId int) (map[string]any, error) {
	var sprintResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "sprint", strconv.Itoa(sprintId)}, nil, nil, &sprintResponse)
	if err != nil {
		return nil, err
	}

	return sprintResponse, nil
}

// GetSprintIssues retrieves issues in a specific sprint.
//
// Parameters:
//   - sprintId: The ID of the sprint
//   - startAt: The starting index of the returned issues
//   - maxResults: The maximum number of issues to return per page
//   - jql: Filters results using a JQL query
//   - validateQuery: Specifies whether to validate the JQL query
//   - fields: The list of fields to return for each issue
//   - expand: A comma-separated list of parameters to expand
//
// Returns:
//   - map[string]any: The issues data
//   - error: An error if the request fails
func (c *JiraClient) GetSprintIssues(sprintId int, startAt int, maxResults int, jql string, validateQuery bool, fields []string, expand string) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", startAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", maxResults, 0)
	utils.SetQueryParam(queryParams, "jql", jql, "")
	utils.SetQueryParam(queryParams, "validateQuery", validateQuery, false)
	utils.SetQueryParam(queryParams, "expand", expand, "")

	if len(fields) > 0 {
		for _, field := range fields {
			queryParams.Add("fields", field)
		}
	}

	var issuesResponse map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "agile", "1.0", "sprint", strconv.Itoa(sprintId), "issue"}, queryParams, nil, &issuesResponse)
	if err != nil {
		return nil, err
	}

	return issuesResponse, nil
}