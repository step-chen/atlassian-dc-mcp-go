package jira

import (
	"net/http"
)

// GetWorklogs retrieves worklogs for a specific issue.
//
// Parameters:
//   - issueKey: The key of the issue
//   - worklogId: Optional worklog ID to retrieve a specific worklog
//
// Returns:
//   - map[string]any: The worklogs data
//   - error: An error if the request fails
func (c *JiraClient) GetWorklogs(issueKey string, worklogId ...string) (map[string]any, error) {

	pathSegments := []string{"rest", "api", "2", "issue", issueKey, "worklog"}

	if len(worklogId) > 0 && worklogId[0] != "" {
		pathSegments = append(pathSegments, worklogId[0])
	}

	var worklogs map[string]any
	err := c.executeRequest(http.MethodGet, pathSegments, nil, nil, &worklogs)
	if err != nil {
		return nil, err
	}

	return worklogs, nil
}