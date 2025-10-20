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
