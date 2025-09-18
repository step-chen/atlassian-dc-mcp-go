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
//   - input: GetPullRequestInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The pull request data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequest(input GetPullRequestInput) (map[string]interface{}, error) {
	var pr map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID)},
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
//   - input: GetPullRequestActivitiesInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The activities data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestActivities(input GetPullRequestActivitiesInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "fromType", input.FromType, "")
	utils.SetQueryParam(queryParams, "fromId", input.FromId, "")
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var activities map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "activities"},
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
//   - input: AddPullRequestCommentInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) AddPullRequestComment(input AddPullRequestCommentInput) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"text": input.CommentText,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var comment map[string]interface{}
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "comments"},
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
//   - input: MergePullRequestInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The merge result data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) MergePullRequest(input MergePullRequestInput) (map[string]interface{}, error) {
	// Create options struct from input fields that are merge options
	options := MergePullRequestOptions{
		AutoMerge:   input.AutoMerge,
		AutoSubject: input.AutoSubject,
		Message:     input.Message,
		StrategyId:  input.StrategyId,
	}
	
	// Handle Version field separately since it's required in the original struct but optional in options
	if input.Version != 0 {
		options.Version = &input.Version
	}
	
	jsonPayload, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "merge"},
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
}

// DeclinePullRequest declines a specific pull request.
//
// This function makes an HTTP POST request to the Bitbucket API to decline
// a specific pull request with optional decline options.
//
// Parameters:
//   - input: DeclinePullRequestInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The decline result data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) DeclinePullRequest(input DeclinePullRequestInput) (map[string]interface{}, error) {
	// Create options struct from input fields that are decline options
	options := DeclinePullRequestOptions{
		Comment: input.Comment,
	}
	
	queryParams := make(url.Values)
	if input.Version != 0 {
		queryParams.Set("version", strconv.Itoa(input.Version))
	}

	jsonPayload, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var result map[string]interface{}
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "decline"},
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
//   - input: GetPullRequestsInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The pull requests data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequests(input GetPullRequestsInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "state", input.State, "")
	utils.SetQueryParam(queryParams, "withAttributes", strconv.FormatBool(input.WithAttributes), "")
	utils.SetQueryParam(queryParams, "at", input.At, "")
	utils.SetQueryParam(queryParams, "withProperties", strconv.FormatBool(input.WithProperties), "")
	utils.SetQueryParam(queryParams, "draft", input.Draft, "")
	utils.SetQueryParam(queryParams, "filterText", input.FilterText, "")
	utils.SetQueryParam(queryParams, "order", input.Order, "")
	utils.SetQueryParam(queryParams, "direction", input.Direction, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var prs map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests"},
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
//   - input: GetPullRequestSuggestionsInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The suggestions data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestSuggestions(input GetPullRequestSuggestionsInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "changesSince", input.ChangesSince, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)

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
//   - input: GetPullRequestJiraIssuesInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The Jira issues data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestJiraIssues(input GetPullRequestJiraIssuesInput) (map[string]interface{}, error) {
	var issues map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "jira", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "issues"},
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
//   - input: GetPullRequestsForUserInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The pull requests data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestsForUser(input GetPullRequestsForUserInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "closedSince", input.ClosedSince, "")
	utils.SetQueryParam(queryParams, "role", input.Role, "")
	utils.SetQueryParam(queryParams, "participantStatus", input.ParticipantStatus, "")
	utils.SetQueryParam(queryParams, "state", input.State, "")
	utils.SetQueryParam(queryParams, "user", input.User, "")
	utils.SetQueryParam(queryParams, "order", input.Order, "")
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)

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
//   - input: GetPullRequestCommentInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestComment(input GetPullRequestCommentInput) (map[string]interface{}, error) {
	var comment map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "comments", input.CommentID},
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
//   - input: GetPullRequestCommentsInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The comments data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestComments(input GetPullRequestCommentsInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)

	utils.SetRequiredPathQueryParam(queryParams, input.Path)
	utils.SetQueryParam(queryParams, "fromHash", input.FromHash, "")
	utils.SetQueryParam(queryParams, "anchorState", input.AnchorState, "")
	utils.SetQueryParam(queryParams, "toHash", input.ToHash, "")
	utils.SetQueryParam(queryParams, "state", input.State, "")
	utils.SetQueryParam(queryParams, "diffType", input.DiffType, "")
	utils.SetQueryParam(queryParams, "diffTypes", input.DiffTypes, "")
	utils.SetQueryParam(queryParams, "states", input.States, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)

	var comments map[string]interface{}
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "comments"},
		queryParams,
		nil,
		&comments,
	); err != nil {
		return nil, err
	}

	return comments, nil
}
