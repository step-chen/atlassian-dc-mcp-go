package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetCommitsInput represents the input parameters for getting commits
type GetCommitsInput struct {
	ProjectKey    string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug      string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Until         string `json:"until,omitempty" jsonschema:"Filter commits until a specific time"`
	Since         string `json:"since,omitempty" jsonschema:"Filter commits since a specific time"`
	Path          string `json:"path,omitempty" jsonschema:"Path to filter commits by"`
	Start         int    `json:"start,omitempty" jsonschema:"The starting index of the returned commits"`
	Limit         int    `json:"limit,omitempty" jsonschema:"The limit of the number of commits to return"`
	Merges        string `json:"merges,omitempty" jsonschema:"Filter merge commits"`
	FollowRenames bool   `json:"followRenames,omitempty" jsonschema:"Follow file renames"`
	IgnoreMissing bool   `json:"ignoreMissing,omitempty" jsonschema:"Ignore missing commits"`
	WithCounts    bool   `json:"withCounts,omitempty" jsonschema:"Include commit counts"`
}

// GetCommitsOutput represents the output for getting commits
type GetCommitsOutput = map[string]interface{}

// GetCommitInput represents the input parameters for getting a specific commit
type GetCommitInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitId   string `json:"commitId" jsonschema:"required,The commit ID"`
	Path       string `json:"path,omitempty" jsonschema:"The path to the file"`
}

// GetCommitOutput represents the output for getting a specific commit
type GetCommitOutput = map[string]interface{}

// GetCommitChangesInput represents the input parameters for getting changes for a specific commit
type GetCommitChangesInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitId     string `json:"commitId" jsonschema:"required,The commit ID"`
	Start        int    `json:"start,omitempty" jsonschema:"The starting index of the returned changes"`
	Limit        int    `json:"limit,omitempty" jsonschema:"The limit of the number of changes to return"`
	WithComments string `json:"withComments,omitempty" jsonschema:"Include comments in the response"`
	Since        string `json:"since,omitempty" jsonschema:"The commit ID or ref to retrieve changes since"`
}

// GetCommitChangesOutput represents the output for getting changes for a specific commit
type GetCommitChangesOutput = map[string]interface{}

// GetCommitCommentsInput represents the input parameters for getting comments on a commit
type GetCommitCommentsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitId   string `json:"commitId" jsonschema:"required,The commit ID"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned comments"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of comments to return"`
	Path       string `json:"path,omitempty" jsonschema:"The path to the file"`
	Since      string `json:"since,omitempty" jsonschema:"The commit ID or ref to retrieve comments since"`
}

// GetCommitCommentsOutput represents the output for getting comments on a commit
type GetCommitCommentsOutput = map[string]interface{}

// GetCommitCommentInput represents the input parameters for getting a specific comment on a commit
type GetCommitCommentInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitId   string `json:"commitId" jsonschema:"required,The commit ID"`
	CommentId  int    `json:"commentId" jsonschema:"required,The comment ID"`
}

// GetCommitCommentOutput represents the output for getting a specific comment on a commit
type GetCommitCommentOutput = map[string]interface{}

// GetCommitDiffStatsSummaryInput represents the input parameters for getting diff statistics summary for a commit
type GetCommitDiffStatsSummaryInput struct {
	ProjectKey  string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug    string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitId    string `json:"commitId" jsonschema:"required,The commit ID"`
	Path        string `json:"path,omitempty" jsonschema:"The path to the file"`
	SrcPath     string `json:"srcPath,omitempty" jsonschema:"Source path for diff"`
	AutoSrcPath string `json:"autoSrcPath,omitempty" jsonschema:"Automatically determine source path"`
	Whitespace  string `json:"whitespace,omitempty" jsonschema:"Whitespace handling options"`
	Since       string `json:"since,omitempty" jsonschema:"The commit ID or ref to retrieve changes since"`
}

// GetCommitDiffStatsSummaryOutput represents the output for getting diff statistics summary for a commit
type GetCommitDiffStatsSummaryOutput = map[string]interface{}

// GetDiffBetweenCommitsInput represents the input parameters for getting diff between commits
type GetDiffBetweenCommitsInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Path         string `json:"path,omitempty" jsonschema:"The file path"`
	From         string `json:"from,omitempty" jsonschema:"The source commit ID or ref"`
	To           string `json:"to,omitempty" jsonschema:"The target commit ID or ref"`
	ContextLines int    `json:"contextLines,omitempty" jsonschema:"Number of context lines to include"`
	SrcPath      string `json:"srcPath,omitempty" jsonschema:"Source path for comparison"`
	Whitespace   string `json:"whitespace,omitempty" jsonschema:"Whitespace handling option"`
	FromRepo     string `json:"fromRepo,omitempty" jsonschema:"The source repository"`
}

// GetDiffBetweenCommitsOutput represents the output for getting diff between commits
type GetDiffBetweenCommitsOutput = map[string]interface{}

// GetJiraIssueCommitsInput represents the input parameters for getting commits related to a Jira issue
type GetJiraIssueCommitsInput struct {
	IssueKey   string `json:"issueKey" jsonschema:"required,The Jira issue key"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned commits"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of commits to return"`
	MaxChanges int    `json:"maxChanges,omitempty" jsonschema:"Maximum number of changes to include"`
}

// GetJiraIssueCommitsOutput represents the output for getting commits related to a Jira issue
type GetJiraIssueCommitsOutput = map[string]interface{}

// getCommitsHandler handles getting commits
func (h *Handler) getCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetCommitsInput) (*mcp.CallToolResult, GetCommitsOutput, error) {
	commits, err := h.client.GetCommits(
		input.ProjectKey,
		input.RepoSlug,
		input.Until,
		input.Since,
		input.Path,
		input.Start,
		input.Limit,
		input.Merges,
		input.FollowRenames,
		input.IgnoreMissing,
		input.WithCounts,
	)
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
func (h *Handler) getCommitHandler(ctx context.Context, req *mcp.CallToolRequest, input GetCommitInput) (*mcp.CallToolResult, GetCommitOutput, error) {
	commit, err := h.client.GetCommit(
		input.ProjectKey,
		input.RepoSlug,
		input.CommitId,
		input.Path,
	)
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
func (h *Handler) getCommitChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input GetCommitChangesInput) (*mcp.CallToolResult, GetCommitChangesOutput, error) {
	changes, err := h.client.GetCommitChanges(
		input.ProjectKey,
		input.RepoSlug,
		input.CommitId,
		input.Start,
		input.Limit,
		input.WithComments,
		input.Since,
	)
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
func (h *Handler) getCommitCommentHandler(ctx context.Context, req *mcp.CallToolRequest, input GetCommitCommentInput) (*mcp.CallToolResult, GetCommitCommentOutput, error) {
	comment, err := h.client.GetCommitComment(
		input.ProjectKey,
		input.RepoSlug,
		input.CommitId,
		input.CommentId,
	)
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
func (h *Handler) getCommitCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetCommitCommentsInput) (*mcp.CallToolResult, GetCommitCommentsOutput, error) {
	comments, err := h.client.GetCommitComments(
		input.ProjectKey,
		input.RepoSlug,
		input.CommitId,
		input.Start,
		input.Limit,
		input.Path,
		input.Since,
	)
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
func (h *Handler) getCommitDiffStatsSummaryHandler(ctx context.Context, req *mcp.CallToolRequest, input GetCommitDiffStatsSummaryInput) (*mcp.CallToolResult, GetCommitDiffStatsSummaryOutput, error) {
	stats, err := h.client.GetCommitDiffStatsSummary(
		input.ProjectKey,
		input.RepoSlug,
		input.CommitId,
		input.Path,
		input.SrcPath,
		input.AutoSrcPath,
		input.Whitespace,
		input.Since,
	)
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
func (h *Handler) getDiffBetweenCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetDiffBetweenCommitsInput) (*mcp.CallToolResult, GetDiffBetweenCommitsOutput, error) {
	diff, err := h.client.GetDiffBetweenCommits(
		input.ProjectKey,
		input.RepoSlug,
		input.Path,
		input.From,
		input.To,
		input.ContextLines,
		input.SrcPath,
		input.Whitespace,
		input.FromRepo,
	)
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

// GetDiffBetweenRevisionsInput represents the input parameters for getting the diff between revisions
type GetDiffBetweenRevisionsInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitId     string `json:"commitId" jsonschema:"required,The commit ID"`
	Path         string `json:"path" jsonschema:"required,The file path"`
	ContextLines int    `json:"contextLines,omitempty" jsonschema:"Number of context lines to include"`
	Since        string `json:"since,omitempty" jsonschema:"Filter changes since a specific time"`
	SrcPath      string `json:"srcPath,omitempty" jsonschema:"Source path for comparison"`
	Whitespace   string `json:"whitespace,omitempty" jsonschema:"Whitespace handling option"`
	Filter       string `json:"filter,omitempty" jsonschema:"Filter option"`
	AutoSrcPath  string `json:"autoSrcPath,omitempty" jsonschema:"Automatically determine source path"`
	WithComments string `json:"withComments,omitempty" jsonschema:"Include comments in response"`
}

// GetDiffBetweenRevisionsOutput represents the output for getting the diff between revisions
type GetDiffBetweenRevisionsOutput struct {
	Diff string `json:"diff"`
}

// getDiffBetweenRevisionsHandler handles getting the diff between revisions
func (h *Handler) getDiffBetweenRevisionsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetDiffBetweenRevisionsInput) (*mcp.CallToolResult, GetDiffBetweenRevisionsOutput, error) {
	diff, err := h.client.GetDiffBetweenRevisions(
		input.ProjectKey,
		input.RepoSlug,
		input.CommitId,
		input.Path,
		input.ContextLines,
		input.Since,
		input.SrcPath,
		input.Whitespace,
		input.Filter,
		input.AutoSrcPath,
		input.WithComments,
	)
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
func (h *Handler) getJiraIssueCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetJiraIssueCommitsInput) (*mcp.CallToolResult, GetJiraIssueCommitsOutput, error) {
	commits, err := h.client.GetJiraIssueCommits(
		input.IssueKey,
		input.Start,
		input.Limit,
		input.MaxChanges,
	)
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

// GetDiffBetweenRevisionsForPathInput represents the input parameters for getting the diff between revisions for a specific path
type GetDiffBetweenRevisionsForPathInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Path         string `json:"path" jsonschema:"required,The file path"`
	Since        string `json:"since,omitempty" jsonschema:"Filter changes since a specific time"`
	Until        string `json:"until,omitempty" jsonschema:"Filter changes until a specific time"`
	ContextLines int    `json:"contextLines,omitempty" jsonschema:"Number of context lines to include"`
	SrcPath      string `json:"srcPath,omitempty" jsonschema:"Source path for comparison"`
	Whitespace   string `json:"whitespace,omitempty" jsonschema:"Whitespace handling option"`
}

// GetDiffBetweenRevisionsForPathOutput represents the output for getting the diff between revisions for a specific path
type GetDiffBetweenRevisionsForPathOutput struct {
	Diff string `json:"diff"`
}

// getDiffBetweenRevisionsForPathHandler handles getting the diff between revisions for a specific path
func (h *Handler) getDiffBetweenRevisionsForPathHandler(ctx context.Context, req *mcp.CallToolRequest, input GetDiffBetweenRevisionsForPathInput) (*mcp.CallToolResult, GetDiffBetweenRevisionsForPathOutput, error) {
	diff, err := h.client.GetDiffBetweenRevisionsForPath(
		input.ProjectKey,
		input.RepoSlug,
		input.Path,
		input.Since,
		input.Until,
		input.ContextLines,
		input.SrcPath,
		input.Whitespace,
	)
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

	mcp.AddTool[GetCommitsInput, GetCommitsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commits",
		Description: "Get commits for a repository",
	}, handler.getCommitsHandler)

	mcp.AddTool[GetCommitInput, GetCommitOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit",
		Description: "Get a specific commit",
	}, handler.getCommitHandler)

	mcp.AddTool[GetCommitChangesInput, GetCommitChangesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_changes",
		Description: "Get changes for a specific commit",
	}, handler.getCommitChangesHandler)

	mcp.AddTool[GetCommitCommentsInput, GetCommitCommentsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_comments",
		Description: "Get comments on a commit",
	}, handler.getCommitCommentsHandler)

	mcp.AddTool[GetCommitCommentInput, GetCommitCommentOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_comment",
		Description: "Get a specific comment on a commit",
	}, handler.getCommitCommentHandler)

	mcp.AddTool[GetCommitDiffStatsSummaryInput, GetCommitDiffStatsSummaryOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_commit_diff_stats_summary",
		Description: "Get diff statistics summary for a commit",
	}, handler.getCommitDiffStatsSummaryHandler)

	mcp.AddTool[GetDiffBetweenCommitsInput, GetDiffBetweenCommitsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_commits",
		Description: "Get the diff between two commits",
	}, handler.getDiffBetweenCommitsHandler)

	mcp.AddTool[GetDiffBetweenRevisionsInput, GetDiffBetweenRevisionsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_revisions",
		Description: "Get the diff between revisions",
	}, handler.getDiffBetweenRevisionsHandler)

	mcp.AddTool[GetJiraIssueCommitsInput, GetJiraIssueCommitsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_jira_issue_commits",
		Description: "Get commits related to a Jira issue",
	}, handler.getJiraIssueCommitsHandler)

	mcp.AddTool[GetDiffBetweenRevisionsForPathInput, GetDiffBetweenRevisionsForPathOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_revisions_for_path",
		Description: "Get the diff between revisions for a specific path",
	}, handler.getDiffBetweenRevisionsForPathHandler)
}
