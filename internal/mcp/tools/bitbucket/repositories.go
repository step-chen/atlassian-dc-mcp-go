package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getRepositoryHandler handles getting a repository
func (h *Handler) getRepositoryHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoryInput) (*mcp.CallToolResult, MapOutput, error) {
	repo, err := h.client.GetRepository(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get repository failed: %w", err)
	}

	return nil, repo, nil
}

// getRepositoriesHandler handles getting repositories
func (h *Handler) getRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoriesInput) (*mcp.CallToolResult, MapOutput, error) {
	repos, err := h.client.GetRepositories(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get repositories failed: %w", err)
	}

	return nil, repos, nil
}

// getProjectRepositoriesHandler handles getting project repositories
func (h *Handler) getProjectRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectRepositoriesInput) (*mcp.CallToolResult, MapOutput, error) {
	repos, err := h.client.GetProjectRepositories(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get project repositories failed: %w", err)
	}

	return nil, repos, nil
}

// getRepositoryLabelsHandler handles getting repository labels
func (h *Handler) getRepositoryLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoryLabelsInput) (*mcp.CallToolResult, MapOutput, error) {
	labels, err := h.client.GetRepositoryLabels(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get repository labels failed: %w", err)
	}

	result := map[string]interface{}{"labels": labels}
	return nil, result, nil
}

// getFileContentHandler handles getting file content
func (h *Handler) getFileContentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetFileContentInput) (*mcp.CallToolResult, ContentOutput, error) {
	content, err := h.client.GetFileContent(input)
	if err != nil {
		return nil, ContentOutput{}, fmt.Errorf("get file content failed: %w", err)
	}

	return nil, ContentOutput{Content: content}, nil
}

// getReadmeHandler handles getting readme
func (h *Handler) getReadmeHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetReadmeInput) (*mcp.CallToolResult, MapOutput, error) {
	readme, err := h.client.GetReadme(input)
	if err != nil {
		return nil, MapOutput{}, fmt.Errorf("get readme failed: %w", err)
	}

	return nil, readme, nil
}

// getFilesHandler handles getting files
func (h *Handler) getFilesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetFilesInput) (*mcp.CallToolResult, MapOutput, error) {
	files, err := h.client.GetFiles(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get files failed: %w", err)
	}

	return nil, files, nil
}

// getChangesHandler handles getting changes
func (h *Handler) getChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetChangesInput) (*mcp.CallToolResult, MapOutput, error) {
	changes, err := h.client.GetChanges(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get changes failed: %w", err)
	}

	return nil, changes, nil
}

// compareChangesHandler handles comparing changes
func (h *Handler) compareChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.CompareChangesInput) (*mcp.CallToolResult, MapOutput, error) {
	changes, err := h.client.CompareChanges(input)
	if err != nil {
		return nil, nil, fmt.Errorf("compare changes failed: %w", err)
	}

	return nil, changes, nil
}

// getForksHandler handles getting forks
func (h *Handler) getForksHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetForksInput) (*mcp.CallToolResult, MapOutput, error) {
	forks, err := h.client.GetForks(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get forks failed: %w", err)
	}

	return nil, forks, nil
}

// getRelatedRepositoriesHandler handles getting related repositories
func (h *Handler) getRelatedRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRelatedRepositoriesInput) (*mcp.CallToolResult, MapOutput, error) {
	repos, err := h.client.GetRelatedRepositories(input)
	if err != nil {
		return nil, nil, fmt.Errorf("get related repositories failed: %w", err)
	}

	return nil, repos, nil
}

// AddRepositoryTools registers the repository-related tools with the MCP server
func AddRepositoryTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	mcp.AddTool[bitbucket.GetRepositoryInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repository",
		Description: "Get a repository",
	}, handler.getRepositoryHandler)

	mcp.AddTool[bitbucket.GetRepositoriesInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repositories",
		Description: "Get a list of repositories",
	}, handler.getRepositoriesHandler)

	mcp.AddTool[bitbucket.GetProjectRepositoriesInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_project_repositories",
		Description: "Get repositories in a project",
	}, handler.getProjectRepositoriesHandler)

	mcp.AddTool[bitbucket.GetRepositoryLabelsInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repository_labels",
		Description: "Get repository labels",
	}, handler.getRepositoryLabelsHandler)

	mcp.AddTool[bitbucket.GetFileContentInput, ContentOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_file_content",
		Description: "Get file content",
	}, handler.getFileContentHandler)

	mcp.AddTool[bitbucket.GetReadmeInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_readme",
		Description: "Get repository readme",
	}, handler.getReadmeHandler)

	mcp.AddTool[bitbucket.GetFilesInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_files",
		Description: "Get files in a repository path",
	}, handler.getFilesHandler)

	mcp.AddTool[bitbucket.GetChangesInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_changes",
		Description: "Get changes in a repository",
	}, handler.getChangesHandler)

	mcp.AddTool[bitbucket.CompareChangesInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_compare_changes",
		Description: "Compare changes between commits",
	}, handler.compareChangesHandler)

	mcp.AddTool[bitbucket.GetForksInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_forks",
		Description: "Get forks of a repository",
	}, handler.getForksHandler)

	mcp.AddTool[bitbucket.GetRelatedRepositoriesInput, MapOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_related_repositories",
		Description: "Get related repositories",
	}, handler.getRelatedRepositoriesHandler)
}
