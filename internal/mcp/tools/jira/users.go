package jira

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCurrentUserHandler handles getting current user
func (h *Handler) getCurrentUserHandler(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	user, err := h.client.GetCurrentUser()
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get current user")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(user)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create current user result")
		return result, nil, err
	}

	return result, user, nil
}

// getUserByNameHandler handles getting user by username
func (h *Handler) getUserByNameHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetUserByNameInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	user, err := h.client.GetUserByName(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get user by name")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(user)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create user by name result")
		return result, nil, err
	}

	return result, user, nil
}

// getUserByKeyHandler handles getting user by key
func (h *Handler) getUserByKeyHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetUserByKeyInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	user, err := h.client.GetUserByKey(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get user by key")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(user)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create user by key result")
		return result, nil, err
	}

	return result, user, nil
}

// searchUsersHandler handles searching users
func (h *Handler) searchUsersHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.SearchUsersInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	users, err := h.client.SearchUsers(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "search users")
		return result, nil, err
	}

	wrappedResult := map[string]interface{}{
		"users": users,
	}

	result, err := tools.CreateToolResult(wrappedResult)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create wrapped search users result")
		return result, nil, err
	}

	return result, wrappedResult, nil
}

// AddUserTools registers the user-related tools with the MCP server
func AddUserTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[struct{}, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_current_user",
		Description: "Get the current user",
	}, handler.getCurrentUserHandler)

	mcp.AddTool[jira.GetUserByNameInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_user_by_name",
		Description: "Get user by username",
	}, handler.getUserByNameHandler)

	mcp.AddTool[jira.GetUserByKeyInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_get_user_by_key",
		Description: "Get user by key",
	}, handler.getUserByKeyHandler)

	mcp.AddTool[jira.SearchUsersInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "jira_search_users",
		Description: "Search for users",
	}, handler.searchUsersHandler)
}
