package bitbucket

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// GetTags retrieves tags with filtering options.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch tags
// with optional filtering by name.
//
// Parameters:
//   - input: GetTagsInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The tags data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetTags(input GetTagsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "filterText", input.FilterText, "")
	client.SetQueryParam(queryParams, "orderBy", input.OrderBy, "")
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "tags"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetTag retrieves a specific tag.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch a specific
// tag identified by its name.
//
// Parameters:
//   - input: GetTagInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The tag data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetTag(input GetTagInput) (types.MapOutput, error) {
	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "tags", input.Name},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
