package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCurrentUserHandler handles getting current user
func (h *Handler) getCurrentUserHandler(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	user, err := h.client.GetCurrentUser()
	if err != nil {
		return nil, nil, fmt.Errorf("get current user failed: %w", err)
	}

	return nil, user, nil
}

// getUserByNameHandler handles getting user by username
func (h *Handler) getUserByNameHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetUserByNameInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	user, err := h.client.GetUserByName(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get user by name failed: %w", err)
	}

	return nil, user, nil
}

// getUserByKeyHandler handles getting user by key
func (h *Handler) getUserByKeyHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetUserByKeyInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	user, err := h.client.GetUserByKey(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get user by key failed: %w", err)
	}

	return nil, user, nil
}

// searchUsersHandler handles searching users
func (h *Handler) searchUsersHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.SearchUsersInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	users, err := h.client.SearchUsers(input)
	if err != nil {
		return nil, nil, fmt.Errorf("search users failed: %w", err)
	}

	wrappedResult := map[string]interface{}{
		"users": users,
	}

	return nil, wrappedResult, nil
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
