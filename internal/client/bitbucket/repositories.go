package bitbucket

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetRepository retrieves details of a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch details
// of a repository identified by its project key and repository slug.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//
// Returns:
//   - map[string]interface{}: The repository data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepository(projectKey, repoSlug string) (map[string]interface{}, error) {
	var repo map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug},
		nil,
		nil,
		&repo,
	); err != nil {
		return nil, err
	}

	return repo, nil
}

// GetRepositories retrieves repositories with various filters.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch repositories
// with various filtering options.
//
// Parameters:
//   - projectName: Filter repositories by project name
//   - projectKey: Filter repositories by project key
//   - name: Filter repositories by name
//   - visibility: Filter repositories by visibility
//   - permission: Filter repositories by permission
//   - state: Filter repositories by state
//   - archived: Filter archived repositories
//   - username: Filter repositories by username
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The repositories data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepositories(projectName, projectKey, name, visibility, permission, state, archived, username string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "projectname", projectName, "")
	utils.SetQueryParam(queryParams, "projectkey", projectKey, "")
	utils.SetQueryParam(queryParams, "name", name, "")
	utils.SetQueryParam(queryParams, "visibility", visibility, "")
	utils.SetQueryParam(queryParams, "permission", permission, "")
	utils.SetQueryParam(queryParams, "state", state, "")
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)

	if archived != "" {
		queryParams.Set("archived", archived)
	}

	var repos map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "repos"},
		queryParams,
		nil,
		&repos,
	); err != nil {
		return nil, err
	}

	return repos, nil
}

// GetProjectRepositories retrieves repositories for a specific project.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch repositories
// associated with a specific project.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The repositories data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjectRepositories(projectKey string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)

	var repos map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos"},
		queryParams,
		nil,
		&repos,
	); err != nil {
		return nil, err
	}

	return repos, nil
}

// GetRepositoryLabels retrieves labels for a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch labels
// associated with a specific repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//
// Returns:
//   - map[string]interface{}: The labels data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepositoryLabels(projectKey, repoSlug string) (map[string]interface{}, error) {
	var labels map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "labels"},
		nil,
		nil,
		&labels,
	); err != nil {
		return nil, err
	}

	return labels, nil
}

// GetFileContent retrieves the content of a file in a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the content
// of a file in a specific repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - path: The path to the file
//   - at: The commit ID or ref to retrieve the file at
//   - size: Include file size information
//   - typeParam: Include file type information
//   - blame: Include blame information
//   - noContent: Skip content retrieval
//
// Returns:
//   - string: The file content as a string
//   - error: An error if the request fails
func (c *BitbucketClient) GetFileContent(projectKey, repoSlug, path, at string, size, typeParam *bool, blame, noContent *bool) (string, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "at", at, "")
	utils.SetQueryParam(queryParams, "size", size, (*bool)(nil))
	utils.SetQueryParam(queryParams, "type", typeParam, (*bool)(nil))
	utils.SetQueryParam(queryParams, "blame", blame, (*bool)(nil))
	utils.SetQueryParam(queryParams, "noContent", noContent, (*bool)(nil))

	return c.executeTextRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "raw", path},
		queryParams,
		nil,
	)
}

// GetFiles retrieves files in a directory of a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch files
// in a specific directory of a repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - path: The path to the directory
//   - at: The commit ID or ref to retrieve files at
//   - start: Starting index for pagination (default: 0)
//   - limit: Maximum number of results to return (default: 25)
//
// Returns:
//   - map[string]interface{}: The files data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetFiles(projectKey, repoSlug, path, at string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "at", at, "")
	utils.SetQueryParam(queryParams, "start", start, 0)
	utils.SetQueryParam(queryParams, "limit", limit, 0)

	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "files", path},
		queryParams,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return result, nil
}

// GetChanges retrieves changes between commits in a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch changes
// between commits in a specific repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - until: The commit ID or ref to compare until
//   - since: The commit ID or ref to compare since
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The changes data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetChanges(projectKey, repoSlug, until, since string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "until", until, "")
	utils.SetQueryParam(queryParams, "since", since, "")
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)

	var changes map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "changes"},
		queryParams,
		nil,
		&changes,
	); err != nil {
		return nil, err
	}

	return changes, nil
}

// CompareChanges compares changes between two commits in a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to compare changes
// between two commits in a specific repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - from: The source commit ID or ref
//   - to: The target commit ID or ref
//   - fromRepo: The source repository
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The changes data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) CompareChanges(projectKey, repoSlug, from, to, fromRepo string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "from", from, "")
	utils.SetQueryParam(queryParams, "to", to, "")
	utils.SetQueryParam(queryParams, "fromRepo", fromRepo, "")
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)

	var changes map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "compare", "changes"},
		queryParams,
		nil,
		&changes,
	); err != nil {
		return nil, err
	}

	return changes, nil
}

// GetForks retrieves forks of a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch forks
// of a specific repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The forks data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetForks(projectKey, repoSlug string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)

	var forks map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "forks"},
		queryParams,
		nil,
		&forks,
	); err != nil {
		return nil, err
	}

	return forks, nil
}

// GetReadme retrieves the README file of a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the README
// file of a specific repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - at: The commit ID or ref to retrieve the README at
//   - markup: Markup format for the response
//   - htmlEscape: HTML escape option
//   - includeHeadingId: Include heading IDs
//   - hardwrap: Hard wrap option
//
// Returns:
//   - string: The README content as a string
//   - error: An error if the request fails
func (c *BitbucketClient) GetReadme(projectKey, repoSlug string, at, markup, htmlEscape, includeHeadingId, hardwrap *string) (string, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "at", at, (*string)(nil))
	utils.SetQueryParam(queryParams, "markup", markup, (*string)(nil))
	utils.SetQueryParam(queryParams, "htmlEscape", htmlEscape, (*string)(nil))
	utils.SetQueryParam(queryParams, "includeHeadingId", includeHeadingId, (*string)(nil))
	utils.SetQueryParam(queryParams, "hardwrap", hardwrap, (*string)(nil))

	return c.executeTextRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "readme"},
		queryParams,
		nil,
	)
}

// GetAttachment retrieves an attachment from a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch an attachment
// from a specific repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - attachmentId: The ID of the attachment to retrieve
//
// Returns:
//   - []byte: The attachment content as bytes
//   - error: An error if the request fails
func (c *BitbucketClient) GetAttachment(projectKey, repoSlug, attachmentId string) ([]byte, error) {
	var content []byte
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "attachments", attachmentId},
		nil,
		nil,
		&content,
	); err != nil {
		return nil, err
	}

	return content, nil
}

// GetAttachmentMetadata retrieves metadata for an attachment.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch metadata
// for an attachment from a specific repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - attachmentId: The ID of the attachment to retrieve metadata for
//
// Returns:
//   - map[string]interface{}: The attachment metadata retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetAttachmentMetadata(projectKey, repoSlug, attachmentId string) (map[string]interface{}, error) {
	var metadata map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "attachments", attachmentId, "metadata"},
		nil,
		nil,
		&metadata,
	); err != nil {
		return nil, err
	}

	return metadata, nil
}

// GetRelatedRepositories retrieves repositories related to a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch repositories
// related to a specific repository.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - start: Starting index for pagination (default: 0)
//   - limit: Maximum number of results to return (default: 25)
//
// Returns:
//   - map[string]interface{}: The related repositories data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRelatedRepositories(projectKey, repoSlug string, start, limit int) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)

	var relatedRepos map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "related"},
		params,
		nil,
		&relatedRepos,
	); err != nil {
		return nil, err
	}

	return relatedRepos, nil
}
