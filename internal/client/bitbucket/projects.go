package bitbucket

import (
	"net/http"
	"net/url"

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
//   - map[string]interface{}: The projects data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjects(input GetProjectsInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "name", input.Name, "")
	utils.SetQueryParam(queryParams, "permission", input.Permission, "")

	var projects map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects"},
		queryParams,
		nil,
		&projects,
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
//   - map[string]interface{}: The project data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProject(input GetProjectInput) (map[string]interface{}, error) {
	var project map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey},
		nil,
		nil,
		&project,
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
//   - map[string]interface{}: The entity link data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjectPrimaryEnhancedEntityLink(input GetProjectPrimaryEnhancedEntityLinkInput) (map[string]interface{}, error) {
	var entityLink map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "jira", "latest", "projects", input.ProjectKey, "primary-enhanced-entitylink"},
		nil,
		nil,
		&entityLink,
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
//   - map[string]interface{}: The tasks data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjectTasks(input GetProjectTasksInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "markup", input.Markup, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var tasks map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "default-tasks", "latest", "projects", input.ProjectKey, "tasks"},
		queryParams,
		nil,
		&tasks,
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
//   - map[string]interface{}: The tasks data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepositoryTasks(input GetRepositoryTasksInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "markup", input.Markup, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var tasks map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "default-tasks", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "tasks"},
		queryParams,
		nil,
		&tasks,
	); err != nil {
		return nil, err
	}

	return tasks, nil
}
