package jira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// GetComments retrieves comments for a specific issue.
//
// Parameters:
//   - input: GetCommentsInput containing issueKey, startAt, maxResults, expand, and orderBy
//
// Returns:
//   - types.MapOutput: The comments data
//   - error: An error if the request fails
func (c *JiraClient) GetComments(input GetCommentsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	client.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)
	client.SetQueryParam(queryParams, "orderBy", input.OrderBy, "")
	client.SetQueryParam(queryParams, "expand", input.Expand, "")

	var output types.MapOutput
	err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "2", "issue", input.IssueKey, "comment"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// AddComment adds a comment to a specific issue.
//
// Parameters:
//   - input: AddCommentInput containing issueKey and comment
//
// Returns:
//   - types.MapOutput: The added comment data
//   - error: An error if the request fails
func (c *JiraClient) AddComment(input AddCommentInput) (types.MapOutput, error) {
	payload := types.MapOutput{}
	client.SetRequestBodyParam(payload, "body", input.Comment)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	var output types.MapOutput
	err = client.ExecuteRequest(
		c.BaseClient,
		http.MethodPost,
		[]string{"rest", "api", "2", "issue", input.IssueKey, "comment"},
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
