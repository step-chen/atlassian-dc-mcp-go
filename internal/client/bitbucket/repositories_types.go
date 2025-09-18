package bitbucket

// GetRepositoriesInput represents the input parameters for getting repositories
type GetRepositoriesInput struct {
	ProjectName string `json:"projectName,omitempty" jsonschema:"Filter repositories by project name"`
	ProjectKey  string `json:"projectKey,omitempty" jsonschema:"Filter repositories by project key"`
	Name        string `json:"name,omitempty" jsonschema:"Filter repositories by name"`
	Visibility  string `json:"visibility,omitempty" jsonschema:"Filter repositories by visibility"`
	Permission  string `json:"permission,omitempty" jsonschema:"Filter repositories by permission"`
	State       string `json:"state,omitempty" jsonschema:"Filter repositories by state"`
	Archived    string `json:"archived,omitempty" jsonschema:"Filter archived repositories"`
	Start       int    `json:"start,omitempty" jsonschema:"The starting index of the returned repositories"`
	Limit       int    `json:"limit,omitempty" jsonschema:"The limit of the number of repositories to return"`
}

// GetRepositoryInput represents the input parameters for getting a specific repository
type GetRepositoryInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetRepositoryAvatarInput represents the input parameters for getting a repository avatar
type GetRepositoryAvatarInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetRepositoryBranchesInput represents the input parameters for getting repository branches
type GetRepositoryBranchesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Base       string `json:"base,omitempty" jsonschema:"The base branch to filter branches"`
	Details    bool   `json:"details,omitempty" jsonschema:"Include details in response"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned branches"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of branches to return"`
}

// GetRepositoryBranchInput represents the input parameters for getting a specific repository branch
type GetRepositoryBranchInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Name       string `json:"name" jsonschema:"required,The branch name"`
}

// GetRepositoryCommitsInput represents the input parameters for getting repository commits
type GetRepositoryCommitsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Until      string `json:"until,omitempty" jsonschema:"Filter commits until a specific time"`
	Since      string `json:"since,omitempty" jsonschema:"Filter commits since a specific time"`
	Path       string `json:"path,omitempty" jsonschema:"Path to filter commits by"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned commits"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of commits to return"`
	Merges     string `json:"merges,omitempty" jsonschema:"Filter merge commits"`
}

// GetRepositoryTagsInput represents the input parameters for getting repository tags
type GetRepositoryTagsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Name       string `json:"name,omitempty" jsonschema:"Filter tags by name"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned tags"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of tags to return"`
}

// GetRepositoryTagInput represents the input parameters for getting a specific repository tag
type GetRepositoryTagInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Name       string `json:"name" jsonschema:"required,The tag name"`
}

// GetRepositoryUsersInput represents the input parameters for getting repository users
type GetRepositoryUsersInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Filter     string `json:"filter,omitempty" jsonschema:"Filter users by name"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned users"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of users to return"`
}

// GetRepositoryLabelsInput represents the input parameters for getting repository labels
type GetRepositoryLabelsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetFileContentInput represents the input parameters for getting file content
type GetFileContentInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	At         string `json:"at,omitempty" jsonschema:"The commit ID or ref to retrieve the file at"`
	Path       string `json:"path" jsonschema:"required,The path to the file"`
	Size       bool   `json:"size,omitempty" jsonschema:"Include file size information"`
	TypeParam  bool   `json:"typeParam,omitempty" jsonschema:"Include file type information"`
	Blame      bool   `json:"blame,omitempty" jsonschema:"Include blame information"`
	NoContent  bool   `json:"noContent,omitempty" jsonschema:"Skip content retrieval"`
}

// GetFilesInput represents the input parameters for getting files
type GetFilesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Path       string `json:"path" jsonschema:"required,The path to the directory"`
	At         string `json:"at,omitempty" jsonschema:"The commit ID or ref to retrieve files at"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
}

// GetChangesInput represents the input parameters for getting changes
type GetChangesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Until      string `json:"until,omitempty" jsonschema:"The commit ID or ref to compare until"`
	Since      string `json:"since,omitempty" jsonschema:"The commit ID or ref to compare since"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
}

// CompareChangesInput represents the input parameters for comparing changes
type CompareChangesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	From       string `json:"from,omitempty" jsonschema:"The source commit ID or ref"`
	To         string `json:"to,omitempty" jsonschema:"The target commit ID or ref"`
	FromRepo   string `json:"fromRepo,omitempty" jsonschema:"The source repository"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
}

// GetForksInput represents the input parameters for getting forks
type GetForksInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
}

// GetReadmeInput represents the input parameters for getting readme
type GetReadmeInput struct {
	ProjectKey       string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug         string `json:"repoSlug" jsonschema:"required,The repository slug"`
	At               string `json:"at,omitempty" jsonschema:"The commit ID or ref to retrieve the README at"`
	Markup           string `json:"markup,omitempty" jsonschema:"Markup format for the response"`
	HtmlEscape       string `json:"htmlEscape,omitempty" jsonschema:"HTML escape option"`
	IncludeHeadingId string `json:"includeHeadingId,omitempty" jsonschema:"Include heading IDs"`
	Hardwrap         string `json:"hardwrap,omitempty" jsonschema:"Hard wrap option"`
}

// GetRelatedRepositoriesInput represents the input parameters for getting related repositories
type GetRelatedRepositoriesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
}

// GetProjectRepositoriesInput represents the input parameters for getting project repositories
type GetProjectRepositoriesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
}
