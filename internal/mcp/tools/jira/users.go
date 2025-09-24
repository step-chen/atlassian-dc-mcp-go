package jira

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCurrentUserHandler handles getting current user
func (h *Handler) getCurrentUserHandler(ctx context.Context, req *mcp.CallToolRequest, input types.EmptyInput) (*mcp.CallToolResult, types.MapOutput, error) {
	user, err := h.client.GetCurrentUser()
	if err != nil {
		return nil, nil, fmt.Errorf("get current user failed: %w", err)
	}

	return nil, user, nil
}

// getUserByNameHandler handles getting user by username
func (h *Handler) getUserByNameHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetUserByNameInput) (*mcp.CallToolResult, types.MapOutput, error) {
	user, err := h.client.GetUserByName(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get user by name failed: %w", err)
	}

	return nil, user, nil
}

// getUserByKeyHandler handles getting user by key
func (h *Handler) getUserByKeyHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.GetUserByKeyInput) (*mcp.CallToolResult, types.MapOutput, error) {
	user, err := h.client.GetUserByKey(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get user by key failed: %w", err)
	}

	return nil, user, nil
}

// searchUsersHandler handles searching users
func (h *Handler) searchUsersHandler(ctx context.Context, req *mcp.CallToolRequest, input jira.SearchUsersInput) (*mcp.CallToolResult, types.MapOutput, error) {
	users, err := h.client.SearchUsers(input)
	if err != nil {
		return nil, nil, fmt.Errorf("search users failed: %w", err)
	}

	wrappedResult := types.MapOutput{
		"users": users,
	}

	return nil, wrappedResult, nil
}

// AddUserTools registers the user-related tools with the MCP server
func AddUserTools(server *mcp.Server, client *jira.JiraClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[types.EmptyInput, types.MapOutput](server, "jira_get_current_user", "Get the current user", handler.getCurrentUserHandler)
	utils.RegisterTool[jira.GetUserByNameInput, types.MapOutput](server, "jira_get_user_by_name", "Get user by username", handler.getUserByNameHandler)
	utils.RegisterTool[jira.GetUserByKeyInput, types.MapOutput](server, "jira_get_user_by_key", "Get user by key", handler.getUserByKeyHandler)
	utils.RegisterTool[jira.SearchUsersInput, types.MapOutput](server, "jira_search_users", "Search for users", handler.searchUsersHandler)
}
