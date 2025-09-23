package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/utils"

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

	utils.RegisterTool[bitbucket.GetCommitsInput, MapOutput](server, "bitbucket_get_commits", "Get commits for a repository", handler.getCommitsHandler)
	utils.RegisterTool[bitbucket.GetCommitInput, MapOutput](server, "bitbucket_get_commit", "Get a specific commit", handler.getCommitHandler)
	utils.RegisterTool[bitbucket.GetCommitChangesInput, MapOutput](server, "bitbucket_get_commit_changes", "Get changes for a specific commit", handler.getCommitChangesHandler)
	utils.RegisterTool[bitbucket.GetCommitCommentsInput, MapOutput](server, "bitbucket_get_commit_comments", "Get comments on a commit", handler.getCommitCommentsHandler)
	utils.RegisterTool[bitbucket.GetCommitCommentInput, MapOutput](server, "bitbucket_get_commit_comment", "Get a specific comment on a commit", handler.getCommitCommentHandler)
	utils.RegisterTool[bitbucket.GetCommitDiffStatsSummaryInput, MapOutput](server, "bitbucket_get_commit_diff_stats_summary", "Get diff statistics summary for a commit", handler.getCommitDiffStatsSummaryHandler)
	utils.RegisterTool[bitbucket.GetDiffBetweenCommitsInput, DiffOutput](server, "bitbucket_get_diff_between_commits", "Get the diff between two commits", handler.getDiffBetweenCommitsHandler)
	utils.RegisterTool[bitbucket.GetDiffBetweenRevisionsInput, DiffOutput](server, "bitbucket_get_diff_between_revisions", "Get the diff between revisions", handler.getDiffBetweenRevisionsHandler)
	utils.RegisterTool[bitbucket.GetJiraIssueCommitsInput, MapOutput](server, "bitbucket_get_jira_issue_commits", "Get commits related to a Jira issue", handler.getJiraIssueCommitsHandler)
	utils.RegisterTool[bitbucket.GetDiffBetweenRevisionsForPathInput, DiffOutput](server, "bitbucket_get_diff_between_revisions_for_path", "Get the diff between revisions for a specific path", handler.getDiffBetweenRevisionsForPathHandler)
}
