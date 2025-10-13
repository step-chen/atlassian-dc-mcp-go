package bitbucket

// CommonInput represents common input parameters for Bitbucket operations
type CommonInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}

// PaginationInput represents pagination parameters
type PaginationInput struct {
	Start int `json:"start,omitempty" jsonschema:"Start number for the page (inclusive). If not passed, first page is assumed"`
	Limit int `json:"limit,omitempty" jsonschema:"Number of items to return. If not passed, a page size of 25 is used"`
}
