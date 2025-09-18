package bitbucket

// GetTagsInput represents the input parameters for getting tags
type GetTagsInput struct {
	CommonInput
	PaginationInput
	FilterText string `json:"filterText,omitempty" jsonschema:"Text to filter tags by"`
	OrderBy    string `json:"orderBy,omitempty" jsonschema:"Field to order tags by"`
}

// GetTagInput represents the input parameters for getting a specific tag
type GetTagInput struct {
	CommonInput
	Name string `json:"name" jsonschema:"required,The name of the tag to retrieve"`
}
