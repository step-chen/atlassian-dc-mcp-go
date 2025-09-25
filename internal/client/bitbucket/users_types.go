package bitbucket

// GetUsersInput represents the input parameters for getting users
type GetUsersInput struct {
	Filter            string            `json:"filter,omitempty" jsonschema:"Text to filter users by"`
	Permission        string            `json:"permission,omitempty" jsonschema:"Filter users by permission"`
	Group             string            `json:"group,omitempty" jsonschema:"Filter users by group"`
	PermissionFilters map[string]string `json:"permissionFilters,omitempty" jsonschema:"Additional permission filters"`
}

// GetUserInput represents the input parameters for getting a specific user
type GetUserInput struct {
	UserSlug string `json:"userslug" jsonschema:"required,The slug of the user to retrieve"`
}

// GetUserProjectsInput represents the input parameters for getting projects for a user
type GetUserProjectsInput struct {
	PaginationInput
	Username   string `json:"username" jsonschema:"required,The username"`
	Permission string `json:"permission,omitempty" jsonschema:"Filter projects by permission"`
}

// GetUserRepositoriesInput represents the input parameters for getting repositories for a user
type GetUserRepositoriesInput struct {
	PaginationInput
	Username   string `json:"username" jsonschema:"required,The username"`
	ProjectKey string `json:"projectKey,omitempty" jsonschema:"Filter repositories by project key"`
	Permission string `json:"permission,omitempty" jsonschema:"Filter repositories by permission"`
}

// GetUserRepositoryInput represents the input parameters for getting a specific repository for a user
type GetUserRepositoryInput struct {
	CommonInput
	Username string `json:"username" jsonschema:"required,The username"`
}
