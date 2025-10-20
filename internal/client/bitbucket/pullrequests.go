package bitbucket

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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
func (c *BitbucketClient) GetPullRequestChanges(input GetPullRequestChangesInput) (types.MapOutput, error) {
	limit := input.Limit
	if limit <= 0 {
		limit = 25
	}
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

// AddPullRequestCommentV2 adds a comment to a pull request with enhanced functionality
//
// This function makes an HTTP POST request to the Bitbucket API to add a comment
// to a pull request. It supports various comment types including general comments,
// replies to existing comments, inline comments, and code suggestions.
//
// Parameters:
//   - input: AddPullRequestCommentV2Input containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) AddPullRequestCommentV2(input AddPullRequestCommentV2Input) (types.MapOutput, error) {
	// Validate input parameters
	if input.CommentText == "" && input.Suggestion == nil {
		return nil, fmt.Errorf("either commentText or suggestion must be provided")
	}

	// Initialize variables for line number and line type
	var lineNumber *int
	var lineType string

	// Set default line type
	if input.LineType != nil {
		lineType = *input.LineType
	} else {
		lineType = "CONTEXT"
	}

	// If code snippet is provided, resolve line number from code
	if input.CodeSnippet != nil {
		// Convert SearchContext from string to struct if provided
		var searchContext *SearchContext
		if input.SearchContext != nil {
			if err := json.Unmarshal([]byte(*input.SearchContext), &searchContext); err != nil {
				return nil, fmt.Errorf("failed to parse search context: %w", err)
			}
		}

		// Create input for resolving line number
		resolveInput := ResolveLineFromCodeInput{
			CommonInput: CommonInput{
				ProjectKey: input.ProjectKey,
				RepoSlug:   input.RepoSlug,
			},
			PullRequestID: input.PullRequestID,
			CodeSnippet:   *input.CodeSnippet,
			FilePath:      input.FilePath,
			LineType:      input.LineType,
			SearchContext: searchContext,
		}

		// Resolve line number from code snippet
		resolvedInfo, err := c.resolveLineFromCode(resolveInput)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve line from code snippet: %w", err)
		}

		lineNumber = &resolvedInfo.LineNumber
		if input.FilePath == nil {
			input.FilePath = &resolvedInfo.FilePath
		}
	} else {
		lineNumber = input.LineNumber
	}

	// Format comment with suggestion if provided
	finalCommentText := input.CommentText
	if input.Suggestion != nil {
		if input.FilePath == nil || lineNumber == nil {
			return nil, fmt.Errorf("suggestions require file_path and line_number to be specified")
		}

		suggestionEndLine := lineNumber
		if input.SuggestionEndLine != nil {
			suggestionEndLine = input.SuggestionEndLine
		}

		// Format code suggestion comment
		finalCommentText = c.formatSuggestionComment(input.CommentText, *input.Suggestion, *lineNumber, *suggestionEndLine)
	}

	// Create the payload
	payload := make(types.MapOutput)
	utils.SetRequestBodyParam(payload, "text", finalCommentText)

	// Handle reply to existing comment
	if input.ParentCommentID != nil {
		parent := make(types.MapOutput)
		parent["id"] = *input.ParentCommentID
		utils.SetRequestBodyParam(payload, "parent", parent)
	}

	// Handle inline comments and code suggestions
	if input.FilePath != nil {
		anchor := make(types.MapOutput)
		utils.SetRequestBodyParam(anchor, "path", *input.FilePath)

		// Handle line-based anchor
		if lineNumber != nil {
			utils.SetRequestBodyParam(anchor, "line", *lineNumber)
			utils.SetRequestBodyParam(anchor, "lineType", lineType)

			// Set file type based on line type
			fileType := "TO"
			if lineType == "REMOVED" {
				fileType = "FROM"
			}
			utils.SetRequestBodyParam(anchor, "fileType", fileType)
		}

		// Handle snippet-based anchor
		if input.CodeSnippet != nil {
			utils.SetRequestBodyParam(anchor, "snippet", *input.CodeSnippet)
			matchStrategy := "strict"
			if input.MatchStrategy != nil {
				matchStrategy = *input.MatchStrategy
			}
			utils.SetRequestBodyParam(anchor, "matchStrategy", matchStrategy)
		}

		utils.SetRequestBodyParam(anchor, "diffType", "EFFECTIVE")
		utils.SetRequestBodyParam(payload, "anchor", anchor)
	}

	// Handle code suggestions
	if input.Suggestion != nil {
		suggestion := make(types.MapOutput)
		utils.SetRequestBodyParam(suggestion, "content", *input.Suggestion)
		if input.SuggestionEndLine != nil {
			utils.SetRequestBodyParam(suggestion, "endLine", *input.SuggestionEndLine)
		}
		utils.SetRequestBodyParam(payload, "suggestion", suggestion)
	}

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

// resolveLineFromCode resolves a line number from a code snippet in a pull request diff
func (c *BitbucketClient) resolveLineFromCode(input ResolveLineFromCodeInput) (ResolvedLineInfo, error) {
	// Get the pull request diff
	diffInput := GetPullRequestDiffStreamInput{
		CommonInput: CommonInput{
			ProjectKey: input.CommonInput.ProjectKey,
			RepoSlug:   input.CommonInput.RepoSlug,
		},
		PullRequestID: input.PullRequestID,
	}

	diffStream, err := c.GetPullRequestDiffStreamRaw(diffInput)
	if err != nil {
		return ResolvedLineInfo{}, fmt.Errorf("failed to get pull request diff: %w", err)
	}
	defer diffStream.Close()

	// Read the diff content
	diffContent, err := io.ReadAll(diffStream)
	if err != nil {
		return ResolvedLineInfo{}, fmt.Errorf("failed to read diff content: %w", err)
	}

	// Split diff into lines
	diffLines := strings.Split(string(diffContent), "\n")

	// Find all matches of the code snippet
	var matches []ResolvedLineInfo
	for i, line := range diffLines {
		// Check if this line matches our code snippet
		if strings.Contains(line, input.CodeSnippet) {
			// Determine line type based on diff line prefix
			var lineType string
			if strings.HasPrefix(line, "+") {
				lineType = "ADDED"
			} else if strings.HasPrefix(line, "-") {
				lineType = "REMOVED"
			} else {
				lineType = "CONTEXT"
			}

			// Only consider matches of the correct line type
			if input.LineType != nil && lineType != *input.LineType {
				continue
			}

			// Calculate line number in the destination file
			lineNumber := c.calculateLineNumber(diffLines, i, lineType)

			// If we have a file path filter, check if this match is in the correct file
			if input.FilePath != nil {
				// Find the file header for this line
				filePath := c.findFilePathForLine(diffLines, i)
				if filePath != *input.FilePath {
					continue
				}

				matches = append(matches, ResolvedLineInfo{
					LineNumber: lineNumber,
					FilePath:   filePath,
					LineType:   lineType,
				})
			} else {
				// If no file path filter, find the file path for this match
				filePath := c.findFilePathForLine(diffLines, i)
				matches = append(matches, ResolvedLineInfo{
					LineNumber: lineNumber,
					FilePath:   filePath,
					LineType:   lineType,
				})
			}
		}
	}

	// Filter matches based on search context if provided
	if input.SearchContext != nil {
		matches = c.filterMatchesWithContext(matches, diffLines, input)
	}

	// Handle match strategy
	if len(matches) == 0 {
		return ResolvedLineInfo{}, fmt.Errorf("no matches found for code snippet")
	}

	matchStrategy := "strict"
	if input.MatchStrategy != nil {
		matchStrategy = *input.MatchStrategy
	}

	switch matchStrategy {
	case "best":
		// For "best" strategy, return the first match (could be enhanced with better logic)
		return matches[0], nil
	case "strict":
		fallthrough
	default:
		if len(matches) > 1 {
			// Collect details about all matches for error message
			var matchDetails []string
			for _, match := range matches {
				matchDetails = append(matchDetails, fmt.Sprintf("line %d in file %s", match.LineNumber, match.FilePath))
			}
			return ResolvedLineInfo{}, fmt.Errorf("multiple matches found for code snippet: %s. Please provide more context or use 'best' match strategy", strings.Join(matchDetails, ", "))
		}
		return matches[0], nil
	}
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
		utils.AcceptText,
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
//   - []RestJiraIssue: The Jira issues data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestJiraIssues(input GetPullRequestJiraIssuesInput) ([]RestJiraIssue, error) {
	var issues []RestJiraIssue
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "jira", "1.0", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "issues"},
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

// calculateLineNumber calculates the line number in the destination file
func (c *BitbucketClient) calculateLineNumber(diffLines []string, matchIndex int, lineType string) int {
	lineNumber := 0

	// Look backwards from the match index to find the nearest hunk header
	for i := matchIndex; i >= 0; i-- {
		line := diffLines[i]
		// Check for hunk header pattern like @@ -10,7 +10,7 @@
		if strings.HasPrefix(line, "@@") && strings.Contains(line, "@@") {
			// Extract destination line number from hunk header
			parts := strings.Split(line, " ")
			if len(parts) >= 4 {
				destPart := parts[2] // +10,7 part
				destPart = strings.TrimPrefix(destPart, "+")
				lineNumStr := strings.Split(destPart, ",")[0]
				if baseLineNum, err := strconv.Atoi(lineNumStr); err == nil {
					// Count lines from the hunk header to our match
					linesSinceHunkStart := matchIndex - i

					// Adjust count based on line types in between
					adjustment := 0
					for j := i + 1; j <= matchIndex; j++ {
						hunkLine := diffLines[j]
						// Skip context and added lines when counting removed lines
						if lineType == "REMOVED" && !strings.HasPrefix(hunkLine, "-") {
							adjustment++
						}
						// Skip context and removed lines when counting added lines
						if lineType == "ADDED" && !strings.HasPrefix(hunkLine, "+") {
							adjustment++
						}
					}

					lineNumber = baseLineNum + linesSinceHunkStart - adjustment - 1
					break
				}
			}
		}
	}

	// Fallback to simple calculation if we couldn't determine from hunk header
	if lineNumber == 0 {
		lineNumber = matchIndex + 1
	}

	return lineNumber
}

// filterMatchesWithContext filters matches based on search context
func (c *BitbucketClient) filterMatchesWithContext(matches []ResolvedLineInfo, diffLines []string, input ResolveLineFromCodeInput) []ResolvedLineInfo {
	if input.SearchContext == nil {
		return matches
	}

	// Filter matches based on context
	var filteredMatches []ResolvedLineInfo
	for _, match := range matches {
		// Find the position of this match in the diff
		matchPosition := c.findMatchPosition(diffLines, match.FilePath, input.CodeSnippet)
		if matchPosition == -1 {
			continue
		}

		// Check context before the match
		contextBeforeMatch := true
		if input.SearchContext.Before != nil {
			contextBeforeMatch = c.checkContextBefore(diffLines, matchPosition, input.SearchContext.Before)
		}

		// Check context after the match
		contextAfterMatch := true
		if input.SearchContext.After != nil {
			contextAfterMatch = c.checkContextAfter(diffLines, matchPosition, input.SearchContext.After)
		}

		// If both context checks pass, include this match
		if contextBeforeMatch && contextAfterMatch {
			filteredMatches = append(filteredMatches, match)
		}
	}

	return filteredMatches
}

// findMatchPosition finds the position of a match in the diff lines
func (c *BitbucketClient) findMatchPosition(diffLines []string, filePath string, codeSnippet string) int {
	// First, find the file section
	fileStart := -1
	for i, line := range diffLines {
		if strings.HasPrefix(line, "+++ b/"+filePath) || strings.HasPrefix(line, "--- a/"+filePath) {
			fileStart = i
			break
		}
	}

	if fileStart == -1 {
		return -1
	}

	// Then, find the line within that file
	for i := fileStart; i < len(diffLines); i++ {
		// Stop if we reach the next file section
		if i > fileStart && (strings.HasPrefix(diffLines[i], "+++ b/") || strings.HasPrefix(diffLines[i], "--- a/")) {
			break
		}

		// Check if this line contains our code snippet
		if strings.Contains(diffLines[i], codeSnippet) {
			return i
		}
	}

	return -1
}

// checkContextBefore checks if the context before a match position matches the expected context
func (c *BitbucketClient) checkContextBefore(diffLines []string, matchPosition int, expectedContext []string) bool {
	// Compare the lines before the match position with the expected context
	for i, contextLine := range expectedContext {
		contextPosition := matchPosition - len(expectedContext) + i
		if contextPosition < 0 || contextPosition >= len(diffLines) {
			return false
		}

		// Compare the content (skip diff prefixes +/- for context lines)
		actualLine := strings.TrimPrefix(strings.TrimPrefix(diffLines[contextPosition], "+"), "-")
		expectedLine := strings.TrimPrefix(strings.TrimPrefix(contextLine, "+"), "-")

		if actualLine != expectedLine {
			return false
		}
	}

	return true
}

// checkContextAfter checks if the context after a match position matches the expected context
func (c *BitbucketClient) checkContextAfter(diffLines []string, matchPosition int, expectedContext []string) bool {
	// Compare the lines after the match position with the expected context
	for i, contextLine := range expectedContext {
		contextPosition := matchPosition + 1 + i
		if contextPosition >= len(diffLines) {
			return false
		}

		// Compare the content (skip diff prefixes +/- for context lines)
		actualLine := strings.TrimPrefix(strings.TrimPrefix(diffLines[contextPosition], "+"), "-")
		expectedLine := strings.TrimPrefix(strings.TrimPrefix(contextLine, "+"), "-")

		if actualLine != expectedLine {
			return false
		}
	}

	return true
}

// findFilePathForLine finds the file path for a given line in the diff
func (c *BitbucketClient) findFilePathForLine(diffLines []string, lineIndex int) string {
	// Look backwards from the line index to find the nearest file header
	for i := lineIndex; i >= 0; i-- {
		line := diffLines[i]
		// Check for diff file header pattern like "+++ b/path/to/file"
		if strings.HasPrefix(line, "+++") {
			// Extract file path (remove "+++ b/")
			filePath := strings.TrimPrefix(line, "+++ b/")
			return filePath
		}
		// Also check for "--- a/path/to/file" in case of removed files
		if strings.HasPrefix(line, "--- a/") {
			// Extract file path (remove "--- a/")
			filePath := strings.TrimPrefix(line, "--- a/")
			return filePath
		}
	}

	// If no file header found, return empty string
	return ""
}

// formatSuggestionComment formats a comment with a code suggestion
func (c *BitbucketClient) formatSuggestionComment(commentText, suggestion string, startLine, endLine int) string {
	// 添加行范围信息（如果是多行建议）
	lineInfo := ""
	if endLine > startLine {
		lineInfo = fmt.Sprintf(" (lines %d-%d)", startLine, endLine)
	}

	// 使用 Bitbucket 可识别的建议格式
	suggestionBlock := fmt.Sprintf("```suggestion\n%s\n```", suggestion)
	if commentText != "" {
		return fmt.Sprintf("%s%s\n\n%s", commentText, lineInfo, suggestionBlock)
	}
	return suggestionBlock
}
