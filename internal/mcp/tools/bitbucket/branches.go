// Package bitbucket provides MCP tools for interacting with Bitbucket.
package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetBranchesInput represents the input parameters for getting branches
type GetBranchesInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	FilterText   string `json:"filterText,omitempty" jsonschema:"Filter text to apply to the branch names"`
	OrderBy      string `json:"orderBy,omitempty" jsonschema:"Field to order branches by"`
	Context      string `json:"context,omitempty" jsonschema:"Context for filtering"`
	Base         string `json:"base,omitempty" jsonschema:"Base branch for comparison"`
	BoostMatches bool   `json:"boostMatches,omitempty" jsonschema:"Boost exact matches"`
	Start        int    `json:"start,omitempty" jsonschema:"The starting index of the returned branches"`
	Limit        int    `json:"limit,omitempty" jsonschema:"The limit of the number of branches to return"`
	Details      bool   `json:"details,omitempty" jsonschema:"Include detailed branch information"`
}

// GetBranchesOutput represents the output for getting branches
type GetBranchesOutput = map[string]interface{}

// getBranchesHandler handles getting branches
func (h *Handler) getBranchesHandler(ctx context.Context, req *mcp.CallToolRequest, input GetBranchesInput) (*mcp.CallToolResult, GetBranchesOutput, error) {
	branches, err := h.client.GetBranches(
		input.ProjectKey,
		input.RepoSlug,
		input.FilterText,
		input.OrderBy,
		input.Context,
		input.Base,
		input.BoostMatches,
		input.Start,
		input.Limit,
		input.Details,
	)
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

// GetDefaultBranchInput represents the input parameters for getting the default branch
type GetDefaultBranchInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetDefaultBranchOutput represents the output for getting the default branch
type GetDefaultBranchOutput = map[string]interface{}

// getDefaultBranchHandler handles getting the default branch
func (h *Handler) getDefaultBranchHandler(ctx context.Context, req *mcp.CallToolRequest, input GetDefaultBranchInput) (*mcp.CallToolResult, GetDefaultBranchOutput, error) {
	branch, err := h.client.GetDefaultBranch(input.ProjectKey, input.RepoSlug)
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

// GetBranchInfoByCommitIdInput represents the input parameters for getting branch information by commit ID
type GetBranchInfoByCommitIdInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitId   string `json:"commitId" jsonschema:"required,The commit ID to retrieve branch information for"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned results"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of results to return"`
}

// GetBranchInfoByCommitIdOutput represents the output for getting branch information by commit ID
type GetBranchInfoByCommitIdOutput = map[string]interface{}

// getBranchInfoByCommitIdHandler handles getting branch information by commit ID
func (h *Handler) getBranchInfoByCommitIdHandler(ctx context.Context, req *mcp.CallToolRequest, input GetBranchInfoByCommitIdInput) (*mcp.CallToolResult, GetBranchInfoByCommitIdOutput, error) {
	branchInfo, err := h.client.GetBranchInfoByCommitId(
		input.ProjectKey,
		input.RepoSlug,
		input.CommitId,
		input.Start,
		input.Limit,
	)
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

	mcp.AddTool[GetBranchesInput, GetBranchesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_branches",
		Description: "Get branches for a repository",
	}, handler.getBranchesHandler)

	mcp.AddTool[GetDefaultBranchInput, GetDefaultBranchOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_default_branch",
		Description: "Get the default branch of a repository",
	}, handler.getDefaultBranchHandler)

	mcp.AddTool[GetBranchInfoByCommitIdInput, GetBranchInfoByCommitIdOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_branch_info_by_commit_id",
		Description: "Get branch information by commit ID",
	}, handler.getBranchInfoByCommitIdHandler)
}
