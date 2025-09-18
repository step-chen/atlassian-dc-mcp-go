package bitbucket

// CommonInput represents common input parameters for Bitbucket operations
type CommonInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// PaginationInput represents pagination parameters
type PaginationInput struct {
	Start int `json:"start,omitempty" jsonschema:"The starting index of the returned branches"`
	Limit int `json:"limit,omitempty" jsonschema:"The limit of the number of branches to return"`
}
