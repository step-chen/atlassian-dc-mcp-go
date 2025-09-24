package bitbucket

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
)

// GetCurrentUser retrieves details of the current user.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch details
// of the currently authenticated user.
//
// Returns:
//   - types.MapOutput: The current user data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCurrentUser() (types.MapOutput, error) {
	var user types.MapOutput
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
// of a user identified by their username.
//
// Parameters:
//   - input: GetUserInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The user data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetUser(input GetUserInput) (types.MapOutput, error) {
	var user types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "users", input.Username},
		nil,
		nil,
		&user,
	); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsers retrieves a list of users.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch a list of users
// based on the provided input parameters.
//
// Parameters:
//   - input: GetUsersInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The users data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetUsers(input GetUsersInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "filter", input.Filter, "")
	utils.SetQueryParam(queryParams, "permission", input.Permission, "")
	utils.SetQueryParam(queryParams, "group", input.Group, "")

	// Handle permissionFilters map
	if input.PermissionFilters != nil {
		for key, value := range input.PermissionFilters {
			if key != "" && value != "" {
				queryParams.Set(key, value)
			}
		}
	}

	var users types.MapOutput
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
