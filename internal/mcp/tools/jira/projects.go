package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getProjectHandler handles getting a Jira project by key.
func (h *Handler) getProjectHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get project", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		return h.client.GetProject(projectKey)
	})
}

// getProjectsHandler handles getting all Jira projects with optional filters and expansions.
func (h *Handler) getProjectsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get projects", func() (interface{}, error) {
		expand, _ := tools.GetStringArg(args, "expand")

		recent := tools.GetIntArg(args, "recent", 0)
		includeArchived := tools.GetBoolArg(args, "includeArchived", false)
		browseArchive := tools.GetBoolArg(args, "browseArchive", false)

		return h.client.GetAllProjects(expand, recent, includeArchived, browseArchive)
	})
}

// AddProjectTools registers the project-related tools with the MCP server.
func AddProjectTools(server *mcp.Server, client *jira.JiraClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_project",
		Description: "Get a specific Jira project by key",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key to identify the project",
				},
			},
			Required: []string{"projectKey"},
		},
	}, handler.getProjectHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_projects",
		Description: "Get all Jira projects with optional filters",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"expand": {
					Type:        "string",
					Description: "Fields to expand in the response",
				},
				"recent": {
					Type:        "integer",
					Description: "Number of recent projects to include in the response",
				},
				"includeArchived": {
					Type:        "boolean",
					Description: "Include archived projects in the results",
				},
				"browseArchive": {
					Type:        "boolean",
					Description: "Include projects in the browse archive",
				},
			},
		},
	}, handler.getProjectsHandler)
}
