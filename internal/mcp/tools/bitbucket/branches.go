// Package bitbucket provides MCP tools for interacting with Bitbucket.
package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getBranchesHandler handles getting branches
func (h *Handler) getBranchesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get branches", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		filterText, _ := tools.GetStringArg(args, "filterText")
		orderBy, _ := tools.GetStringArg(args, "orderBy")
		context, _ := tools.GetStringArg(args, "context")
		base, _ := tools.GetStringArg(args, "base")

		boostMatches := tools.GetBoolArg(args, "boostMatches", false)

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		details := tools.GetBoolArg(args, "details", false)

		return h.client.GetBranches(projectKey, repoSlug, filterText, orderBy, context, base, boostMatches, start, limit, details)
	})
}

// getDefaultBranchHandler handles getting the default branch
func (h *Handler) getDefaultBranchHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get default branch", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		return h.client.GetDefaultBranch(projectKey, repoSlug)
	})
}

// getBranchInfoByCommitIdHandler handles getting branch information by commit ID
func (h *Handler) getBranchInfoByCommitIdHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get branch info by commit id", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		commitId, ok := tools.GetStringArg(args, "commitId")
		if !ok {
			return nil, fmt.Errorf("missing or invalid commitId parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 25)

		return h.client.GetBranchInfoByCommitId(projectKey, repoSlug, commitId, start, limit)
	})
}

// AddBranchTools registers the branch-related tools with the MCP server
func AddBranchTools(server *mcp.Server, client *bitbucket.BitbucketClient) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_branches",
		Description: "Get branches for a repository",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"filterText": {
					Type:        "string",
					Description: "Filter text to apply to the branch names",
				},
				"orderBy": {
					Type:        "string",
					Description: "Field to order branches by",
				},
				"context": {
					Type:        "string",
					Description: "Context for filtering",
				},
				"base": {
					Type:        "string",
					Description: "Base branch for comparison",
				},
				"boostMatches": {
					Type:        "boolean",
					Description: "Boost exact matches",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned branches",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of branches to return",
				},
				"details": {
					Type:        "boolean",
					Description: "Include detailed branch information",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getBranchesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_default_branch",
		Description: "Get the default branch of a repository",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getDefaultBranchHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_branch_info_by_commit_id",
		Description: "Get branch information by commit ID",
		InputSchema: &jsonschema.Schema{
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"commitId": {
					Type:        "string",
					Description: "The commit ID to retrieve branch information for",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned results",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of results to return",
				},
			},
			Required: []string{"projectKey", "repoSlug", "commitId"},
		},
	}, handler.getBranchInfoByCommitIdHandler)

}