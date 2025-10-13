package jira

import (
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
)

// GetUserByName retrieves a user by their username.
//
// Parameters:
//   - input: GetUserByNameInput containing username
//
// Returns:
//   - types.MapOutput: The user data
//   - error: An error if the request fails
func (c *JiraClient) GetUserByName(input GetUserByNameInput) (types.MapOutput, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "username", input.Username, "")

	var user types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "user"}, queryParams, nil, &user, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByKey retrieves a user by their key.
//
// Parameters:
//   - input: GetUserByKeyInput containing key
//
// Returns:
//   - types.MapOutput: The user data
//   - error: An error if the request fails
func (c *JiraClient) GetUserByKey(input GetUserByKeyInput) (types.MapOutput, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "key", input.Key, "")

	var user types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "user"}, queryParams, nil, &user, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// SearchUsers searches for users based on a query.
//
// Parameters:
//   - input: SearchUsersInput containing query, startAt, and maxResults
//
// Returns:
//   - []types.MapOutput: The users data
//   - error: An error if the request fails
func (c *JiraClient) SearchUsers(input SearchUsersInput) ([]types.MapOutput, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "username", input.Query, "")
	utils.SetQueryParam(queryParams, "startAt", input.StartAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", input.MaxResults, 0)

	var users []types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "user", "search"}, queryParams, nil, &users, utils.AcceptJSON)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetCurrentUser retrieves the current user.
//
// Parameters:
//   - input: GetCurrentUserInput (no parameters needed)
//
// Returns:
//   - types.MapOutput: The current user data
//   - error: An error if the request fails
func (c *JiraClient) GetCurrentUser() (types.MapOutput, error) {

	req, err := utils.BuildHttpRequest(
		http.MethodGet,
		c.Config.URL,
		[]string{"rest", "api", "2", "myself"},
		nil,
		nil,
		c.Config.Token,
		utils.AcceptJSON,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	var user types.MapOutput
	if err := utils.HandleHTTPResponse(resp, "jira", &user); err != nil {
		return nil, err
	}

	return user, nil
}
