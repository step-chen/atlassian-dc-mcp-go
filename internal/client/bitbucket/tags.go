package bitbucket

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetTags retrieves tags for a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch tags
// for a specific repository with filtering and ordering options.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - filterText: Text to filter tags by
//   - orderBy: Field to order tags by
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The tags data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetTags(projectKey, repoSlug string, filterText string, orderBy string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "filterText", filterText, "")
	utils.SetQueryParam(queryParams, "orderBy", orderBy, "")
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)

	var tags map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "tags"},
		queryParams,
		nil,
		&tags,
	); err != nil {
		return nil, err
	}

	return tags, nil
}

// GetTag retrieves details of a specific tag.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch details
// of a specific tag.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - tagName: The name of the tag to retrieve
//
// Returns:
//   - map[string]interface{}: The tag data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetTag(projectKey, repoSlug, tagName string) (map[string]interface{}, error) {
	var tag map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "tags", tagName},
		nil,
		nil,
		&tag,
	); err != nil {
		return nil, err
	}

	return tag, nil
}
