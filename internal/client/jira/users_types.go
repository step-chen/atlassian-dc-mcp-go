package jira

// GetUserByNameInput represents the input parameters for getting a user by name
type GetUserByNameInput struct {
	Username string `json:"username" jsonschema:"required,The username of the user to retrieve"`
}

// GetUserByKeyInput represents the input parameters for getting a user by key
type GetUserByKeyInput struct {
	Key string `json:"key" jsonschema:"required,The key of the user to retrieve"`
}

// SearchUsersInput represents the input parameters for searching users
type SearchUsersInput struct {
	Query string `json:"query" jsonschema:"required,The search query"`
	PaginationInput
}
