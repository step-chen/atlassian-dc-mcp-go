package bitbucket

// GetProjectsInput represents the input parameters for getting projects
type GetProjectsInput struct {
	Name       string `json:"name,omitempty" jsonschema:"Filter projects by name"`
	Permission string `json:"permission,omitempty" jsonschema:"Filter projects by permission"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
}

// GetProjectInput represents the input parameters for getting a specific project
type GetProjectInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The unique key of the project to retrieve"`
}

// GetProjectPrimaryEnhancedEntityLinkInput represents the input parameters for getting project primary enhanced entity link
type GetProjectPrimaryEnhancedEntityLinkInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The unique key of the project"`
}

// GetProjectTasksInput represents the input parameters for getting project tasks
type GetProjectTasksInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The unique key of the project"`
	Markup     string `json:"markup,omitempty" jsonschema:"Markup formatting option"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
}

// GetRepositoryTasksInput represents the input parameters for getting repository tasks
type GetRepositoryTasksInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The unique key of the project"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Markup     string `json:"markup,omitempty" jsonschema:"Markup formatting option"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
}

// GetProjectAvatarInput represents the input parameters for getting a project avatar
type GetProjectAvatarInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
}

// GetProjectBranchModelInput represents the input parameters for getting a project branch model
type GetProjectBranchModelInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetProjectBranchModelSettingsInput represents the input parameters for getting project branch model settings
type GetProjectBranchModelSettingsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetProjectBranchRestrictionsInput represents the input parameters for getting project branch restrictions
type GetProjectBranchRestrictionsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned restrictions"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of restrictions to return"`
}

// GetProjectBranchingModelInput represents the input parameters for getting project branching model
type GetProjectBranchingModelInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetProjectCommitsInput represents the input parameters for getting project commits
type GetProjectCommitsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Until      string `json:"until,omitempty" jsonschema:"Filter commits until a specific time"`
	Since      string `json:"since,omitempty" jsonschema:"Filter commits since a specific time"`
	Path       string `json:"path,omitempty" jsonschema:"Path to filter commits by"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned commits"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of commits to return"`
	Merges     string `json:"merges,omitempty" jsonschema:"Filter merge commits"`
}

// GetProjectGroupsInput represents the input parameters for getting project groups
type GetProjectGroupsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	Filter     string `json:"filter,omitempty" jsonschema:"Filter groups by name"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned groups"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of groups to return"`
}

// GetProjectPermissionsInput represents the input parameters for getting project permissions
type GetProjectPermissionsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	Filter     string `json:"filter,omitempty" jsonschema:"Filter permissions by name"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned permissions"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of permissions to return"`
}

// GetProjectRepoHookInput represents the input parameters for getting a project repository hook
type GetProjectRepoHookInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetProjectRepoHooksInput represents the input parameters for getting project repository hooks
type GetProjectRepoHooksInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned hooks"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of hooks to return"`
}

// GetProjectRepoSettingsInput represents the input parameters for getting project repository settings
type GetProjectRepoSettingsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetProjectRepositoryInput represents the input parameters for getting a specific project repository
type GetProjectRepositoryInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetProjectRepositoryAvatarInput represents the input parameters for getting a project repository avatar
type GetProjectRepositoryAvatarInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// GetProjectRepositoryBranchesInput represents the input parameters for getting project repository branches
type GetProjectRepositoryBranchesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Base       string `json:"base,omitempty" jsonschema:"The base branch to filter branches"`
	Details    bool   `json:"details,omitempty" jsonschema:"Include details in response"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned branches"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of branches to return"`
}

// GetProjectRepositoryBranchInput represents the input parameters for getting a specific project repository branch
type GetProjectRepositoryBranchInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	BranchName string `json:"branchName" jsonschema:"required,The branch name"`
}

// GetProjectRepositoryCommitsInput represents the input parameters for getting project repository commits
type GetProjectRepositoryCommitsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Until      string `json:"until,omitempty" jsonschema:"Filter commits until a specific time"`
	Since      string `json:"since,omitempty" jsonschema:"Filter commits since a specific time"`
	Path       string `json:"path,omitempty" jsonschema:"Path to filter commits by"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned commits"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of commits to return"`
	Merges     string `json:"merges,omitempty" jsonschema:"Filter merge commits"`
}

// GetProjectRepositoryTagsInput represents the input parameters for getting project repository tags
type GetProjectRepositoryTagsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Name       string `json:"name,omitempty" jsonschema:"Filter tags by name"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned tags"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of tags to return"`
}

// GetProjectRepositoryTagInput represents the input parameters for getting a specific project repository tag
type GetProjectRepositoryTagInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	TagName    string `json:"tagName" jsonschema:"required,The tag name"`
}

// GetProjectRepositoryUsersInput represents the input parameters for getting project repository users
type GetProjectRepositoryUsersInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Filter     string `json:"filter,omitempty" jsonschema:"Filter users by name"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned users"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of users to return"`
}

// GetProjectSettingsInput represents the input parameters for getting project settings
type GetProjectSettingsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
}

// GetProjectUsersInput represents the input parameters for getting project users
type GetProjectUsersInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	Filter     string `json:"filter,omitempty" jsonschema:"Filter users by name"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned users"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of users to return"`
}