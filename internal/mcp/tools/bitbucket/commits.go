package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getCommitsHandler handles getting commits
func (h *Handler) getCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get commits", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		until, _ := tools.GetStringArg(args, "until")
		since, _ := tools.GetStringArg(args, "since")
		path, _ := tools.GetStringArg(args, "path")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		merges, _ := tools.GetStringArg(args, "merges")

		followRenames := tools.GetBoolArg(args, "followRenames", false)

		ignoreMissing := tools.GetBoolArg(args, "ignoreMissing", false)

		withCounts := tools.GetBoolArg(args, "withCounts", false)

		return h.client.GetCommits(projectKey, repoSlug, until, since, path, start, limit, merges, followRenames, ignoreMissing, withCounts)
	})
}

// getCommitHandler handles getting a specific commit
func (h *Handler) getCommitHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get commit", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		commitId, ok := tools.GetStringArg(args, "commitId")
		if !ok {
			return nil, fmt.Errorf("missing or invalid commitId parameter")
		}

		path, _ := tools.GetStringArg(args, "path")

		return h.client.GetCommit(projectKey, repoSlug, commitId, path)
	})
}

// getCommitChangesHandler handles getting changes for a specific commit
func (h *Handler) getCommitChangesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get commit changes", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		commitId, ok := tools.GetStringArg(args, "commitId")
		if !ok {
			return nil, fmt.Errorf("missing or invalid commitId parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		withComments, _ := tools.GetStringArg(args, "withComments")
		since, _ := tools.GetStringArg(args, "since")

		return h.client.GetCommitChanges(projectKey, repoSlug, commitId, start, limit, withComments, since)
	})
}

// getCommitCommentHandler handles getting a specific comment on a commit
func (h *Handler) getCommitCommentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get commit comment", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		commitId, ok := tools.GetStringArg(args, "commitId")
		if !ok {
			return nil, fmt.Errorf("missing or invalid commitId parameter")
		}

		commentId := tools.GetIntArg(args, "commentId", 0)
		if commentId <= 0 {
			return nil, fmt.Errorf("missing or invalid commentId parameter")
		}

		return h.client.GetCommitComment(projectKey, repoSlug, commitId, commentId)
	})
}

// getCommitCommentsHandler handles getting comments on a commit
func (h *Handler) getCommitCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get commit comments", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		commitId, ok := tools.GetStringArg(args, "commitId")
		if !ok {
			return nil, fmt.Errorf("missing or invalid commitId parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		path, _ := tools.GetStringArg(args, "path")

		since, _ := tools.GetStringArg(args, "since")

		return h.client.GetCommitComments(projectKey, repoSlug, commitId, start, limit, path, since)
	})
}

// getCommitDiffStatsSummaryHandler handles getting diff statistics summary for a commit
func (h *Handler) getCommitDiffStatsSummaryHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get commit diff stats summary", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		commitId, ok := tools.GetStringArg(args, "commitId")
		if !ok {
			return nil, fmt.Errorf("missing or invalid commitId parameter")
		}

		path, ok := tools.GetStringArg(args, "path")
		if !ok {
			return nil, fmt.Errorf("missing or invalid path parameter")
		}

		srcPath, _ := tools.GetStringArg(args, "srcPath")
		autoSrcPath, _ := tools.GetStringArg(args, "autoSrcPath")
		whitespace, _ := tools.GetStringArg(args, "whitespace")
		since, _ := tools.GetStringArg(args, "since")

		return h.client.GetCommitDiffStatsSummary(projectKey, repoSlug, commitId, path, srcPath, autoSrcPath, whitespace, since)
	})
}

// getDiffBetweenCommitsHandler handles getting diff between commits
func (h *Handler) getDiffBetweenCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get diff between commits", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		path, _ := tools.GetStringArg(args, "path")

		from, ok := tools.GetStringArg(args, "from")
		if !ok {
			return nil, fmt.Errorf("missing or invalid from parameter")
		}

		to, ok := tools.GetStringArg(args, "to")
		if !ok {
			return nil, fmt.Errorf("missing or invalid to parameter")
		}

		contextLines := tools.GetIntArg(args, "contextLines", 0)

		srcPath, _ := tools.GetStringArg(args, "srcPath")
		whitespace, _ := tools.GetStringArg(args, "whitespace")
		fromRepo, _ := tools.GetStringArg(args, "fromRepo")

		return h.client.GetDiffBetweenCommits(projectKey, repoSlug, path, from, to, contextLines, srcPath, whitespace, fromRepo)
	})
}

// getDiffBetweenRevisionsHandler handles getting diff between revisions
func (h *Handler) getDiffBetweenRevisionsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get diff between revisions", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		commitId, ok := tools.GetStringArg(args, "commitId")
		if !ok {
			return nil, fmt.Errorf("missing or invalid commitId parameter")
		}

		path, ok := tools.GetStringArg(args, "path")
		if !ok {
			return nil, fmt.Errorf("missing or invalid path parameter")
		}

		contextLines := tools.GetIntArg(args, "contextLines", 0)

		since, _ := tools.GetStringArg(args, "since")
		srcPath, _ := tools.GetStringArg(args, "srcPath")
		whitespace, _ := tools.GetStringArg(args, "whitespace")
		filter, _ := tools.GetStringArg(args, "filter")
		autoSrcPath, _ := tools.GetStringArg(args, "autoSrcPath")

		withComments, _ := tools.GetStringArg(args, "withComments")

		return h.client.GetDiffBetweenRevisions(projectKey, repoSlug, commitId, path, contextLines, since, srcPath, whitespace, filter, autoSrcPath, withComments)
	})
}

// getJiraIssueCommitsHandler handles getting commits for a Jira issue
func (h *Handler) getJiraIssueCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get Jira issue commits", func() (interface{}, error) {
		issueKey, ok := tools.GetStringArg(args, "issueKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid issueKey parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)
		maxChanges := tools.GetIntArg(args, "maxChanges", 0)

		return h.client.GetJiraIssueCommits(issueKey, start, limit, maxChanges)
	})
}

// getDiffBetweenRevisionsForPathHandler handles getting diff between revisions for a path
func (h *Handler) getDiffBetweenRevisionsForPathHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get diff between revisions for path", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		path, ok := tools.GetStringArg(args, "path")
		if !ok {
			return nil, fmt.Errorf("missing or invalid path parameter")
		}

		since, _ := tools.GetStringArg(args, "since")
		until, _ := tools.GetStringArg(args, "until")

		contextLines := tools.GetIntArg(args, "contextLines", 0)

		srcPath, _ := tools.GetStringArg(args, "srcPath")
		whitespace, _ := tools.GetStringArg(args, "whitespace")

		return h.client.GetDiffBetweenRevisionsForPath(projectKey, repoSlug, path, since, until, contextLines, srcPath, whitespace)
	})
}

// AddCommitTools registers the commit-related tools with the MCP server
func AddCommitTools(server *mcp.Server, client *bitbucket.BitbucketClient, hasWritePermission bool) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_commits",
		Description: "Get commits for a repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"until": {
					Type:        "string",
					Description: "Commit ID or ref to retrieve commits until",
				},
				"since": {
					Type:        "string",
					Description: "Commit ID or ref to retrieve commits since",
				},
				"path": {
					Type:        "string",
					Description: "Path to filter commits by",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned commits",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of commits to return",
				},
				"merges": {
					Type:        "string",
					Description: "Filter merge commits",
				},
				"followRenames": {
					Type:        "boolean",
					Description: "Follow file renames",
				},
				"ignoreMissing": {
					Type:        "boolean",
					Description: "Ignore missing commits",
				},
				"withCounts": {
					Type:        "boolean",
					Description: "Include commit counts",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getCommitsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_commit",
		Description: "Get a specific commit",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"commitId": {
					Type:        "string",
					Description: "The commit ID to retrieve",
				},
				"path": {
					Type:        "string",
					Description: "Filter commit details by file path",
				},
			},
			Required: []string{"projectKey", "repoSlug", "commitId"},
		},
	}, handler.getCommitHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_commit_changes",
		Description: "Get changes for a specific commit",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"commitId": {
					Type:        "string",
					Description: "The commit ID to retrieve changes for",
				},
				"since": {
					Type:        "string",
					Description: "Filter changes since a specific time",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned changes",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of changes to return",
				},
				"withComments": {
					Type:        "boolean",
					Description: "Include comments in response",
				},
			},
			Required: []string{"projectKey", "repoSlug", "commitId"},
		},
	}, handler.getCommitChangesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_commit_comments",
		Description: "Get comments on a commit",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"commitId": {
					Type:        "string",
					Description: "The commit ID",
				},
				"path": {
					Type:        "string",
					Description: "Filter comments by file path",
				},
				"since": {
					Type:        "string",
					Description: "Filter comments since a specific time",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned comments",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of comments to return",
				},
			},
			Required: []string{"projectKey", "repoSlug", "commitId"},
		},
	}, handler.getCommitCommentsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_commit_comment",
		Description: "Get a specific comment on a commit",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"commitId": {
					Type:        "string",
					Description: "The commit ID",
				},
				"commentId": {
					Type:        "integer",
					Description: "The comment ID to retrieve",
				},
			},
			Required: []string{"projectKey", "repoSlug", "commitId", "commentId"},
		},
	}, handler.getCommitCommentHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_commit_diff_stats_summary",
		Description: "Get diff statistics summary for a commit",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"commitId": {
					Type:        "string",
					Description: "The commit ID",
				},
				"path": {
					Type:        "string",
					Description: "The file path",
				},
				"srcPath": {
					Type:        "string",
					Description: "Source path for comparison",
				},
				"autoSrcPath": {
					Type:        "string",
					Description: "Automatically determine source path",
				},
				"whitespace": {
					Type:        "string",
					Description: "Whitespace handling option",
				},
				"since": {
					Type:        "string",
					Description: "Filter changes since a specific time",
				},
			},
			Required: []string{"projectKey", "repoSlug", "commitId", "path"},
		},
	}, handler.getCommitDiffStatsSummaryHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_commits",
		Description: "Get the diff between two commits",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"path": {
					Type:        "string",
					Description: "The file path",
				},
				"from": {
					Type:        "string",
					Description: "The source commit ID or ref",
				},
				"to": {
					Type:        "string",
					Description: "The target commit ID or ref",
				},
				"contextLines": {
					Type:        "integer",
					Description: "Number of context lines to include",
				},
				"srcPath": {
					Type:        "string",
					Description: "Source path for comparison",
				},
				"whitespace": {
					Type:        "string",
					Description: "Whitespace handling option",
				},
				"fromRepo": {
					Type:        "string",
					Description: "The source repository",
				},
			},
			Required: []string{"projectKey", "repoSlug", "from", "to"},
		},
	}, handler.getDiffBetweenCommitsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_revisions",
		Description: "Get the diff between revisions",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"commitId": {
					Type:        "string",
					Description: "The commit ID",
				},
				"path": {
					Type:        "string",
					Description: "The file path",
				},
				"contextLines": {
					Type:        "integer",
					Description: "Number of context lines to include",
				},
				"since": {
					Type:        "string",
					Description: "Filter changes since a specific time",
				},
				"srcPath": {
					Type:        "string",
					Description: "Source path for comparison",
				},
				"whitespace": {
					Type:        "string",
					Description: "Whitespace handling option",
				},
				"filter": {
					Type:        "string",
					Description: "Filter option",
				},
				"autoSrcPath": {
					Type:        "string",
					Description: "Automatically determine source path",
				},
				"withComments": {
					Type:        "string",
					Description: "Include comments in response",
				},
			},
			Required: []string{"projectKey", "repoSlug", "commitId", "path"},
		},
	}, handler.getDiffBetweenRevisionsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_jira_issue_commits",
		Description: "Get commits related to a Jira issue",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"issueKey": {
					Type:        "string",
					Description: "The Jira issue key",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned commits",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of commits to return",
				},
				"maxChanges": {
					Type:        "integer",
					Description: "Maximum number of changes to include",
				},
			},
			Required: []string{"issueKey"},
		},
	}, handler.getJiraIssueCommitsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_diff_between_revisions_for_path",
		Description: "Get the diff between revisions for a specific path",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"path": {
					Type:        "string",
					Description: "The file path",
				},
				"since": {
					Type:        "string",
					Description: "Filter changes since a specific time",
				},
				"until": {
					Type:        "string",
					Description: "Filter changes until a specific time",
				},
				"contextLines": {
					Type:        "integer",
					Description: "Number of context lines to include",
				},
				"srcPath": {
					Type:        "string",
					Description: "Source path for comparison",
				},
				"whitespace": {
					Type:        "string",
					Description: "Whitespace handling option",
				},
			},
			Required: []string{"projectKey", "repoSlug", "path"},
		},
	}, handler.getDiffBetweenRevisionsForPathHandler)
}
