package bitbucket

// PaginationInput represents pagination parameters
type PaginationInput struct {
	Start int `json:"start,omitempty" jsonschema:"The starting index of the returned branches"`
	Limit int `json:"limit,omitempty" jsonschema:"The limit of the number of branches to return"`
}

// GetBranchesInput represents the input parameters for getting branches
type GetBranchesInput struct {
	CommonInput
	PaginationInput
	Base          string `json:"base,omitempty" jsonschema:"The base branch to filter branches"`
	Details       bool   `json:"details,omitempty" jsonschema:"Include details in response"`
	FilterText    string `json:"filterText,omitempty" jsonschema:"Text to filter branches by"`
	OrderBy       string `json:"orderBy,omitempty" jsonschema:"Order of branches"`
	Context       string `json:"context,omitempty" jsonschema:"The context of the branches"`
	BoostMatches  bool   `json:"boostMatches,omitempty" jsonschema:"Boost matching branches"`
}

// GetBranchInput represents the input parameters for getting a specific branch
type GetBranchInput struct {
	CommonInput
	PaginationInput
	CommitId string `json:"commitId" jsonschema:"required,The commit ID"`
}

// GetDefaultBranchInput represents the input parameters for getting the default branch
type GetDefaultBranchInput struct {
	CommonInput
}