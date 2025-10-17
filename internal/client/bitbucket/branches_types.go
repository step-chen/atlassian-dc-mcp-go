package bitbucket

// GetBranchesInput represents the input parameters for getting branches
type GetBranchesInput struct {
	CommonInput
	PaginationInput
	Base         string `json:"base,omitempty" jsonschema:"The base branch to filter branches"`
	Details      bool   `json:"details,omitempty" jsonschema:"Include details in response"`
	FilterText   string `json:"filterText,omitempty" jsonschema:"Text to filter branches by"`
	OrderBy      string `json:"orderBy,omitempty" jsonschema:"Order of branches"`
	Context      string `json:"context,omitempty" jsonschema:"The context of the branches"`
	BoostMatches bool   `json:"boostMatches,omitempty" jsonschema:"Boost matching branches"`
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

// CreateBranchInput represents the input parameters for creating a branch
type CreateBranchInput struct {
	CommonInput
	Name       string `json:"name" jsonschema:"required,Name of the branch to be created"`
	StartPoint string `json:"startPoint" jsonschema:"required,Commit ID from which the branch is created"`
}