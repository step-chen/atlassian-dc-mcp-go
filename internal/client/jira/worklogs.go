package jira

import (
	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// GetWorklogs retrieves worklogs for a specific issue.
//
// Parameters:
//   - input: GetWorklogsInput containing issueKey and optional worklogId
//
// Returns:
//   - types.MapOutput: The worklogs data
//   - error: An error if the request fails
func (c *JiraClient) GetWorklogs(input GetWorklogsInput) (types.MapOutput, error) {

	pathSegments := []string{"rest", "api", "2", "issue", input.IssueKey, "worklog"}

	if input.WorklogId != "" {
		pathSegments = append(pathSegments, input.WorklogId)
	}

	var worklogs types.MapOutput
	err := c.executeRequest(http.MethodGet, pathSegments, nil, nil, &worklogs, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return worklogs, nil
}

// AddWorklog adds a new worklog entry to an issue.
//
// Parameters:
//   - input: AddWorklogInput containing issueKey and worklog details
//
// Returns:
//   - types.MapOutput: The created worklog data
//   - error: An error if the request fails
func (c *JiraClient) AddWorklog(input AddWorklogInput) (types.MapOutput, error) {
	pathSegments := []string{"rest", "api", "2", "issue", input.IssueKey, "worklog"}

	// Prepare query parameters
	queryParams := url.Values{}
	utils.SetQueryParam(queryParams, "newEstimate", input.NewEstimate, "")
	utils.SetQueryParam(queryParams, "adjustEstimate", input.AdjustEstimate, "")
	utils.SetQueryParam(queryParams, "reduceBy", input.ReduceBy, "")

	// Prepare request body
	requestBody := make(map[string]interface{})
	utils.SetRequestBodyParam(requestBody, "timeSpent", input.TimeSpent)
	utils.SetRequestBodyParam(requestBody, "comment", input.Comment)
	utils.SetRequestBodyParam(requestBody, "started", input.Started)

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	var worklog types.MapOutput
	err = c.executeRequest(http.MethodPost, pathSegments, queryParams, body, &worklog, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return worklog, nil
}
