package jira

import (
	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// SearchIssues searches for issues using JQL.
//
// Parameters:
//   - input: SearchIssuesInput containing jql, projectKeyOrId, orderBy, statuses, maxResults, startAt, and fields
//
// Returns:
//   - types.MapOutput: The search results
//   - error: An error if the request fails
func (c *JiraClient) SearchIssues(ctx context.Context, input SearchIssuesInput) (*Issues, error) {

	finalJQL := input.JQL
	if finalJQL == "" {
		var jqlParts []string
		if input.ProjectKeyOrId != "" {
			jqlParts = append(jqlParts, fmt.Sprintf("project = '%s'", input.ProjectKeyOrId))
		}
		if len(input.Statuses) > 0 {
			quotedStatuses := make([]string, len(input.Statuses))
			for i, s := range input.Statuses {
				quotedStatuses[i] = fmt.Sprintf("'%s'", s)
			}
			jqlParts = append(jqlParts, fmt.Sprintf("status in (%s)", strings.Join(quotedStatuses, ", ")))
		}
		finalJQL = strings.Join(jqlParts, " AND ")

		if input.OrderBy != "" {
			finalJQL = fmt.Sprintf("%s ORDER BY %s", finalJQL, input.OrderBy)
		}
	}

	payload := types.MapOutput{}
	client.SetRequestBodyParam(payload, "jql", finalJQL)
	client.SetRequestBodyParam(payload, "maxResults", input.MaxResults)
	client.SetRequestBodyParam(payload, "startAt", input.StartAt)
	client.SetRequestBodyParam(payload, "fields", input.Fields)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var output *Issues
	err = client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodPost,
		[]any{"rest", "api", "2", "search"},
		nil,
		jsonPayload,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}
