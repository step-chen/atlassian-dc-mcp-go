package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getPullRequestsHandler handles getting pull requests
func (h *Handler) getPullRequestsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestsInput) (*mcp.CallToolResult, MapOutput, error) {
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
func (h *Handler) getPullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestInput) (*mcp.CallToolResult, MapOutput, error) {
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
func (h *Handler) getPullRequestActivitiesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestActivitiesInput) (*mcp.CallToolResult, MapOutput, error) {
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
func (h *Handler) getPullRequestCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitsInput) (*mcp.CallToolResult, MapOutput, error) {
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
func (h *Handler) getPullRequestCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestCommentsInput) (*mcp.CallToolResult, MapOutput, error) {
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
func (h *Handler) getPullRequestDiffHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDiffBetweenCommitsInput) (*mcp.CallToolResult, DiffOutput, error) {
	// Using GetDiffBetweenCommits method as a substitute for GetPullRequestDiff
	diff, err := h.client.GetDiffBetweenCommits(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request diff")
		return result, DiffOutput{}, err
	}

	result, err := tools.CreateToolResult(diff)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request diff result")
		return result, DiffOutput{}, err
	}

	return result, DiffOutput{Diff: diff}, nil
}

// mergePullRequestHandler handles merging a pull request
func (h *Handler) mergePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.MergePullRequestInput) (*mcp.CallToolResult, MapOutput, error) {
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
func (h *Handler) declinePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.DeclinePullRequestInput) (*mcp.CallToolResult, MapOutput, error) {
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

	mcp.AddTool[bitbucket.GetPullRequestsInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_requests",
		Description: "Get a list of pull requests",
	}, handler.getPullRequestsHandler)

	mcp.AddTool[bitbucket.GetPullRequestInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request",
		Description: "Get a specific pull request",
	}, handler.getPullRequestHandler)

	mcp.AddTool[bitbucket.GetPullRequestActivitiesInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_activities",
		Description: "Get activities for a specific pull request",
	}, handler.getPullRequestActivitiesHandler)

	mcp.AddTool[bitbucket.GetCommitsInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_commits",
		Description: "Get commits for a specific pull request",
	}, handler.getPullRequestCommitsHandler)

	mcp.AddTool[bitbucket.GetPullRequestCommentsInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_comments",
		Description: "Get comments for a specific pull request",
	}, handler.getPullRequestCommentsHandler)

	mcp.AddTool[bitbucket.GetDiffBetweenCommitsInput, DiffOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_diff",
		Description: "Get diff of a pull request",
	}, handler.getPullRequestDiffHandler)

	if permissions["bitbucket_merge_pull_request"] {
		mcp.AddTool[bitbucket.MergePullRequestInput, MapOutput](server, &mcp.Tool{
			Name:        "bitbucket_merge_pull_request",
			Description: "Merge a pull request",
		}, handler.mergePullRequestHandler)
	}

	if permissions["bitbucket_decline_pull_request"] {
		mcp.AddTool[bitbucket.DeclinePullRequestInput, MapOutput](server, &mcp.Tool{
			Name:        "bitbucket_decline_pull_request",
			Description: "Decline a pull request",
		}, handler.declinePullRequestHandler)
	}
}
