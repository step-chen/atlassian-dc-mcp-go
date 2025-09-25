package bitbucket

// GetUsersInput represents the input parameters for getting users
type GetUsersInput struct {
	Filter            string            `json:"filter,omitempty" jsonschema:"Text to filter users by"`
	Permission        string            `json:"permission,omitempty" jsonschema:"Filter users by permission"`
	Group             string            `json:"group,omitempty" jsonschema:"Filter users by group"`
	PermissionFilters map[string]string `json:"permissionFilters,omitempty" jsonschema:"Additional permission filters"`
}

// GetUserInput represents the input parameters for getting a specific user
type GetUserInput struct {
	UserSlug string `json:"userslug" jsonschema:"required,The slug of the user to retrieve"`
}
