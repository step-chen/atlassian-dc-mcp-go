package bitbucket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetPullRequest retrieves details of a specific pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch details
// of a pull request identified by its project key, repository slug, and pull request ID.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - pullRequestID: The ID of the pull request to retrieve
//
// Returns:
//   - map[string]interface{}: The pull request data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequest(projectKey, repoSlug string, pullRequestID int) (map[string]interface{}, error) {
	var pr map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "pull-requests", strconv.Itoa(pullRequestID)},
		nil,
		nil,
		&pr,
	); err != nil {
		return nil, err
	}

	return pr, nil
}

// GetPullRequestActivities retrieves activities for a specific pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch activities
// for a specific pull request with optional filtering.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - pullRequestID: The ID of the pull request
//   - fromType: Filter activities by type
//   - fromId: Filter activities by ID
//   - start: Starting index for pagination (default: 0)
//   - limit: Maximum number of results to return (default: 25)
//
// Returns:
//   - map[string]interface{}: The activities data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestActivities(projectKey, repoSlug string, pullRequestID int, fromType, fromId string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "fromType", fromType, "")
	utils.SetQueryParam(queryParams, "fromId", fromId, "")
	utils.SetQueryParam(queryParams, "start", start, 0)
	utils.SetQueryParam(queryParams, "limit", limit, 0)

	var activities map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "pull-requests", strconv.Itoa(pullRequestID), "activities"},
		queryParams,
		nil,
		&activities,
	); err != nil {
		return nil, err
	}

	return activities, nil
}

// AddPullRequestComment adds a comment to a specific pull request.
//
// This function makes an HTTP POST request to the Bitbucket API to add a comment
// to a specific pull request.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - pullRequestID: The ID of the pull request
//   - commentText: The text of the comment to add
//
// Returns:
//   - map[string]interface{}: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) AddPullRequestComment(projectKey, repoSlug string, pullRequestID int, commentText string) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"text": commentText,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var comment map[string]interface{}
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "pull-requests", strconv.Itoa(pullRequestID), "comments"},
		nil,
		jsonPayload,
		&comment,
	); err != nil {
		return nil, err
	}

	return comment, nil
}

// MergePullRequestOptions represents the options for merging a pull request.
type MergePullRequestOptions struct {
	Version     *int    `json:"version,omitempty"`
	AutoMerge   *bool   `json:"autoMerge,omitempty"`
	AutoSubject *string `json:"autoSubject,omitempty"`
	Message     *string `json:"message,omitempty"`
	StrategyId  *string `json:"strategyId,omitempty"`
}

// MergePullRequest merges a specific pull request.
//
// This function makes an HTTP POST request to the Bitbucket API to merge
// a specific pull request with optional merge options.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - pullRequestID: The ID of the pull request to merge
//   - options: The merge options
//
// Returns:
//   - map[string]interface{}: The merge result data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) MergePullRequest(projectKey, repoSlug string, pullRequestID int, options MergePullRequestOptions) (map[string]interface{}, error) {
	jsonPayload, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "pull-requests", strconv.Itoa(pullRequestID), "merge"},
		nil,
		jsonPayload,
		&result,
	); err != nil {
		return nil, err
	}

	return result, nil
}

// DeclinePullRequestOptions represents the options for declining a pull request.
type DeclinePullRequestOptions struct {
	Comment *string `json:"comment,omitempty"`
	Version *int    `json:"version,omitempty"`
}

// DeclinePullRequest declines a specific pull request.
//
// This function makes an HTTP POST request to the Bitbucket API to decline
// a specific pull request with optional decline options.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - pullRequestID: The ID of the pull request to decline
//   - version: The version of the pull request
//   - options: The decline options
//
// Returns:
//   - map[string]interface{}: The decline result data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) DeclinePullRequest(projectKey, repoSlug, pullRequestID string, version string, options *DeclinePullRequestOptions) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	if version != "" {
		queryParams.Set("version", version)
	}

	jsonPayload, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "pull-requests", pullRequestID, "decline"},
		queryParams,
		jsonPayload,
		&result,
	); err != nil {
		return nil, err
	}

	return result, nil
}

// GetPullRequests retrieves pull requests for a specific repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch pull requests
// for a specific repository with various filtering options.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - state: Filter pull requests by state
//   - withAttributes: Include attributes in response
//   - at: Filter pull requests by ref
//   - withProperties: Include properties in response
//   - draft: Filter by draft status
//   - filterText: Filter by text
//   - order: Sort order
//   - direction: Sort direction
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The pull requests data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequests(projectKey, repoSlug, state, withAttributes, at, withProperties, draft, filterText, order, direction string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "state", state, "")
	utils.SetQueryParam(queryParams, "withAttributes", withAttributes, "")
	utils.SetQueryParam(queryParams, "at", at, "")
	utils.SetQueryParam(queryParams, "withProperties", withProperties, "")
	utils.SetQueryParam(queryParams, "draft", draft, "")
	utils.SetQueryParam(queryParams, "filterText", filterText, "")
	utils.SetQueryParam(queryParams, "order", order, "")
	utils.SetQueryParam(queryParams, "direction", direction, "")
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)

	var prs map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "pull-requests"},
		queryParams,
		nil,
		&prs,
	); err != nil {
		return nil, err
	}

	return prs, nil
}

// GetPullRequestSuggestions retrieves pull request suggestions.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch pull request
// suggestions based on changes since a specific commit.
//
// Parameters:
//   - changesSince: The commit ID or ref to compare since
//   - limit: Maximum number of results to return (default: 25)
//
// Returns:
//   - map[string]interface{}: The suggestions data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestSuggestions(changesSince string, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "changesSince", changesSince, "")
	utils.SetQueryParam(queryParams, "limit", limit, 0)

	var suggestions map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "dashboard", "pull-request-suggestions"},
		queryParams,
		nil,
		&suggestions,
	); err != nil {
		return nil, err
	}

	return suggestions, nil
}

// GetPullRequestJiraIssues retrieves Jira issues linked to a pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch Jira issues
// linked to a specific pull request.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - pullRequestID: The ID of the pull request
//
// Returns:
//   - map[string]interface{}: The Jira issues data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestJiraIssues(projectKey, repoSlug string, pullRequestID int) (map[string]interface{}, error) {
	var issues map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "jira", "latest", "projects", projectKey, "repos", repoSlug, "pull-requests", strconv.Itoa(pullRequestID), "issues"},
		nil,
		nil,
		&issues,
	); err != nil {
		return nil, err
	}

	return issues, nil
}

// GetPullRequestsForUser retrieves pull requests for a specific user.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch pull requests
// for a specific user with various filtering options.
//
// Parameters:
//   - closedSince: Filter pull requests closed since a specific time
//   - role: Filter by user role
//   - participantStatus: Filter by participant status
//   - state: Filter pull requests by state
//   - user: Filter by user
//   - order: Sort order
//   - start: Starting index for pagination (default: 0)
//   - limit: Maximum number of results to return (default: 25)
//
// Returns:
//   - map[string]interface{}: The pull requests data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestsForUser(closedSince string, role string, participantStatus string, state string, user string, order string, start int, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "closedSince", closedSince, "")
	utils.SetQueryParam(queryParams, "role", role, "")
	utils.SetQueryParam(queryParams, "participantStatus", participantStatus, "")
	utils.SetQueryParam(queryParams, "state", state, "")
	utils.SetQueryParam(queryParams, "user", user, "")
	utils.SetQueryParam(queryParams, "order", order, "")
	utils.SetQueryParam(queryParams, "start", start, 0)
	utils.SetQueryParam(queryParams, "limit", limit, 0)

	var pullRequests map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "dashboard", "pull-requests"},
		queryParams,
		nil,
		&pullRequests,
	); err != nil {
		return nil, err
	}

	return pullRequests, nil
}

// GetPullRequestComment retrieves a specific comment on a pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch a specific
// comment on a pull request.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - pullRequestID: The ID of the pull request
//   - commentID: The ID of the comment to retrieve
//
// Returns:
//   - map[string]interface{}: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestComment(projectKey, repoSlug string, pullRequestID int, commentID string) (map[string]interface{}, error) {
	var comment map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "pull-requests", strconv.Itoa(pullRequestID), "comments", commentID},
		nil,
		nil,
		&comment,
	); err != nil {
		return nil, err
	}

	return comment, nil
}

// GetPullRequestComments retrieves comments on a pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch comments
// on a pull request with various filtering options.
//
// Parameters:
//   - projectKey: The unique key of the project
//   - repoSlug: The repository slug
//   - pullRequestID: The ID of the pull request
//   - path: Filter comments by file path
//   - fromHash: Filter comments by from hash
//   - anchorState: Filter comments by anchor state
//   - toHash: Filter comments by to hash
//   - state: Filter comments by state
//   - diffType: Filter comments by diff type
//   - diffTypes: Filter comments by diff types
//   - states: Filter comments by states
//   - limit: Maximum number of results to return (default: 25)
//   - start: Starting index for pagination (default: 0)
//
// Returns:
//   - map[string]interface{}: The comments data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestComments(projectKey, repoSlug string, pullRequestID int, path, fromHash, anchorState, toHash, state, diffType, diffTypes, states string, start, limit int) (map[string]interface{}, error) {
	queryParams := make(url.Values)

	utils.SetRequiredPathQueryParam(queryParams, path)
	utils.SetQueryParam(queryParams, "fromHash", fromHash, "")
	utils.SetQueryParam(queryParams, "anchorState", anchorState, "")
	utils.SetQueryParam(queryParams, "toHash", toHash, "")
	utils.SetQueryParam(queryParams, "state", state, "")
	utils.SetQueryParam(queryParams, "diffType", diffType, "")
	utils.SetQueryParam(queryParams, "diffTypes", diffTypes, "")
	utils.SetQueryParam(queryParams, "states", states, "")
	utils.SetQueryParam(queryParams, "limit", limit, 0)
	utils.SetQueryParam(queryParams, "start", start, 0)

	var comments map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", projectKey, "repos", repoSlug, "pull-requests", strconv.Itoa(pullRequestID), "comments"},
		queryParams,
		nil,
		&comments,
	); err != nil {
		return nil, err
	}

	return comments, nil
}
