package jira

import (
	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
	"context"
	"net/http"
)

// GetPriorities retrieves all priorities.
// Parameters:
//   - input: The input for retrieving priorities (empty struct)
//
// Returns:
//   - []types.MapOutput: The priorities data
//   - error: An error if the request fails
func (c *JiraClient) GetPriorities(ctx context.Context) ([]types.MapOutput, error) {
	var outputs []types.MapOutput
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "2", "priority"},
		nil,
		nil,
		client.AcceptJSON,
		&outputs,
	)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}
