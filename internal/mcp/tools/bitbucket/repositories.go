package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/utils"

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

	utils.RegisterTool[bitbucket.GetRepositoryInput, MapOutput](server, "bitbucket_get_repository", "Get a repository", handler.getRepositoryHandler)
	utils.RegisterTool[bitbucket.GetRepositoriesInput, MapOutput](server, "bitbucket_get_repositories", "Get a list of repositories", handler.getRepositoriesHandler)
	utils.RegisterTool[bitbucket.GetProjectRepositoriesInput, MapOutput](server, "bitbucket_get_project_repositories", "Get repositories in a project", handler.getProjectRepositoriesHandler)
	utils.RegisterTool[bitbucket.GetRepositoryLabelsInput, MapOutput](server, "bitbucket_get_repository_labels", "Get repository labels", handler.getRepositoryLabelsHandler)
	utils.RegisterTool[bitbucket.GetFileContentInput, ContentOutput](server, "bitbucket_get_file_content", "Get file content", handler.getFileContentHandler)
	utils.RegisterTool[bitbucket.GetReadmeInput, MapOutput](server, "bitbucket_get_readme", "Get repository readme", handler.getReadmeHandler)
	utils.RegisterTool[bitbucket.GetFilesInput, MapOutput](server, "bitbucket_get_files", "Get files in a repository path", handler.getFilesHandler)
	utils.RegisterTool[bitbucket.GetChangesInput, MapOutput](server, "bitbucket_get_changes", "Get changes in a repository", handler.getChangesHandler)
	utils.RegisterTool[bitbucket.CompareChangesInput, MapOutput](server, "bitbucket_compare_changes", "Compare changes between commits", handler.compareChangesHandler)
	utils.RegisterTool[bitbucket.GetForksInput, MapOutput](server, "bitbucket_get_forks", "Get forks of a repository", handler.getForksHandler)
	utils.RegisterTool[bitbucket.GetRelatedRepositoriesInput, MapOutput](server, "bitbucket_get_related_repositories", "Get related repositories", handler.getRelatedRepositoriesHandler)
}
