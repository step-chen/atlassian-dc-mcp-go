package jira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetComments retrieves comments for a specific issue.
//
// Parameters:
//   - issueKey: The key of the issue
//   - startAt: The index of the first item to return
//   - maxResults: The maximum number of items to return per page
//   - expand: Use expand to include additional information about comments
//   - orderBy: Ordering of comments by creation date
//
// Returns:
//   - map[string]any: The comments data
//   - error: An error if the request fails
func (c *JiraClient) GetComments(issueKey string, startAt, maxResults int, expand, orderBy string) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", startAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", maxResults, 0)
	utils.SetQueryParam(queryParams, "expand", expand, "")
	utils.SetQueryParam(queryParams, "orderBy", orderBy, "")

	var comments map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "issue", issueKey, "comment"}, queryParams, nil, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// AddComment adds a comment to a specific issue.
//
// Parameters:
//   - issueKey: The key of the issue
//   - comment: The comment text to add
//
// Returns:
//   - map[string]any: The added comment data
//   - error: An error if the request fails
func (c *JiraClient) AddComment(issueKey, comment string) (map[string]any, error) {

	payload := map[string]any{
		"body": comment,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var result map[string]any
	err = c.executeRequest(http.MethodPost, []string{"rest", "api", "2", "issue", issueKey, "comment"}, nil, jsonPayload, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}