package bitbucket

import (
	"net/http"
	"net/url"
	"strconv"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// GetCommits retrieves a list of commits from a repository.
// Parameters:
//   - input: The input for retrieving commits
//
// Returns:
//   - types.MapOutput: The commits data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommits(input GetCommitsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "until", input.Until, "")
	client.SetQueryParam(queryParams, "since", input.Since, "")
	client.SetQueryParam(queryParams, "path", input.Path, "")
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "merges", input.Merges, "")
	client.SetQueryParam(queryParams, "followRenames", input.FollowRenames, false)
	client.SetQueryParam(queryParams, "ignoreMissing", input.IgnoreMissing, false)
	client.SetQueryParam(queryParams, "withCounts", input.WithCounts, false)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetPullRequestCommits retrieves commits for the specified pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch commits
// for a specific pull request with various filtering options.
// The authenticated user must have REPO_READ permission for the repository that
// this pull request targets to call this resource.
//
// Parameters:
//   - input: GetPullRequestCommitsInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The commits data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestCommits(input GetPullRequestCommitsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "avatarScheme", input.AvatarScheme, "")
	client.SetQueryParam(queryParams, "withCounts", input.WithCounts, "")
	client.SetQueryParam(queryParams, "avatarSize", input.AvatarSize, "")
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "commits"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetCommit retrieves details of a specific commit.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch details
// of a specific commit.
//
// Parameters:
//   - input: GetCommitInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The commit data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommit(input GetCommitInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "path", input.Path, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetCommitChanges retrieves changes for a specific commit.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch changes
// made in a specific commit.
//
// Parameters:
//   - input: GetCommitChangesInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The changes data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitChanges(input GetCommitChangesInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "withComments", input.WithComments, "")
	client.SetQueryParam(queryParams, "since", input.Since, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID, "changes"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetCommitDiffStatsSummary retrieves diff statistics summary for a specific commit.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch diff statistics
// summary for a specific commit and file path.
//
// Parameters:
//   - input: GetCommitDiffStatsSummaryInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The diff statistics summary retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitDiffStatsSummary(input GetCommitDiffStatsSummaryInput) (types.MapOutput, error) {
	queryParams := url.Values{}

	client.SetQueryParam(queryParams, "srcPath", input.SrcPath, "")
	client.SetQueryParam(queryParams, "autoSrcPath", input.AutoSrcPath, "")
	client.SetQueryParam(queryParams, "whitespace", input.Whitespace, "")
	client.SetQueryParam(queryParams, "since", input.Since, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID, "diff-stats-summary", input.Path},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetDiffBetweenCommits retrieves the diff between two commits.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the diff
// between two commits in a specific repository.
//
// Parameters:
//   - input: GetDiffBetweenCommitsInput containing the parameters for the request
//
// Returns:
//   - string: The diff content as a string
//   - error: An error if the request fails
func (c *BitbucketClient) GetDiffBetweenCommits(input GetDiffBetweenCommitsInput) (string, error) {
	queryParams := url.Values{}

	client.SetQueryParam(queryParams, "contextLines", input.ContextLines, 0)
	client.SetQueryParam(queryParams, "from", input.From, "")
	client.SetQueryParam(queryParams, "to", input.To, "")
	client.SetQueryParam(queryParams, "srcPath", input.SrcPath, "")
	client.SetQueryParam(queryParams, "whitespace", input.Whitespace, "")
	client.SetQueryParam(queryParams, "fromRepo", input.FromRepo, "")

	var pathSegments = []string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "compare", "diff" + input.Path}

	var diff string
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		pathSegments,
		queryParams,
		nil,
		client.AcceptJSON,
		&diff,
	); err != nil {
		return "", err
	}

	return diff, nil
}

// GetDiffBetweenRevisions retrieves the diff between revisions.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the diff
// between revisions for a specific file path.
//
// Parameters:
//   - input: GetDiffBetweenRevisionsInput containing the parameters for the request
//
// Returns:
//   - string: The diff content as a string
//   - error: An error if the request fails
func (c *BitbucketClient) GetDiffBetweenRevisions(input GetDiffBetweenRevisionsInput) (string, error) {
	queryParams := url.Values{}

	client.SetQueryParam(queryParams, "contextLines", input.ContextLines, 0)
	client.SetQueryParam(queryParams, "since", input.Since, "")
	client.SetQueryParam(queryParams, "srcPath", input.SrcPath, "")
	client.SetQueryParam(queryParams, "whitespace", input.Whitespace, "")
	client.SetQueryParam(queryParams, "filter", input.Filter, "")
	client.SetQueryParam(queryParams, "autoSrcPath", input.AutoSrcPath, "")
	client.SetQueryParam(queryParams, "withComments", input.WithComments, "")

	var diff string
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID, "diff", input.Path},
		queryParams,
		nil,
		client.AcceptJSON,
		&diff,
	); err != nil {
		return "", err
	}

	return diff, nil
}

// GetCommitComment retrieves a specific comment on a commit.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch a specific
// comment on a commit.
//
// Parameters:
//   - input: GetCommitCommentInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitComment(input GetCommitCommentInput) (types.MapOutput, error) {
	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID, "comments", strconv.Itoa(input.CommentID)},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetCommitComments retrieves comments on a specific commit.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch comments
// on a specific commit with filtering options.
//
// Parameters:
//   - input: GetCommitCommentsInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The comments data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitComments(input GetCommitCommentsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "path", input.Path, "")
	client.SetQueryParam(queryParams, "since", input.Since, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID, "comments"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetJiraIssueCommits retrieves commits related to a Jira issue.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch commits
// related to a specific Jira issue.
//
// Parameters:
//   - input: GetJiraIssueCommitsInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The commits data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetJiraIssueCommits(input GetJiraIssueCommitsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "maxChanges", input.MaxChanges, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "jira", "latest", "issues", input.IssueKey, "commits"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetDiffBetweenRevisionsForPath retrieves the diff between revisions for a specific path.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the diff
// between revisions for a specific file path.
//
// Parameters:
//   - input: GetDiffBetweenRevisionsForPathInput containing the parameters for the request
//
// Returns:
//   - string: The diff content as a string
//   - error: An error if the request fails
func (c *BitbucketClient) GetDiffBetweenRevisionsForPath(input GetDiffBetweenRevisionsForPathInput) (string, error) {
	queryParams := url.Values{}

	client.SetQueryParam(queryParams, "contextLines", input.ContextLines, 0)
	client.SetQueryParam(queryParams, "since", input.Since, "")
	client.SetQueryParam(queryParams, "until", input.Until, "")
	client.SetQueryParam(queryParams, "srcPath", input.SrcPath, "")
	client.SetQueryParam(queryParams, "whitespace", input.Whitespace, "")

	var diff string
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "diff", input.Path},
		queryParams,
		nil,
		client.AcceptJSON,
		&diff,
	); err != nil {
		return "", err
	}

	return diff, nil
}
