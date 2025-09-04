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
//   - jql: The JQL query string
//   - projectKeyOrId: The project key or ID to filter by
//   - orderBy: The field to order results by
//   - statuses: The statuses to filter by
//   - maxResults: The maximum number of issues to return per page
//   - startAt: The index of the first item to return
//   - fields: The list of fields to return for each issue
//
// Returns:
//   - map[string]any: The search results
//   - error: An error if the request fails
func (c *JiraClient) SearchIssues(jql, projectKeyOrId, orderBy string, statuses []string, maxResults, startAt int, fields []string) (map[string]any, error) {

	finalJQL := jql
	if finalJQL == "" {
		var jqlParts []string
		if projectKeyOrId != "" {
			jqlParts = append(jqlParts, fmt.Sprintf("project = '%s'", projectKeyOrId))
		}
		if len(statuses) > 0 {
			quotedStatuses := make([]string, len(statuses))
			for i, s := range statuses {
				quotedStatuses[i] = fmt.Sprintf("'%s'", s)
			}
			jqlParts = append(jqlParts, fmt.Sprintf("status in (%s)", strings.Join(quotedStatuses, ", ")))
		}
		finalJQL = strings.Join(jqlParts, " AND ")

		if orderBy != "" {
			finalJQL = fmt.Sprintf("%s ORDER BY %s", finalJQL, orderBy)
		}
	}

	payload := map[string]any{
		"jql":        finalJQL,
		"maxResults": maxResults,
		"startAt":    startAt,
	}

	if len(fields) > 0 {
		payload["fields"] = fields
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