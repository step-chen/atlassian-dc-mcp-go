package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetCurrentUserOutput represents the output for getting the current user
type GetCurrentUserOutput = map[string]interface{}

// GetUserOutput represents the output for getting a user
type GetUserOutput struct {
	User map[string]interface{} `json:"user" jsonschema:"the user details"`
}

// GetUsersOutput represents the output for getting users
type GetUsersOutput = map[string]interface{}

// getCurrentUserHandler handles getting the current Bitbucket user
func (h *Handler) getCurrentUserHandler(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, GetCurrentUserOutput, error) {
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

// getUserHandler handles getting a user
func (h *Handler) getUserHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetUserInput) (*mcp.CallToolResult, GetUserOutput, error) {
	user, err := h.client.GetUser(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get user")
		return result, GetUserOutput{}, err
	}

	result, err := tools.CreateToolResult(user)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create user result")
		return result, GetUserOutput{}, err
	}

	return result, GetUserOutput{User: user}, nil
}

// getUsersHandler handles getting Bitbucket users
func (h *Handler) getUsersHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetUsersInput) (*mcp.CallToolResult, GetUsersOutput, error) {
	users, err := h.client.GetUsers(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get users")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(users)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create users result")
		return result, nil, err
	}

	return result, users, nil
}

// AddUserTools registers the user-related tools with the MCP server
func AddUserTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[struct{}, GetCurrentUserOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_current_user",
		Description: "Get current Bitbucket user",
	}, handler.getCurrentUserHandler)

	mcp.AddTool[bitbucket.GetUserInput, GetUserOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_user",
		Description: "Get a Bitbucket user",
	}, handler.getUserHandler)

	mcp.AddTool[bitbucket.GetUsersInput, GetUsersOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_users",
		Description: "Get a list of Bitbucket users",
	}, handler.getUsersHandler)
}
