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
//   - name: Filter projects by name
//   - permission: Filter projects by permission
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The projects data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjects(name string, permission string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)
	utils.SetQueryParam(queryParams, "name", name, "")
	utils.SetQueryParam(queryParams, "permission", permission, "")

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
//   - projectKey: The unique key of the project to retrieve
//
// Returns:
//   - map[string]interface{}: The project data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProject(projectKey string) (map[string]interface{}, error) {
	var project map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey},
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
//   - projectKey: The unique key of the project
//
// Returns:
//   - map[string]interface{}: The entity link data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjectPrimaryEnhancedEntityLink(projectKey string) (map[string]interface{}, error) {
	var entityLink map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "jira", "latest", "projects", projectKey, "primary-enhanced-entitylink"},
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
//   - projectKey: The unique key of the project
//   - markup: Markup format for the response
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The tasks data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetProjectTasks(projectKey string, markup string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "markup", markup, "")
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)

	var tasks map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "default-tasks", "latest", "projects", projectKey, "tasks"},
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
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - markup: Markup format for the response
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The tasks data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetRepositoryTasks(projectKey, repoSlug string, markup string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "markup", markup, "")
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)

	var tasks map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "default-tasks", "latest", "projects", projectKey, "repos", repoSlug, "tasks"},
		queryParams,
		nil,
		&tasks,
	); err != nil {
		return nil, err
	}

	return tasks, nil
}
