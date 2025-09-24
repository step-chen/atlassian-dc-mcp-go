package jira

import (
	"atlassian-dc-mcp-go/internal/types"
	"net/http"
)

// GetIssueTypes retrieves all issue types.
//
// Returns:
//   - []types.MapOutput: The issue types data
//   - error: An error if the request fails
func (c *JiraClient) GetIssueTypes() ([]types.MapOutput, error) {
	var issueTypes []types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "issuetype"}, nil, nil, &issueTypes)
	if err != nil {
		return nil, err
	}

	return issueTypes, nil
}
