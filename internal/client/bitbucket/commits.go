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
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - until: The commit ID or ref to retrieve commits until
//   - since: The commit ID or ref to retrieve commits since
//   - path: Filter commits by file path
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//   - merges: Filter merge commits
//   - followRenames: Follow file renames
//   - ignoreMissing: Ignore missing commits
//   - withCounts: Include commit counts
//
// Returns:
//   - map[string]interface{}: The commits data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommits(projectKey, repoSlug, until, since, path string, start, limit int, merges string, followRenames, ignoreMissing, withCounts bool) (map[string]interface{}, error) {
	queryParams := url.Values{}
	utils.SetQueryParam(queryParams, "until", until, "")
	utils.SetQueryParam(queryParams, "since", since, "")
	utils.SetQueryParam(queryParams, "path", path, "")
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)
	utils.SetQueryParam(queryParams, "merges", merges, "")
	utils.SetQueryParam(queryParams, "followRenames", followRenames, false)
	utils.SetQueryParam(queryParams, "ignoreMissing", ignoreMissing, false)
	utils.SetQueryParam(queryParams, "withCounts", withCounts, false)

	var commits map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "commits"},
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
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - commitID: The ID of the commit to retrieve
//   - path: Filter commit details by file path
//
// Returns:
//   - map[string]interface{}: The commit data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommit(projectKey, repoSlug, commitID, path string) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "path", path, "")

	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "commits", commitID},
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
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - commitID: The ID of the commit to retrieve changes for
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//   - withComments: Include comments in response
//   - since: Filter changes since a specific time
//
// Returns:
//   - map[string]interface{}: The changes data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitChanges(projectKey, repoSlug, commitID string, start, limit int, withComments, since string) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)
	utils.SetQueryParam(queryParams, "withComments", withComments, "")
	utils.SetQueryParam(queryParams, "since", since, "")

	var changes map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "commits", commitID, "changes"},
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
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - commitID: The ID of the commit
//   - path: The file path
//   - srcPath: Source path for comparison
//   - autoSrcPath: Automatically determine source path
//   - whitespace: Whitespace handling option
//   - since: Filter changes since a specific time
//
// Returns:
//   - map[string]interface{}: The diff statistics summary retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitDiffStatsSummary(projectKey, repoSlug, commitID, path string, srcPath, autoSrcPath, whitespace, since string) (map[string]interface{}, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "srcPath", srcPath, "")
	utils.SetQueryParam(queryParams, "autoSrcPath", autoSrcPath, "")
	utils.SetQueryParam(queryParams, "whitespace", whitespace, "")
	utils.SetQueryParam(queryParams, "since", since, "")

	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "commits", commitID, "diff-stats-summary", path},
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
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - path: The file path
//   - from: The source commit ID or ref
//   - to: The target commit ID or ref
//   - contextLines: Number of context lines to include
//   - srcPath: Source path for comparison
//   - whitespace: Whitespace handling option
//   - fromRepo: The source repository
//
// Returns:
//   - string: The diff content as a string
//   - error: An error if the request fails
func (c *BitbucketClient) GetDiffBetweenCommits(projectKey, repoSlug, path, from, to string, contextLines int, srcPath, whitespace, fromRepo string) (string, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "contextLines", contextLines, 0)

	utils.SetQueryParam(queryParams, "from", from, "")
	utils.SetQueryParam(queryParams, "to", to, "")
	utils.SetQueryParam(queryParams, "srcPath", srcPath, "")
	utils.SetQueryParam(queryParams, "whitespace", whitespace, "")
	utils.SetQueryParam(queryParams, "fromRepo", fromRepo, "")

	var pathSegments []string
	if path != "" {
		pathSegments = []string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "compare", "diff" + path}
	} else {
		pathSegments = []string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "compare", "diff"}
	}

	return c.executeTextRequest(
		http.MethodGet,
		pathSegments,
		queryParams,
		nil,
	)
}

// GetDiffBetweenRevisions retrieves the diff between revisions.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch the diff
// between revisions for a specific file path.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - commitID: The commit ID
//   - path: The file path
//   - contextLines: Number of context lines to include
//   - since: Filter changes since a specific time
//   - srcPath: Source path for comparison
//   - whitespace: Whitespace handling option
//   - filter: Filter option
//   - autoSrcPath: Automatically determine source path
//   - withComments: Include comments in response
//
// Returns:
//   - string: The diff content as a string
//   - error: An error if the request fails
func (c *BitbucketClient) GetDiffBetweenRevisions(projectKey, repoSlug, commitID, path string, contextLines int, since, srcPath, whitespace, filter, autoSrcPath, withComments string) (string, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "contextLines", contextLines, 0)
	utils.SetQueryParam(queryParams, "since", since, "")
	utils.SetQueryParam(queryParams, "srcPath", srcPath, "")
	utils.SetQueryParam(queryParams, "whitespace", whitespace, "")
	utils.SetQueryParam(queryParams, "filter", filter, "")
	utils.SetQueryParam(queryParams, "autoSrcPath", autoSrcPath, "")
	utils.SetQueryParam(queryParams, "withComments", withComments, "")

	return c.executeTextRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "commits", commitID, "diff", path},
		queryParams,
		nil,
	)
}

// GetCommitComment retrieves a specific comment on a commit.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch a specific
// comment on a commit.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - commitID: The ID of the commit
//   - commentID: The ID of the comment to retrieve
//
// Returns:
//   - map[string]interface{}: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitComment(projectKey, repoSlug, commitID string, commentID int) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "commits", commitID, "comments", strconv.Itoa(commentID)},
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
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - commitID: The ID of the commit
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//   - path: Filter comments by file path
//   - since: Filter comments since a specific time
//
// Returns:
//   - map[string]interface{}: The comments data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetCommitComments(projectKey, repoSlug, commitID string, start, limit int, path, since string) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)
	utils.SetQueryParam(queryParams, "path", path, "")
	utils.SetQueryParam(queryParams, "since", since, "")

	var comments map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "commits", commitID, "comments"},
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
//   - issueKey: The Jira issue key
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//   - maxChanges: Maximum number of changes to include
//
// Returns:
//   - map[string]interface{}: The commits data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetJiraIssueCommits(issueKey string, start, limit, maxChanges int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)
	utils.SetQueryParam(queryParams, "maxChanges", maxChanges, 0)

	var commits map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "jira", "latest", "issues", issueKey, "commits"},
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
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - path: The file path
//   - since: Filter changes since a specific time
//   - until: Filter changes until a specific time
//   - contextLines: Number of context lines to include
//   - srcPath: Source path for comparison
//   - whitespace: Whitespace handling option
//
// Returns:
//   - string: The diff content as a string
//   - error: An error if the request fails
func (c *BitbucketClient) GetDiffBetweenRevisionsForPath(projectKey, repoSlug, path, since, until string, contextLines int, srcPath, whitespace string) (string, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "contextLines", contextLines, 0)
	utils.SetQueryParam(queryParams, "since", since, "")
	utils.SetQueryParam(queryParams, "until", until, "")
	utils.SetQueryParam(queryParams, "srcPath", srcPath, "")
	utils.SetQueryParam(queryParams, "whitespace", whitespace, "")

	return c.executeTextRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "diff", path},
		queryParams,
		nil,
	)
}
