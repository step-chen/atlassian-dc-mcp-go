package jira

import (
	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
	"net/http"
)

// GetPriorities retrieves all priorities.
//
// Returns:
//   - []types.MapOutput: The priorities data
//   - error: An error if the request fails
func (c *JiraClient) GetPriorities() ([]types.MapOutput, error) {
	var priorities []types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "priority"}, nil, nil, &priorities, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return priorities, nil
}
