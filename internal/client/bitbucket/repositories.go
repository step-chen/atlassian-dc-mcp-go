package bitbucket

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
)

// GetRepository retrieves details of a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch details
// of a repository identified by its project key and repository slug.
//
// Parameters:
//   - input: GetRepositoryInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The repository data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepository(input GetRepositoryInput) (types.MapOutput, error) {
	var repository types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug},
		nil,
		nil,
		&repository,
	); err != nil {
		return nil, err
	}

	return repository, nil
}

// GetRepositories retrieves repositories based on input parameters.
//
// This function makes an HTTP GET request to the Bitbucket API to retrieve
// repositories based on the provided input parameters.
//
// Parameters:
//   - input: GetRepositoriesInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The repositories data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepositories(input GetRepositoriesInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "projectName", input.ProjectName, "")
	utils.SetQueryParam(queryParams, "projectKey", input.ProjectKey, "")
	utils.SetQueryParam(queryParams, "name", input.Name, "")
	utils.SetQueryParam(queryParams, "visibility", input.Visibility, "")
	utils.SetQueryParam(queryParams, "permission", input.Permission, "")
	utils.SetQueryParam(queryParams, "state", input.State, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "archived", input.Archived, "")

	var repositories types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "repos"},
		queryParams,
		nil,
		&repositories,
	); err != nil {
		return nil, err
	}

	return repositories, nil
}

// GetProjectRepositories retrieves repositories for a specific project.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch repositories
// for a specific project with optional filtering by name.
//
// Parameters:
//   - input: GetProjectRepositoriesInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The repositories data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjectRepositories(input GetProjectRepositoriesInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var repositories types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos"},
		queryParams,
		nil,
		&repositories,
	); err != nil {
		return nil, err
	}

	return repositories, nil
}

// GetRepositoryLabels retrieves labels for a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch labels
// for a specific repository.
//
// Parameters:
//   - input: GetRepositoryLabelsInput containing the parameters for the request
//
// Returns:
//   - []string: The labels retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepositoryLabels(input GetRepositoryLabelsInput) ([]string, error) {
	var labels []string
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "labels"},
		nil,
		nil,
		&labels,
	); err != nil {
		return nil, err
	}

	return labels, nil
}

// GetFileContent retrieves the content of a file from a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the content
// of a specific file from a repository at a given commit or branch.
//
// Parameters:
//   - input: GetFileContentInput containing the parameters for the request
//
// Returns:
//   - []byte: The file content retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetFileContent(input GetFileContentInput) ([]byte, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "at", input.At, "")
	utils.SetQueryParam(queryParams, "size", input.Size, false)
	utils.SetQueryParam(queryParams, "type", input.TypeParam, false)
	utils.SetQueryParam(queryParams, "blame", input.Blame, false)
	utils.SetQueryParam(queryParams, "noContent", input.NoContent, false)

	var content []byte
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "raw", input.Path},
		queryParams,
		nil,
		&content,
	); err != nil {
		return nil, err
	}

	return content, nil
}

// GetFiles retrieves a list of files from a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch a list
// of files from a specific repository at a given commit or branch.
//
// Parameters:
//   - input: GetFilesInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The files data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetFiles(input GetFilesInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "at", input.At, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var files types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "files"},
		queryParams,
		nil,
		&files,
	); err != nil {
		return nil, err
	}

	return files, nil
}

// GetChanges retrieves changes in a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch changes
// in a specific repository at a given commit or branch.
//
// Parameters:
//   - input: GetChangesInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The changes data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetChanges(input GetChangesInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "until", input.Until, "")
	utils.SetQueryParam(queryParams, "since", input.Since, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var changes types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "changes"},
		queryParams,
		nil,
		&changes,
	); err != nil {
		return nil, err
	}

	return changes, nil
}

// CompareChanges compares changes between two commits or branches.
//
// This function makes an HTTP GET request to the Bitbucket API to compare changes
// between two commits or branches in a specific repository.
//
// Parameters:
//   - input: CompareChangesInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The comparison data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) CompareChanges(input CompareChangesInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "from", input.From, "")
	utils.SetQueryParam(queryParams, "to", input.To, "")
	utils.SetQueryParam(queryParams, "fromRepo", input.FromRepo, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var changes types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "compare", "changes"},
		queryParams,
		nil,
		&changes,
	); err != nil {
		return nil, err
	}

	return changes, nil
}

// GetForks retrieves forks of a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch forks
// of a specific repository.
//
// Parameters:
//   - input: GetForksInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The forks data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetForks(input GetForksInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var forks types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "forks"},
		queryParams,
		nil,
		&forks,
	); err != nil {
		return nil, err
	}

	return forks, nil
}

// GetReadme retrieves the README file for a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the README
// file for a specific repository at a given commit or branch.
//
// Parameters:
//   - input: GetReadmeInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The README data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetReadme(input GetReadmeInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "at", input.At, "")
	utils.SetQueryParam(queryParams, "markup", input.Markup, "")
	utils.SetQueryParam(queryParams, "htmlEscape", input.HtmlEscape, "")
	utils.SetQueryParam(queryParams, "includeHeadingId", input.IncludeHeadingId, "")
	utils.SetQueryParam(queryParams, "hardwrap", input.Hardwrap, "")

	var readme types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "readme"},
		queryParams,
		nil,
		&readme,
	); err != nil {
		return nil, err
	}

	return readme, nil
}

// GetRelatedRepositories retrieves repositories related to a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch repositories
// related to a specific repository.
//
// Parameters:
//   - input: GetRelatedRepositoriesInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The related repositories data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRelatedRepositories(input GetRelatedRepositoriesInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var repositories types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "related"},
		queryParams,
		nil,
		&repositories,
	); err != nil {
		return nil, err
	}

	return repositories, nil
}
