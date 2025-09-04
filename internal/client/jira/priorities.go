package jira

import (
	"net/http"
)

// GetPriorities retrieves all priorities.
//
// Returns:
//   - []map[string]any: The priorities data
//   - error: An error if the request fails
func (c *JiraClient) GetPriorities() ([]map[string]any, error) {
	var priorities []map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "priority"}, nil, nil, &priorities)
	if err != nil {
		return nil, err
	}

	return priorities, nil
}