// Package bitbucket provides MCP tools for interacting with Bitbucket.
package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getBranchesHandler handles getting branches
func (h *Handler) getBranchesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetBranchesInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	branches, err := h.client.GetBranches(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get branches")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(branches)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create branches result")
		return result, nil, err
	}

	return result, branches, nil
}

// getDefaultBranchHandler handles getting the default branch
func (h *Handler) getDefaultBranchHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDefaultBranchInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	branch, err := h.client.GetDefaultBranch(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get default branch")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(branch)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create default branch result")
		return result, nil, err
	}

	return result, branch, nil
}

// getBranchInfoByCommitIdHandler handles getting branch information by commit ID
func (h *Handler) getBranchHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetBranchInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	branchInfo, err := h.client.GetBranch(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get branch info by commit id")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(branchInfo)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create branch info result")
		return result, nil, err
	}

	return result, branchInfo, nil
}

// AddBranchTools registers the branch-related tools with the MCP server
func AddBranchTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[bitbucket.GetBranchesInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "bitbucket_get_branches",
		Description: "Get branches for a repository",
	}, handler.getBranchesHandler)

	mcp.AddTool[bitbucket.GetDefaultBranchInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "bitbucket_get_default_branch",
		Description: "Get the default branch of a repository",
	}, handler.getDefaultBranchHandler)

	mcp.AddTool[bitbucket.GetBranchInput, map[string]interface{}](server, &mcp.Tool{
		Name:        "bitbucket_get_branch_info_by_commit_id",
		Description: "Get branch information by commit ID",
	}, handler.getBranchHandler)
}
