package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getIssueTypesHandler handles getting Jira issue types
func (h *Handler) getIssueTypesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get issue types", func() (interface{}, error) {
		issueTypes, err := h.client.GetIssueTypes()
		if err != nil {
			return nil, err
		}

		resultMap := map[string]interface{}{
			"issueTypes": issueTypes,
		}
		return resultMap, nil
	})
}

// AddIssueTypeTools registers the issue type-related tools with the MCP server
func AddIssueTypeTools(server *mcp.Server, client *jira.JiraClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_issue_types",
		Description: "Get Jira issue types",
		InputSchema: &jsonschema.Schema{
			Type:       "object",
			Properties: map[string]*jsonschema.Schema{},
		},
	}, handler.getIssueTypesHandler)
}
