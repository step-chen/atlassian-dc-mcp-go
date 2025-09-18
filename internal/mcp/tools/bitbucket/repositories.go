package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetRepositoryOutput represents the output for getting a repository
type GetRepositoryOutput = map[string]interface{}

// getRepositoryHandler handles getting a repository
func (h *Handler) getRepositoryHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoryInput) (*mcp.CallToolResult, GetRepositoryOutput, error) {
	repo, err := h.client.GetRepository(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get repository")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(repo)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create repository result")
		return result, nil, err
	}

	return result, repo, nil
}

// GetRepositoriesOutput represents the output for getting repositories
type GetRepositoriesOutput = map[string]interface{}

// getRepositoriesHandler handles getting repositories
func (h *Handler) getRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoriesInput) (*mcp.CallToolResult, GetRepositoriesOutput, error) {
	repos, err := h.client.GetRepositories(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get repositories")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(repos)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create repositories result")
		return result, nil, err
	}

	return result, repos, nil
}

// GetProjectRepositoriesOutput represents the output for getting project repositories
type GetProjectRepositoriesOutput = map[string]interface{}

// getProjectRepositoriesHandler handles getting project repositories
func (h *Handler) getProjectRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectRepositoriesInput) (*mcp.CallToolResult, GetProjectRepositoriesOutput, error) {
	repos, err := h.client.GetProjectRepositories(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get project repositories")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(repos)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create project repositories result")
		return result, nil, err
	}

	return result, repos, nil
}

// GetRepositoryLabelsOutput represents the output for getting repository labels
type GetRepositoryLabelsOutput = map[string]interface{}

// getRepositoryLabelsHandler handles getting repository labels
func (h *Handler) getRepositoryLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoryLabelsInput) (*mcp.CallToolResult, GetRepositoryLabelsOutput, error) {
	labels, err := h.client.GetRepositoryLabels(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get repository labels")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(labels)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create repository labels result")
		return result, nil, err
	}

	return result, GetRepositoryLabelsOutput{"labels": labels}, nil
}

// GetFileContentOutput represents the output for getting file content
type GetFileContentOutput struct {
	Content []byte `json:"content" jsonschema:"the file content"`
}

// getFileContentHandler handles getting file content
func (h *Handler) getFileContentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetFileContentInput) (*mcp.CallToolResult, GetFileContentOutput, error) {
	content, err := h.client.GetFileContent(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get file content")
		return result, GetFileContentOutput{}, err
	}

	result, err := tools.CreateToolResult(content)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create file content result")
		return result, GetFileContentOutput{}, err
	}

	return result, GetFileContentOutput{Content: content}, nil
}

// GetFilesOutput represents the output for getting files
type GetFilesOutput = map[string]interface{}

// getFilesHandler handles getting files
func (h *Handler) getFilesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetFilesInput) (*mcp.CallToolResult, GetFilesOutput, error) {
	files, err := h.client.GetFiles(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get files")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(files)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create files result")
		return result, nil, err
	}

	return result, files, nil
}

// GetChangesOutput represents the output for getting changes
type GetChangesOutput = map[string]interface{}

// getChangesHandler handles getting changes
func (h *Handler) getChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetChangesInput) (*mcp.CallToolResult, GetChangesOutput, error) {
	changes, err := h.client.GetChanges(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get changes")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(changes)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create changes result")
		return result, nil, err
	}

	return result, changes, nil
}

// CompareChangesOutput represents the output for comparing changes
type CompareChangesOutput = map[string]interface{}

// compareChangesHandler handles comparing changes
func (h *Handler) compareChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.CompareChangesInput) (*mcp.CallToolResult, CompareChangesOutput, error) {
	changes, err := h.client.CompareChanges(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "compare changes")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(changes)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create compare changes result")
		return result, nil, err
	}

	return result, changes, nil
}

// GetForksOutput represents the output for getting forks
type GetForksOutput = map[string]interface{}

// getForksHandler handles getting forks
func (h *Handler) getForksHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetForksInput) (*mcp.CallToolResult, GetForksOutput, error) {
	forks, err := h.client.GetForks(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get forks")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(forks)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create forks result")
		return result, nil, err
	}

	return result, forks, nil
}

// GetReadmeOutput represents the output for getting readme
type GetReadmeOutput struct {
	Content []byte `json:"content" jsonschema:"the readme content"`
}

// getReadmeHandler handles getting readme
func (h *Handler) getReadmeHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetReadmeInput) (*mcp.CallToolResult, GetReadmeOutput, error) {
	readme, err := h.client.GetReadme(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get readme")
		return result, GetReadmeOutput{}, err
	}

	result, err := tools.CreateToolResult(readme)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create readme result")
		return result, GetReadmeOutput{}, err
	}

	return result, GetReadmeOutput{Content: []byte(readme["raw"].(string))}, nil
}

// GetRelatedRepositoriesOutput represents the output for getting related repositories
type GetRelatedRepositoriesOutput = map[string]interface{}

// getRelatedRepositoriesHandler handles getting related repositories
func (h *Handler) getRelatedRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRelatedRepositoriesInput) (*mcp.CallToolResult, GetRelatedRepositoriesOutput, error) {
	repos, err := h.client.GetRelatedRepositories(input)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get related repositories")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(repos)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create related repositories result")
		return result, nil, err
	}

	return result, repos, nil
}

// AddRepositoryTools registers the repository-related tools with the MCP server
func AddRepositoryTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[bitbucket.GetRepositoryInput, GetRepositoryOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repository",
		Description: "Get a specific repository by project key and repository slug",
	}, handler.getRepositoryHandler)

	mcp.AddTool[bitbucket.GetRepositoriesInput, GetRepositoriesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repositories",
		Description: "Get repositories with various filters",
	}, handler.getRepositoriesHandler)

	mcp.AddTool[bitbucket.GetProjectRepositoriesInput, GetProjectRepositoriesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_project_repositories",
		Description: "Get repositories for a specific project",
	}, handler.getProjectRepositoriesHandler)

	mcp.AddTool[bitbucket.GetRepositoryLabelsInput, GetRepositoryLabelsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repository_labels",
		Description: "Get labels for a specific repository",
	}, handler.getRepositoryLabelsHandler)

	mcp.AddTool[bitbucket.GetFileContentInput, GetFileContentOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_file_content",
		Description: "Get file content from a repository",
	}, handler.getFileContentHandler)

	mcp.AddTool[bitbucket.GetFilesInput, GetFilesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_files",
		Description: "Get files from a repository",
	}, handler.getFilesHandler)

	mcp.AddTool[bitbucket.GetChangesInput, GetChangesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_changes",
		Description: "Get changes for a repository",
	}, handler.getChangesHandler)

	mcp.AddTool[bitbucket.CompareChangesInput, CompareChangesOutput](server, &mcp.Tool{
		Name:        "bitbucket_compare_changes",
		Description: "Compare changes between commits",
	}, handler.compareChangesHandler)

	mcp.AddTool[bitbucket.GetForksInput, GetForksOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_forks",
		Description: "Get forks of a repository",
	}, handler.getForksHandler)

	mcp.AddTool[bitbucket.GetReadmeInput, GetReadmeOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_readme",
		Description: "Get README file of a repository",
	}, handler.getReadmeHandler)

	mcp.AddTool[bitbucket.GetRelatedRepositoriesInput, GetRelatedRepositoriesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_related_repositories",
		Description: "Get related repositories",
	}, handler.getRelatedRepositoriesHandler)
}