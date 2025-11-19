package jira

import (
	"context"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
)

// GetProject retrieves a specific project by its key.
//
// Parameters:
//   - projectKey: The key of the project to retrieve
//
// Returns:
//   - types.MapOutput: The project data
//   - error: An error if the request fails
func (c *JiraClient) GetProject(ctx context.Context, input GetProjectInput) (*Project, error) {
	var output *Project
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "2", "project", input.ProjectKey},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// GetAllProjects retrieves all projects with filtering options.
//
// Parameters:
//   - input: The input for retrieving projects
//
// Returns:
//   - []types.MapOutput: The projects data
//   - error: An error if the request fails
func (c *JiraClient) GetAllProjects(ctx context.Context, input GetAllProjectsInput) ([]Project, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "expand", input.Expand, "")
	client.SetQueryParam(queryParams, "recent", input.Recent, 0)
	client.SetQueryParam(queryParams, "includeArchived", input.IncludeArchived, false)
	client.SetQueryParam(queryParams, "browseArchive", input.BrowseArchive, false)

	var outputs []Project
	err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "2", "project"},
		queryParams,
		nil,
		client.AcceptJSON,
		&outputs,
	)
	if err != nil {
		return nil, err
	}

	return outputs, nil
}
