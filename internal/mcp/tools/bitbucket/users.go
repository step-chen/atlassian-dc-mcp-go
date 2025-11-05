package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getUserHandler handles getting a user
func (h *Handler) getUserHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetUserInput) (*mcp.CallToolResult, GetUserOutput, error) {
	user, err := h.client.GetUser(ctx, input)
	if err != nil {
		return nil, GetUserOutput{}, fmt.Errorf("get user failed: %w", err)
	}

	return nil, GetUserOutput{User: user}, nil
}

// getUsersHandler handles getting Bitbucket users
func (h *Handler) getUsersHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetUsersInput) (*mcp.CallToolResult, types.MapOutput, error) {
	users, err := h.client.GetUsers(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get users failed: %w", err)
	}

	return nil, users, nil
}

// AddUserTools registers the user-related tools with the MCP server
func AddUserTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[bitbucket.GetUserInput, GetUserOutput](server, "bitbucket_get_user", "Get a Bitbucket user", handler.getUserHandler)
	utils.RegisterTool[bitbucket.GetUsersInput, types.MapOutput](server, "bitbucket_get_users", "Get a list of Bitbucket users", handler.getUsersHandler)
}
