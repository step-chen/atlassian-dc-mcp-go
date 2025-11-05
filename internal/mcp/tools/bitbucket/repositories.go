package bitbucket

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// getRepositoryHandler handles getting a repository
func (h *Handler) getRepositoryHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoryInput) (*mcp.CallToolResult, types.MapOutput, error) {
	repo, err := h.client.GetRepository(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get repository failed: %w", err)
	}

	return nil, repo, nil
}

// getRepositoriesHandler handles getting repositories
func (h *Handler) getRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoriesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	repos, err := h.client.GetRepositories(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get repositories failed: %w", err)
	}

	return nil, repos, nil
}

// getProjectRepositoriesHandler handles getting project repositories
func (h *Handler) getProjectRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetProjectRepositoriesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	repos, err := h.client.GetProjectRepositories(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get project repositories failed: %w", err)
	}

	return nil, repos, nil
}

// getRepositoryLabelsHandler handles getting repository labels
func (h *Handler) getRepositoryLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRepositoryLabelsInput) (*mcp.CallToolResult, types.MapOutput, error) {
	labels, err := h.client.GetRepositoryLabels(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get repository labels failed: %w", err)
	}

	result := types.MapOutput{"labels": labels}
	return nil, result, nil
}

// getFileContentHandler handles getting file content
func (h *Handler) getFileContentHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetFileContentInput) (*mcp.CallToolResult, ContentOutput, error) {
	content, err := h.client.GetFileContent(ctx, input)
	if err != nil {
		return nil, ContentOutput{}, fmt.Errorf("get file content failed: %w", err)
	}

	return nil, ContentOutput{Content: string(content)}, nil
}

// getReadmeHandler handles getting readme
func (h *Handler) getReadmeHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetReadmeInput) (*mcp.CallToolResult, types.MapOutput, error) {
	readme, err := h.client.GetReadme(ctx, input)
	if err != nil {
		return nil, types.MapOutput{}, fmt.Errorf("get readme failed: %w", err)
	}

	return nil, readme, nil
}

// getFilesHandler handles getting files
func (h *Handler) getFilesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetFilesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	files, err := h.client.GetFiles(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get files failed: %w", err)
	}

	return nil, files, nil
}

// getChangesHandler handles getting changes
func (h *Handler) getChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetChangesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	changes, err := h.client.GetChanges(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get changes failed: %w", err)
	}

	return nil, changes, nil
}

// compareChangesHandler handles comparing changes
func (h *Handler) compareChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.CompareChangesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	changes, err := h.client.CompareChanges(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("compare changes failed: %w", err)
	}

	return nil, changes, nil
}

// getForksHandler handles getting forks
func (h *Handler) getForksHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetForksInput) (*mcp.CallToolResult, types.MapOutput, error) {
	forks, err := h.client.GetForks(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get forks failed: %w", err)
	}

	return nil, forks, nil
}

// getRelatedRepositoriesHandler handles getting related repositories
func (h *Handler) getRelatedRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.GetRelatedRepositoriesInput) (*mcp.CallToolResult, types.MapOutput, error) {
	repos, err := h.client.GetRelatedRepositories(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("get related repositories failed: %w", err)
	}

	return nil, repos, nil
}

// AddRepositoryTools registers the repository-related tools with the MCP server
func AddRepositoryTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[bitbucket.GetRepositoryInput, types.MapOutput](server, "bitbucket_get_repository", "Get a repository", handler.getRepositoryHandler)
	utils.RegisterTool[bitbucket.GetRepositoriesInput, types.MapOutput](server, "bitbucket_get_repositories", "Get a list of repositories", handler.getRepositoriesHandler)
	utils.RegisterTool[bitbucket.GetProjectRepositoriesInput, types.MapOutput](server, "bitbucket_get_project_repositories", "Get repositories in a project", handler.getProjectRepositoriesHandler)
	utils.RegisterTool[bitbucket.GetRepositoryLabelsInput, types.MapOutput](server, "bitbucket_get_repository_labels", "Get repository labels", handler.getRepositoryLabelsHandler)
	utils.RegisterTool[bitbucket.GetFileContentInput, ContentOutput](server, "bitbucket_get_file_content", "Get file content", handler.getFileContentHandler)
	utils.RegisterTool[bitbucket.GetReadmeInput, types.MapOutput](server, "bitbucket_get_readme", "Get repository readme", handler.getReadmeHandler)
	utils.RegisterTool[bitbucket.GetFilesInput, types.MapOutput](server, "bitbucket_get_files", "Get files in a repository path", handler.getFilesHandler)
	utils.RegisterTool[bitbucket.GetChangesInput, types.MapOutput](server, "bitbucket_get_changes", "Get changes in a repository", handler.getChangesHandler)
	utils.RegisterTool[bitbucket.CompareChangesInput, types.MapOutput](server, "bitbucket_compare_changes", "Compare changes between commits", handler.compareChangesHandler)
	utils.RegisterTool[bitbucket.GetForksInput, types.MapOutput](server, "bitbucket_get_forks", "Get forks of a repository", handler.getForksHandler)
	utils.RegisterTool[bitbucket.GetRelatedRepositoriesInput, types.MapOutput](server, "bitbucket_get_related_repositories", "Get related repositories", handler.getRelatedRepositoriesHandler)
}
