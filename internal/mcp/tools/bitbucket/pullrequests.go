package bitbucket

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"regexp"
	"strings"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getPullRequestsHandler handles getting pull requests
func (h *Handler) getPullRequestsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	pullRequests, err := h.client.GetPullRequests(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull requests failed: %w", err)
	}

	return nil, pullRequests, nil
}

// getPullRequestHandler handles getting a specific pull request
func (h *Handler) getPullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestInput) (*mcp.CallToolResult, types.MapOutput, error) {
	pullRequest, err := h.client.GetPullRequest(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request failed: %w", err)
	}

	return nil, pullRequest, nil
}

// getPullRequestActivitiesHandler handles getting pull request activities
func (h *Handler) getPullRequestActivitiesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestActivitiesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	activities, err := h.client.GetPullRequestActivities(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request activities failed: %w", err)
	}

	return nil, activities, nil
}

// getPullRequestCommentsHandler handles getting pull request comments
func (h *Handler) getPullRequestCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestCommentsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	comments, err := h.client.GetPullRequestComments(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request comments failed: %w", err)
	}

	return nil, comments, nil
}

// mergePullRequestHandler handles merging a pull request
func (h *Handler) mergePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.MergePullRequestInput) (*mcp.CallToolResult, types.MapOutput, error) {
	result, err := h.client.MergePullRequest(input)
	if err != nil {
		return nil, nil, fmt.Errorf("merge pull request failed: %w", err)
	}

	return nil, result, nil
}

// declinePullRequestHandler handles declining a pull request
func (h *Handler) declinePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.DeclinePullRequestInput) (*mcp.CallToolResult, types.MapOutput, error) {
	result, err := h.client.DeclinePullRequest(input)
	if err != nil {
		return nil, nil, fmt.Errorf("decline pull request failed: %w", err)
	}

	return nil, result, nil
}

// addPullRequestCommentHandler handles adding an enhanced comment to a pull request
func (h *Handler) addPullRequestCommentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.AddPullRequestCommentInput) (*mcp.CallToolResult, types.MapOutput, error) {
	comment, err := h.client.AddPullRequestComment(input)
	if err != nil {
		isInlineComment := input.FilePath != nil && input.LineNumber != nil
		commentType := ""
		if isInlineComment {
			commentType = "inline "
		}
		return nil, nil, fmt.Errorf("error adding %scomment to pull request: %w", commentType, err)
	}

	return nil, comment, nil
}

// getPullRequestChangesHandler handles getting pull request changes
func (h *Handler) getPullRequestChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestChangesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	changes, err := h.client.GetPullRequestChanges(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request changes failed: %w", err)
	}

	return nil, changes, nil
}

// getPullRequestSuggestionsHandler handles getting pull request suggestions
func (h *Handler) getPullRequestSuggestionsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestSuggestionsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	suggestions, err := h.client.GetPullRequestSuggestions(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request suggestions failed: %w", err)
	}

	return nil, suggestions, nil
}

// getPullRequestJiraIssuesHandler handles getting Jira issues linked to a pull request
func (h *Handler) getPullRequestJiraIssuesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestJiraIssuesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	issues, err := h.client.GetPullRequestJiraIssues(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request Jira issues failed: %w", err)
	}

	// Convert slice to map for MCP response
	result := types.MapOutput{
		"issues": issues,
	}

	return nil, result, nil
}

// getPullRequestsForUserHandler handles getting pull requests for a user
func (h *Handler) getPullRequestsForUserHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestsForUserInput) (*mcp.CallToolResult, types.MapOutput, error) {
	pullRequests, err := h.client.GetPullRequestsForUser(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull requests for user failed: %w", err)
	}

	return nil, pullRequests, nil
}

// getPullRequestCommentHandler handles getting a specific comment on a pull request
func (h *Handler) getPullRequestCommentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestCommentInput) (*mcp.CallToolResult, types.MapOutput, error) {
	comment, err := h.client.GetPullRequestComment(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request comment failed: %w", err)
	}

	return nil, comment, nil
}

// matchAny checks if a string matches any of the glob patterns.
func matchAny(patterns []string, s string) (bool, error) {
	for _, pattern := range patterns {
		matched, err := filepath.Match(pattern, s)
		if err != nil {
			return false, err
		}
		if matched {
			return true, nil
		}
	}
	return false, nil
}

// getPullRequestDiffStreamHandler handles streaming the diff for a pull request
func (h *Handler) getPullRequestDiffStreamHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestDiffStreamInput) (*mcp.CallToolResult, DiffOutput, error) {
	stream, err := h.client.GetPullRequestDiffStreamRaw(input)
	if err != nil {
		return nil, DiffOutput{}, fmt.Errorf("get pull request diff stream failed: %w", err)
	}
	defer stream.Close()

	hasIncludePatterns := len(input.IncludePatterns) > 0
	hasExcludePatterns := len(input.ExcludePatterns) > 0

	result := &mcp.CallToolResult{
		Content: []mcp.Content{},
	}

	scanner := bufio.NewScanner(stream)
	var currentDiff strings.Builder
	var currentFilePath string
	// Regex to extract file path from "diff --git a/path/to/file b/path/to/file"
	re := regexp.MustCompile(`^diff --git a/(.+) b/(.+)`)

	processChunk := func() error {
		if currentDiff.Len() == 0 || currentFilePath == "" {
			return nil
		}

		// Default to include if no include patterns are given
		included := !hasIncludePatterns
		if hasIncludePatterns {
			match, err := matchAny(input.IncludePatterns, currentFilePath)
			if err != nil {
				return fmt.Errorf("error matching include pattern: %w", err)
			}
			if match {
				included = true
			}
		}

		excluded := false
		if hasExcludePatterns {
			match, err := matchAny(input.ExcludePatterns, currentFilePath)
			if err != nil {
				return fmt.Errorf("error matching exclude pattern: %w", err)
			}
			if match {
				excluded = true
			}
		}

		if included && !excluded {
			result.Content = append(result.Content, &mcp.TextContent{
				Text: currentDiff.String(),
			})
		}

		return nil
	}

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return nil, DiffOutput{}, ctx.Err()
		default:
		}

		line := scanner.Text()
		if matches := re.FindStringSubmatch(line); len(matches) > 2 {
			// This is a new file diff header. Process the previous chunk.
			if err := processChunk(); err != nil {
				return nil, DiffOutput{}, err
			}

			// Start a new chunk
			currentDiff.Reset()
			// The file path is in the second capture group (b path)
			currentFilePath = matches[2]
		}

		currentDiff.WriteString(line)
		currentDiff.WriteString("\n")
	}

	// Process the last chunk after the loop
	if err := processChunk(); err != nil {
		return nil, DiffOutput{}, err
	}

	if err := scanner.Err(); err != nil {
		return nil, DiffOutput{}, fmt.Errorf("reading pull request diff stream failed: %w", err)
	}

	return result, DiffOutput{}, nil
}

// getPullRequestDiffHandler handles getting the diff for a specific file in a pull request
func (h *Handler) getPullRequestDiffHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestDiffInput) (*mcp.CallToolResult, DiffOutput, error) {
	stream, err := h.client.GetPullRequestDiff(input)
	if err != nil {
		return nil, DiffOutput{}, fmt.Errorf("get pull request diff failed: %w", err)
	}
	defer stream.Close()

	// Read the entire stream into a string
	var buf bytes.Buffer
	_, err = io.Copy(&buf, stream)
	if err != nil {
		return nil, DiffOutput{}, fmt.Errorf("reading pull request diff failed: %w", err)
	}

	return nil, DiffOutput{Diff: buf.String()}, nil
}

// testPullRequestCanMergeHandler handles testing if a pull request can be merged
func (h *Handler) testPullRequestCanMergeHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.TestPullRequestCanMergeInput) (*mcp.CallToolResult, types.MapOutput, error) {
	mergeStatus, err := h.client.TestPullRequestCanMerge(input)
	if err != nil {
		return nil, nil, fmt.Errorf("test pull request can merge failed: %w", err)
	}

	return nil, mergeStatus, nil
}

// setPullRequestApproved handles approving a pull request
func (h *Handler) setPullRequestApproved(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.UpdatePullRequestWithoutStatusInput) (*mcp.CallToolResult, types.MapOutput, error) {
	// Create the full input with status set to APPROVED
	fullInput := bitbucket.UpdatePullRequestStatusInput{
		UpdatePullRequestWithoutStatusInput: input,
		Status:                              "APPROVED",
	}

	participant, err := h.client.UpdatePullRequestParticipantStatus(fullInput)
	if err != nil {
		return nil, nil, fmt.Errorf("approve pull request failed: %w", err)
	}

	return nil, participant, nil
}

// setPullRequestNeedsWork handles requesting changes for a pull request
func (h *Handler) setPullRequestNeedsWork(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.UpdatePullRequestWithoutStatusInput) (*mcp.CallToolResult, types.MapOutput, error) {
	// Create the full input with status set to NEEDS_WORK
	fullInput := bitbucket.UpdatePullRequestStatusInput{
		UpdatePullRequestWithoutStatusInput: input,
		Status:                              "NEEDS_WORK",
	}

	participant, err := h.client.UpdatePullRequestParticipantStatus(fullInput)
	if err != nil {
		return nil, nil, fmt.Errorf("request changes for pull request failed: %w", err)
	}

	return nil, participant, nil
}

// setPullRequestUnapproved handles resetting a pull request approval
func (h *Handler) setPullRequestUnapproved(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.UpdatePullRequestWithoutStatusInput) (*mcp.CallToolResult, types.MapOutput, error) {
	// Create the full input with status set to UNAPPROVED
	fullInput := bitbucket.UpdatePullRequestStatusInput{
		UpdatePullRequestWithoutStatusInput: input,
		Status:                              "UNAPPROVED",
	}

	participant, err := h.client.UpdatePullRequestParticipantStatus(fullInput)
	if err != nil {
		return nil, nil, fmt.Errorf("reset pull request approval failed: %w", err)
	}

	return nil, participant, nil
}

// AddPullRequestTools registers the pull request-related tools with the MCP server
func AddPullRequestTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[bitbucket.GetPullRequestsInput, types.MapOutput](server, "bitbucket_get_pull_requests", "Get a list of pull requests", handler.getPullRequestsHandler)
	utils.RegisterTool[bitbucket.GetPullRequestInput, types.MapOutput](server, "bitbucket_get_pull_request", "Get a specific pull request", handler.getPullRequestHandler)
	utils.RegisterTool[bitbucket.GetPullRequestActivitiesInput, types.MapOutput](server, "bitbucket_get_pull_request_activities", "Get activities for a specific pull request", handler.getPullRequestActivitiesHandler)
	utils.RegisterTool[bitbucket.GetPullRequestCommentsInput, types.MapOutput](server, "bitbucket_get_pull_request_comments", "Get comments for a specific pull request", handler.getPullRequestCommentsHandler)
	utils.RegisterTool[bitbucket.GetPullRequestChangesInput, types.MapOutput](server, "bitbucket_get_pull_request_changes", "Get changes for a specific pull request", handler.getPullRequestChangesHandler)
	utils.RegisterTool[bitbucket.GetPullRequestDiffStreamInput, DiffOutput](server, "bitbucket_get_pull_request_diff_stream", "Stream the diff for a pull request", handler.getPullRequestDiffStreamHandler)
	utils.RegisterTool[bitbucket.TestPullRequestCanMergeInput, types.MapOutput](server, "bitbucket_test_pull_request_can_merge", "Test if a pull request can be merged", handler.testPullRequestCanMergeHandler)
	utils.RegisterTool[bitbucket.GetPullRequestSuggestionsInput, types.MapOutput](server, "bitbucket_get_pull_request_suggestions", "Get pull request suggestions", handler.getPullRequestSuggestionsHandler)
	utils.RegisterTool[bitbucket.GetPullRequestJiraIssuesInput, types.MapOutput](server, "bitbucket_get_pull_request_jira_issues", "Get Jira issues linked to a pull request", handler.getPullRequestJiraIssuesHandler)
	utils.RegisterTool[bitbucket.GetPullRequestsForUserInput, types.MapOutput](server, "bitbucket_get_pull_requests_for_user", "Get pull requests for a specific user", handler.getPullRequestsForUserHandler)
	utils.RegisterTool[bitbucket.GetPullRequestCommentInput, types.MapOutput](server, "bitbucket_get_pull_request_comment", "Get a specific comment on a pull request", handler.getPullRequestCommentHandler)

	// Register specific tools for each status
	if permissions["bitbucket_update_pull_request_status"] {
		utils.RegisterTool[bitbucket.UpdatePullRequestWithoutStatusInput, types.MapOutput](server, "bitbucket_approve_pull_request", "Set pull request status to approved", handler.setPullRequestApproved)
		utils.RegisterTool[bitbucket.UpdatePullRequestWithoutStatusInput, types.MapOutput](server, "bitbucket_request_changes_pull_request", "Set pull request status to needs work", handler.setPullRequestNeedsWork)
		utils.RegisterTool[bitbucket.UpdatePullRequestWithoutStatusInput, types.MapOutput](server, "bitbucket_reset_pull_request_approval", "Set pull request status to unapproved", handler.setPullRequestUnapproved)
	}

	utils.RegisterTool[bitbucket.GetPullRequestDiffInput, DiffOutput](server, "bitbucket_get_pull_request_diff", "Get the diff for a specific file in a pull request", handler.getPullRequestDiffHandler)

	if permissions["bitbucket_merge_pull_request"] {
		utils.RegisterTool[bitbucket.MergePullRequestInput, types.MapOutput](server, "bitbucket_merge_pull_request", "Merge a pull request", handler.mergePullRequestHandler)
	}

	if permissions["bitbucket_decline_pull_request"] {
		utils.RegisterTool[bitbucket.DeclinePullRequestInput, types.MapOutput](server, "bitbucket_decline_pull_request", "Decline a pull request", handler.declinePullRequestHandler)
	}

	if permissions["bitbucket_add_pull_request_comment"] {
		utils.RegisterTool[bitbucket.AddPullRequestCommentInput, types.MapOutput](server, "bitbucket_add_pull_request_comment", "Add an enhanced comment to a pull request. Supports general comments, replies, inline comments, and code suggestions", handler.addPullRequestCommentHandler)
	}
}
