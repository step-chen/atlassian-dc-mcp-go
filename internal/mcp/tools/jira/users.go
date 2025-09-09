package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCurrentUserHandler handles getting current user
func (h *Handler) getCurrentUserHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get current user", func() (interface{}, error) {
		return h.client.GetCurrentUser()
	})
}

// getUserByNameHandler handles getting user by username
func (h *Handler) getUserByNameHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	username, ok := tools.GetStringArg(args, "username")
	if !ok {
		return nil, nil, fmt.Errorf("missing or invalid username parameter")
	}

	return tools.HandleToolOperation("get user by name", func() (interface{}, error) {
		return h.client.GetUserByName(username)
	})
}

// getUserByKeyHandler handles getting user by key
func (h *Handler) getUserByKeyHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	key, ok := tools.GetStringArg(args, "key")
	if !ok {
		return nil, nil, fmt.Errorf("missing or invalid key parameter")
	}

	return tools.HandleToolOperation("get user by key", func() (interface{}, error) {
		return h.client.GetUserByKey(key)
	})
}

// searchUsersHandler handles searching users
func (h *Handler) searchUsersHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	query, _ := tools.GetStringArg(args, "query")
	startAt := tools.GetIntArg(args, "startAt", 0)
	maxResults := tools.GetIntArg(args, "maxResults", 50)

	return tools.HandleToolOperation("search users", func() (interface{}, error) {
		return h.client.SearchUsers(query, startAt, maxResults)
	})
}

// AddUserTools registers the user-related tools with the MCP server
func AddUserTools(server *mcp.Server, client *jira.JiraClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_current_user",
		Description: "Get current Jira user",
		InputSchema: &jsonschema.Schema{
			Type:       "object",
			Properties: map[string]*jsonschema.Schema{},
		},
	}, handler.getCurrentUserHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_user_by_name",
		Description: "Get a Jira user by username",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"username": {
					Type:        "string",
					Description: "The username of the user to retrieve",
				},
			},
			Required: []string{"username"},
		},
	}, handler.getUserByNameHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_get_user_by_key",
		Description: "Get a Jira user by key",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"key": {
					Type:        "string",
					Description: "The key of the user to retrieve",
				},
			},
			Required: []string{"key"},
		},
	}, handler.getUserByKeyHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "jira_search_users",
		Description: "Search for Jira users",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"query": {
					Type:        "string",
					Description: "Search query string",
				},
				"startAt": {
					Type:        "integer",
					Description: "Starting index for results",
				},
				"maxResults": {
					Type:        "integer",
					Description: "Maximum number of results to return",
				},
			},
			Required: []string{"query"},
		},
	}, handler.searchUsersHandler)
}