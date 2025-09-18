// Package bitbucket provides a client for interacting with Bitbucket Data Center APIs.
package bitbucket

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetBranches retrieves branches with filtering options.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch branches
// with various filtering options.
//
// Parameters:
//   - input: GetBranchesInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The branches data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetBranches(input GetBranchesInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "base", input.Base, "")
	utils.SetQueryParam(queryParams, "details", input.Details, false)
	utils.SetQueryParam(queryParams, "filterText", input.FilterText, "")
	utils.SetQueryParam(queryParams, "orderBy", input.OrderBy, "")
	utils.SetQueryParam(queryParams, "context", input.Context, "")
	utils.SetQueryParam(queryParams, "boostMatches", input.BoostMatches, false)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var branches map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "branches"},
		queryParams,
		nil,
		&branches,
	); err != nil {
		return nil, err
	}

	return branches, nil
}

// GetDefaultBranch retrieves the default branch for a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the default
// branch for a repository.
//
// Parameters:
//   - input: GetDefaultBranchInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The default branch data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetDefaultBranch(input GetDefaultBranchInput) (map[string]interface{}, error) {
	var branch map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "default-branch"},
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
// information by commit ID.
//
// Parameters:
//   - input: GetBranchInfoByCommitIdInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The branch information retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetBranch(input GetBranchInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var branch map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "branch-utils", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "branches", "info", input.CommitId},
		queryParams,
		nil,
		&branch,
	); err != nil {
		return nil, err
	}

	return branch, nil
}
