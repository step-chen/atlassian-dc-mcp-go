package bitbucket

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"atlassian-dc-mcp-go/internal/types"
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
//   - types.MapOutput: The pull request data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequest(input GetPullRequestInput) (types.MapOutput, error) {
	var pr types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID)},
		nil,
		nil,
		&pr,
		utils.AcceptJSON,
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
//   - types.MapOutput: The activities data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestActivities(input GetPullRequestActivitiesInput) (types.MapOutput, error) {
	queryParams := make(url.Values)

	utils.SetQueryParam(queryParams, "fromType", input.FromType, "")
	utils.SetQueryParam(queryParams, "fromId", input.FromId, "")
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var activities types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "activities"},
		queryParams,
		nil,
		&activities,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return activities, nil
}

// GetPullRequestChanges retrieves changes for a specific pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch changes
// for a specific pull request.
//
// Parameters:
//   - input: GetPullRequestChangesInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The changes data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestChanges(input GetPullRequestChangesInput) (map[string]interface{}, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "changeScope", string(input.ChangeScope), "")
	utils.SetQueryParam(queryParams, "limit", strconv.Itoa(input.Limit), 0)
	utils.SetQueryParam(queryParams, "sinceId", input.SinceId, "")
	utils.SetQueryParam(queryParams, "start", strconv.Itoa(input.Start), 0)
	utils.SetQueryParam(queryParams, "untilId", input.UntilId, "")
	utils.SetQueryParam(queryParams, "withComments", strconv.FormatBool(input.WithComments), false)

	var changes types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "changes"},
		queryParams,
		nil,
		&changes,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return changes, nil
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
//   - types.MapOutput: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) AddPullRequestComment(input AddPullRequestCommentInput) (types.MapOutput, error) {
	payload := make(types.MapOutput)
	utils.SetRequestBodyParam(payload, "text", input.CommentText)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal comment: %w", err)
	}

	var comment types.MapOutput
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "comments"},
		nil,
		jsonPayload,
		&comment,
		utils.AcceptJSON,
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
//   - types.MapOutput: The merge result data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) MergePullRequest(input MergePullRequestInput) (types.MapOutput, error) {
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

	var result types.MapOutput
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "merge"},
		nil,
		jsonPayload,
		&result,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return result, nil
}

// DeclinePullRequestOptions represents the options for declining a pull request.
type DeclinePullRequestOptions struct {
	Comment *string `json:"comment,omitempty"`
}

// GetPullRequestDiff retrieves a diff for a specific file within a pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch a diff
// for a specific file within a pull request.
//
// Parameters:
//   - input: GetPullRequestDiffInput containing the parameters for the request
//
// Returns:
//   - io.ReadCloser: A reader to stream the diff content
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestDiff(input GetPullRequestDiffInput) (io.ReadCloser, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "srcPath", input.SrcPath, "")
	utils.SetQueryParam(queryParams, "contextLines", input.ContextLines, "")
	utils.SetQueryParam(queryParams, "sinceId", input.SinceId, "")
	utils.SetQueryParam(queryParams, "untilId", input.UntilId, "")
	utils.SetQueryParam(queryParams, "whitespace", input.Whitespace, "")
	utils.SetQueryParam(queryParams, "withComments", input.WithComments, "")
	utils.SetQueryParam(queryParams, "diffType", input.DiffType, "")
	utils.SetQueryParam(queryParams, "avatarScheme", input.AvatarScheme, "")
	utils.SetQueryParam(queryParams, "avatarSize", input.AvatarSize, "")

	return c.executeStreamRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "diff", input.Path},
		queryParams,
		nil,
		utils.AcceptJSON,
	)
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
//   - types.MapOutput: The decline result data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) DeclinePullRequest(input DeclinePullRequestInput) (types.MapOutput, error) {
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

	var result types.MapOutput
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "decline"},
		queryParams,
		jsonPayload,
		&result,
		utils.AcceptJSON,
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
//   - types.MapOutput: The pull requests data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequests(input GetPullRequestsInput) (types.MapOutput, error) {
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

	var prs types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests"},
		queryParams,
		nil,
		&prs,
		utils.AcceptJSON,
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
//   - types.MapOutput: The suggestions data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestSuggestions(input GetPullRequestSuggestionsInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "changesSince", input.ChangesSince, "")
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var suggestions types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "dashboard", "pull-request-suggestions"},
		queryParams,
		nil,
		&suggestions,
		utils.AcceptJSON,
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
//   - types.MapOutput: The Jira issues data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestJiraIssues(input GetPullRequestJiraIssuesInput) (types.MapOutput, error) {
	var issues types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "jira", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "issues"},
		nil,
		nil,
		&issues,
		utils.AcceptJSON,
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
//   - types.MapOutput: The pull requests data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestsForUser(input GetPullRequestsForUserInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "closedSince", input.ClosedSince, "")
	utils.SetQueryParam(queryParams, "role", input.Role, "")
	utils.SetQueryParam(queryParams, "participantStatus", input.ParticipantStatus, "")
	utils.SetQueryParam(queryParams, "state", input.State, "")
	utils.SetQueryParam(queryParams, "user", input.User, "")
	utils.SetQueryParam(queryParams, "order", input.Order, "")
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var pullRequests types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "dashboard", "pull-requests"},
		queryParams,
		nil,
		&pullRequests,
		utils.AcceptJSON,
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
//   - types.MapOutput: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestComment(input GetPullRequestCommentInput) (types.MapOutput, error) {
	var comment types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "comments", input.CommentID},
		nil,
		nil,
		&comment,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return comment, nil
}

// UpdatePullRequestParticipantStatus updates a participant's status for a pull request.
//
// This function makes an HTTP PUT request to the Bitbucket API to update the current
// user's status for a pull request. Implicitly adds the user as a participant if they
// are not already. If the current user is the author, this method will fail.
//
// Parameters:
//   - input: UpdatePullRequestStatusInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The participant data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) UpdatePullRequestParticipantStatus(input UpdatePullRequestStatusInput) (types.MapOutput, error) {
	switch input.Status {
	case "UNAPPROVED", "NEEDS_WORK", "APPROVED":
	default:
		return nil, fmt.Errorf("invalid status value: %s, valid values are: UNAPPROVED, NEEDS_WORK, APPROVED", input.Status)
	}

	// Create the payload with participant data
	payload := make(types.MapOutput)
	utils.SetRequestBodyParam(payload, "status", input.Status)
	utils.SetRequestBodyParam(payload, "lastReviewedCommit", input.LastReviewedCommit)
	if input.Version != nil {
		utils.SetRequestBodyParam(payload, "version", *input.Version)
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var participant types.MapOutput
	if err := c.executeRequest(
		http.MethodPut,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "participants", input.UserSlug},
		nil,
		jsonPayload,
		&participant,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return participant, nil
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
//   - types.MapOutput: The comments data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestComments(input GetPullRequestCommentsInput) (types.MapOutput, error) {
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

	var comments types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "comments"},
		queryParams,
		nil,
		&comments,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return comments, nil
}

// GetPullRequestDiffStreamRaw streams the raw diff for a pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to stream the raw diff
// for a specific pull request. The authenticated user must have REPO_READ permission
// for the repository that this pull request targets.
//
// Parameters:
//   - input: GetPullRequestDiffStreamInput containing the parameters for the request
//
// Returns:
//   - io.ReadCloser: A reader that can be used to stream the diff content
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestDiffStreamRaw(input GetPullRequestDiffStreamInput) (io.ReadCloser, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "contextLines", strconv.Itoa(input.ContextLines), 0)
	utils.SetQueryParam(queryParams, "whitespace", input.Whitespace, "")

	return c.executeStreamRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID) + ".diff"},
		queryParams,
		nil,
		utils.AcceptText,
	)
}

// GetPullRequestMergeStatus retrieves merge status for a specific pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch merge status
// for a specific pull request.
//
// Parameters:
//   - input: TestPullRequestCanMergeInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The merge status data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) TestPullRequestCanMerge(input TestPullRequestCanMergeInput) (types.MapOutput, error) {
	var mergeStatus types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "merge"},
		nil,
		nil,
		&mergeStatus,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return mergeStatus, nil
}
