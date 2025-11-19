package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getProjectHandler handles getting a Jira project by key.
func (h *Handler) getProjectHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetProjectInput) (*mcp.CallToolResult, *jira.Project, error) {
	project, err := h.client.GetProject(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get project failed: %w", err)
	}

	return nil, project, nil
}

// getProjectsHandler handles getting all Jira projects with optional filters and expansions.
func (h *Handler) getProjectsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetAllProjectsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	projects, err := h.client.GetAllProjects(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get projects failed: %w", err)
	}

	wrappedResult := types.MapOutput{
		"projects": projects,
	}

	return nil, wrappedResult, nil
}

// AddProjectTools registers the project-related tools with the MCP server.
func AddProjectTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[jira.GetProjectInput, *jira.Project](server, "jira_get_project", "Get a specific Jira project by key", handler.getProjectHandler)
	utils.RegisterTool[jira.GetAllProjectsInput, types.MapOutput](server, "jira_get_projects", "Get all Jira projects with optional filters", handler.getProjectsHandler)
}
