package bitbucket

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
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
func (c *BitbucketClient) GetProjects(input GetProjectsInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "name", input.Name, "")
	utils.SetQueryParam(queryParams, "permission", input.Permission, "")

	var projects types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects"},
		queryParams,
		nil,
		&projects,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return projects, nil
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
func (c *BitbucketClient) GetProject(input GetProjectInput) (types.MapOutput, error) {
	var project types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey},
		nil,
		nil,
		&project,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return project, nil
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
func (c *BitbucketClient) GetProjectPrimaryEnhancedEntityLink(input GetProjectPrimaryEnhancedEntityLinkInput) (types.MapOutput, error) {
	var entityLink types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "jira", "latest", "projects", input.ProjectKey, "primary-enhanced-entitylink"},
		nil,
		nil,
		&entityLink,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return entityLink, nil
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
func (c *BitbucketClient) GetProjectTasks(input GetProjectTasksInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "markup", input.Markup, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var tasks types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "default-tasks", "latest", "projects", input.ProjectKey, "tasks"},
		queryParams,
		nil,
		&tasks,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return tasks, nil
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
func (c *BitbucketClient) GetRepositoryTasks(input GetRepositoryTasksInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "markup", input.Markup, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var tasks types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "default-tasks", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "tasks"},
		queryParams,
		nil,
		&tasks,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return tasks, nil
}
