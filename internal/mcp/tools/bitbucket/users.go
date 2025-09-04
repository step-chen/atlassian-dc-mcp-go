package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"
	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCurrentUserHandler handles getting the current Bitbucket user
func (h *Handler) getCurrentUserHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get current user", func() (interface{}, error) {
		return h.client.GetCurrentUser()
	})
}

// getUserHandler handles getting a Bitbucket user
func (h *Handler) getUserHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get user", func() (interface{}, error) {
		userSlug, ok := tools.GetStringArg(args, "userSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid userSlug parameter")
		}

		return h.client.GetUser(userSlug)
	})
}

// getUsersHandler handles getting Bitbucket users
func (h *Handler) getUsersHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get users", func() (interface{}, error) {
		filter, _ := tools.GetStringArg(args, "filter")
		permission, _ := tools.GetStringArg(args, "permission")
		group, _ := tools.GetStringArg(args, "group")

		// Create an empty permission filter map
		permissionFilters := make(map[string]string)

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		// Add pagination to permission filters
		if start > 0 {
			permissionFilters["start"] = fmt.Sprintf("%d", start)
		}
		if limit > 0 {
			permissionFilters["limit"] = fmt.Sprintf("%d", limit)
		}

		return h.client.GetUsers(filter, permission, group, permissionFilters)
	})
}

// AddUserTools registers the user-related tools with the MCP server
func AddUserTools(server *mcp.Server, client *bitbucket.BitbucketClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_current_user",
		Description: "Get current Bitbucket user",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{},
		},
	}, handler.getCurrentUserHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_user",
		Description: "Get a Bitbucket user",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"userSlug": {
					Type:        "string",
					Description: "The slug of the user to retrieve",
				},
			},
			Required: []string{"userSlug"},
		},
	}, handler.getUserHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_users",
		Description: "Get Bitbucket users",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"filter": {
					Type:        "string",
					Description: "Filter users by name",
				},
				"permission": {
					Type:        "string",
					Description: "Filter users by permission",
				},
				"group": {
					Type:        "string",
					Description: "Filter users by group",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of users to return",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned users",
				},
			},
		},
	}, handler.getUsersHandler)
}