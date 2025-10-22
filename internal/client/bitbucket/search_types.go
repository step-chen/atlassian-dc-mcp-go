package bitbucket

// SearchCodeInput represents the input parameters for code search
type SearchCodeInput struct {
	CommonInput
	SearchQuery   string  `json:"search_query" jsonschema:"required,The search query"`
	SearchContext string  `json:"search_context,omitempty" jsonschema:"The search context (assignment, declaration, usage, exact, any)"`
	FilePattern   *string `json:"file_pattern,omitempty" jsonschema:"File pattern to limit search"`
	Limit         int     `json:"limit,omitempty" jsonschema:"Maximum number of results to return"`
	Start         int     `json:"start,omitempty" jsonschema:"Start index for pagination"`
}

// BitbucketServerSearchRequest represents the request structure for Bitbucket Server search
type BitbucketServerSearchRequest struct {
	Query    string                       `json:"query"`
	Entities BitbucketSearchRequestEntities `json:"entities"`
}

// BitbucketSearchRequestEntities represents the entities to search in
type BitbucketSearchRequestEntities struct {
	Code BitbucketCodeEntity `json:"code"`
}

// BitbucketCodeEntity represents the code entity search parameters
type BitbucketCodeEntity struct {
	Start int `json:"start"`
	Limit int `json:"limit"`
}