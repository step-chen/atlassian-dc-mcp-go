package bitbucket

// GetRepositoriesInput represents the input parameters for getting repositories
type GetRepositoriesInput struct {
	PaginationInput
	ProjectName string `json:"projectName,omitempty" jsonschema:"Filter repositories by project name"`
	ProjectKey  string `json:"projectKey,omitempty" jsonschema:"Filter repositories by project key"`
	Name        string `json:"name,omitempty" jsonschema:"Filter repositories by name"`
	Visibility  string `json:"visibility,omitempty" jsonschema:"Filter repositories by visibility"`
	Permission  string `json:"permission,omitempty" jsonschema:"Filter repositories by permission"`
	State       string `json:"state,omitempty" jsonschema:"Filter repositories by state"`
	Archived    string `json:"archived,omitempty" jsonschema:"Filter archived repositories"`
}

// GetRepositoryInput represents the input parameters for getting a specific repository
type GetRepositoryInput struct {
	CommonInput
}

// GetRepositoryLabelsInput represents the input parameters for getting repository labels
type GetRepositoryLabelsInput struct {
	CommonInput
}

// GetFileContentInput represents the input parameters for getting file content
type GetFileContentInput struct {
	CommonInput
	Path        string `json:"path" jsonschema:"required,The file path"`
	At          string `json:"at,omitempty" jsonschema:"The commit ID or ref to retrieve the file from"`
	Markup      string `json:"markup,omitempty" jsonschema:"Markup formatting option"`
	HtmlEscape  string `json:"htmlEscape,omitempty" jsonschema:"HTML escape option"`
	IncludeHeadingId string `json:"includeHeadingId,omitempty" jsonschema:"Include heading IDs"`
	Hardwrap    string `json:"hardwrap,omitempty" jsonschema:"Hard wrap option"`
}

// GetFilesInput represents the input parameters for getting files
type GetFilesInput struct {
	CommonInput
	PaginationInput
	Path string `json:"path,omitempty" jsonschema:"The directory path to list files from"`
	At   string `json:"at,omitempty" jsonschema:"The commit ID or ref to retrieve the files from"`
}

// GetChangesInput represents the input parameters for getting changes
type GetChangesInput struct {
	CommonInput
	PaginationInput
	Until string `json:"until,omitempty" jsonschema:"The commit ID or ref to compare until"`
	Since string `json:"since,omitempty" jsonschema:"The commit ID or ref to compare since"`
}

// CompareChangesInput represents the input parameters for comparing changes
type CompareChangesInput struct {
	CommonInput
	PaginationInput
	From     string `json:"from,omitempty" jsonschema:"The source commit ID or ref"`
	To       string `json:"to,omitempty" jsonschema:"The target commit ID or ref"`
	FromRepo string `json:"fromRepo,omitempty" jsonschema:"The source repository"`
}

// GetForksInput represents the input parameters for getting forks
type GetForksInput struct {
	CommonInput
	PaginationInput
}

// GetReadmeInput represents the input parameters for getting readme
type GetReadmeInput struct {
	CommonInput
	At               string `json:"at,omitempty" jsonschema:"The commit ID or ref to retrieve the README at"`
	Markup           string `json:"markup,omitempty" jsonschema:"Markup format for the response"`
	HtmlEscape       string `json:"htmlEscape,omitempty" jsonschema:"HTML escape option"`
	IncludeHeadingId string `json:"includeHeadingId,omitempty" jsonschema:"Include heading IDs"`
	Hardwrap         string `json:"hardwrap,omitempty" jsonschema:"Hard wrap option"`
}

// GetRelatedRepositoriesInput represents the input parameters for getting related repositories
type GetRelatedRepositoriesInput struct {
	CommonInput
	PaginationInput
}

// GetProjectRepositoriesInput represents the input parameters for getting project repositories
type GetProjectRepositoriesInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	PaginationInput
}
