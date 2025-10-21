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

	"github.com/sourcegraph/go-diff/diff"
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

// AddPullRequestComment adds a comment to a pull request with enhanced functionality
//
// This function makes an HTTP POST request to the Bitbucket API to add a comment
// to a pull request. It supports various comment types including general comments,
// replies to existing comments, inline comments, and code suggestions.
//
// Parameters:
//   - input: AddPullRequestCommentInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The comment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) AddPullRequestComment(input AddPullRequestCommentInput) (types.MapOutput, error) {
	if input.CommentText == "" && input.Suggestion == nil {
		return nil, fmt.Errorf("either commentText or suggestion must be provided")
	}

	payload, err := c.buildCommentPayload(input)
	if err != nil {
		return nil, err
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal comment: %w", err)
	}

	var comment types.MapOutput
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(int(input.PullRequestID)), "comments"},
		nil,
		jsonPayload,
		&comment,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return comment, nil
}

// buildCommentPayload builds the payload for adding a comment to a pull request.
func (c *BitbucketClient) buildCommentPayload(input AddPullRequestCommentInput) (*CommentPayload, error) {
	var lineNumber *int
	var lineType string

	if input.LineType != nil {
		lineType = *input.LineType
	} else {
		lineType = "CONTEXT"
	}

	if input.CodeSnippet != nil {
		resolvedInfo, err := c.resolveLineNumber(input)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve line from code snippet: %w", err)
		}
		lineNumber = &resolvedInfo.LineNumber
		if input.FilePath == nil {
			input.FilePath = &resolvedInfo.FilePath
		}
	} else if input.LineNumber != nil {
		ln, err := strconv.Atoi(*input.LineNumber)
		if err != nil {
			return nil, fmt.Errorf("invalid lineNumber: %w", err)
		}
		lineNumber = &ln
	}

	finalCommentText := input.CommentText
	if input.Suggestion != nil {
		if input.FilePath == nil || lineNumber == nil {
			return nil, fmt.Errorf("suggestions require file_path and line_number to be specified")
		}
		suggestionEndLine := lineNumber
		if input.SuggestionEndLine != nil {
			sel, err := strconv.Atoi(*input.SuggestionEndLine)
			if err != nil {
				return nil, fmt.Errorf("invalid suggestionEndLine: %w", err)
			}
			suggestionEndLine = &sel
		}
		finalCommentText = c.formatSuggestionComment(input.CommentText, *input.Suggestion, *lineNumber, *suggestionEndLine)
	}

	payload := &CommentPayload{
		Text: finalCommentText,
	}

	if input.ParentCommentID != nil {
		pci, err := strconv.Atoi(*input.ParentCommentID)
		if err != nil {
			return nil, fmt.Errorf("invalid parentCommentId: %w", err)
		}
		payload.Parent = &ParentID{ID: pci}
	}

	if input.FilePath != nil {
		payload.Anchor = &Anchor{
			Path:     *input.FilePath,
			DiffType: "EFFECTIVE",
		}

		if lineNumber != nil {
			payload.Anchor.Line = lineNumber
			payload.Anchor.LineType = lineType
			payload.Anchor.FileType = "TO"
			if lineType == "REMOVED" {
				payload.Anchor.FileType = "FROM"
			}
		}

		if input.CodeSnippet != nil {
			payload.Anchor.Snippet = input.CodeSnippet
			matchStrategy := "strict"
			if input.MatchStrategy != nil {
				matchStrategy = *input.MatchStrategy
			}
			payload.Anchor.MatchStrategy = matchStrategy
		}
	}

	if input.Suggestion != nil {
		payload.Suggestion = &Suggestion{
			Content: *input.Suggestion,
		}
		if input.SuggestionEndLine != nil {
			sel, err := strconv.Atoi(*input.SuggestionEndLine)
			if err != nil {
				return nil, fmt.Errorf("invalid suggestionEndLine: %w", err)
			}
			payload.Suggestion.EndLine = &sel
		}
	}

	return payload, nil
}

// resolveLineNumber resolves a line number from a code snippet in a pull request diff.
func (c *BitbucketClient) resolveLineNumber(input AddPullRequestCommentInput) (ResolvedLineInfo, error) {
	var searchContext *SearchContext
	if input.SearchContext != nil {
		if err := json.Unmarshal([]byte(*input.SearchContext), &searchContext); err != nil {
			return ResolvedLineInfo{}, fmt.Errorf("failed to parse search context: %w", err)
		}
	}

	resolveInput := ResolveLineFromCodeInput{
		CommonInput: CommonInput{
			ProjectKey: input.ProjectKey,
			RepoSlug:   input.RepoSlug,
		},
		PullRequestID: int(input.PullRequestID),
		CodeSnippet:   *input.CodeSnippet,
		FilePath:      input.FilePath,
		LineType:      input.LineType,
		SearchContext: searchContext,
	}

	return c.resolveLineFromCode(resolveInput)
}

// resolveLineFromCode resolves a line number from a code snippet in a pull request diff.
// It's a multi-step process:
// 1. Fetch and parse the diff for the pull request.
// 2. Find all potential matches for the given code snippet within the diff.
// 3. Calculate the confidence of each match and select the best one based on the chosen strategy.
func (c *BitbucketClient) resolveLineFromCode(input ResolveLineFromCodeInput) (ResolvedLineInfo, error) {
	fileDiffs, err := c.getAndParseDiff(input)
	if err != nil {
		return ResolvedLineInfo{}, err
	}

	matches := c.findMatchesInDiff(fileDiffs, input)

	return c.selectBestMatch(matches, input)
}

// getAndParseDiff fetches the pull request diff and parses it into a structured format.
// It efficiently fetches only the required file's diff if a file path is provided.
func (c *BitbucketClient) getAndParseDiff(input ResolveLineFromCodeInput) ([]*diff.FileDiff, error) {
	var diffStream io.ReadCloser
	var err error

	if input.FilePath != nil && *input.FilePath != "" {
		contextLines := 10000
		diffInput := GetPullRequestDiffInput{
			CommonInput:   input.CommonInput,
			PullRequestID: input.PullRequestID,
			Path:          *input.FilePath,
			ContextLines:  &contextLines, // Use a large number to ensure we get the whole file's diff
		}
		diffStream, err = c.GetPullRequestDiff(diffInput)
	} else {
		diffInput := GetPullRequestDiffStreamInput{
			CommonInput:   input.CommonInput,
			PullRequestID: input.PullRequestID,
		}
		diffStream, err = c.GetPullRequestDiffStreamRaw(diffInput)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get pull request diff: %w", err)
	}
	defer diffStream.Close()

	// TODO: Reading the entire diff into memory can be a bottleneck for very large PRs.
	// Consider using a streaming parser or limiting the read size if this becomes an issue.
	diffContent, err := io.ReadAll(diffStream)
	if err != nil {
		return nil, fmt.Errorf("failed to read diff content: %w", err)
	}

	fileDiffs, err := diff.ParseMultiFileDiff(diffContent)
	if err != nil {
		return nil, fmt.Errorf("failed to parse diff: %w", err)
	}
	return fileDiffs, nil
}

// findMatchesInDiff scans through parsed file diffs to find all occurrences of a code snippet.
func (c *BitbucketClient) findMatchesInDiff(fileDiffs []*diff.FileDiff, input ResolveLineFromCodeInput) []ResolvedLineInfo {
	var matches []ResolvedLineInfo
	for _, fileDiff := range fileDiffs {
		filePath := fileDiff.NewName
		if input.FilePath != nil && filePath != *input.FilePath {
			continue
		}

		for _, hunk := range fileDiff.Hunks {
			hunkLines := strings.Split(string(hunk.Body), "\n")
			currentDestLine := hunk.NewStartLine

			for i, line := range hunkLines {
				var lineType, lineContent string
				switch {
				case strings.HasPrefix(line, "+"):
					lineType = "ADDED"
					lineContent = strings.TrimPrefix(line, "+")
				case strings.HasPrefix(line, "-"):
					lineType = "REMOVED"
					lineContent = strings.TrimPrefix(line, "-")
				default:
					lineType = "CONTEXT"
					lineContent = strings.TrimPrefix(line, " ")
				}

				if strings.Contains(lineContent, input.CodeSnippet) {
					if input.LineType == nil || *input.LineType == lineType {
						match := ResolvedLineInfo{
							LineNumber:    int(currentDestLine),
							FilePath:      filePath,
							LineType:      lineType,
							hunkBody:      hunkLines,
							hunkLineIndex: i,
						}
						matches = append(matches, match)
					}
				}

				if lineType != "REMOVED" {
					currentDestLine++
				}
			}
		}
	}
	return matches
}

// selectBestMatch filters matches based on context and applies the matching strategy.
func (c *BitbucketClient) selectBestMatch(matches []ResolvedLineInfo, input ResolveLineFromCodeInput) (ResolvedLineInfo, error) {
	if len(matches) == 0 {
		return ResolvedLineInfo{}, fmt.Errorf("no matches found for code snippet")
	}

	// Calculate confidence for all matches first
	for i := range matches {
		matches[i].Confidence = c.calculateConfidence(&matches[i], input.SearchContext)
	}

	// Filter out matches with zero confidence if a search context was provided
	if input.SearchContext != nil {
		var filteredMatches []ResolvedLineInfo
		for _, match := range matches {
			if match.Confidence > 0 {
				filteredMatches = append(filteredMatches, match)
			}
		}
		matches = filteredMatches
	}

	if len(matches) == 0 {
		return ResolvedLineInfo{}, fmt.Errorf("no matches found for code snippet after applying search context")
	}

	matchStrategy := "strict"
	if input.MatchStrategy != nil {
		matchStrategy = *input.MatchStrategy
	}

	if matchStrategy == "best" {
		var bestMatch ResolvedLineInfo
		maxConfidence := -1
		for _, match := range matches {
			if match.Confidence > maxConfidence {
				maxConfidence = match.Confidence
				bestMatch = match
			}
		}
		return bestMatch, nil
	}

	// "strict" strategy is the default
	if len(matches) > 1 {
		var matchDetails []string
		for _, match := range matches {
			matchDetails = append(matchDetails, fmt.Sprintf("line %d in file %s (confidence: %d)", match.LineNumber, match.FilePath, match.Confidence))
		}
		return ResolvedLineInfo{}, fmt.Errorf("multiple matches found for code snippet: %s. Please provide more context or use 'best' match strategy", strings.Join(matchDetails, ", "))
	}
	return matches[0], nil
}

// calculateConfidence calculates a confidence score for a match based on its context.
func (c *BitbucketClient) calculateConfidence(match *ResolvedLineInfo, searchContext *SearchContext) int {
	if searchContext == nil {
		return 1 // Base confidence for any match without context
	}

	confidence := 1
	hunkLines := match.hunkBody
	matchIndex := match.hunkLineIndex

	// Check context before
	if len(searchContext.Before) > 0 {
		for i, expectedLine := range searchContext.Before {
			actualIndex := matchIndex - len(searchContext.Before) + i
			if actualIndex >= 0 && actualIndex < len(hunkLines) {
				actualLine := strings.TrimLeft(hunkLines[actualIndex], "+- ")
				if strings.Contains(actualLine, expectedLine) {
					confidence++
				}
			}
		}
	}

	// Check context after
	if len(searchContext.After) > 0 {
		for i, expectedLine := range searchContext.After {
			actualIndex := matchIndex + 1 + i
			if actualIndex < len(hunkLines) {
				actualLine := strings.TrimLeft(hunkLines[actualIndex], "+- ")
				if strings.Contains(actualLine, expectedLine) {
					confidence++
				}
			}
		}
	}

	return confidence
}

// MergePullRequestOptions represents the options for merging a pull request.
type MergePullRequestOptions struct {
	Version     *int    `json:"version,omitempty"`
	AutoMerge   *bool   `json:"autoMerge,omitempty"`
	AutoSubject *string `json:"autoSubject,omitempty"`
	Message     *string `json:"message,omitempty"`
	StrategyId  *string `json:"strategy,omitempty"`
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
	if input.ContextLines != nil {
		queryParams.Set("contextLines", strconv.Itoa(*input.ContextLines))
	}
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
	if input.ContextLines != nil {
		queryParams.Set("contextLines", strconv.Itoa(*input.ContextLines))
	}
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
