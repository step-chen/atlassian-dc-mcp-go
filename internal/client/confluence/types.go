package confluence

// PaginationInput represents pagination parameters
type PaginationInput struct {
	Start int `json:"start,omitempty" jsonschema:"The starting index of the returned results"`
	Limit int `json:"limit,omitempty" jsonschema:"The limit of the number of results to return"`
}
