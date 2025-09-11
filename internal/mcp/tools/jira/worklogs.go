package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getWorklogsHandler handles getting worklogs for a Jira issue
func (h *Handler) getWorklogsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get worklogs", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		worklogId, _ := tools.GetStringArg(args, "worklogId")

		if worklogId != "" {
			return h.client.GetWorklogs(issueKey, worklogId)
		} else {
			return h.client.GetWorklogs(issueKey)
		}
	})
}

// AddWorklogTools registers the worklog-related tools with the MCP server
func AddWorklogTools(server *mcp.Server, client *jira.JiraClient, hasWritePermission bool) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_worklogs",
		Description: "Get worklogs for a Jira issue or a specific worklog by ID",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"issueKey": {
					Type:        "string",
					Description: "The key of the issue to get worklogs for",
				},
				"worklogId": {
					Type:        "string",
					Description: "Optional worklog ID to get a specific worklog",
				},
			},
			Required: []string{"issueKey"},
		},
	}, handler.getWorklogsHandler)
}