package jira

// GetProjectInput represents the input parameters for getting a project
type GetProjectInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The key of the project to retrieve"`
}

// GetAllProjectsInput represents the input parameters for getting all projects
type GetAllProjectsInput struct {
	Expand          string `json:"expand,omitempty" jsonschema:"Parameters to expand in the response"`
	Recent          int    `json:"recent,omitempty" jsonschema:"The number of recent projects to return"`
	IncludeArchived bool   `json:"includeArchived,omitempty" jsonschema:"Whether to include archived projects"`
	BrowseArchive   bool   `json:"browseArchive,omitempty" jsonschema:"Whether to include projects in the archive browser"`
}
