package jira

import (
	"atlassian-dc-mcp-go/internal/types"
	"net/http"
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
	err := c.executeRequest(http.MethodGet, pathSegments, nil, nil, &worklogs)
	if err != nil {
		return nil, err
	}

	return worklogs, nil
}
