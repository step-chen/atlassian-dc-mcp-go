package bitbucket

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
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
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "filterText", input.FilterText, "")
	utils.SetQueryParam(queryParams, "orderBy", input.OrderBy, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var tags types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "tags"},
		queryParams,
		nil,
		&tags,
	); err != nil {
		return nil, err
	}

	return tags, nil
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
	var tag types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "tags", input.Name},
		nil,
		nil,
		&tag,
	); err != nil {
		return nil, err
	}

	return tag, nil
}
