package bitbucket

// GetTagsInput represents the input parameters for getting tags
type GetTagsInput struct {
	CommonInput
	FilterText string `json:"filterText,omitempty" jsonschema:"Text to filter tags by"`
	OrderBy    string `json:"orderBy,omitempty" jsonschema:"Field to order tags by"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
}

// GetTagInput represents the input parameters for getting a specific tag
type GetTagInput struct {
	CommonInput
	Name string `json:"name" jsonschema:"required,The name of the tag to retrieve"`
}