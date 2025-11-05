package bitbucket

import (
	"context"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// GetProjects retrieves a list of projects from Bitbucket.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch projects
// with optional filtering by name and permission.
//
// Parameters:
//   - input: GetProjectsInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The projects data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjects(ctx context.Context, input GetProjectsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "name", input.Name, "")
	client.SetQueryParam(queryParams, "permission", input.Permission, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetProject retrieves details of a specific project from Bitbucket.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch details
// of a project identified by its project key.
//
// Parameters:
//   - input: GetProjectInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The project data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProject(ctx context.Context, input GetProjectInput) (types.MapOutput, error) {
	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetProjectPrimaryEnhancedEntityLink retrieves the primary enhanced entity link for a project.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the primary
// enhanced entity link for a specific project.
//
// Parameters:
//   - input: GetProjectPrimaryEnhancedEntityLinkInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The entity link data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjectPrimaryEnhancedEntityLink(ctx context.Context, input GetProjectPrimaryEnhancedEntityLinkInput) (types.MapOutput, error) {
	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "jira", "latest", "projects", input.ProjectKey, "primary-enhanced-entitylink"},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetProjectTasks retrieves tasks associated with a specific project.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch tasks
// for a specific project with optional markup formatting.
//
// Parameters:
//   - input: GetProjectTasksInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The tasks data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjectTasks(ctx context.Context, input GetProjectTasksInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "markup", input.Markup, "")
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "default-tasks", "latest", "projects", input.ProjectKey, "tasks"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetRepositoryTasks retrieves tasks associated with a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch tasks
// for a specific repository with optional markup formatting.
//
// Parameters:
//   - input: GetRepositoryTasksInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The tasks data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepositoryTasks(ctx context.Context, input GetRepositoryTasksInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "markup", input.Markup, "")
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "default-tasks", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "tasks"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
