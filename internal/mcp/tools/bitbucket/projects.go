package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getProjectsHandler handles getting projects
func (h *Handler) getProjectsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get projects", func() (interface{}, error) {
		name, _ := tools.GetStringArg(args, "name")
		permission, _ := tools.GetStringArg(args, "permission")

		limit := tools.GetIntArg(args, "limit", 25)
		start := tools.GetIntArg(args, "start", 0)

		return h.client.GetProjects(name, permission, start, limit)
	})
}

// getProjectHandler handles getting a specific project
func (h *Handler) getProjectHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get project", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		return h.client.GetProject(projectKey)
	})
}

// getProjectPrimaryEnhancedEntityLinkHandler handles getting the primary enhanced entity link for a project
func (h *Handler) getProjectPrimaryEnhancedEntityLinkHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get project primary enhanced entity link", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		return h.client.GetProjectPrimaryEnhancedEntityLink(projectKey)
	})
}

// getProjectTasksHandler handles getting tasks for a specific project
func (h *Handler) getProjectTasksHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	projectKey, ok := tools.GetStringArg(args, "projectKey")
	if !ok {
		return tools.HandleToolError(fmt.Errorf("missing or invalid projectKey parameter"), "get project tasks")
	}

	markup, _ := tools.GetStringArg(args, "markup")

	limit := tools.GetIntArg(args, "limit", 25)
	start := tools.GetIntArg(args, "start", 0)

	tasks, err := h.client.GetProjectTasks(projectKey, markup, start, limit)
	if err != nil {
		return tools.HandleToolError(fmt.Errorf("failed to get project tasks: %w", err), "get project tasks")
	}

	result, err := tools.CreateToolResult(tasks)
	if err != nil {
		return tools.HandleToolError(fmt.Errorf("failed to create result: %w", err), "get project tasks")
	}

	return result, tasks, nil
}

// getRepositoryTasksHandler handles getting tasks for a specific repository
func (h *Handler) getRepositoryTasksHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	projectKey, ok := tools.GetStringArg(args, "projectKey")
	if !ok {
		return tools.HandleToolError(fmt.Errorf("missing or invalid projectKey parameter"), "get repository tasks")
	}

	repoSlug, ok := tools.GetStringArg(args, "repoSlug")
	if !ok {
		return tools.HandleToolError(fmt.Errorf("missing or invalid repoSlug parameter"), "get repository tasks")
	}

	markup, _ := tools.GetStringArg(args, "markup")

	limit := tools.GetIntArg(args, "limit", 25)
	start := tools.GetIntArg(args, "start", 0)

	tasks, err := h.client.GetRepositoryTasks(projectKey, repoSlug, markup, start, limit)
	if err != nil {
		return tools.HandleToolError(fmt.Errorf("failed to get repository tasks: %w", err), "get repository tasks")
	}

	result, err := tools.CreateToolResult(tasks)
	if err != nil {
		return tools.HandleToolError(fmt.Errorf("failed to create result: %w", err), "get repository tasks")
	}

	return result, tasks, nil
}

// AddProjectTools registers the project-related tools with the MCP server
func AddProjectTools(server *mcp.Server, client *bitbucket.BitbucketClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_projects",
		Description: "Get projects",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"name": {
					Type:        "string",
					Description: "Filter projects by name",
				},
				"permission": {
					Type:        "string",
					Description: "Filter projects by permission",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of projects to return",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned projects",
				},
			},
		},
	}, handler.getProjectsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_project",
		Description: "Get a specific project by key",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
			},
			Required: []string{"projectKey"},
		},
	}, handler.getProjectHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_project_primary_enhanced_entity_link",
		Description: "Get project's primary enhanced entity link",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
			},
			Required: []string{"projectKey"},
		},
	}, handler.getProjectPrimaryEnhancedEntityLinkHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_project_tasks",
		Description: "Get tasks for a specific project",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"markup": {
					Type:        "string",
					Description: "Markup format for the response",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of tasks to return",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned tasks",
				},
			},
			Required: []string{"projectKey"},
		},
	}, handler.getProjectTasksHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_repository_tasks",
		Description: "Get tasks for a specific repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"markup": {
					Type:        "string",
					Description: "Markup format for the response",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of tasks to return",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned tasks",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getRepositoryTasksHandler)
}
