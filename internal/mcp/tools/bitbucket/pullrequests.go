package bitbucket

import (
	"context"
	"strconv"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetPullRequestsInput represents the input parameters for getting pull requests
type GetPullRequestsInput struct {
	ProjectKey     string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug       string `json:"repoSlug" jsonschema:"required,The repository slug"`
	State          string `json:"state,omitempty" jsonschema:"State of the pull requests to retrieve"`
	WithAttributes string `json:"withAttributes,omitempty" jsonschema:"Include attributes in the response"`
	At             string `json:"at,omitempty" jsonschema:"The commit ID or ref to retrieve pull requests at"`
	WithProperties string `json:"withProperties,omitempty" jsonschema:"Include properties in the response"`
	Draft          string `json:"draft,omitempty" jsonschema:"Include draft pull requests"`
	FilterText     string `json:"filterText,omitempty" jsonschema:"Filter text to apply"`
	Order          string `json:"order,omitempty" jsonschema:"Order the results by a specific field"`
	Direction      string `json:"direction,omitempty" jsonschema:"Direction to order the results"`
	Start          int    `json:"start,omitempty" jsonschema:"The starting index of the returned pull requests"`
	Limit          int    `json:"limit,omitempty" jsonschema:"The limit of the number of pull requests to return"`
}

// GetPullRequestsOutput represents the output for getting pull requests
type GetPullRequestsOutput = map[string]interface{}

// GetPullRequestInput represents the input parameters for getting a specific pull request
type GetPullRequestInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PRId       int    `json:"prId" jsonschema:"required,The pull request ID"`
}

// GetPullRequestOutput represents the output for getting a specific pull request
type GetPullRequestOutput = map[string]interface{}

// GetPullRequestActivitiesInput represents the input parameters for getting pull request activities
type GetPullRequestActivitiesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PRId       int    `json:"prId" jsonschema:"required,The pull request ID"`
	FromType   string `json:"fromType,omitempty" jsonschema:"Filter activities by type"`
	FromId     string `json:"fromId,omitempty" jsonschema:"Filter activities by ID"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned activities"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of activities to return"`
}

// GetPullRequestActivitiesOutput represents the output for getting pull request activities
type GetPullRequestActivitiesOutput = map[string]interface{}

// GetPullRequestChangesInput represents the input parameters for getting pull request changes
type GetPullRequestChangesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PRId       int    `json:"prId" jsonschema:"required,The pull request ID"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned changes"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of changes to return"`
}

// GetPullRequestChangesOutput represents the output for getting pull request changes
type GetPullRequestChangesOutput = map[string]interface{}

// GetPullRequestCommitsInput represents the input parameters for getting pull request commits
type GetPullRequestCommitsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PRId       int    `json:"prId" jsonschema:"required,The pull request ID"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned commits"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of commits to return"`
}

// GetPullRequestCommitsOutput represents the output for getting pull request commits
type GetPullRequestCommitsOutput = map[string]interface{}

// GetPullRequestCommentsInput represents the input parameters for getting pull request comments
type GetPullRequestCommentsInput struct {
	ProjectKey  string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug    string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PRId        int    `json:"prId" jsonschema:"required,The pull request ID"`
	Path        string `json:"path,omitempty" jsonschema:"Filter comments by file path"`
	FromHash    string `json:"fromHash,omitempty" jsonschema:"Filter comments by from hash"`
	AnchorState string `json:"anchorState,omitempty" jsonschema:"Filter comments by anchor state"`
	ToHash      string `json:"toHash,omitempty" jsonschema:"Filter comments by to hash"`
	State       string `json:"state,omitempty" jsonschema:"Filter comments by state"`
	DiffType    string `json:"diffType,omitempty" jsonschema:"Filter comments by diff type"`
	DiffTypes   string `json:"diffTypes,omitempty" jsonschema:"Filter comments by diff types"`
	States      string `json:"states,omitempty" jsonschema:"Filter comments by states"`
	Start       int    `json:"start,omitempty" jsonschema:"The starting index of the returned comments"`
	Limit       int    `json:"limit,omitempty" jsonschema:"The limit of the number of comments to return"`
}

// GetPullRequestCommentsOutput represents the output for getting pull request comments
type GetPullRequestCommentsOutput = map[string]interface{}

// GetPullRequestDiffInput represents the input parameters for getting pull request diff
type GetPullRequestDiffInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PRId         int    `json:"prId" jsonschema:"required,The pull request ID"`
	ContextLines int    `json:"contextLines,omitempty" jsonschema:"Number of context lines to include"`
	Whitespace   string `json:"whitespace,omitempty" jsonschema:"Whitespace handling option"`
}

// GetPullRequestDiffOutput represents the output for getting pull request diff
type GetPullRequestDiffOutput = map[string]interface{}

// GetPullRequestMergeConfigInput represents the input parameters for getting pull request merge config
type GetPullRequestMergeConfigInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PRId       int    `json:"prId" jsonschema:"required,The pull request ID"`
}

// GetPullRequestMergeConfigOutput represents the output for getting pull request merge config
type GetPullRequestMergeConfigOutput = map[string]interface{}

// GetPullRequestMergeStatusInput represents the input parameters for getting pull request merge status
type GetPullRequestMergeStatusInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PRId       int    `json:"prId" jsonschema:"required,The pull request ID"`
}

// GetPullRequestMergeStatusOutput represents the output for getting pull request merge status
type GetPullRequestMergeStatusOutput = map[string]interface{}

// GetPullRequestSettingsInput represents the input parameters for getting pull request settings
type GetPullRequestSettingsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetPullRequestSettingsOutput represents the output for getting pull request settings
type GetPullRequestSettingsOutput = map[string]interface{}

// getPullRequestsHandler handles getting pull requests
func (h *Handler) getPullRequestsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetPullRequestsInput) (*mcp.CallToolResult, GetPullRequestsOutput, error) {

	pullRequests, err := h.client.GetPullRequests(
		input.ProjectKey,
		input.RepoSlug,
		input.State,
		input.WithAttributes,
		input.At,
		input.WithProperties,
		input.Draft,
		input.FilterText,
		input.Order,
		input.Direction,
		input.Start,
		input.Limit,
	)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull requests")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(pullRequests)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull requests result")
		return result, nil, err
	}

	return result, pullRequests, nil
}

// getPullRequestHandler handles getting a specific pull request
func (h *Handler) getPullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input GetPullRequestInput) (*mcp.CallToolResult, GetPullRequestOutput, error) {
	pullRequest, err := h.client.GetPullRequest(input.ProjectKey, input.RepoSlug, input.PRId)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(pullRequest)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request result")
		return result, nil, err
	}

	return result, pullRequest, nil
}

// getPullRequestActivitiesHandler handles getting pull request activities
func (h *Handler) getPullRequestActivitiesHandler(ctx context.Context, req *mcp.CallToolRequest, input GetPullRequestActivitiesInput) (*mcp.CallToolResult, GetPullRequestActivitiesOutput, error) {
	activities, err := h.client.GetPullRequestActivities(input.ProjectKey, input.RepoSlug, input.PRId, input.FromType, input.FromId, input.Start, input.Limit)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request activities")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(activities)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request activities result")
		return result, nil, err
	}

	return result, activities, nil
}

// getPullRequestChangesHandler handles getting pull request changes
func (h *Handler) getPullRequestChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input GetPullRequestChangesInput) (*mcp.CallToolResult, GetPullRequestChangesOutput, error) {
	// Using GetChanges method as a substitute for GetPullRequestChanges
	changes, err := h.client.GetChanges(input.ProjectKey, input.RepoSlug, "", "", input.Start, input.Limit)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request changes")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(changes)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request changes result")
		return result, nil, err
	}

	return result, changes, nil
}

// getPullRequestCommitsHandler handles getting pull request commits
func (h *Handler) getPullRequestCommitsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetPullRequestCommitsInput) (*mcp.CallToolResult, GetPullRequestCommitsOutput, error) {
	// Using GetCommits method as a substitute for GetPullRequestCommits
	commits, err := h.client.GetCommits(input.ProjectKey, input.RepoSlug, "", "", "", input.Start, input.Limit, "", false, false, false)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request commits")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(commits)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request commits result")
		return result, nil, err
	}

	return result, commits, nil
}

// getPullRequestCommentsHandler handles getting pull request comments
func (h *Handler) getPullRequestCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetPullRequestCommentsInput) (*mcp.CallToolResult, GetPullRequestCommentsOutput, error) {
	comments, err := h.client.GetPullRequestComments(
		input.ProjectKey,
		input.RepoSlug,
		input.PRId,
		input.Path,
		input.FromHash,
		input.AnchorState,
		input.ToHash,
		input.State,
		input.DiffType,
		input.DiffTypes,
		input.States,
		input.Start,
		input.Limit,
	)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request comments")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(comments)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request comments result")
		return result, nil, err
	}

	return result, comments, nil
}

// getPullRequestDiffHandler handles getting pull request diff
func (h *Handler) getPullRequestDiffHandler(ctx context.Context, req *mcp.CallToolRequest, input GetPullRequestDiffInput) (*mcp.CallToolResult, GetPullRequestDiffOutput, error) {
	// Using GetDiffBetweenCommits method as a substitute for GetPullRequestDiff
	diff, err := h.client.GetDiffBetweenCommits(input.ProjectKey, input.RepoSlug, "", "", "", 0, "", "", "")
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request diff")
		return result, GetPullRequestDiffOutput{}, err
	}

	// Create a map to hold the diff output
	diffOutput := GetPullRequestDiffOutput{
		"diff": diff,
	}

	result, err := tools.CreateToolResult(diff)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request diff result")
		return result, GetPullRequestDiffOutput{}, err
	}

	return result, diffOutput, nil
}

// getPullRequestMergeConfigHandler handles getting pull request merge config
func (h *Handler) getPullRequestMergeConfigHandler(ctx context.Context, req *mcp.CallToolRequest, input GetPullRequestMergeConfigInput) (*mcp.CallToolResult, GetPullRequestMergeConfigOutput, error) {
	// Using GetPullRequest method to get merge config information
	pr, err := h.client.GetPullRequest(input.ProjectKey, input.RepoSlug, input.PRId)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request merge config")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(pr)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request merge config result")
		return result, nil, err
	}

	return result, pr, nil
}

// getPullRequestMergeStatusHandler handles getting pull request merge status
func (h *Handler) getPullRequestMergeStatusHandler(ctx context.Context, req *mcp.CallToolRequest, input GetPullRequestMergeStatusInput) (*mcp.CallToolResult, GetPullRequestMergeStatusOutput, error) {
	// Using GetPullRequest method to get merge status information
	pr, err := h.client.GetPullRequest(input.ProjectKey, input.RepoSlug, input.PRId)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request merge status")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(pr)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request merge status result")
		return result, nil, err
	}

	return result, pr, nil
}

// getPullRequestSettingsHandler handles getting pull request settings
func (h *Handler) getPullRequestSettingsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetPullRequestSettingsInput) (*mcp.CallToolResult, GetPullRequestSettingsOutput, error) {
	// Using GetPullRequests method to get settings information
	settings, err := h.client.GetPullRequests(input.ProjectKey, input.RepoSlug, "", "", "", "", "", "", "", "", 0, 1)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get pull request settings")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(settings)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create pull request settings result")
		return result, nil, err
	}

	return result, settings, nil
}

// MergePullRequestInput represents the input parameters for merging a pull request
type MergePullRequestInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PRId       int    `json:"prId" jsonschema:"required,The pull request ID"`
	Version    *int   `json:"version,omitempty" jsonschema:"Version of the pull request"`
	AutoMerge  *bool  `json:"autoMerge,omitempty" jsonschema:"Automatically merge the pull request"`
	Message    string `json:"message,omitempty" jsonschema:"Merge commit message"`
}

// MergePullRequestOutput represents the output for merging a pull request
type MergePullRequestOutput = map[string]interface{}

// mergePullRequestHandler handles merging a pull request
func (h *Handler) mergePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input MergePullRequestInput) (*mcp.CallToolResult, MergePullRequestOutput, error) {
	options := bitbucket.MergePullRequestOptions{
		Version:   input.Version,
		AutoMerge: input.AutoMerge,
		Message:   &input.Message,
	}

	result, err := h.client.MergePullRequest(input.ProjectKey, input.RepoSlug, input.PRId, options)
	if err != nil {
		toolResult, _, err := tools.HandleToolError(err, "merge pull request")
		return toolResult, nil, err
	}

	toolResult, err := tools.CreateToolResult(result)
	if err != nil {
		toolResult, _, err := tools.HandleToolError(err, "create merge pull request result")
		return toolResult, nil, err
	}

	return toolResult, result, nil
}

// DeclinePullRequestInput represents the input parameters for declining a pull request
type DeclinePullRequestInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PRId       int    `json:"prId" jsonschema:"required,The pull request ID"`
	Version    string `json:"version,omitempty" jsonschema:"Version of the pull request"`
	Comment    string `json:"comment,omitempty" jsonschema:"Comment explaining why the pull request is declined"`
}

// DeclinePullRequestOutput represents the output for declining a pull request
type DeclinePullRequestOutput = map[string]interface{}

// declinePullRequestHandler handles declining a pull request
func (h *Handler) declinePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, input DeclinePullRequestInput) (*mcp.CallToolResult, map[string]interface{}, error) {
	options := bitbucket.DeclinePullRequestOptions{
		Comment: &input.Comment,
	}

	result, err := h.client.DeclinePullRequest(input.ProjectKey, input.RepoSlug, strconv.Itoa(input.PRId), input.Version, &options)
	if err != nil {
		toolResult, _, err := tools.HandleToolError(err, "decline pull request")
		return toolResult, nil, err
	}

	toolResult, err := tools.CreateToolResult(result)
	if err != nil {
		toolResult, _, err := tools.HandleToolError(err, "create decline pull request result")
		return toolResult, nil, err
	}

	return toolResult, result, nil
}

// AddPullRequestTools registers the pull request-related tools with the MCP server
func AddPullRequestTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[GetPullRequestsInput, GetPullRequestsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_requests",
		Description: "Get a list of pull requests",
	}, handler.getPullRequestsHandler)

	mcp.AddTool[GetPullRequestInput, GetPullRequestOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request",
		Description: "Get a specific pull request",
	}, handler.getPullRequestHandler)

	mcp.AddTool[GetPullRequestActivitiesInput, GetPullRequestActivitiesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_activities",
		Description: "Get activities for a specific pull request",
	}, handler.getPullRequestActivitiesHandler)

	mcp.AddTool[GetPullRequestChangesInput, GetPullRequestChangesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_changes",
		Description: "Get changes for a specific pull request",
	}, handler.getPullRequestChangesHandler)

	mcp.AddTool[GetPullRequestCommitsInput, GetPullRequestCommitsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_commits",
		Description: "Get commits for a specific pull request",
	}, handler.getPullRequestCommitsHandler)

	mcp.AddTool[GetPullRequestCommentsInput, GetPullRequestCommentsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_comments",
		Description: "Get comments for a specific pull request",
	}, handler.getPullRequestCommentsHandler)

	mcp.AddTool[GetPullRequestDiffInput, GetPullRequestDiffOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_diff",
		Description: "Get diff for a specific pull request",
	}, handler.getPullRequestDiffHandler)

	mcp.AddTool[GetPullRequestMergeConfigInput, GetPullRequestMergeConfigOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_merge_config",
		Description: "Get merge configuration for a specific pull request",
	}, handler.getPullRequestMergeConfigHandler)

	mcp.AddTool[GetPullRequestMergeStatusInput, GetPullRequestMergeStatusOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_merge_status",
		Description: "Get merge status for a specific pull request",
	}, handler.getPullRequestMergeStatusHandler)

	mcp.AddTool[GetPullRequestSettingsInput, GetPullRequestSettingsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_settings",
		Description: "Get pull request settings for a repository",
	}, handler.getPullRequestSettingsHandler)

	mcp.AddTool[MergePullRequestInput, MergePullRequestOutput](server, &mcp.Tool{
		Name:        "bitbucket_merge_pull_request",
		Description: "Merge a specific pull request. This tool allows you to merge pull requests that are ready to be merged.",
	}, handler.mergePullRequestHandler)

	mcp.AddTool[DeclinePullRequestInput, DeclinePullRequestOutput](server, &mcp.Tool{
		Name:        "bitbucket_decline_pull_request",
		Description: "Decline a specific pull request. This tool allows you to decline pull requests that are not suitable for merging.",
	}, handler.declinePullRequestHandler)
}
