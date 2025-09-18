package bitbucket

import (
	"net/http"
	"net/url"
	"strconv"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetCommits retrieves commits for a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch commits
// for a specific repository with various filtering options.
//
// Parameters:
//   - input: GetCommitsInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The commits data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommits(input GetCommitsInput) (map[string]interface{}, error) {
	queryParams := url.Values{}
	utils.SetQueryParam(queryParams, "until", input.Until, "")
	utils.SetQueryParam(queryParams, "since", input.Since, "")
	utils.SetQueryParam(queryParams, "path", input.Path, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "merges", input.Merges, "")
	utils.SetQueryParam(queryParams, "followRenames", input.FollowRenames, false)
	utils.SetQueryParam(queryParams, "ignoreMissing", input.IgnoreMissing, false)
	utils.SetQueryParam(queryParams, "withCounts", input.WithCounts, false)

	var commits map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits"},
		queryParams,
		nil,
		&commits,
	); err != nil {
		return nil, err
	}

	return commits, nil
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
//   - map[string]interface{}: The commit data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommit(input GetCommitInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "path", input.Path, "")

	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID},
		queryParams,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return result, nil
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
//   - map[string]interface{}: The changes data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitChanges(input GetCommitChangesInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "withComments", input.WithComments, "")
	utils.SetQueryParam(queryParams, "since", input.Since, "")

	var changes map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID, "changes"},
		queryParams,
		nil,
		&changes,
	); err != nil {
		return nil, err
	}

	return changes, nil
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
//   - map[string]interface{}: The diff statistics summary retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitDiffStatsSummary(input GetCommitDiffStatsSummaryInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "srcPath", input.SrcPath, "")
	utils.SetQueryParam(queryParams, "autoSrcPath", input.AutoSrcPath, "")
	utils.SetQueryParam(queryParams, "whitespace", input.Whitespace, "")
	utils.SetQueryParam(queryParams, "since", input.Since, "")

	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID, "diff-stats-summary", input.Path},
		queryParams,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return result, nil
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
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "contextLines", input.ContextLines, 0)
	utils.SetQueryParam(queryParams, "from", input.From, "")
	utils.SetQueryParam(queryParams, "to", input.To, "")
	utils.SetQueryParam(queryParams, "srcPath", input.SrcPath, "")
	utils.SetQueryParam(queryParams, "whitespace", input.Whitespace, "")
	utils.SetQueryParam(queryParams, "fromRepo", input.FromRepo, "")

	var pathSegments []string
	if input.Path != "" {
		pathSegments = []string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "compare", "diff" + input.Path}
	} else {
		pathSegments = []string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "compare", "diff"}
	}

	var diff string
	if err := c.executeRequest(
		http.MethodGet,
		pathSegments,
		queryParams,
		nil,
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
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "contextLines", input.ContextLines, 0)
	utils.SetQueryParam(queryParams, "since", input.Since, "")
	utils.SetQueryParam(queryParams, "srcPath", input.SrcPath, "")
	utils.SetQueryParam(queryParams, "whitespace", input.Whitespace, "")
	utils.SetQueryParam(queryParams, "filter", input.Filter, "")
	utils.SetQueryParam(queryParams, "autoSrcPath", input.AutoSrcPath, "")
	utils.SetQueryParam(queryParams, "withComments", input.WithComments, "")

	var diff string
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID, "diff", input.Path},
		queryParams,
		nil,
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
//   - map[string]interface{}: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitComment(input GetCommitCommentInput) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID, "comments", strconv.Itoa(input.CommentID)},
		nil,
		nil,
		&result,
	); err != nil {
		return nil, err
	}

	return result, nil
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
//   - map[string]interface{}: The comments data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitComments(input GetCommitCommentsInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "path", input.Path, "")
	utils.SetQueryParam(queryParams, "since", input.Since, "")

	var comments map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "commits", input.CommitID, "comments"},
		queryParams,
		nil,
		&comments,
	); err != nil {
		return nil, err
	}

	return comments, nil
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
//   - map[string]interface{}: The commits data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetJiraIssueCommits(input GetJiraIssueCommitsInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "maxChanges", input.MaxChanges, 0)

	var commits map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "jira", "latest", "issues", input.IssueKey, "commits"},
		queryParams,
		nil,
		&commits,
	); err != nil {
		return nil, err
	}

	return commits, nil
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
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "contextLines", input.ContextLines, 0)
	utils.SetQueryParam(queryParams, "since", input.Since, "")
	utils.SetQueryParam(queryParams, "until", input.Until, "")
	utils.SetQueryParam(queryParams, "srcPath", input.SrcPath, "")
	utils.SetQueryParam(queryParams, "whitespace", input.Whitespace, "")

	var diff string
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "diff", input.Path},
		queryParams,
		nil,
		&diff,
	); err != nil {
		return "", err
	}

	return diff, nil
}
