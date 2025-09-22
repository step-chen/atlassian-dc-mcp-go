package confluence

// SearchInput represents the input parameters for searching content
type SearchInput struct {
	PaginationInput
	CQL                   string   `json:"cql" jsonschema:"required,The CQL query string"`
	CQLContext            string   `json:"cqlcontext,omitempty" jsonschema:"The context for the CQL query"`
	Excerpt               string   `json:"excerpt,omitempty" jsonschema:"The excerpt format"`
	Expand                []string `json:"expand,omitempty" jsonschema:"Fields to expand in the response"`
	IncludeArchivedSpaces bool     `json:"includeArchivedSpaces,omitempty" jsonschema:"Whether to include archived spaces in the search"`
}
