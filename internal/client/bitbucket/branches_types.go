package bitbucket

// GetBranchesInput represents the input parameters for getting branches
type GetBranchesInput struct {
	ProjectKey    string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug      string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Base          string `json:"base,omitempty" jsonschema:"The base branch to filter branches"`
	Details       bool   `json:"details,omitempty" jsonschema:"Include details in response"`
	FilterText    string `json:"filterText,omitempty" jsonschema:"Text to filter branches by"`
	OrderBy       string `json:"orderBy,omitempty" jsonschema:"Order of branches"`
	Context       string `json:"context,omitempty" jsonschema:"The context of the branches"`
	BoostMatches  bool   `json:"boostMatches,omitempty" jsonschema:"Boost matching branches"`
	Start         int    `json:"start,omitempty" jsonschema:"The starting index of the returned branches"`
	Limit         int    `json:"limit,omitempty" jsonschema:"The limit of the number of branches to return"`
}

// GetBranchInput represents the input parameters for getting a specific branch
type GetBranchInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitId   string `json:"commitId" jsonschema:"required,The commit ID"`
	Start      int    `json:"start,omitempty" jsonschema:"The starting index of the returned branches"`
	Limit      int    `json:"limit,omitempty" jsonschema:"The limit of the number of branches to return"`
}

// GetDefaultBranchInput represents the input parameters for getting the default branch
type GetDefaultBranchInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
}