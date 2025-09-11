package bitbucket

import (
	"context"
	"fmt"
	"strconv"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getPullRequestsHandler handles getting pull requests
func (h *Handler) getPullRequestsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get pull requests", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		state, _ := tools.GetStringArg(args, "state")
		withAttributes, _ := tools.GetStringArg(args, "withAttributes")
		at, _ := tools.GetStringArg(args, "at")
		withProperties, _ := tools.GetStringArg(args, "withProperties")
		draft, _ := tools.GetStringArg(args, "draft")
		filterText, _ := tools.GetStringArg(args, "filterText")
		order, _ := tools.GetStringArg(args, "order")
		direction, _ := tools.GetStringArg(args, "direction")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetPullRequests(projectKey, repoSlug, state, withAttributes, at, withProperties, draft, filterText, order, direction, start, limit)
	})
}

// getPullRequestHandler handles getting a specific pull request
func (h *Handler) getPullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get pull request", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		pullRequestId := tools.GetIntArg(args, "pullRequestId", 0)
		if pullRequestId <= 0 {
			return nil, fmt.Errorf("missing or invalid pullRequestId parameter")
		}

		return h.client.GetPullRequest(projectKey, repoSlug, pullRequestId)
	})
}

// getPullRequestActivitiesHandler handles getting activities for a specific pull request
func (h *Handler) getPullRequestActivitiesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get pull request activities", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		pullRequestId := tools.GetIntArg(args, "pullRequestId", 0)
		if pullRequestId <= 0 {
			return nil, fmt.Errorf("missing or invalid pullRequestId parameter")
		}

		fromType, _ := tools.GetStringArg(args, "fromType")
		fromId, _ := tools.GetStringArg(args, "fromId")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetPullRequestActivities(projectKey, repoSlug, pullRequestId, fromType, fromId, start, limit)
	})
}

// addPullRequestCommentHandler handles adding a comment to a specific pull request
func (h *Handler) addPullRequestCommentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("add pull request comment", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		pullRequestId := tools.GetIntArg(args, "pullRequestId", 0)
		if pullRequestId <= 0 {
			return nil, fmt.Errorf("missing or invalid pullRequestId parameter")
		}

		commentText, ok := tools.GetStringArg(args, "commentText")
		if !ok {
			return nil, fmt.Errorf("missing or invalid commentText parameter")
		}

		return h.client.AddPullRequestComment(projectKey, repoSlug, pullRequestId, commentText)
	})
}

// mergePullRequestHandler handles merging a specific pull request
func (h *Handler) mergePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("merge pull request", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		pullRequestId := tools.GetIntArg(args, "pullRequestId", 0)
		if pullRequestId <= 0 {
			return nil, fmt.Errorf("missing or invalid pullRequestId parameter")
		}

		// Prepare merge options
		options := bitbucket.MergePullRequestOptions{}

		// Handle version parameter
		version := tools.GetIntArg(args, "version", 0)
		if version != 0 {
			options.Version = &version
		}

		// Handle autoMerge parameter
		autoMerge := tools.GetBoolArg(args, "autoMerge", false)
		options.AutoMerge = &autoMerge

		// Handle autoSubject parameter
		if val, ok := tools.GetStringArg(args, "autoSubject"); ok {
			options.AutoSubject = &val
		}

		// Handle message parameter
		if val, ok := tools.GetStringArg(args, "message"); ok {
			options.Message = &val
		}

		// Handle strategyId parameter
		if val, ok := tools.GetStringArg(args, "strategyId"); ok {
			options.StrategyId = &val
		}

		return h.client.MergePullRequest(projectKey, repoSlug, pullRequestId, options)
	})
}

// declinePullRequestHandler handles declining a specific pull request
func (h *Handler) declinePullRequestHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("decline pull request", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		// Get pullRequestId as integer and convert to string for the client method
		pullRequestId := tools.GetIntArg(args, "pullRequestId", 0)
		if pullRequestId <= 0 {
			return nil, fmt.Errorf("missing or invalid pullRequestId parameter")
		}
		pullRequestIdStr := strconv.Itoa(pullRequestId)

		version, _ := tools.GetStringArg(args, "version")

		// Prepare decline options
		options := &bitbucket.DeclinePullRequestOptions{}

		// Handle version parameter in options
		optionsVersion := tools.GetIntArg(args, "optionsVersion", 0)
		if optionsVersion != 0 {
			options.Version = &optionsVersion
		}

		// Handle comment parameter
		if val, ok := tools.GetStringArg(args, "comment"); ok {
			options.Comment = &val
		}

		return h.client.DeclinePullRequest(projectKey, repoSlug, pullRequestIdStr, version, options)
	})
}

// getPullRequestJiraIssuesHandler handles getting Jira issues for a specific pull request
func (h *Handler) getPullRequestJiraIssuesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get pull request Jira issues", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		pullRequestId := tools.GetIntArg(args, "pullRequestId", 0)
		if pullRequestId <= 0 {
			return nil, fmt.Errorf("missing or invalid pullRequestId parameter")
		}

		return h.client.GetPullRequestJiraIssues(projectKey, repoSlug, pullRequestId)
	})
}

// getPullRequestCommentHandler handles getting a specific comment on a pull request
func (h *Handler) getPullRequestCommentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get pull request comment", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		pullRequestId := tools.GetIntArg(args, "pullRequestId", 0)
		if pullRequestId <= 0 {
			return nil, fmt.Errorf("missing or invalid pullRequestId parameter")
		}

		commentId, ok := tools.GetStringArg(args, "commentId")
		if !ok {
			return nil, fmt.Errorf("missing or invalid commentId parameter")
		}

		return h.client.GetPullRequestComment(projectKey, repoSlug, pullRequestId, commentId)
	})
}

// getPullRequestCommentsHandler handles getting comments on a pull request
func (h *Handler) getPullRequestCommentsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get pull request comments", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		pullRequestId := tools.GetIntArg(args, "pullRequestId", 0)
		if pullRequestId <= 0 {
			return nil, fmt.Errorf("missing or invalid pullRequestId parameter")
		}

		path, _ := tools.GetStringArg(args, "path")
		fromHash, _ := tools.GetStringArg(args, "fromHash")
		anchorState, _ := tools.GetStringArg(args, "anchorState")
		toHash, _ := tools.GetStringArg(args, "toHash")
		state, _ := tools.GetStringArg(args, "state")
		diffType, _ := tools.GetStringArg(args, "diffType")
		diffTypes, _ := tools.GetStringArg(args, "diffTypes")
		states, _ := tools.GetStringArg(args, "states")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetPullRequestComments(projectKey, repoSlug, pullRequestId, path, fromHash, anchorState, toHash, state, diffType, diffTypes, states, start, limit)
	})
}

// getPullRequestSuggestionsHandler handles getting pull request suggestions
func (h *Handler) getPullRequestSuggestionsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get pull request suggestions", func() (interface{}, error) {
		changesSince, _ := tools.GetStringArg(args, "changesSince")

		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetPullRequestSuggestions(changesSince, limit)
	})
}

// getPullRequestsForUserHandler handles getting pull requests for a specific user
func (h *Handler) getPullRequestsForUserHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get pull requests for user", func() (interface{}, error) {
		closedSince, _ := tools.GetStringArg(args, "closedSince")
		role, _ := tools.GetStringArg(args, "role")
		participantStatus, _ := tools.GetStringArg(args, "participantStatus")
		state, _ := tools.GetStringArg(args, "state")
		user, _ := tools.GetStringArg(args, "user")
		order, _ := tools.GetStringArg(args, "order")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetPullRequestsForUser(closedSince, role, participantStatus, state, user, order, start, limit)
	})
}

// AddPullRequestTools registers the pull request-related tools with the MCP server
func AddPullRequestTools(server *mcp.Server, client *bitbucket.BitbucketClient, hasWritePermission bool) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request",
		Description: "Retrieve details of a specific pull request. This tool provides comprehensive information about a pull request including its status, participants, and comments.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "Key of the project containing the pull request (e.g., 'PROJ').",
				},
				"repoSlug": {
					Type:        "string",
					Description: "Slug of the repository containing the pull request (e.g., 'my-repo').",
				},
				"pullRequestId": {
					Type:        "integer",
					Description: "ID of the pull request to retrieve.",
				},
			},
			Required: []string{"projectKey", "repoSlug", "pullRequestId"},
		},
	}, handler.getPullRequestHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_pull_requests",
		Description: "List pull requests in a repository with optional filtering. This tool allows you to retrieve multiple pull requests based on various criteria.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "Key of the project containing the pull requests (e.g., 'PROJ').",
				},
				"repoSlug": {
					Type:        "string",
					Description: "Slug of the repository containing the pull requests (e.g., 'my-repo').",
				},
				"state": {
					Type:        "string",
					Description: "Filter pull requests by state (e.g., 'OPEN', 'MERGED', 'DECLINED').",
				},
				"withAttributes": {
					Type:        "string",
					Description: "Include additional attributes in the response.",
				},
				"at": {
					Type:        "string",
					Description: "Filter pull requests by ref.",
				},
				"withProperties": {
					Type:        "string",
					Description: "Include properties in the response.",
				},
				"draft": {
					Type:        "string",
					Description: "Filter by draft status.",
				},
				"filterText": {
					Type:        "string",
					Description: "Filter by text.",
				},
				"order": {
					Type:        "string",
					Description: "Sort order.",
				},
				"direction": {
					Type:        "string",
					Description: "Sort direction.",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned pull requests (for pagination). Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of pull requests to return (for pagination). Default: 25, Max: 100",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getPullRequestsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_activities",
		Description: "Retrieve activities related to a pull request. This includes comments, approvals, and other events associated with the pull request.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "Key of the project containing the pull request (e.g., 'PROJ').",
				},
				"repoSlug": {
					Type:        "string",
					Description: "Slug of the repository containing the pull request (e.g., 'my-repo').",
				},
				"pullRequestId": {
					Type:        "integer",
					Description: "ID of the pull request to get activities for.",
				},
				"fromType": {
					Type:        "string",
					Description: "Filter activities by type.",
				},
				"fromId": {
					Type:        "string",
					Description: "Filter activities by ID.",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned activities (for pagination). Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of activities to return (for pagination). Default: 25, Max: 100",
				},
			},
			Required: []string{"projectKey", "repoSlug", "pullRequestId"},
		},
	}, handler.getPullRequestActivitiesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_comments",
		Description: "Retrieve comments on a pull request. This tool allows you to get detailed information about all comments made on a specific pull request.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "Key of the project containing the pull request (e.g., 'PROJ').",
				},
				"repoSlug": {
					Type:        "string",
					Description: "Slug of the repository containing the pull request (e.g., 'my-repo').",
				},
				"pullRequestId": {
					Type:        "integer",
					Description: "ID of the pull request to get comments for.",
				},
				"path": {
					Type:        "string",
					Description: "Filter comments by file path.",
				},
				"fromHash": {
					Type:        "string",
					Description: "Filter comments by from hash.",
				},
				"anchorState": {
					Type:        "string",
					Description: "Filter comments by anchor state.",
				},
				"toHash": {
					Type:        "string",
					Description: "Filter comments by to hash.",
				},
				"state": {
					Type:        "string",
					Description: "Filter comments by state.",
				},
				"diffType": {
					Type:        "string",
					Description: "Filter comments by diff type.",
				},
				"diffTypes": {
					Type:        "string",
					Description: "Filter comments by diff types.",
				},
				"states": {
					Type:        "string",
					Description: "Filter comments by states.",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned comments (for pagination). Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of comments to return (for pagination). Default: 25, Max: 100",
				},
			},
			Required: []string{"projectKey", "repoSlug", "pullRequestId"},
		},
	}, handler.getPullRequestCommentsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_suggestions",
		Description: "Get pull request suggestions based on recent changes. This tool helps identify potential pull requests that could be created based on recent commits.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"changesSince": {
					Type:        "string",
					Description: "Timestamp to filter changes since (format: ISO 8601).",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of suggestions to return. Default: 25, Max: 100",
				},
			},
			Required: []string{},
		},
	}, handler.getPullRequestSuggestionsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_pull_requests_for_user",
		Description: "Get pull requests associated with a specific user. This tool allows you to retrieve pull requests where the user is involved in various roles.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"user": {
					Type:        "string",
					Description: "Username of the user to filter by.",
				},
				"state": {
					Type:        "string",
					Description: "State of the pull request (e.g., 'OPEN', 'MERGED', 'DECLINED').",
				},
				"role": {
					Type:        "string",
					Description: "Role of the user (e.g., 'AUTHOR', 'REVIEWER', 'PARTICIPANT').",
				},
				"participantStatus": {
					Type:        "string",
					Description: "Status of the participant (e.g., 'APPROVED', 'UNAPPROVED').",
				},
				"order": {
					Type:        "string",
					Description: "Order of the pull requests (e.g., 'NEWEST', 'OLDEST').",
				},
				"closedSince": {
					Type:        "string",
					Description: "Timestamp to filter closed pull requests since (format: ISO 8601).",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned pull requests (for pagination). Default: 0",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of pull requests to return (for pagination). Default: 25, Max: 100",
				},
			},
			Required: []string{},
		},
	}, handler.getPullRequestsForUserHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_jira_issues",
		Description: "Get Jira issues referenced in a pull request. This tool retrieves all Jira issues that are mentioned in a pull request.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "Key of the project containing the pull request (e.g., 'PROJ').",
				},
				"repoSlug": {
					Type:        "string",
					Description: "Slug of the repository containing the pull request (e.g., 'my-repo').",
				},
				"pullRequestId": {
					Type:        "integer",
					Description: "ID of the pull request to get Jira issues for.",
				},
			},
			Required: []string{"projectKey", "repoSlug", "pullRequestId"},
		},
	}, handler.getPullRequestJiraIssuesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_pull_request_comment",
		Description: "Get a specific comment on a pull request. This tool retrieves detailed information about a specific comment on a pull request.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "Key of the project containing the pull request (e.g., 'PROJ').",
				},
				"repoSlug": {
					Type:        "string",
					Description: "Slug of the repository containing the pull request (e.g., 'my-repo').",
				},
				"pullRequestId": {
					Type:        "integer",
					Description: "ID of the pull request containing the comment.",
				},
				"commentId": {
					Type:        "string",
					Description: "ID of the comment to retrieve.",
				},
			},
			Required: []string{"projectKey", "repoSlug", "pullRequestId", "commentId"},
		},
	}, handler.getPullRequestCommentHandler)

	// Only register write tools if write permission is enabled
	if hasWritePermission {
		mcp.AddTool(server, &mcp.Tool{
			Name:        "bitbucket_add_pull_request_comment",
			Description: "Add a comment to a specific pull request. This tool allows you to add comments to pull requests for discussion and feedback.",
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"projectKey": {
						Type:        "string",
						Description: "Key of the project containing the pull request (e.g., 'PROJ').",
					},
					"repoSlug": {
						Type:        "string",
						Description: "Slug of the repository containing the pull request (e.g., 'my-repo').",
					},
					"pullRequestId": {
						Type:        "integer",
						Description: "ID of the pull request to add comment to.",
					},
					"commentText": {
						Type:        "string",
						Description: "Text of the comment to add.",
					},
				},
				Required: []string{"projectKey", "repoSlug", "pullRequestId", "commentText"},
			},
		}, handler.addPullRequestCommentHandler)

		mcp.AddTool(server, &mcp.Tool{
			Name:        "bitbucket_merge_pull_request",
			Description: "Merge a specific pull request. This tool allows you to merge pull requests after review and approval.",
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"projectKey": {
						Type:        "string",
						Description: "Key of the project containing the pull request (e.g., 'PROJ').",
					},
					"repoSlug": {
						Type:        "string",
						Description: "Slug of the repository containing the pull request (e.g., 'my-repo').",
					},
					"pullRequestId": {
						Type:        "integer",
						Description: "ID of the pull request to merge.",
					},
					"version": {
						Type:        "integer",
						Description: "Version of the pull request to merge.",
					},
					"autoMerge": {
						Type:        "boolean",
						Description: "Whether to auto-merge the pull request.",
					},
					"autoSubject": {
						Type:        "string",
						Description: "Auto-generated subject for the merge commit.",
					},
					"message": {
						Type:        "string",
						Description: "Commit message for the merge.",
					},
					"strategyId": {
						Type:        "string",
						Description: "Merge strategy to use.",
					},
				},
				Required: []string{"projectKey", "repoSlug", "pullRequestId"},
			},
		}, handler.mergePullRequestHandler)

		mcp.AddTool(server, &mcp.Tool{
			Name:        "bitbucket_decline_pull_request",
			Description: "Decline a specific pull request. This tool allows you to decline pull requests that are not suitable for merging.",
			InputSchema: &jsonschema.Schema{
				Type: "object",
				Properties: map[string]*jsonschema.Schema{
					"projectKey": {
						Type:        "string",
						Description: "Key of the project containing the pull request (e.g., 'PROJ').",
					},
					"repoSlug": {
						Type:        "string",
						Description: "Slug of the repository containing the pull request (e.g., 'my-repo').",
					},
					"pullRequestId": {
						Type:        "integer",
						Description: "ID of the pull request to decline.",
					},
					"version": {
						Type:        "string",
						Description: "Version of the pull request to decline.",
					},
					"comment": {
						Type:        "string",
						Description: "Comment explaining why the pull request is declined.",
					},
					"optionsVersion": {
						Type:        "integer",
						Description: "Version in the decline options.",
					},
				},
				Required: []string{"projectKey", "repoSlug", "pullRequestId"},
			},
		}, handler.declinePullRequestHandler)
	}
}
