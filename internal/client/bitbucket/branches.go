// Package bitbucket provides a client for interacting with Bitbucket Data Center APIs.
package bitbucket

import (
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetBranches retrieves branches for a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch branches
// for a specific repository with various filtering and ordering options.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - filterText: Text to filter branches by
//   - orderBy: Field to order branches by
//   - context: Context for filtering
//   - base: Base branch for comparison
//   - boostMatches: Boost exact matches
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//   - details: Include detailed branch information
//
// Returns:
//   - map[string]interface{}: The branches data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetBranches(projectKey, repoSlug string, filterText, orderBy, context, base string, boostMatches bool, start, limit int, details bool) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "filterText", filterText, "")
	utils.SetQueryParam(queryParams, "orderBy", orderBy, "")
	utils.SetQueryParam(queryParams, "context", context, "")
	utils.SetQueryParam(queryParams, "base", base, "")
	utils.SetQueryParam(queryParams, "boostMatches", boostMatches, false)
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)
	utils.SetQueryParam(queryParams, "details", details, false)

	var branches map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "branches"},
		queryParams,
		nil,
		&branches,
	); err != nil {
		return nil, err
	}

	return branches, nil
}

// GetDefaultBranch retrieves the default branch of a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the default
// branch of a specific repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//
// Returns:
//   - map[string]interface{}: The default branch data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetDefaultBranch(projectKey, repoSlug string) (map[string]interface{}, error) {
	var branch map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "default-branch"},
		nil,
		nil,
		&branch,
	); err != nil {
		return nil, err
	}

	return branch, nil
}

// GetBranchInfoByCommitId retrieves branch information by commit ID.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch branch
// information for a specific commit ID.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - commitId: The commit ID to retrieve branch information for
//   - start: Starting index for pagination (default: 0)
//   - limit: Maximum number of results to return (default: 25)
//
// Returns:
//   - map[string]interface{}: The branch information retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetBranchInfoByCommitId(projectKey, repoSlug, commitId string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "start", start, 0)
	utils.SetQueryParam(queryParams, "limit", limit, 0)

	var branchInfo map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "branch-utils", "latest", "projects", projectKey, "repos", repoSlug, "branches", "info", commitId},
		queryParams,
		nil,
		&branchInfo,
	); err != nil {
		return nil, fmt.Errorf("failed to get branch info by commit id: %w", err)
	}

	return branchInfo, nil
}
