# Development Guide for Atlassian DC MCP Go

This guide explains how to develop new interfaces in the Atlassian DC MCP Go project using AI coding assistants.

## Table of Contents
1. [Overview](#overview)
2. [Understanding the Codebase Structure](#understanding-the-codebase-structure)
3. [Adding a New Bitbucket API Endpoint - Example: Get Pull Request Changes](#adding-a-new-bitbucket-api-endpoint---example-get-pull-request-changes)
   - [Step 1: Define Input/Output Types](#step-1-define-inputoutput-types)
   - [Step 2: Implement Client Method](#step-2-implement-client-method)
   - [Step 3: Implement MCP Handler](#step-3-implement-mcp-handler)
   - [Step 4: Register the Tool](#step-4-register-the-tool)
4. [Using AI Coding Assistants for Development](#using-ai-coding-assistants-for-development)
   - [AI Coding Assistant Prompt Template](#ai-coding-assistant-prompt-template)
5. [Testing](#testing)
6. [Best Practices](#best-practices)

## Overview

This project follows a clean architecture pattern with separate layers for:
- Client layer: Direct API calls to Atlassian products
- MCP layer: Handlers that interface between the MCP protocol and the client
- Types layer: Shared data structures

## Understanding the Codebase Structure

The project is organized as follows:
```
internal/
├── client/              # Direct API clients for Atlassian products
│   ├── bitbucket/
│   ├── confluence/
│   └── jira/
├── mcp/                 # MCP protocol handlers
│   └── tools/
│       ├── bitbucket/
│       ├── confluence/
│       └── jira/
└── types/               # Shared data types
```

Each Atlassian product has:
1. A client implementation with methods for each API endpoint
2. Type definitions for inputs and outputs
3. MCP handlers that register tools with the MCP server

## Adding a New Bitbucket API Endpoint - Example: Get Pull Request Changes

Let's walk through implementing the "Get Pull Request Changes" endpoint as documented in [Bitbucket REST API](https://developer.atlassian.com/server/bitbucket/rest/v1000/api-group-pull-requests/#api-api-latest-projects-projectkey-repos-repositoryslug-pull-requests-pullrequestid-changes-get).

### Step 1: Define Input/Output Types

First, we need to define the input parameters for the API call in [internal/client/bitbucket/pullrequests_types.go](file:///home/stephen/workspace/atlassian-dc-mcp-go/internal/client/bitbucket/pullrequests_types.go):

```go
// GetPullRequestChangesInput represents the input parameters for getting pull request changes
type GetPullRequestChangesInput struct {
	CommonInput
	PaginationInput
	PullRequestID int `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	// Additional query parameters for pull request changes
	SinceId     string `json:"sinceId,omitempty" jsonschema:"The change ID from which to start retrieving changes"`
	ChangeScope string `json:"changeScope,omitempty" jsonschema:"The scope of changes to retrieve (e.g. UNREVIEWED)"`
	UntilId     string `json:"untilId,omitempty" jsonschema:"The change ID until which to retrieve changes"`
	WithComments bool   `json:"withComments,omitempty" jsonschema:"Include comments in the response"`
	Start       int    `json:"start,omitempty" jsonschema:"Pagination start position (ignored by server)"`
	Limit       int    `json:"limit,omitempty" jsonschema:"Maximum number of changes to retrieve"`
}
```

For output, we typically use `types.MapOutput` which is `map[string]interface{}`.

### Step 2: Implement Client Method

Next, implement the client method in [internal/client/bitbucket/pullrequests.go](file:///home/stephen/workspace/atlassian-dc-mcp-go/internal/client/bitbucket/pullrequests.go):

```go
// GetPullRequestChanges retrieves changes for a specific pull request.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch changes
// for a specific pull request with optional filtering.
//
// Parameters:
//   - input: GetPullRequestChangesInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The changes data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetPullRequestChanges(input GetPullRequestChangesInput) (types.MapOutput, error) {
	// Validate required parameters
	if input.ProjectKey == "" {
		return nil, fmt.Errorf("project key is required")
	}
	if input.RepoSlug == "" {
		return nil, fmt.Errorf("repository slug is required")
	}
	if input.PullRequestID <= 0 {
		return nil, fmt.Errorf("pull request ID must be a positive integer")
	}

	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "start", input.Start, 0)
	utils.SetQueryParam(queryParams, "limit", input.Limit, 0)
	utils.SetQueryParam(queryParams, "sinceId", input.SinceId, "")
	utils.SetQueryParam(queryParams, "untilId", input.UntilId, "")
	utils.SetQueryParam(queryParams, "changeScope", input.ChangeScope, "")
	if input.WithComments {
		queryParams.Set("withComments", "true")
	}

	var changes types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", strconv.Itoa(input.PullRequestID), "changes"},
		queryParams,
		nil,
		&changes,
	); err != nil {
		return nil, fmt.Errorf("failed to get pull request changes: %w", err)
	}

	return changes, nil
}
```

### Step 3: Implement MCP Handler

In [internal/mcp/tools/bitbucket/pullrequests.go](file:///home/stephen/workspace/atlassian-dc-mcp-go/internal/mcp/tools/bitbucket/pullrequests.go), add the handler method:

```go
// getPullRequestChangesHandler handles getting changes in a pull request
func (h *Handler) getPullRequestChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetPullRequestChangesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	changes, err := h.client.GetPullRequestChanges(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get pull request changes failed: %w", err)
	}

	return nil, changes, nil
}
```

### Step 4: Register the Tool

Finally, register the tool in the [AddPullRequestTools](file:///home/stephen/workspace/atlassian-dc-mcp-go/internal/mcp/tools/bitbucket/pullrequests.go#L204-L235) function in [internal/mcp/tools/bitbucket/pullrequests.go](file:///home/stephen/workspace/atlassian-dc-mcp-go/internal/mcp/tools/bitbucket/pullrequests.go):

```go
utils.RegisterTool[bitbucket.GetPullRequestChangesInput, types.MapOutput](server, "bitbucket_get_pull_request_changes", "Get changes for a specific pull request", handler.getPullRequestChangesHandler)
```

## Using AI Coding Assistants for Development

AI coding assistants can help with various development tasks:

1. **Code Generation**: You can ask AI coding assistants to generate code based on specifications or examples.
2. **Code Explanation**: If you're not sure what a piece of code does, ask AI coding assistants to explain it.
3. **Bug Finding**: When you encounter issues, AI coding assistants can help identify potential bugs.
4. **Refactoring Suggestions**: AI coding assistants can suggest improvements to make code more efficient or readable.
5. **Documentation**: AI coding assistants can help write documentation for your code.

To use AI coding assistants effectively:
1. Provide clear, specific instructions about what you want to implement
2. Share relevant code snippets or documentation
3. Ask for explanations if generated code is unclear
4. Verify that generated code follows project conventions

### AI Coding Assistant Prompt Template

When asking AI coding assistants to implement a new Bitbucket API endpoint, use the following prompt template:

```
Implement a new Bitbucket API endpoint for [endpoint name] following the project conventions:

1. Input struct (in internal/client/bitbucket/pullrequests_types.go):
- Use CommonInput and PaginationInput when appropriate
- Add json tags with snake_case naming
- Add jsonschema tags with descriptions
- Mark required fields with "required," in the jsonschema tag

2. Client method (in internal/client/bitbucket/pullrequests.go):
- Validate required parameters
- Use queryParams for GET parameters
- Use utils.SetQueryParam for optional parameters
- Use c.executeRequest for the HTTP call
- Return types.MapOutput and error
- Add comprehensive Go documentation

3. MCP Handler (in internal/mcp/tools/bitbucket/pullrequests.go):
- Follow the pattern of existing handlers
- Return (*mcp.CallToolResult, types.MapOutput, error)
- Use fmt.Errorf with error wrapping (%w)

4. Tool Registration (in internal/mcp/tools/bitbucket/pullrequests.go):
- Use utils.RegisterTool in AddPullRequestTools function
- Follow naming convention: "bitbucket_[action]_[object]"
- Provide a clear, concise description

API documentation: [link to API documentation]
```

For example, to implement the "Get Pull Request Changes" endpoint, you would use:

```
Implement a new Bitbucket API endpoint for getting pull request changes following the project conventions:

1. Input struct (in internal/client/bitbucket/pullrequests_types.go):
- Use CommonInput and PaginationInput when appropriate
- Add json tags with snake_case naming
- Add jsonschema tags with descriptions
- Mark required fields with "required," in the jsonschema tag
- Fields needed: PullRequestID (required), SinceId, ChangeScope, UntilId, WithComments, Start, Limit

2. Client method (in internal/client/bitbucket/pullrequests.go):
- Validate required parameters (ProjectKey, RepoSlug, PullRequestID)
- Use queryParams for GET parameters
- Use utils.SetQueryParam for optional parameters
- Use c.executeRequest for the HTTP call
- Return types.MapOutput and error
- Add comprehensive Go documentation
- Path: /rest/api/latest/projects/{projectKey}/repos/{repoSlug}/pull-requests/{pullRequestId}/changes

3. MCP Handler (in internal/mcp/tools/bitbucket/pullrequests.go):
- Follow the pattern of existing handlers
- Return (*mcp.CallToolResult, types.MapOutput, error)
- Use fmt.Errorf with error wrapping (%w)

4. Tool Registration (in internal/mcp/tools/bitbucket/pullrequests.go):
- Use utils.RegisterTool in AddPullRequestTools function
- Follow naming convention: "bitbucket_[action]_[object]"
- Provide a clear, concise description

API documentation: https://developer.atlassian.com/server/bitbucket/rest/v1000/api-group-pull-requests/#api-api-latest-projects-projectkey-repos-repositoryslug-pull-requests-pullrequestid-changes-get
```

## Testing

All new functionality should include appropriate tests:

1. For client methods, add tests in the corresponding `_test.go` files
2. For MCP handlers, ensure they properly call client methods and handle errors
3. Test edge cases and error conditions

## Best Practices

1. **Follow existing patterns**: Look at how similar functionality is implemented in the codebase
2. **Use proper error handling**: Always wrap errors with context using `fmt.Errorf("message: %w", err)`
3. **Validate inputs**: Check that required parameters are provided
4. **Use jsonschema tags**: These are used to generate documentation for the tools
5. **Write clear comments**: Document public functions and complex logic
6. **Keep functions focused**: Each function should have a single responsibility