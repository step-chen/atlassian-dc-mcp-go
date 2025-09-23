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
//   - input: GetCommentsInput containing issueKey, startAt, maxResults, expand, and orderBy
//
// Returns:
//   - map[string]any: The comments data
//   - error: An error if the request fails
func (c *JiraClient) GetComments(input GetCommentsInput) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	utils.SetQueryParam(queryParams, "expand", input.Expand, "")
	utils.SetQueryParam(queryParams, "orderBy", input.OrderBy, "")

	var comments map[string]any
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "issue", input.IssueKey, "comment"}, queryParams, nil, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// AddComment adds a comment to a specific issue.
//
// Parameters:
//   - input: AddCommentInput containing issueKey and comment
//
// Returns:
//   - map[string]any: The added comment data
//   - error: An error if the request fails
func (c *JiraClient) AddComment(input AddCommentInput) (map[string]any, error) {

	payload := map[string]any{
		"body": input.Comment,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var result map[string]any
	err = c.executeRequest(http.MethodPost, []string{"rest", "api", "2", "issue", input.IssueKey, "comment"}, nil, jsonPayload, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}