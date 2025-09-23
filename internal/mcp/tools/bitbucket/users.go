package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/utils"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCurrentUserHandler handles getting the current Bitbucket user
func (h *Handler) getCurrentUserHandler(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, MapOutput, error) {
	user, err := h.client.GetCurrentUser()
	if err != nil {
		return nil, nil, fmt.Errorf("get current user failed: %w", err)
	}

	return nil, user, nil
}

// getUserHandler handles getting a user
func (h *Handler) getUserHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetUserInput) (*mcp.CallToolResult, GetUserOutput, error) {
	user, err := h.client.GetUser(input)
	if err != nil {
		return nil, GetUserOutput{}, fmt.Errorf("get user failed: %w", err)
	}

	return nil, GetUserOutput{User: user}, nil
}

// getUsersHandler handles getting Bitbucket users
func (h *Handler) getUsersHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetUsersInput) (*mcp.CallToolResult, MapOutput, error) {
	users, err := h.client.GetUsers(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get users failed: %w", err)
	}

	return nil, users, nil
}

// AddUserTools registers the user-related tools with the MCP server
func AddUserTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[struct{}, MapOutput](server, "bitbucket_get_current_user", "Get current Bitbucket user", handler.getCurrentUserHandler)
	utils.RegisterTool[bitbucket.GetUserInput, GetUserOutput](server, "bitbucket_get_user", "Get a Bitbucket user", handler.getUserHandler)
	utils.RegisterTool[bitbucket.GetUsersInput, MapOutput](server, "bitbucket_get_users", "Get a list of Bitbucket users", handler.getUsersHandler)
}