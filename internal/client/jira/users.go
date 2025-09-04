package jira

import (
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetUserByName retrieves a user by their username.
//
// Parameters:
//   - username: The username of the user to retrieve
//
// Returns:
//   - map[string]any: The user data
//   - error: An error if the request fails
func (c *JiraClient) GetUserByName(username string) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "username", username, "")

	var user map[string]interface{}
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "user"}, queryParams, nil, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByKey retrieves a user by their key.
//
// Parameters:
//   - key: The key of the user to retrieve
//
// Returns:
//   - map[string]any: The user data
//   - error: An error if the request fails
func (c *JiraClient) GetUserByKey(key string) (map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "key", key, "")

	var user map[string]interface{}
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "user"}, queryParams, nil, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// SearchUsers searches for users based on a query.
//
// Parameters:
//   - query: The search query
//   - startAt: The index of the first item to return
//   - maxResults: The maximum number of items to return per page
//
// Returns:
//   - []map[string]any: The users data
//   - error: An error if the request fails
func (c *JiraClient) SearchUsers(query string, startAt, maxResults int) ([]map[string]any, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "username", query, "")
	utils.SetQueryParam(queryParams, "startAt", startAt, 0)
	utils.SetQueryParam(queryParams, "maxResults", maxResults, 0)

	var users []map[string]interface{}
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "user", "search"}, queryParams, nil, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetCurrentUser retrieves the current user.
//
// Returns:
//   - map[string]any: The current user data
//   - error: An error if the request fails
func (c *JiraClient) GetCurrentUser() (map[string]any, error) {

	req, err := utils.BuildHttpRequest(
		http.MethodGet,
		c.Config.URL,
		[]string{"rest", "api", "2", "myself"},
		nil,
		nil,
		c.Config.Token,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	var user map[string]interface{}
	if err := utils.HandleHTTPResponse(resp, "jira", &user); err != nil {
		return nil, err
	}

	return user, nil
}