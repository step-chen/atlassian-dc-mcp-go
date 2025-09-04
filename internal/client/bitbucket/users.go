package bitbucket

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetCurrentUser retrieves details of the current user.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch details
// of the currently authenticated user.
//
// Returns:
//   - map[string]interface{}: The user data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCurrentUser() (map[string]interface{}, error) {
	var user map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "users"},
		nil,
		nil,
		&user,
	); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves details of a specific user.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch details
// of a user identified by their user slug.
//
// Parameters:
//   - userSlug: The slug of the user to retrieve
//
// Returns:
//   - map[string]interface{}: The user data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetUser(userSlug string) (map[string]interface{}, error) {
	var user map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "users", userSlug},
		nil,
		nil,
		&user,
	); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsers retrieves a list of users with filtering options.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch users
// with various filtering options.
//
// Parameters:
//   - filter: Text to filter users by
//   - permission: Filter users by permission
//   - group: Filter users by group
//   - permissionFilters: Additional permission filters
//
// Returns:
//   - map[string]interface{}: The users data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetUsers(filter, permission, group string, permissionFilters map[string]string) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "filter", filter, "")
	utils.SetQueryParam(queryParams, "permission", permission, "")
	utils.SetQueryParam(queryParams, "group", group, "")

	for key, value := range permissionFilters {
		if key != "" && value != "" {
			queryParams.Set(key, value)
		}
	}

	var users map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "users"},
		queryParams,
		nil,
		&users,
	); err != nil {
		return nil, err
	}

	return users, nil
}