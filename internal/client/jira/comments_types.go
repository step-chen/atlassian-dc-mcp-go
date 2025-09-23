package jira

// GetCommentsInput represents the input parameters for getting comments
type GetCommentsInput struct {
	PaginationInput
	IssueKey string `json:"issueKey" jsonschema:"required,The key of the issue"`
	Expand   string `json:"expand,omitempty" jsonschema:"Use expand to include additional information about comments"`
	OrderBy  string `json:"orderBy,omitempty" jsonschema:"Ordering of comments by creation date"`
}

// AddCommentInput represents the input parameters for adding a comment
type AddCommentInput struct {
	IssueKey string `json:"issueKey" jsonschema:"required,The key of the issue"`
	Comment  string `json:"comment" jsonschema:"required,The comment text to add"`
}
