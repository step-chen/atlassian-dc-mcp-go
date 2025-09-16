package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetProjectsInput represents the input parameters for getting projects
type GetProjectsInput struct {
	Name       string `json:"name,omitempty" jsonschema:"Filter projects by name"`
	Permission string `json:"permission,omitempty" jsonschema:"Filter projects by permission"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned projects"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of projects to return"`
}

// GetProjectsOutput represents the output for getting projects
type GetProjectsOutput = map[string]interface{}

// GetProjectInput represents the input parameters for getting a specific project
type GetProjectInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
}

// GetProjectOutput represents the output for getting a specific project
type GetProjectOutput = map[string]interface{}

// GetProjectPrimaryEnhancedEntityLinkInput represents the input parameters for getting project's primary enhanced entity link
type GetProjectPrimaryEnhancedEntityLinkInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
}

// GetProjectPrimaryEnhancedEntityLinkOutput represents the output for getting project's primary enhanced entity link
type GetProjectPrimaryEnhancedEntityLinkOutput = map[string]interface{}

// GetProjectTasksInput represents the input parameters for getting tasks for a specific project
type GetProjectTasksInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	Markup     string `json:"markup,omitempty" jsonschema:"Markup format for the response"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned tasks"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of tasks to return"`
}

// GetProjectTasksOutput represents the output for getting tasks for a specific project
type GetProjectTasksOutput = map[string]interface{}

// GetRepositoryTasksInput represents the input parameters for getting tasks for a specific repository
type GetRepositoryTasksInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Markup     string `json:"markup,omitempty" jsonschema:"Markup format for the response"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned tasks"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of tasks to return"`
}

// GetRepositoryTasksOutput represents the output for getting tasks for a specific repository
type GetRepositoryTasksOutput = map[string]interface{}

// getProjectsHandler handles getting projects
func (h *Handler) getProjectsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetProjectsInput) (*mcp.CallToolResult, GetProjectsOutput, error) {
	projects, err := h.client.GetProjects(input.Name, input.Permission, input.Start, input.Limit)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get projects")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(projects)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create projects result")
		return result, nil, err
	}

	return result, projects, nil
}

// getProjectHandler handles getting a specific project
func (h *Handler) getProjectHandler(ctx context.Context, req *mcp.CallToolRequest, input GetProjectInput) (*mcp.CallToolResult, GetProjectOutput, error) {
	project, err := h.client.GetProject(input.ProjectKey)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get project")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(project)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create project result")
		return result, nil, err
	}

	return result, project, nil
}

// getProjectPrimaryEnhancedEntityLinkHandler handles getting the primary enhanced entity link for a project
func (h *Handler) getProjectPrimaryEnhancedEntityLinkHandler(ctx context.Context, req *mcp.CallToolRequest, input GetProjectPrimaryEnhancedEntityLinkInput) (*mcp.CallToolResult, GetProjectPrimaryEnhancedEntityLinkOutput, error) {
	link, err := h.client.GetProjectPrimaryEnhancedEntityLink(input.ProjectKey)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get project primary enhanced entity link")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(link)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create project primary enhanced entity link result")
		return result, nil, err
	}

	return result, link, nil
}

// getProjectTasksHandler handles getting tasks for a specific project
func (h *Handler) getProjectTasksHandler(ctx context.Context, req *mcp.CallToolRequest, input GetProjectTasksInput) (*mcp.CallToolResult, GetProjectTasksOutput, error) {
	tasks, err := h.client.GetProjectTasks(input.ProjectKey, input.Markup, input.Start, input.Limit)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get project tasks")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(tasks)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create project tasks result")
		return result, nil, err
	}

	return result, tasks, nil
}

// getRepositoryTasksHandler handles getting tasks for a specific repository
func (h *Handler) getRepositoryTasksHandler(ctx context.Context, req *mcp.CallToolRequest, input GetRepositoryTasksInput) (*mcp.CallToolResult, GetRepositoryTasksOutput, error) {
	tasks, err := h.client.GetRepositoryTasks(input.ProjectKey, input.RepoSlug, input.Markup, input.Start, input.Limit)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get repository tasks")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(tasks)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create repository tasks result")
		return result, nil, err
	}

	return result, tasks, nil
}

// AddProjectTools registers the project-related tools with the MCP server
func AddProjectTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[GetProjectsInput, GetProjectsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_projects",
		Description: "Get a list of projects",
	}, handler.getProjectsHandler)

	mcp.AddTool[GetProjectInput, GetProjectOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_project",
		Description: "Get a specific project by project key",
	}, handler.getProjectHandler)

	mcp.AddTool[GetProjectPrimaryEnhancedEntityLinkInput, GetProjectPrimaryEnhancedEntityLinkOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_project_primary_enhanced_entity_link",
		Description: "Get project's primary enhanced entity link",
	}, handler.getProjectPrimaryEnhancedEntityLinkHandler)

	mcp.AddTool[GetProjectTasksInput, GetProjectTasksOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_project_tasks",
		Description: "Get tasks for a specific project",
	}, handler.getProjectTasksHandler)

	mcp.AddTool[GetRepositoryTasksInput, GetRepositoryTasksOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repository_tasks",
		Description: "Get tasks for a specific repository",
	}, handler.getRepositoryTasksHandler)
}
