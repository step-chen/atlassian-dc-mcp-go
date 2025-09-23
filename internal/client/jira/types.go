package jira

// EmptyInput represents an empty input struct
type EmptyInput struct{}

// PaginationInput represents pagination parameters
type PaginationInput struct {
	StartAt    int `json:"startAt,omitempty" jsonschema:"The index of the first item to return"`
	MaxResults int `json:"maxResults,omitempty" jsonschema:"The maximum number of items to return per page"`
}