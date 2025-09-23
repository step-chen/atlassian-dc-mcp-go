// Package bitbucket provides MCP tools for interacting with Bitbucket.
package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getBranchesHandler handles getting branches
func (h *Handler) getBranchesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetBranchesInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	branches, err := h.client.GetBranches(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get branches failed: %w", err)
	}

	return nil, branches, nil
}

// getDefaultBranchHandler handles getting the default branch
func (h *Handler) getDefaultBranchHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDefaultBranchInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	branch, err := h.client.GetDefaultBranch(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get default branch failed: %w", err)
	}

	return nil, branch, nil
}

// getBranchInfoByCommitIdHandler handles getting branch information by commit ID
func (h *Handler) getBranchHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetBranchInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	branchInfo, err := h.client.GetBranch(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get branch info by commit id failed: %w", err)
	}

	return nil, branchInfo, nil
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
