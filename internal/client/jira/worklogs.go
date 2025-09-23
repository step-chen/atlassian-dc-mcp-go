package jira

import (
	"net/http"
)

// GetWorklogs retrieves worklogs for a specific issue.
//
// Parameters:
//   - input: GetWorklogsInput containing issueKey and optional worklogId
//
// Returns:
//   - map[string]any: The worklogs data
//   - error: An error if the request fails
func (c *JiraClient) GetWorklogs(input GetWorklogsInput) (map[string]any, error) {

	pathSegments := []string{"rest", "api", "2", "issue", input.IssueKey, "worklog"}

	if input.WorklogId != "" {
		pathSegments = append(pathSegments, input.WorklogId)
	}

	var worklogs map[string]any
	err := c.executeRequest(http.MethodGet, pathSegments, nil, nil, &worklogs)
	if err != nil {
		return nil, err
	}

	return worklogs, nil
}