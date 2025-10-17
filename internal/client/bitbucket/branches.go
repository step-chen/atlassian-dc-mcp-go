// Package bitbucket provides a client for interacting with Bitbucket Data Center APIs.
package bitbucket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
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
//   - types.MapOutput: The branches data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetBranches(input GetBranchesInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "base", input.Base, "")
	utils.SetQueryParam(queryParams, "details", input.Details, false)
	utils.SetQueryParam(queryParams, "filterText", input.FilterText, "")
	utils.SetQueryParam(queryParams, "orderBy", input.OrderBy, "")
	utils.SetQueryParam(queryParams, "context", input.Context, "")
	utils.SetQueryParam(queryParams, "boostMatches", input.BoostMatches, false)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var branches types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "branches"},
		queryParams,
		nil,
		&branches,
		utils.AcceptJSON,
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
//   - types.MapOutput: The default branch data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetDefaultBranch(input GetDefaultBranchInput) (types.MapOutput, error) {
	var branch types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "default-branch"},
		nil,
		nil,
		&branch,
		utils.AcceptJSON,
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
//   - types.MapOutput: The branch information retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetBranch(input GetBranchInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var branch types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "branch-utils", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "branches", "info", input.CommitId},
		queryParams,
		nil,
		&branch,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return branch, nil
}

// CreateBranch creates a new branch in the specified repository.
//
// This function makes an HTTP POST request to the Bitbucket API to create a new branch.
// The authenticated user must have an effective REPO_WRITE permission to call this resource.
// If branch permissions are set up in the repository, the authenticated user must also
// have access to the branch name that is to be created.
//
// Parameters:
//   - input: CreateBranchInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The created branch data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) CreateBranch(input CreateBranchInput) (types.MapOutput, error) {
	payload := make(types.MapOutput)
	utils.SetRequestBodyParam(payload, "name", input.Name)
	utils.SetRequestBodyParam(payload, "startPoint", input.StartPoint)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal branch data: %w", err)
	}

	var branch types.MapOutput
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "branch-utils", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "branches"},
		nil,
		jsonPayload,
		&branch,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return branch, nil
}
