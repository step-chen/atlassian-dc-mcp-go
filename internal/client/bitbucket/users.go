package bitbucket

import (
	"context"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

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
func (c *BitbucketClient) GetUser(ctx context.Context, input GetUserInput) (types.MapOutput, error) {
	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "users", input.UserSlug},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
func (c *BitbucketClient) GetUsers(ctx context.Context, input GetUsersInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "filter", input.Filter, "")
	client.SetQueryParam(queryParams, "permission", input.Permission, "")
	client.SetQueryParam(queryParams, "group", input.Group, "")

	// Handle permissionFilters map
	if input.PermissionFilters != nil {
		for key, value := range input.PermissionFilters {
			if key != "" && value != "" {
				queryParams.Set(key, value)
			}
		}
	}

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "users"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
