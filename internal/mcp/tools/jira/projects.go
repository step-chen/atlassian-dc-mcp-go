package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getProjectHandler handles getting a Jira project by key.
func (h *Handler) getProjectHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetProjectInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	project, err := h.client.GetProject(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get project failed: %w", err)
	}

	return nil, project, nil
}

// getProjectsHandler handles getting all Jira projects with optional filters and expansions.
func (h *Handler) getProjectsHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetAllProjectsInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	projects, err := h.client.GetAllProjects(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get projects failed: %w", err)
	}

	wrappedResult := map[string]interface{}{
		"projects": projects,
	}

	return nil, wrappedResult, nil
}

// AddProjectTools registers the project-related tools with the MCP server.
func AddProjectTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[jira.GetProjectInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_project",
		Description: "Get a specific Jira project by key",
	}, handler.getProjectHandler)

	mcp.AddTool[jira.GetAllProjectsInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_projects",
		Description: "Get all Jira projects with optional filters",
	}, handler.getProjectsHandler)
}
