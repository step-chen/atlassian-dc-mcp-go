package bitbucket

// GetProjectsInput represents the input parameters for getting projects
type GetProjectsInput struct {
	PaginationInput
	Name       string `json:"name,omitempty" jsonschema:"Filter projects by name"`
	Permission string `json:"permission,omitempty" jsonschema:"Filter projects by permission"`
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
	PaginationInput
	ProjectKey string `json:"projectKey" jsonschema:"required,The unique key of the project"`
	Markup     string `json:"markup,omitempty" jsonschema:"Markup formatting option"`
}

// GetRepositoryTasksInput represents the input parameters for getting repository tasks
type GetRepositoryTasksInput struct {
	CommonInput
	PaginationInput
	Markup string `json:"markup,omitempty" jsonschema:"Markup formatting option"`
}

// GetProjectAvatarInput represents the input parameters for getting a project avatar
type GetProjectAvatarInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
}

// GetProjectBranchModelInput represents the input parameters for getting a project branch model
type GetProjectBranchModelInput struct {
	CommonInput
}

// GetProjectBranchModelSettingsInput represents the input parameters for getting project branch model settings
type GetProjectBranchModelSettingsInput struct {
	CommonInput
}

// GetProjectBranchRestrictionsInput represents the input parameters for getting project branch restrictions
type GetProjectBranchRestrictionsInput struct {
	CommonInput
	PaginationInput
}

// GetProjectBranchingModelInput represents the input parameters for getting project branching model
type GetProjectBranchingModelInput struct {
	CommonInput
}

// GetProjectCommitsInput represents the input parameters for getting project commits
type GetProjectCommitsInput struct {
	CommonInput
	PaginationInput
	Until  string `json:"until,omitempty" jsonschema:"Filter commits until a specific time"`
	Since  string `json:"since,omitempty" jsonschema:"Filter commits since a specific time"`
	Path   string `json:"path,omitempty" jsonschema:"Path to filter commits by"`
	Merges string `json:"merges,omitempty" jsonschema:"Filter merge commits"`
}

// GetProjectGroupsInput represents the input parameters for getting project groups
type GetProjectGroupsInput struct {
	PaginationInput
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	Filter     string `json:"filter,omitempty" jsonschema:"Filter groups by name"`
}

// GetProjectPermissionsInput represents the input parameters for getting project permissions
type GetProjectPermissionsInput struct {
	PaginationInput
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	Filter     string `json:"filter,omitempty" jsonschema:"Filter permissions by name"`
}

// GetProjectRepoHookInput represents the input parameters for getting a project repository hook
type GetProjectRepoHookInput struct {
	CommonInput
}

// GetProjectRepoHooksInput represents the input parameters for getting project repository hooks
type GetProjectRepoHooksInput struct {
	CommonInput
	PaginationInput
}

// GetProjectRepoSettingsInput represents the input parameters for getting project repository settings
type GetProjectRepoSettingsInput struct {
	CommonInput
}

// GetProjectRepositoryInput represents the input parameters for getting a specific project repository
type GetProjectRepositoryInput struct {
	CommonInput
}

// GetProjectRepositoryAvatarInput represents the input parameters for getting a project repository avatar
type GetProjectRepositoryAvatarInput struct {
	CommonInput
}

// GetProjectRepositoryBranchesInput represents the input parameters for getting project repository branches
type GetProjectRepositoryBranchesInput struct {
	CommonInput
	PaginationInput
	Base    string `json:"base,omitempty" jsonschema:"The base branch to filter branches"`
	Details bool   `json:"details,omitempty" jsonschema:"Include details in response"`
}

// GetProjectRepositoryBranchInput represents the input parameters for getting a specific project repository branch
type GetProjectRepositoryBranchInput struct {
	CommonInput
	BranchName string `json:"branchName" jsonschema:"required,The branch name"`
}

// GetProjectRepositoryCommitsInput represents the input parameters for getting project repository commits
type GetProjectRepositoryCommitsInput struct {
	CommonInput
	Until  string `json:"until,omitempty" jsonschema:"Filter commits until a specific time"`
	Since  string `json:"since,omitempty" jsonschema:"Filter commits since a specific time"`
	Path   string `json:"path,omitempty" jsonschema:"Path to filter commits by"`
	Start  int    `json:"start,omitempty" jsonschema:"The starting index of the returned commits"`
	Limit  int    `json:"limit,omitempty" jsonschema:"The limit of the number of commits to return"`
	Merges string `json:"merges,omitempty" jsonschema:"Filter merge commits"`
}

// GetProjectRepositoryTagsInput represents the input parameters for getting project repository tags
type GetProjectRepositoryTagsInput struct {
	CommonInput
	PaginationInput
	Name string `json:"name,omitempty" jsonschema:"Filter tags by name"`
}

// GetProjectRepositoryTagInput represents the input parameters for getting a specific project repository tag
type GetProjectRepositoryTagInput struct {
	CommonInput
	TagName string `json:"tagName" jsonschema:"required,The tag name"`
}

// GetProjectRepositoryUsersInput represents the input parameters for getting project repository users
type GetProjectRepositoryUsersInput struct {
	CommonInput
	PaginationInput
	Filter string `json:"filter,omitempty" jsonschema:"Filter users by name"`
}

// GetProjectSettingsInput represents the input parameters for getting project settings
type GetProjectSettingsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
}

// GetProjectUsersInput represents the input parameters for getting project users
type GetProjectUsersInput struct {
	PaginationInput
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	Filter     string `json:"filter,omitempty" jsonschema:"Filter users by name"`
}
