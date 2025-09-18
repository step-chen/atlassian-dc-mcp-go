package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetProjectsOutput represents the output for getting projects
type GetProjectsOutput = map[string]interface{}

// GetProjectOutput represents the output for getting a specific project
type GetProjectOutput = map[string]interface{}

// GetProjectPrimaryEnhancedEntityLinkOutput represents the output for getting project's primary enhanced entity link
type GetProjectPrimaryEnhancedEntityLinkOutput = map[string]interface{}

// GetProjectTasksOutput represents the output for getting tasks for a specific project
type GetProjectTasksOutput = map[string]interface{}

// GetRepositoryTasksOutput represents the output for getting tasks for a specific repository
type GetRepositoryTasksOutput = map[string]interface{}

// getProjectsHandler handles getting projects
func (h *Handler) getProjectsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectsInput) (*mcp.CallToolResult, GetProjectsOutput, error) {
	projects, err := h.client.GetProjects(input)
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
func (h *Handler) getProjectHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectInput) (*mcp.CallToolResult, GetProjectOutput, error) {
	project, err := h.client.GetProject(input)
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
func (h *Handler) getProjectPrimaryEnhancedEntityLinkHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectPrimaryEnhancedEntityLinkInput) (*mcp.CallToolResult, GetProjectPrimaryEnhancedEntityLinkOutput, error) {
	link, err := h.client.GetProjectPrimaryEnhancedEntityLink(input)
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
func (h *Handler) getProjectTasksHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectTasksInput) (*mcp.CallToolResult, GetProjectTasksOutput, error) {
	tasks, err := h.client.GetProjectTasks(input)
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
func (h *Handler) getRepositoryTasksHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoryTasksInput) (*mcp.CallToolResult, GetRepositoryTasksOutput, error) {
	tasks, err := h.client.GetRepositoryTasks(input)
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

	mcp.AddTool[bitbucket.GetProjectsInput, GetProjectsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_projects",
		Description: "Get a list of projects",
	}, handler.getProjectsHandler)

	mcp.AddTool[bitbucket.GetProjectInput, GetProjectOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_project",
		Description: "Get a specific project by project key",
	}, handler.getProjectHandler)

	mcp.AddTool[bitbucket.GetProjectPrimaryEnhancedEntityLinkInput, GetProjectPrimaryEnhancedEntityLinkOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_project_primary_enhanced_entity_link",
		Description: "Get project's primary enhanced entity link",
	}, handler.getProjectPrimaryEnhancedEntityLinkHandler)

	mcp.AddTool[bitbucket.GetProjectTasksInput, GetProjectTasksOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_project_tasks",
		Description: "Get tasks for a specific project",
	}, handler.getProjectTasksHandler)

	mcp.AddTool[bitbucket.GetRepositoryTasksInput, GetRepositoryTasksOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repository_tasks",
		Description: "Get tasks for a specific repository",
	}, handler.getRepositoryTasksHandler)
}