package bitbucket

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetRepositoryInput represents the input parameters for getting a repository
type GetRepositoryInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetRepositoryOutput represents the output for getting a repository
type GetRepositoryOutput = map[string]interface{}

// getRepositoryHandler handles getting a repository
func (h *Handler) getRepositoryHandler(ctx context.Context, req *mcp.CallToolRequest, input GetRepositoryInput) (*mcp.CallToolResult, GetRepositoryOutput, error) {
	repo, err := h.client.GetRepository(input.ProjectKey, input.RepoSlug)
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

// GetRepositoriesInput represents the input parameters for getting repositories
type GetRepositoriesInput struct {
	ProjectKey    string `json:"projectKey" jsonschema:"required,The project key"`
	ProjectName   string `json:"projectName,omitempty" jsonschema:"Filter repositories by project name"`
	Name          string `json:"name,omitempty" jsonschema:"Filter repositories by name"`
	Visibility    string `json:"visibility,omitempty" jsonschema:"Filter repositories by visibility"`
	Permission    string `json:"permission,omitempty" jsonschema:"Filter repositories by permission"`
	State         string `json:"state,omitempty" jsonschema:"Filter repositories by state"`
	Archived      string `json:"archived,omitempty" jsonschema:"Filter archived repositories"`
	Username      string `json:"username,omitempty" jsonschema:"Filter repositories by username"`
	Start         int    `json:"start,omitempty" jsonschema:"The starting index of the returned repositories"`
	Limit         int    `json:"limit,omitempty" jsonschema:"The limit of the number of repositories to return"`
}

// GetRepositoriesOutput represents the output for getting repositories
type GetRepositoriesOutput = map[string]interface{}

// getRepositoriesHandler handles getting repositories
func (h *Handler) getRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input GetRepositoriesInput) (*mcp.CallToolResult, GetRepositoriesOutput, error) {
	repos, err := h.client.GetRepositories(
		input.ProjectName,
		input.ProjectKey,
		input.Name,
		input.Visibility,
		input.Permission,
		input.State,
		input.Archived,
		input.Username,
		input.Start,
		input.Limit,
	)
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

// GetProjectRepositoriesInput represents the input parameters for getting repositories for a specific project
type GetProjectRepositoriesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned repositories"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of repositories to return"`
}

// GetProjectRepositoriesOutput represents the output for getting repositories for a specific project
type GetProjectRepositoriesOutput = map[string]interface{}

// getProjectRepositoriesHandler handles getting repositories for a specific project
func (h *Handler) getProjectRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input GetProjectRepositoriesInput) (*mcp.CallToolResult, GetProjectRepositoriesOutput, error) {
	repos, err := h.client.GetProjectRepositories(input.ProjectKey, input.Start, input.Limit)
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

// GetRepositoryLabelsInput represents the input parameters for getting repository labels
type GetRepositoryLabelsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetRepositoryLabelsOutput represents the output for getting repository labels
type GetRepositoryLabelsOutput = map[string]interface{}

// getRepositoryLabelsHandler handles getting repository labels
func (h *Handler) getRepositoryLabelsHandler(ctx context.Context, req *mcp.CallToolRequest, input GetRepositoryLabelsInput) (*mcp.CallToolResult, GetRepositoryLabelsOutput, error) {
	labels, err := h.client.GetRepositoryLabels(input.ProjectKey, input.RepoSlug)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get repository labels")
		return result, nil, err
	}

	result, err := tools.CreateToolResult(labels)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create repository labels result")
		return result, nil, err
	}

	return result, labels, nil
}

// GetFilesInput represents the input parameters for getting files in a directory
type GetFilesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Path       string `json:"path,omitempty" jsonschema:"The path to the directory"`
	At         string `json:"at,omitempty" jsonschema:"The commit ID or ref to retrieve files at"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned files"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of files to return"`
}

// GetFilesOutput represents the output for getting files in a directory
type GetFilesOutput = map[string]interface{}

// GetChangesInput represents the input parameters for getting changes between commits
type GetChangesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	From       string `json:"from,omitempty" jsonschema:"The commit ID or ref to compare from"`
	To         string `json:"to,omitempty" jsonschema:"The commit ID or ref to compare to"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned changes"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of changes to return"`
}

// GetChangesOutput represents the output for getting changes between commits
type GetChangesOutput = map[string]interface{}

// CompareChangesInput represents the input parameters for comparing changes between commits
type CompareChangesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	From       string `json:"from,omitempty" jsonschema:"The commit ID or ref to compare from"`
	To         string `json:"to,omitempty" jsonschema:"The commit ID or ref to compare to"`
	FromRepo   string `json:"fromRepo,omitempty" jsonschema:"The source repository"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned changes"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of changes to return"`
}

// CompareChangesOutput represents the output for comparing changes between commits
type CompareChangesOutput = map[string]interface{}

// GetForksInput represents the input parameters for getting forks of a repository
type GetForksInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned forks"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of forks to return"`
}

// GetForksOutput represents the output for getting forks of a repository
type GetForksOutput = map[string]interface{}

// GetReadmeInput represents the input parameters for getting the README file of a repository
type GetReadmeInput struct {
	ProjectKey       string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug         string `json:"repoSlug" jsonschema:"required,The repository slug"`
	At               string `json:"at,omitempty" jsonschema:"The commit ID or ref to retrieve the README at"`
	Markup           string `json:"markup,omitempty" jsonschema:"Markup format for the README"`
	HtmlEscape       string `json:"htmlEscape,omitempty" jsonschema:"HTML escape option"`
	IncludeHeadingId string `json:"includeHeadingId,omitempty" jsonschema:"Include heading IDs"`
	Hardwrap         string `json:"hardwrap,omitempty" jsonschema:"Hard wrap option"`
}

// GetReadmeOutput represents the output for getting the README file of a repository
type GetReadmeOutput struct {
	Readme string `json:"readme" jsonschema:"The README content"`
}

// GetRelatedRepositoriesInput represents the input parameters for getting related repositories
type GetRelatedRepositoriesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned repositories"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of repositories to return"`
}

// GetRelatedRepositoriesOutput represents the output for getting related repositories
type GetRelatedRepositoriesOutput = map[string]interface{}

// GetFileContentInput represents the input parameters for getting file content
type GetFileContentInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Path       string `json:"path" jsonschema:"required,The path to the file"`
	At         string `json:"at,omitempty" jsonschema:"The commit ID or ref to retrieve the file at"`
	Size       bool   `json:"size,omitempty" jsonschema:"Include file size information"`
	Type       bool   `json:"type,omitempty" jsonschema:"Include file type information"`
	Blame      bool   `json:"blame,omitempty" jsonschema:"Include blame information"`
	NoContent  bool   `json:"noContent,omitempty" jsonschema:"Skip content retrieval"`
}

// GetFileContentOutput represents the output for getting file content
type GetFileContentOutput struct {
	Content string `json:"content" jsonschema:"The file content"`
}

// getFileContentHandler handles getting file content
func (h *Handler) getFileContentHandler(ctx context.Context, req *mcp.CallToolRequest, input GetFileContentInput) (*mcp.CallToolResult, GetFileContentOutput, error) {
	content, err := h.client.GetFileContent(
		input.ProjectKey,
		input.RepoSlug,
		input.Path,
		input.At,
		&input.Size,
		&input.Type,
		&input.Blame,
		&input.NoContent,
	)
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

// getFilesHandler handles getting files in a directory
func (h *Handler) getFilesHandler(ctx context.Context, req *mcp.CallToolRequest, input GetFilesInput) (*mcp.CallToolResult, GetFilesOutput, error) {
	files, err := h.client.GetFiles(input.ProjectKey, input.RepoSlug, input.Path, input.At, input.Start, input.Limit)
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

// getChangesHandler handles getting changes between commits
func (h *Handler) getChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input GetChangesInput) (*mcp.CallToolResult, GetChangesOutput, error) {
	changes, err := h.client.GetChanges(input.ProjectKey, input.RepoSlug, input.From, input.To, input.Start, input.Limit)
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

// compareChangesHandler handles comparing changes between commits
func (h *Handler) compareChangesHandler(ctx context.Context, req *mcp.CallToolRequest, input CompareChangesInput) (*mcp.CallToolResult, CompareChangesOutput, error) {
	changes, err := h.client.CompareChanges(input.ProjectKey, input.RepoSlug, input.From, input.To, input.FromRepo, input.Start, input.Limit)
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

// getForksHandler handles getting forks of a repository
func (h *Handler) getForksHandler(ctx context.Context, req *mcp.CallToolRequest, input GetForksInput) (*mcp.CallToolResult, GetForksOutput, error) {
	forks, err := h.client.GetForks(input.ProjectKey, input.RepoSlug, input.Start, input.Limit)
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

// getReadmeHandler handles getting the README file of a repository
func (h *Handler) getReadmeHandler(ctx context.Context, req *mcp.CallToolRequest, input GetReadmeInput) (*mcp.CallToolResult, GetReadmeOutput, error) {
	readme, err := h.client.GetReadme(
		input.ProjectKey,
		input.RepoSlug,
		&input.At,
		&input.Markup,
		&input.HtmlEscape,
		&input.IncludeHeadingId,
		&input.Hardwrap,
	)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "get readme")
		return result, GetReadmeOutput{}, err
	}

	result, err := tools.CreateToolResult(readme)
	if err != nil {
		result, _, err := tools.HandleToolError(err, "create readme result")
		return result, GetReadmeOutput{}, err
	}

	return result, GetReadmeOutput{Readme: readme}, nil
}

// getRelatedRepositoriesHandler handles getting related repositories
func (h *Handler) getRelatedRepositoriesHandler(ctx context.Context, req *mcp.CallToolRequest, input GetRelatedRepositoriesInput) (*mcp.CallToolResult, GetRelatedRepositoriesOutput, error) {
	repos, err := h.client.GetRelatedRepositories(input.ProjectKey, input.RepoSlug, input.Start, input.Limit)
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

	mcp.AddTool[GetRepositoryInput, GetRepositoryOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repository",
		Description: "Get a specific repository",
	}, handler.getRepositoryHandler)

	mcp.AddTool[GetRepositoriesInput, GetRepositoriesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repositories",
		Description: "Get a list of repositories",
	}, handler.getRepositoriesHandler)

	mcp.AddTool[GetProjectRepositoriesInput, GetProjectRepositoriesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_project_repositories",
		Description: "Get repositories for a specific project",
	}, handler.getProjectRepositoriesHandler)

	mcp.AddTool[GetRepositoryLabelsInput, GetRepositoryLabelsOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_repository_labels",
		Description: "Get labels for a specific repository",
	}, handler.getRepositoryLabelsHandler)

	mcp.AddTool[GetFileContentInput, GetFileContentOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_file_content",
		Description: "Get the content of a file in a repository",
	}, handler.getFileContentHandler)

	mcp.AddTool[GetFilesInput, GetFilesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_files",
		Description: "Get files in a directory of a repository",
	}, handler.getFilesHandler)

	mcp.AddTool[GetChangesInput, GetChangesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_changes",
		Description: "Get changes between commits in a repository",
	}, handler.getChangesHandler)

	mcp.AddTool[CompareChangesInput, CompareChangesOutput](server, &mcp.Tool{
		Name:        "bitbucket_compare_changes",
		Description: "Compare changes between two commits in a repository",
	}, handler.compareChangesHandler)

	mcp.AddTool[GetForksInput, GetForksOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_forks",
		Description: "Get forks of a specific repository",
	}, handler.getForksHandler)

	mcp.AddTool[GetReadmeInput, GetReadmeOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_readme",
		Description: "Get the README file of a repository",
	}, handler.getReadmeHandler)

	mcp.AddTool[GetRelatedRepositoriesInput, GetRelatedRepositoriesOutput](server, &mcp.Tool{
		Name:        "bitbucket_get_related_repositories",
		Description: "Get repositories related to a repository",
	}, handler.getRelatedRepositoriesHandler)
}