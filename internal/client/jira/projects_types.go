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

type Project struct {
	Archived        bool            `json:"archived,omitempty"`
	AssigneeType    string          `json:"assigneeType,omitempty"`
	Description     string          `json:"description,omitempty"`
	Expand          string          `json:"expand,omitempty"`
	ID              string          `json:"id,omitempty"`
	Key             string          `json:"key,omitempty"`
	Name            string          `json:"name,omitempty"`
	Self            string          `json:"self,omitempty"`
	Components      []Component     `json:"components,omitempty"`
	IssueTypes      []IssueType     `json:"issueTypes,omitempty"`
	Lead            User            `json:"lead,omitempty"`
	ProjectCategory ProjectCategory `json:"projectCategory,omitempty"`
	Versions        []Version       `json:"versions,omitempty"`
}

type Component struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
