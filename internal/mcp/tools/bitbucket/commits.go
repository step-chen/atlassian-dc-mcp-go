package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetCommitsOutput represents the output for getting commits
type GetCommitsOutput = map[string]interface{}

// GetCommitOutput represents the output for getting a specific commit
type GetCommitOutput = map[string]interface{}

// GetCommitChangesOutput represents the output for getting changes for a specific commit
type GetCommitChangesOutput = map[string]interface{}

// GetCommitCommentsOutput represents the output for getting comments on a commit
type GetCommitCommentsOutput = map[string]interface{}

// GetCommitCommentOutput represents the output for getting a specific comment on a commit
type GetCommitCommentOutput = map[string]interface{}

// GetCommitDiffStatsSummaryOutput represents the output for getting diff statistics summary for a commit
type GetCommitDiffStatsSummaryOutput = map[string]interface{}

// GetDiffBetweenCommitsOutput represents the output for getting diff between commits
type GetDiffBetweenCommitsOutput = map[string]interface{}

// GetJiraIssueCommitsOutput represents the output for getting commits related to a Jira issue
type GetJiraIssueCommitsOutput = map[string]interface{}

// GetDiffBetweenRevisionsOutput represents the output for getting the diff between revisions
type GetDiffBetweenRevisionsOutput struct {
	Diff string `json:"diff"`
}

// GetDiffBetweenRevisionsForPathOutput represents the output for getting the diff between revisions for a specific path
type GetDiffBetweenRevisionsForPathOutput struct {
	Diff string `json:"diff"`
}

// getCommitsHandler handles getting commits
func (h *Handler) getCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitsInput) (*mcp.CallToolResult, GetCommitsOutput, error) {
	commits, err := h.client.GetCommits(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get commits")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(commits)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create commits result")
		return result, nil, err
	}

	return result, commits, nil
}

// getCommitHandler handles getting a specific commit
func (h *Handler) getCommitHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitInput) (*mcp.CallToolResult, GetCommitOutput, error) {
	commit, err := h.client.GetCommit(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get commit")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(commit)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create commit result")
		return result, nil, err
	}

	return result, commit, nil
}

// getCommitChangesHandler handles getting changes for a specific commit
func (h *Handler) getCommitChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitChangesInput) (*mcp.CallToolResult, GetCommitChangesOutput, error) {
	changes, err := h.client.GetCommitChanges(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get commit changes")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(changes)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create commit changes result")
		return result, nil, err
	}

	return result, changes, nil
}

// getCommitCommentHandler handles getting a specific comment on a commit
func (h *Handler) getCommitCommentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitCommentInput) (*mcp.CallToolResult, GetCommitCommentOutput, error) {
	comment, err := h.client.GetCommitComment(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get commit comment")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(comment)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create commit comment result")
		return result, nil, err
	}

	return result, comment, nil
}

// getCommitCommentsHandler handles getting comments on a commit
func (h *Handler) getCommitCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitCommentsInput) (*mcp.CallToolResult, GetCommitCommentsOutput, error) {
	comments, err := h.client.GetCommitComments(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get commit comments")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(comments)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create commit comments result")
		return result, nil, err
	}

	return result, comments, nil
}

// getCommitDiffStatsSummaryHandler handles getting diff statistics summary for a commit
func (h *Handler) getCommitDiffStatsSummaryHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitDiffStatsSummaryInput) (*mcp.CallToolResult, GetCommitDiffStatsSummaryOutput, error) {
	stats, err := h.client.GetCommitDiffStatsSummary(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get commit diff stats summary")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(stats)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create commit diff stats summary result")
		return result, nil, err
	}

	return result, stats, nil
}

// getDiffBetweenCommitsHandler handles getting diff between commits
func (h *Handler) getDiffBetweenCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDiffBetweenCommitsInput) (*mcp.CallToolResult, GetDiffBetweenCommitsOutput, error) {
	diff, err := h.client.GetDiffBetweenCommits(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get diff between commits")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(diff)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create diff between commits result")
		return result, nil, err
	}

	return result, GetDiffBetweenCommitsOutput{"diff": diff}, nil
}

// getDiffBetweenRevisionsHandler handles getting the diff between revisions
func (h *Handler) getDiffBetweenRevisionsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDiffBetweenRevisionsInput) (*mcp.CallToolResult, GetDiffBetweenRevisionsOutput, error) {
	diff, err := h.client.GetDiffBetweenRevisions(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get diff between revisions")
		return result, GetDiffBetweenRevisionsOutput{}, err
	}

	result, err := tools.CreateToolResult(diff)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create diff between revisions result")
		return result, GetDiffBetweenRevisionsOutput{}, err
	}

	return result, GetDiffBetweenRevisionsOutput{Diff: diff}, nil
}

// getJiraIssueCommitsHandler handles getting commits related to a Jira issue
func (h *Handler) getJiraIssueCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetJiraIssueCommitsInput) (*mcp.CallToolResult, GetJiraIssueCommitsOutput, error) {
	commits, err := h.client.GetJiraIssueCommits(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get Jira issue commits")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(commits)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create Jira issue commits result")
		return result, nil, err
	}

	return result, commits, nil
}

// getDiffBetweenRevisionsForPathHandler handles getting the diff between revisions for a specific path
func (h *Handler) getDiffBetweenRevisionsForPathHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDiffBetweenRevisionsForPathInput) (*mcp.CallToolResult, GetDiffBetweenRevisionsForPathOutput, error) {
	diff, err := h.client.GetDiffBetweenRevisionsForPath(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get diff between revisions for path")
		return result, GetDiffBetweenRevisionsForPathOutput{}, err
	}

	result, err := tools.CreateToolResult(diff)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create diff between revisions for path result")
		return result, GetDiffBetweenRevisionsForPathOutput{}, err
	}

	return result, GetDiffBetweenRevisionsForPathOutput{Diff: diff}, nil
}

// AddCommitTools registers the commit-related tools with the MCP server
func AddCommitTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[bitbucket.GetCommitsInput, GetCommitsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commits",
		Description: "Get commits for a repository",
	}, handler.getCommitsHandler)

	mcp.AddTool[bitbucket.GetCommitInput, GetCommitOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit",
		Description: "Get a specific commit",
	}, handler.getCommitHandler)

	mcp.AddTool[bitbucket.GetCommitChangesInput, GetCommitChangesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_changes",
		Description: "Get changes for a specific commit",
	}, handler.getCommitChangesHandler)

	mcp.AddTool[bitbucket.GetCommitCommentsInput, GetCommitCommentsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_comments",
		Description: "Get comments on a commit",
	}, handler.getCommitCommentsHandler)

	mcp.AddTool[bitbucket.GetCommitCommentInput, GetCommitCommentOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_comment",
		Description: "Get a specific comment on a commit",
	}, handler.getCommitCommentHandler)

	mcp.AddTool[bitbucket.GetCommitDiffStatsSummaryInput, GetCommitDiffStatsSummaryOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_diff_stats_summary",
		Description: "Get diff statistics summary for a commit",
	}, handler.getCommitDiffStatsSummaryHandler)

	mcp.AddTool[bitbucket.GetDiffBetweenCommitsInput, GetDiffBetweenCommitsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_commits",
		Description: "Get the diff between two commits",
	}, handler.getDiffBetweenCommitsHandler)

	mcp.AddTool[bitbucket.GetDiffBetweenRevisionsInput, GetDiffBetweenRevisionsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_revisions",
		Description: "Get the diff between revisions",
	}, handler.getDiffBetweenRevisionsHandler)

	mcp.AddTool[bitbucket.GetJiraIssueCommitsInput, GetJiraIssueCommitsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_jira_issue_commits",
		Description: "Get commits related to a Jira issue",
	}, handler.getJiraIssueCommitsHandler)

	mcp.AddTool[bitbucket.GetDiffBetweenRevisionsForPathInput, GetDiffBetweenRevisionsForPathOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_revisions_for_path",
		Description: "Get the diff between revisions for a specific path",
	}, handler.getDiffBetweenRevisionsForPathHandler)
}
