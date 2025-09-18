package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetPullRequestsOutput represents the output for getting pull requests
type GetPullRequestsOutput = map[string]interface{}

// GetPullRequestOutput represents the output for getting a specific pull request
type GetPullRequestOutput = map[string]interface{}

// GetPullRequestActivitiesOutput represents the output for getting pull request activities
type GetPullRequestActivitiesOutput = map[string]interface{}

// GetPullRequestChangesOutput represents the output for getting pull request changes
type GetPullRequestChangesOutput = map[string]interface{}

// GetPullRequestCommitsOutput represents the output for getting pull request commits
type GetPullRequestCommitsOutput = map[string]interface{}

// GetPullRequestCommentsOutput represents the output for getting pull request comments
type GetPullRequestCommentsOutput = map[string]interface{}

// GetPullRequestDiffOutput represents the output for getting pull request diff
type GetPullRequestDiffOutput = map[string]interface{}

// GetPullRequestMergeConfigOutput represents the output for getting pull request merge config
type GetPullRequestMergeConfigOutput = map[string]interface{}

// GetPullRequestMergeStatusOutput represents the output for getting pull request merge status

// GetPullRequestSettingsOutput represents the output for getting pull request settings
type GetPullRequestSettingsOutput = map[string]interface{}

// MergePullRequestOutput represents the output for merging a pull request
type MergePullRequestOutput = map[string]interface{}

// DeclinePullRequestOutput represents the output for declining a pull request
type DeclinePullRequestOutput = map[string]interface{}

// getPullRequestsHandler handles getting pull requests
func (h *Handler) getPullRequestsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestsInput) (*mcp.CallToolResult, GetPullRequestsOutput, error) {
	pullRequests, err := h.client.GetPullRequests(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull requests")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(pullRequests)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull requests result")
		return result, nil, err
	}

	return result, pullRequests, nil
}

// getPullRequestHandler handles getting a specific pull request
func (h *Handler) getPullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestInput) (*mcp.CallToolResult, GetPullRequestOutput, error) {
	pullRequest, err := h.client.GetPullRequest(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(pullRequest)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request result")
		return result, nil, err
	}

	return result, pullRequest, nil
}

// getPullRequestActivitiesHandler handles getting pull request activities
func (h *Handler) getPullRequestActivitiesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestActivitiesInput) (*mcp.CallToolResult, GetPullRequestActivitiesOutput, error) {
	activities, err := h.client.GetPullRequestActivities(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request activities")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(activities)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request activities result")
		return result, nil, err
	}

	return result, activities, nil
}

// getPullRequestCommitsHandler handles getting pull request commits
func (h *Handler) getPullRequestCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitsInput) (*mcp.CallToolResult, GetPullRequestCommitsOutput, error) {
	// Using GetCommits method as a substitute for GetPullRequestCommits
	commits, err := h.client.GetCommits(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request commits")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(commits)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request commits result")
		return result, nil, err
	}

	return result, commits, nil
}

// getPullRequestCommentsHandler handles getting pull request comments
func (h *Handler) getPullRequestCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestCommentsInput) (*mcp.CallToolResult, GetPullRequestCommentsOutput, error) {
	comments, err := h.client.GetPullRequestComments(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request comments")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(comments)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request comments result")
		return result, nil, err
	}

	return result, comments, nil
}

// getPullRequestDiffHandler handles getting pull request diff
func (h *Handler) getPullRequestDiffHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDiffBetweenCommitsInput) (*mcp.CallToolResult, GetPullRequestDiffOutput, error) {
	// Using GetDiffBetweenCommits method as a substitute for GetPullRequestDiff
	diff, err := h.client.GetDiffBetweenCommits(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request diff")
		return result, GetPullRequestDiffOutput{}, err
	}

	// Create a map to hold the diff output
	diffOutput := GetPullRequestDiffOutput{
		"diff": diff,
	}

	result, err := tools.CreateToolResult(diff)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request diff result")
		return result, GetPullRequestDiffOutput{}, err
	}

	return result, diffOutput, nil
}

// mergePullRequestHandler handles merging a pull request
func (h *Handler) mergePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.MergePullRequestInput) (*mcp.CallToolResult, MergePullRequestOutput, error) {
	result, err := h.client.MergePullRequest(input)
	if err != nil {
		toolResult, _, err := tools.HandleToolError(err, "merge pull request")
		return toolResult, nil, err
	}

	toolResult, err := tools.CreateToolResult(result)
	if err != nil {
		toolResult, _, err := tools.HandleToolError(err, "create merge pull request result")
		return toolResult, nil, err
	}

	return toolResult, result, nil
}

// declinePullRequestHandler handles declining a pull request
func (h *Handler) declinePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.DeclinePullRequestInput) (*mcp.CallToolResult, DeclinePullRequestOutput, error) {
	result, err := h.client.DeclinePullRequest(input)
	if err != nil {
		toolResult, _, err := tools.HandleToolError(err, "decline pull request")
		return toolResult, nil, err
	}

	toolResult, err := tools.CreateToolResult(result)
	if err != nil {
		toolResult, _, err := tools.HandleToolError(err, "create decline pull request result")
		return toolResult, nil, err
	}

	return toolResult, result, nil
}

// AddPullRequestTools registers the pull request-related tools with the MCP server
func AddPullRequestTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[bitbucket.GetPullRequestsInput, GetPullRequestsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_requests",
		Description: "Get a list of pull requests",
	}, handler.getPullRequestsHandler)

	mcp.AddTool[bitbucket.GetPullRequestInput, GetPullRequestOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request",
		Description: "Get a specific pull request",
	}, handler.getPullRequestHandler)

	mcp.AddTool[bitbucket.GetPullRequestActivitiesInput, GetPullRequestActivitiesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_activities",
		Description: "Get activities for a specific pull request",
	}, handler.getPullRequestActivitiesHandler)

	mcp.AddTool[bitbucket.GetCommitsInput, GetPullRequestCommitsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_commits",
		Description: "Get commits for a specific pull request",
	}, handler.getPullRequestCommitsHandler)

	mcp.AddTool[bitbucket.GetPullRequestCommentsInput, GetPullRequestCommentsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_comments",
		Description: "Get comments for a specific pull request",
	}, handler.getPullRequestCommentsHandler)

	mcp.AddTool[bitbucket.GetDiffBetweenCommitsInput, GetPullRequestDiffOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_diff",
		Description: "Get diff of a pull request",
	}, handler.getPullRequestDiffHandler)

	if permissions["bitbucket_merge_pull_request"] {
		mcp.AddTool[bitbucket.MergePullRequestInput, MergePullRequestOutput](server, &mcp.Tool{
			Name:        "bitbucket_merge_pull_request",
			Description: "Merge a pull request",
		}, handler.mergePullRequestHandler)
	}

	if permissions["bitbucket_decline_pull_request"] {
		mcp.AddTool[bitbucket.DeclinePullRequestInput, DeclinePullRequestOutput](server, &mcp.Tool{
			Name:        "bitbucket_decline_pull_request",
			Description: "Decline a pull request",
		}, handler.declinePullRequestHandler)
	}
}
