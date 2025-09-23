package jira

// GetBoardsInput represents the input parameters for getting boards
type GetBoardsInput struct {
	PaginationInput
	Name           string `json:"name,omitempty" jsonschema:"Filters results to boards that match the specified name"`
	ProjectKeyOrId string `json:"projectKeyOrId,omitempty" jsonschema:"Filters results to boards that match the specified project key or ID"`
	BoardType      string `json:"boardType,omitempty" jsonschema:"Filters results to boards of the specified type"`
}

// GetBoardInput represents the input parameters for getting a board
type GetBoardInput struct {
	Id int `json:"id" jsonschema:"required,The ID of the board to retrieve"`
}

// GetBoardBacklogInput represents the input parameters for getting the backlog of a board
type GetBoardBacklogInput struct {
	PaginationInput
	BoardId       int      `json:"boardId" jsonschema:"required,The ID of the board"`
	JQL           string   `json:"jql,omitempty" jsonschema:"Filters results using a JQL query"`
	ValidateQuery bool     `json:"validateQuery,omitempty" jsonschema:"Specifies whether to validate the JQL query"`
	Fields        []string `json:"fields,omitempty" jsonschema:"The list of fields to return for each issue"`
	Expand        string   `json:"expand,omitempty" jsonschema:"A comma-separated list of parameters to expand"`
}

// GetBoardEpicsInput represents the input parameters for getting the epics associated with a board
type GetBoardEpicsInput struct {
	PaginationInput
	BoardId int    `json:"boardId" jsonschema:"required,The ID of the board"`
	Done    string `json:"done,omitempty" jsonschema:"Filters results to epics that are either done or not done"`
}

// GetBoardSprintsInput represents the input parameters for getting the sprints associated with a board
type GetBoardSprintsInput struct {
	PaginationInput
	BoardId int    `json:"boardId" jsonschema:"required,The ID of the board"`
	State   string `json:"state,omitempty" jsonschema:"Filters results to sprints in the specified states"`
}

// GetSprintInput represents the input parameters for getting a sprint
type GetSprintInput struct {
	SprintId int `json:"sprintId" jsonschema:"required,The ID of the sprint to retrieve"`
}

// GetSprintIssuesInput represents the input parameters for getting issues in a sprint
type GetSprintIssuesInput struct {
	PaginationInput
	SprintId      int      `json:"sprintId" jsonschema:"required,The ID of the sprint"`
	JQL           string   `json:"jql,omitempty" jsonschema:"Filters results using a JQL query"`
	ValidateQuery bool     `json:"validateQuery,omitempty" jsonschema:"Specifies whether to validate the JQL query"`
	Fields        []string `json:"fields,omitempty" jsonschema:"The list of fields to return for each issue"`
	Expand        string   `json:"expand,omitempty" jsonschema:"A comma-separated list of parameters to expand"`
}
