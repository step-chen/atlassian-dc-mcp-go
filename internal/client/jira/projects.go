package jira

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
)

// GetProject retrieves a specific project by its key.
//
// Parameters:
//   - projectKey: The key of the project to retrieve
//
// Returns:
//   - types.MapOutput: The project data
//   - error: An error if the request fails
func (c *JiraClient) GetProject(input GetProjectInput) (types.MapOutput, error) {
	var project types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "project", input.ProjectKey}, nil, nil, &project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// GetAllProjects retrieves all projects with filtering options.
//
// Parameters:
//   - expand: Parameters to expand in the response
//   - recent: The number of recent projects to return
//   - includeArchived: Whether to include archived projects
//   - browseArchive: Whether to include projects in the archive browser
//
// Returns:
//   - []types.MapOutput: The projects data
//   - error: An error if the request fails
func (c *JiraClient) GetAllProjects(input GetAllProjectsInput) ([]types.MapOutput, error) {

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "expand", input.Expand, "")
	utils.SetQueryParam(queryParams, "recent", input.Recent, 0)
	utils.SetQueryParam(queryParams, "includeArchived", input.IncludeArchived, false)
	utils.SetQueryParam(queryParams, "browseArchive", input.BrowseArchive, false)

	var projects []types.MapOutput
	err := c.executeRequest(http.MethodGet, []string{"rest", "api", "2", "project"}, queryParams, nil, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
