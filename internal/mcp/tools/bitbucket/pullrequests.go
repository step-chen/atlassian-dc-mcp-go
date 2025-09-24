package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/utils"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getPullRequestsHandler handles getting pull requests
func (h *Handler) getPullRequestsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestsInput) (*mcp.CallToolResult, MapOutput, error) {
	pullRequests, err := h.client.GetPullRequests(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull requests failed: %w", err)
	}

	return nil, pullRequests, nil
}

// getPullRequestHandler handles getting a specific pull request
func (h *Handler) getPullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestInput) (*mcp.CallToolResult, MapOutput, error) {
	pullRequest, err := h.client.GetPullRequest(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request failed: %w", err)
	}

	return nil, pullRequest, nil
}

// getPullRequestActivitiesHandler handles getting pull request activities
func (h *Handler) getPullRequestActivitiesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestActivitiesInput) (*mcp.CallToolResult, MapOutput, error) {
	activities, err := h.client.GetPullRequestActivities(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request activities failed: %w", err)
	}

	return nil, activities, nil
}

// getPullRequestCommitsHandler handles getting pull request commits
func (h *Handler) getPullRequestCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitsInput) (*mcp.CallToolResult, MapOutput, error) {
	// Using GetCommits method as a substitute for GetPullRequestCommits
	commits, err := h.client.GetCommits(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request commits failed: %w", err)
	}

	return nil, commits, nil
}

// getPullRequestCommentsHandler handles getting pull request comments
func (h *Handler) getPullRequestCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestCommentsInput) (*mcp.CallToolResult, MapOutput, error) {
	comments, err := h.client.GetPullRequestComments(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request comments failed: %w", err)
	}

	return nil, comments, nil
}

// getPullRequestDiffHandler handles getting pull request diff
func (h *Handler) getPullRequestDiffHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDiffBetweenCommitsInput) (*mcp.CallToolResult, DiffOutput, error) {
	// Using GetDiffBetweenCommits method as a substitute for GetPullRequestDiff
	diff, err := h.client.GetDiffBetweenCommits(input)
	if err != nil {
		return nil, DiffOutput{}, fmt.Errorf("get pull request diff failed: %w", err)
	}

	return nil, DiffOutput{Diff: string(diff)}, nil
}

// mergePullRequestHandler handles merging a pull request
func (h *Handler) mergePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.MergePullRequestInput) (*mcp.CallToolResult, MapOutput, error) {
	result, err := h.client.MergePullRequest(input)
	if err != nil {
		return nil, nil, fmt.Errorf("merge pull request failed: %w", err)
	}

	return nil, result, nil
}

// declinePullRequestHandler handles declining a pull request
func (h *Handler) declinePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.DeclinePullRequestInput) (*mcp.CallToolResult, MapOutput, error) {
	result, err := h.client.DeclinePullRequest(input)
	if err != nil {
		return nil, nil, fmt.Errorf("decline pull request failed: %w", err)
	}

	return nil, result, nil
}

// addPullRequestCommentHandler handles adding a comment to a pull request
func (h *Handler) addPullRequestCommentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.AddPullRequestCommentInput) (*mcp.CallToolResult, MapOutput, error) {
	comment, err := h.client.AddPullRequestComment(input)
	if err != nil {
		return nil, nil, fmt.Errorf("add pull request comment failed: %w", err)
	}

	return nil, comment, nil
}

// getPullRequestSuggestionsHandler handles getting pull request suggestions
func (h *Handler) getPullRequestSuggestionsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestSuggestionsInput) (*mcp.CallToolResult, MapOutput, error) {
	suggestions, err := h.client.GetPullRequestSuggestions(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request suggestions failed: %w", err)
	}

	return nil, suggestions, nil
}

// getPullRequestJiraIssuesHandler handles getting Jira issues linked to a pull request
func (h *Handler) getPullRequestJiraIssuesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestJiraIssuesInput) (*mcp.CallToolResult, MapOutput, error) {
	issues, err := h.client.GetPullRequestJiraIssues(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request Jira issues failed: %w", err)
	}

	return nil, issues, nil
}

// getPullRequestsForUserHandler handles getting pull requests for a user
func (h *Handler) getPullRequestsForUserHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestsForUserInput) (*mcp.CallToolResult, MapOutput, error) {
	pullRequests, err := h.client.GetPullRequestsForUser(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull requests for user failed: %w", err)
	}

	return nil, pullRequests, nil
}

// getPullRequestCommentHandler handles getting a specific comment on a pull request
func (h *Handler) getPullRequestCommentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestCommentInput) (*mcp.CallToolResult, MapOutput, error) {
	comment, err := h.client.GetPullRequestComment(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request comment failed: %w", err)
	}

	return nil, comment, nil
}

// AddPullRequestTools registers the pull request-related tools with the MCP server
func AddPullRequestTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[bitbucket.GetPullRequestsInput, MapOutput](server, "bitbucket_get_pull_requests", "Get a list of pull requests", handler.getPullRequestsHandler)
	utils.RegisterTool[bitbucket.GetPullRequestInput, MapOutput](server, "bitbucket_get_pull_request", "Get a specific pull request", handler.getPullRequestHandler)
	utils.RegisterTool[bitbucket.GetPullRequestActivitiesInput, MapOutput](server, "bitbucket_get_pull_request_activities", "Get activities for a specific pull request", handler.getPullRequestActivitiesHandler)
	utils.RegisterTool[bitbucket.GetCommitsInput, MapOutput](server, "bitbucket_get_pull_request_commits", "Get commits for a specific pull request", handler.getPullRequestCommitsHandler)
	utils.RegisterTool[bitbucket.GetPullRequestCommentsInput, MapOutput](server, "bitbucket_get_pull_request_comments", "Get comments for a specific pull request", handler.getPullRequestCommentsHandler)
	utils.RegisterTool[bitbucket.GetDiffBetweenCommitsInput, DiffOutput](server, "bitbucket_get_pull_request_diff", "Get diff of a pull request", handler.getPullRequestDiffHandler)
	utils.RegisterTool[bitbucket.GetPullRequestSuggestionsInput, MapOutput](server, "bitbucket_get_pull_request_suggestions", "Get pull request suggestions", handler.getPullRequestSuggestionsHandler)
	utils.RegisterTool[bitbucket.GetPullRequestJiraIssuesInput, MapOutput](server, "bitbucket_get_pull_request_jira_issues", "Get Jira issues linked to a pull request", handler.getPullRequestJiraIssuesHandler)
	utils.RegisterTool[bitbucket.GetPullRequestsForUserInput, MapOutput](server, "bitbucket_get_pull_requests_for_user", "Get pull requests for a specific user", handler.getPullRequestsForUserHandler)
	utils.RegisterTool[bitbucket.GetPullRequestCommentInput, MapOutput](server, "bitbucket_get_pull_request_comment", "Get a specific comment on a pull request", handler.getPullRequestCommentHandler)

	if permissions["bitbucket_merge_pull_request"] {
		utils.RegisterTool[bitbucket.MergePullRequestInput, MapOutput](server, "bitbucket_merge_pull_request", "Merge a pull request", handler.mergePullRequestHandler)
	}

	if permissions["bitbucket_decline_pull_request"] {
		utils.RegisterTool[bitbucket.DeclinePullRequestInput, MapOutput](server, "bitbucket_decline_pull_request", "Decline a pull request", handler.declinePullRequestHandler)
	}

	if permissions["bitbucket_add_pull_request_comment"] {
		utils.RegisterTool[bitbucket.AddPullRequestCommentInput, MapOutput](server, "bitbucket_add_pull_request_comment", "Add a comment to a pull request", handler.addPullRequestCommentHandler)
	}
}
