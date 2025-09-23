package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCommitsHandler handles getting commits
func (h *Handler) getCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitsInput) (*mcp.CallToolResult, MapOutput, error) {
	commits, err := h.client.GetCommits(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get commits failed: %w", err)
	}

	return nil, commits, nil
}

// getCommitHandler handles getting a specific commit
func (h *Handler) getCommitHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitInput) (*mcp.CallToolResult, MapOutput, error) {
	commit, err := h.client.GetCommit(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get commit failed: %w", err)
	}

	return nil, commit, nil
}

// getCommitChangesHandler handles getting changes for a specific commit
func (h *Handler) getCommitChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitChangesInput) (*mcp.CallToolResult, MapOutput, error) {
	changes, err := h.client.GetCommitChanges(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get commit changes failed: %w", err)
	}

	return nil, changes, nil
}

// getCommitCommentHandler handles getting a specific comment on a commit
func (h *Handler) getCommitCommentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitCommentInput) (*mcp.CallToolResult, MapOutput, error) {
	comment, err := h.client.GetCommitComment(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get commit comment failed: %w", err)
	}

	return nil, comment, nil
}

// getCommitCommentsHandler handles getting comments on a commit
func (h *Handler) getCommitCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitCommentsInput) (*mcp.CallToolResult, MapOutput, error) {
	comments, err := h.client.GetCommitComments(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get commit comments failed: %w", err)
	}

	return nil, comments, nil
}

// getCommitDiffStatsSummaryHandler handles getting diff statistics summary for a commit
func (h *Handler) getCommitDiffStatsSummaryHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetCommitDiffStatsSummaryInput) (*mcp.CallToolResult, MapOutput, error) {
	stats, err := h.client.GetCommitDiffStatsSummary(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get commit diff stats summary failed: %w", err)
	}

	return nil, stats, nil
}

// getDiffBetweenCommitsHandler handles getting diff between commits
func (h *Handler) getDiffBetweenCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDiffBetweenCommitsInput) (*mcp.CallToolResult, DiffOutput, error) {
	diff, err := h.client.GetDiffBetweenCommits(input)
	if err != nil {
		return nil, DiffOutput{}, fmt.Errorf("get diff between commits failed: %w", err)
	}

	return nil, DiffOutput{Diff: diff}, nil
}

// getDiffBetweenRevisionsHandler handles getting the diff between revisions
func (h *Handler) getDiffBetweenRevisionsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDiffBetweenRevisionsInput) (*mcp.CallToolResult, DiffOutput, error) {
	diff, err := h.client.GetDiffBetweenRevisions(input)
	if err != nil {
		return nil, DiffOutput{}, fmt.Errorf("get diff between revisions failed: %w", err)
	}

	return nil, DiffOutput{Diff: diff}, nil
}

// getJiraIssueCommitsHandler handles getting commits related to a Jira issue
func (h *Handler) getJiraIssueCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetJiraIssueCommitsInput) (*mcp.CallToolResult, MapOutput, error) {
	commits, err := h.client.GetJiraIssueCommits(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get Jira issue commits failed: %w", err)
	}

	return nil, commits, nil
}

// getDiffBetweenRevisionsForPathHandler handles getting the diff between revisions for a specific path
func (h *Handler) getDiffBetweenRevisionsForPathHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetDiffBetweenRevisionsForPathInput) (*mcp.CallToolResult, DiffOutput, error) {
	diff, err := h.client.GetDiffBetweenRevisionsForPath(input)
	if err != nil {
		return nil, DiffOutput{}, fmt.Errorf("get diff between revisions for path failed: %w", err)
	}

	return nil, DiffOutput{Diff: diff}, nil
}

// AddCommitTools registers the commit-related tools with the MCP server
func AddCommitTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[bitbucket.GetCommitsInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commits",
		Description: "Get commits for a repository",
	}, handler.getCommitsHandler)

	mcp.AddTool[bitbucket.GetCommitInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit",
		Description: "Get a specific commit",
	}, handler.getCommitHandler)

	mcp.AddTool[bitbucket.GetCommitChangesInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_changes",
		Description: "Get changes for a specific commit",
	}, handler.getCommitChangesHandler)

	mcp.AddTool[bitbucket.GetCommitCommentsInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_comments",
		Description: "Get comments on a commit",
	}, handler.getCommitCommentsHandler)

	mcp.AddTool[bitbucket.GetCommitCommentInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_comment",
		Description: "Get a specific comment on a commit",
	}, handler.getCommitCommentHandler)

	mcp.AddTool[bitbucket.GetCommitDiffStatsSummaryInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_diff_stats_summary",
		Description: "Get diff statistics summary for a commit",
	}, handler.getCommitDiffStatsSummaryHandler)

	mcp.AddTool[bitbucket.GetDiffBetweenCommitsInput, DiffOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_commits",
		Description: "Get the diff between two commits",
	}, handler.getDiffBetweenCommitsHandler)

	mcp.AddTool[bitbucket.GetDiffBetweenRevisionsInput, DiffOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_revisions",
		Description: "Get the diff between revisions",
	}, handler.getDiffBetweenRevisionsHandler)

	mcp.AddTool[bitbucket.GetJiraIssueCommitsInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_jira_issue_commits",
		Description: "Get commits related to a Jira issue",
	}, handler.getJiraIssueCommitsHandler)

	mcp.AddTool[bitbucket.GetDiffBetweenRevisionsForPathInput, DiffOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_revisions_for_path",
		Description: "Get the diff between revisions for a specific path",
	}, handler.getDiffBetweenRevisionsForPathHandler)
}
