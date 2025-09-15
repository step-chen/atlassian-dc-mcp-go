package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getRepositoryHandler handles getting a repository
func (h *Handler) getRepositoryHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get repository", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		return h.client.GetRepository(projectKey, repoSlug)
	})
}

// getRepositoriesHandler handles getting repositories
func (h *Handler) getRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get repositories", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		projectName, _ := tools.GetStringArg(args, "projectName")
		permission, _ := tools.GetStringArg(args, "permission")
		name, _ := tools.GetStringArg(args, "name")
		visibility, _ := tools.GetStringArg(args, "visibility")
		state, _ := tools.GetStringArg(args, "state")
		archived, _ := tools.GetStringArg(args, "archived")
		username, _ := tools.GetStringArg(args, "username")

		return h.client.GetRepositories(projectName, projectKey, name, visibility, permission, state, archived, username, start, limit)
	})
}

// getProjectRepositoriesHandler handles getting repositories for a specific project
func (h *Handler) getProjectRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get project repositories", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetProjectRepositories(projectKey, start, limit)
	})
}

// getRepositoryLabelsHandler handles getting labels for a repository
func (h *Handler) getRepositoryLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get repository labels", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		return h.client.GetRepositoryLabels(projectKey, repoSlug)
	})
}

// getFileContentHandler handles getting the content of a file in a repository
func (h *Handler) getFileContentHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get file content", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		path, ok := tools.GetStringArg(args, "path")
		if !ok {
			return nil, fmt.Errorf("missing or invalid path parameter")
		}

		at, _ := tools.GetStringArg(args, "at")

		var size, typeParam, blame, noContent *bool

		if val, ok := args["size"].(bool); ok {
			size = &val
		}

		if val, ok := args["type"].(bool); ok {
			typeParam = &val
		}

		if val, ok := args["blame"].(bool); ok {
			blame = &val
		}

		if val, ok := args["noContent"].(bool); ok {
			noContent = &val
		}

		return h.client.GetFileContent(projectKey, repoSlug, path, at, size, typeParam, blame, noContent)
	})
}

// getFilesHandler handles getting files in a directory of a repository
func (h *Handler) getFilesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get files", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		path, _ := tools.GetStringArg(args, "path")
		at, _ := tools.GetStringArg(args, "at")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetFiles(projectKey, repoSlug, path, at, start, limit)
	})
}

// getChangesHandler handles getting changes between commits in a repository
func (h *Handler) getChangesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get changes", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		until, _ := tools.GetStringArg(args, "until")
		since, _ := tools.GetStringArg(args, "since")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetChanges(projectKey, repoSlug, until, since, limit, start)
	})
}

// compareChangesHandler handles comparing changes between two commits in a repository
func (h *Handler) compareChangesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("compare changes", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		from, ok := tools.GetStringArg(args, "from")
		if !ok {
			return nil, fmt.Errorf("missing or invalid from parameter")
		}

		to, ok := tools.GetStringArg(args, "to")
		if !ok {
			return nil, fmt.Errorf("missing or invalid to parameter")
		}

		fromRepo, _ := tools.GetStringArg(args, "fromRepo")

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.CompareChanges(projectKey, repoSlug, from, to, fromRepo, limit, start)
	})
}

// getForksHandler handles getting forks of a specific repository
func (h *Handler) getForksHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get forks", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetForks(projectKey, repoSlug, start, limit)
	})
}

// getReadmeHandler handles getting the README file of a repository
func (h *Handler) getReadmeHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get readme", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		// Handle string parameters that can be nil
		var at, markup, htmlEscape, includeHeadingId, hardwrap *string

		if val, ok := tools.GetStringArg(args, "at"); ok {
			at = &val
		}

		if val, ok := tools.GetStringArg(args, "markup"); ok {
			markup = &val
		}

		if val, ok := tools.GetStringArg(args, "htmlEscape"); ok {
			htmlEscape = &val
		}

		if val, ok := tools.GetStringArg(args, "includeHeadingId"); ok {
			includeHeadingId = &val
		}

		if val, ok := tools.GetStringArg(args, "hardwrap"); ok {
			hardwrap = &val
		}

		return h.client.GetReadme(projectKey, repoSlug, at, markup, htmlEscape, includeHeadingId, hardwrap)
	})
}

// getRelatedRepositoriesHandler handles getting repositories related to a repository
func (h *Handler) getRelatedRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, args map[string]interface{}) (*mcp.CallToolResult, map[string]interface{}, error) {
	return tools.HandleToolOperation("get related repositories", func() (interface{}, error) {
		projectKey, ok := tools.GetStringArg(args, "projectKey")
		if !ok {
			return nil, fmt.Errorf("missing or invalid projectKey parameter")
		}

		repoSlug, ok := tools.GetStringArg(args, "repoSlug")
		if !ok {
			return nil, fmt.Errorf("missing or invalid repoSlug parameter")
		}

		start := tools.GetIntArg(args, "start", 0)
		limit := tools.GetIntArg(args, "limit", 10)

		return h.client.GetRelatedRepositories(projectKey, repoSlug, start, limit)
	})
}

// AddRepositoryTools registers the repository-related tools with the MCP server
func AddRepositoryTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_repository",
		Description: "Get a specific repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getRepositoryHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_repositories",
		Description: "Get repositories with various filters",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"projectName": {
					Type:        "string",
					Description: "Filter repositories by project name",
				},
				"name": {
					Type:        "string",
					Description: "Filter repositories by name",
				},
				"visibility": {
					Type:        "string",
					Description: "Filter repositories by visibility",
				},
				"permission": {
					Type:        "string",
					Description: "Filter repositories by permission",
				},
				"state": {
					Type:        "string",
					Description: "Filter repositories by state",
				},
				"archived": {
					Type:        "string",
					Description: "Filter archived repositories",
				},
				"username": {
					Type:        "string",
					Description: "Filter repositories by username",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned repositories",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of repositories to return",
				},
			},
			Required: []string{"projectKey"},
		},
	}, handler.getRepositoriesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_project_repositories",
		Description: "Get repositories for a specific project",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned repositories",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of repositories to return",
				},
			},
			Required: []string{"projectKey"},
		},
	}, handler.getProjectRepositoriesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_repository_labels",
		Description: "Get labels for a specific repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getRepositoryLabelsHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_file_content",
		Description: "Get the content of a file in a repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"path": {
					Type:        "string",
					Description: "The path to the file",
				},
				"at": {
					Type:        "string",
					Description: "The commit ID or ref to retrieve the file at",
				},
				"blame": {
					Type:        "boolean",
					Description: "Include blame information",
				},
				"noContent": {
					Type:        "boolean",
					Description: "Skip content retrieval",
				},
				"size": {
					Type:        "boolean",
					Description: "Include file size information",
				},
				"type": {
					Type:        "boolean",
					Description: "Include file type information",
				},
			},
			Required: []string{"projectKey", "repoSlug", "path"},
		},
	}, handler.getFileContentHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_files",
		Description: "Get files in a directory of a repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"path": {
					Type:        "string",
					Description: "The path to the directory",
				},
				"at": {
					Type:        "string",
					Description: "The commit ID or ref to retrieve files at",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned files",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of files to return",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getFilesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_changes",
		Description: "Get changes between commits in a repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"since": {
					Type:        "string",
					Description: "The commit ID or ref to compare since",
				},
				"until": {
					Type:        "string",
					Description: "The commit ID or ref to compare until",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned changes",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of changes to return",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getChangesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_compare_changes",
		Description: "Compare changes between two commits in a repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"from": {
					Type:        "string",
					Description: "The source commit ID or ref",
				},
				"to": {
					Type:        "string",
					Description: "The target commit ID or ref",
				},
				"fromRepo": {
					Type:        "string",
					Description: "The source repository",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned changes",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of changes to return",
				},
			},
			Required: []string{"projectKey", "repoSlug", "from", "to"},
		},
	}, handler.compareChangesHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_forks",
		Description: "Get forks of a specific repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned forks",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of forks to return",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getForksHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_readme",
		Description: "Get the README file of a repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"at": {
					Type:        "string",
					Description: "The commit ID or ref to retrieve the README at",
				},
				"markup": {
					Type:        "string",
					Description: "Markup format for the response",
				},
				"htmlEscape": {
					Type:        "string",
					Description: "HTML escape option",
				},
				"includeHeadingId": {
					Type:        "string",
					Description: "Include heading IDs",
				},
				"hardwrap": {
					Type:        "string",
					Description: "Hard wrap option",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getReadmeHandler)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bitbucket_get_related_repositories",
		Description: "Get repositories related to a repository",
		InputSchema: &jsonschema.Schema{
			Type: "object",
			Properties: map[string]*jsonschema.Schema{
				"projectKey": {
					Type:        "string",
					Description: "The project key",
				},
				"repoSlug": {
					Type:        "string",
					Description: "The repository slug",
				},
				"start": {
					Type:        "integer",
					Description: "The starting index of the returned repositories",
				},
				"limit": {
					Type:        "integer",
					Description: "The limit of the number of repositories to return",
				},
			},
			Required: []string{"projectKey", "repoSlug"},
		},
	}, handler.getRelatedRepositoriesHandler)
}