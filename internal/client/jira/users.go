package jira

import (
	"context"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// GetUserByName retrieves a user by their username.
//
// Parameters:
//   - input: GetUserByNameInput containing username
//
// Returns:
//   - types.MapOutput: The user data
//   - error: An error if the request fails
func (c *JiraClient) GetUserByName(ctx context.Context, input GetUserByNameInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "username", input.Username, "")

	var output types.MapOutput
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "2", "user"},
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

// GetUserByKey retrieves a user by their key.
//
// Parameters:
//   - input: GetUserByKeyInput containing key
//
// Returns:
//   - types.MapOutput: The user data
//   - error: An error if the request fails
func (c *JiraClient) GetUserByKey(ctx context.Context, input GetUserByKeyInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "key", input.Key, "")

	var output types.MapOutput
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "2", "user"},
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

// SearchUsers searches for users based on a query.
//
// Parameters:
//   - input: SearchUsersInput containing query, startAt, and maxResults
//
// Returns:
//   - []types.MapOutput: The users data
//   - error: An error if the request fails
func (c *JiraClient) SearchUsers(ctx context.Context, input SearchUsersInput) ([]types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "username", input.Query, "")
	client.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	client.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)

	var outputs []types.MapOutput
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "2", "user", "search"},
		queryParams,
		nil,
		client.AcceptJSON,
		&outputs,
	)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}

// GetCurrentUser retrieves the current user.
//
// Parameters:
//   - input: GetCurrentUserInput (no parameters needed)
//
// Returns:
//   - types.MapOutput: The current user data
//   - error: An error if the request fails
func (c *JiraClient) GetCurrentUser(ctx context.Context) (types.MapOutput, error) {
	var output types.MapOutput
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "2", "myself"},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}
