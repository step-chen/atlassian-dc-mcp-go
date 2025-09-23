package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getProjectsHandler handles getting projects
func (h *Handler) getProjectsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	projects, err := h.client.GetProjects(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get projects failed: %w", err)
	}

	return nil, projects, nil
}

// getProjectHandler handles getting a specific project
func (h *Handler) getProjectHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	project, err := h.client.GetProject(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get project failed: %w", err)
	}

	return nil, project, nil
}

// getProjectPrimaryEnhancedEntityLinkHandler handles getting the primary enhanced entity link for a project
func (h *Handler) getProjectPrimaryEnhancedEntityLinkHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectPrimaryEnhancedEntityLinkInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	link, err := h.client.GetProjectPrimaryEnhancedEntityLink(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get project primary enhanced entity link failed: %w", err)
	}

	return nil, link, nil
}

// getProjectTasksHandler handles getting tasks for a specific project
func (h *Handler) getProjectTasksHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectTasksInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	tasks, err := h.client.GetProjectTasks(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get project tasks failed: %w", err)
	}

	return nil, tasks, nil
}

// getRepositoryTasksHandler handles getting tasks for a specific repository
func (h *Handler) getRepositoryTasksHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoryTasksInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	tasks, err := h.client.GetRepositoryTasks(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get repository tasks failed: %w", err)
	}

	return nil, tasks, nil
}

// AddProjectTools registers the project-related tools with the MCP server
func AddProjectTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[bitbucket.GetProjectsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "bitbucket_get_projects",
		Description: "Get a list of projects",
	}, handler.getProjectsHandler)

	mcp.AddTool[bitbucket.GetProjectInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "bitbucket_get_project",
		Description: "Get a specific project by project key",
	}, handler.getProjectHandler)

	mcp.AddTool[bitbucket.GetProjectPrimaryEnhancedEntityLinkInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "bitbucket_get_project_primary_enhanced_entity_link",
		Description: "Get project's primary enhanced entity link",
	}, handler.getProjectPrimaryEnhancedEntityLinkHandler)

	mcp.AddTool[bitbucket.GetProjectTasksInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "bitbucket_get_project_tasks",
		Description: "Get tasks for a specific project",
	}, handler.getProjectTasksHandler)

	mcp.AddTool[bitbucket.GetRepositoryTasksInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "bitbucket_get_repository_tasks",
		Description: "Get tasks for a specific repository",
	}, handler.getRepositoryTasksHandler)
}
