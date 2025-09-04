package jira

import (
	"net/http"
)

// GetIssueTypes retrieves all issue types.
//
// Returns:
//   - []map[string]any: The issue types data
//   - error: An error if the request fails
func (c *JiraClient) GetIssueTypes() ([]map[string]any, error) {
	var issueTypes []map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "issuetype"}, nil, nil, &issueTypes)
	if err != nil {
		return nil, err
	}

	return issueTypes, nil
}