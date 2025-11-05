package bitbucket

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// GetRepository retrieves a specific repository by project key and repository slug.
// Parameters:
//   - input: The input for retrieving repositories
//
// Returns:
//   - types.MapOutput: The repository data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepository(ctx context.Context, input GetRepositoryInput) (types.MapOutput, error) {
	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetRepositories retrieves a list of repositories.
// Parameters:
//   - input: The input for retrieving repositories
//
// Returns:
//   - types.MapOutput: The repositories data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepositories(ctx context.Context, input GetRepositoriesInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "projectName", input.ProjectName, "")
	client.SetQueryParam(queryParams, "projectKey", input.ProjectKey, "")
	client.SetQueryParam(queryParams, "name", input.Name, "")
	client.SetQueryParam(queryParams, "visibility", input.Visibility, "")
	client.SetQueryParam(queryParams, "permission", input.Permission, "")
	client.SetQueryParam(queryParams, "state", input.State, "")
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "archived", input.Archived, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "repos"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
func (c *BitbucketClient) GetProjectRepositories(ctx context.Context, input GetProjectRepositoriesInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetRepositoryLabels retrieves labels for a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch labels
// associated with a repository identified by its project key and repository slug.
//
// Parameters:
//   - input: GetRepositoryLabelsInput containing the parameters for the request
//
// Returns:
//   - []string: The labels retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepositoryLabels(ctx context.Context, input GetRepositoryLabelsInput) ([]string, error) {
	var labels []string
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "labels"},
		nil,
		nil,
		client.AcceptJSON,
		&labels,
	); err != nil {
		return nil, err
	}

	return labels, nil
}

// GetFileContent retrieves the content of a file from a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the raw content
// of a file from a repository identified by its project key, repository slug, and file path.
//
// Parameters:
//   - input: GetFileContentInput containing the parameters for the request
//
// Returns:
//   - []byte: The file content retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetFileContent(ctx context.Context, input GetFileContentInput) ([]byte, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "at", input.At, "")
	client.SetQueryParam(queryParams, "markup", input.Markup, "")
	client.SetQueryParam(queryParams, "htmlEscape", input.HtmlEscape, "")
	client.SetQueryParam(queryParams, "includeHeadingId", input.IncludeHeadingId, "")
	client.SetQueryParam(queryParams, "hardwrap", input.Hardwrap, "")

	respBody, err := client.ExecuteStream(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "raw", input.Path},
		queryParams,
		nil,
		client.AcceptText,
		0,
	)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	content, err := io.ReadAll(respBody)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return content, nil
}

// GetFiles retrieves a list of files from a specific repository.
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
func (c *BitbucketClient) GetFiles(ctx context.Context, input GetFilesInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "at", input.At, "")
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)

	pathSegments := []any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "files"}
	if input.Path != "" {
		pathSegments = append(pathSegments, input.Path)
	}

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		pathSegments,
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
func (c *BitbucketClient) GetChanges(ctx context.Context, input GetChangesInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "until", input.Until, "")
	client.SetQueryParam(queryParams, "since", input.Since, "")
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "changes"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
func (c *BitbucketClient) CompareChanges(ctx context.Context, input CompareChangesInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "from", input.From, "")
	client.SetQueryParam(queryParams, "to", input.To, "")
	client.SetQueryParam(queryParams, "fromRepo", input.FromRepo, "")
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "compare", "changes"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
func (c *BitbucketClient) GetForks(ctx context.Context, input GetForksInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "forks"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
func (c *BitbucketClient) GetReadme(ctx context.Context, input GetReadmeInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "at", input.At, "")
	client.SetQueryParam(queryParams, "markup", input.Markup, "")
	client.SetQueryParam(queryParams, "htmlEscape", input.HtmlEscape, "")
	client.SetQueryParam(queryParams, "includeHeadingId", input.IncludeHeadingId, "")
	client.SetQueryParam(queryParams, "hardwrap", input.Hardwrap, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "readme"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
func (c *BitbucketClient) GetRelatedRepositories(ctx context.Context, input GetRelatedRepositoriesInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "related"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
