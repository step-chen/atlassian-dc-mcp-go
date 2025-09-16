package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetCurrentUserOutput represents the output for getting the current user
type GetCurrentUserOutput = map[string]interface{}

// GetUserInput represents the input parameters for getting a user
type GetUserInput struct {
	UserSlug string `json:"userSlug" jsonschema:"required,The slug of the user to retrieve"`
}

// GetUserOutput represents the output for getting a user
type GetUserOutput = map[string]interface{}

// GetUsersInput represents the input parameters for getting users
type GetUsersInput struct {
	Filter     string `json:"filter,omitempty" jsonschema:"Filter users by name"`
	Permission string `json:"permission,omitempty" jsonschema:"Filter users by permission"`
	Group      string `json:"group,omitempty" jsonschema:"Filter users by group"`
	Start      int    `json:"start,omitempty" jsonschema:"Start index for pagination"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of users to return"`
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

// getUserHandler handles getting a Bitbucket user
func (h *Handler) getUserHandler(ctx context.Context, req *mcp.CallToolRequest, input GetUserInput) (*mcp.CallToolResult, GetUserOutput, error) {
	user, err := h.client.GetUser(input.UserSlug)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get user")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(user)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create user result")
		return result, nil, err
	}

	return result, user, nil
}

// getUsersHandler handles getting Bitbucket users
func (h *Handler) getUsersHandler(ctx context.Context, req *mcp.CallToolRequest, input GetUsersInput) (*mcp.CallToolResult, GetUsersOutput, error) {
	// Create an empty permission filter map
	permissionFilters := make(map[string]string)

	// Add pagination to permission filters
	if input.Start > 0 {
		permissionFilters["start"] = fmt.Sprintf("%d", input.Start)
	}
	if input.Limit > 0 {
		permissionFilters["limit"] = fmt.Sprintf("%d", input.Limit)
	}

	users, err := h.client.GetUsers(input.Filter, input.Permission, input.Group, permissionFilters)
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

	mcp.AddTool[GetUserInput, GetUserOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_user",
		Description: "Get a Bitbucket user",
	}, handler.getUserHandler)

	mcp.AddTool[GetUsersInput, GetUsersOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_users",
		Description: "Get a list of Bitbucket users",
	}, handler.getUsersHandler)
}
