package jira

import (
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
//   - map[string]any: The search results
//   - error: An error if the request fails
func (c *JiraClient) SearchIssues(input SearchIssuesInput) (map[string]any, error) {

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

	payload := map[string]any{
		"jql":        finalJQL,
		"maxResults": input.MaxResults,
		"startAt":    input.StartAt,
	}

	if len(input.Fields) > 0 {
		payload["fields"] = input.Fields
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var result map[string]interface{}
	err = c.executeRequest(http.MethodPost, []string{"rest", "api", "2", "search"}, nil, jsonPayload, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}