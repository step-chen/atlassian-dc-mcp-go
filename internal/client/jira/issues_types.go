package jira

import "atlassian-dc-mcp-go/internal/types"

// GetIssueInput represents the input parameters for getting an issue
type GetIssueInput struct {
	IssueKey string   `json:"issueKey" jsonschema:"required,The key of the issue to retrieve"`
	Fields   []string `json:"fields,omitempty" jsonschema:"The list of fields to return for the issue"`
}

// CreateIssueInput represents the input parameters for creating an issue
type CreateIssueInput struct {
	ProjectKey  string `json:"projectKey" jsonschema:"required,The key of the project to create the issue in"`
	Summary     string `json:"summary" jsonschema:"required,The summary of the issue"`
	IssueType   string `json:"issueType" jsonschema:"required,The type of the issue"`
	Description string `json:"description,omitempty" jsonschema:"The description of the issue"`
	Priority    string `json:"priority,omitempty" jsonschema:"The priority of the issue"`
}

// CreateIssueWithPayloadInput represents the input parameters for creating an issue with a custom payload
type CreateIssueWithPayloadInput struct {
	Payload       types.MapOutput `json:"payload" jsonschema:"required,The payload containing issue data"`
	UpdateHistory bool            `json:"updateHistory,omitempty" jsonschema:"Whether to update the user's history"`
}

// CreateSubTaskInput represents the input parameters for creating a sub-task
type CreateSubTaskInput struct {
	ParentKeyOrID string `json:"parentKeyOrID" jsonschema:"required,The key or ID of the parent issue"`
	ProjectKey    string `json:"projectKey" jsonschema:"required,The key of the project to create the sub-task in"`
	Summary       string `json:"summary" jsonschema:"required,The summary of the sub-task"`
	IssueType     string `json:"issueType" jsonschema:"required,The type of the sub-task"`
	Description   string `json:"description,omitempty" jsonschema:"The description of the sub-task"`
	Priority      string `json:"priority,omitempty" jsonschema:"The priority of the sub-task"`
}

// UpdateIssueInput represents the input parameters for updating an issue
type UpdateIssueInput struct {
	IssueKey string          `json:"issueKey" jsonschema:"required,The key of the issue to update"`
	Updates  types.MapOutput `json:"updates" jsonschema:"required,The fields to update"`
}

// UpdateIssueWithOptionsInput represents the input parameters for updating an issue with additional options
type UpdateIssueWithOptionsInput struct {
	IssueKey string            `json:"issueKey" jsonschema:"required,The key of the issue to update"`
	Updates  types.MapOutput   `json:"updates" jsonschema:"required,The fields to update"`
	Options  map[string]string `json:"options,omitempty" jsonschema:"Additional options for the update"`
}

// GetTransitionsInput represents the input parameters for getting transitions available for an issue
type GetTransitionsInput struct {
	IssueKey string `json:"issueKey" jsonschema:"required,The key of the issue"`
}

// TransitionIssueInput represents the input parameters for transitioning an issue to a new status
type TransitionIssueInput struct {
	IssueKey     string `json:"issueKey" jsonschema:"required,The key of the issue to transition"`
	TransitionID string `json:"transitionID" jsonschema:"required,The ID of the transition to apply"`
}

// GetSubtasksInput represents the input parameters for getting sub-tasks of an issue
type GetSubtasksInput struct {
	IssueKey string `json:"issueKey" jsonschema:"required,The key of the issue"`
}

// GetAgileIssueInput represents the input parameters for getting an agile issue
type GetAgileIssueInput struct {
	IssueIdOrKey  string   `json:"issueIdOrKey" jsonschema:"required,The ID or key of the issue"`
	Expand        string   `json:"expand,omitempty" jsonschema:"Parameters to expand in the response"`
	Fields        []string `json:"fields,omitempty" jsonschema:"The list of fields to return for the issue"`
	UpdateHistory bool     `json:"updateHistory,omitempty" jsonschema:"Whether to update the user's history"`
}

// GetIssueEstimationForBoardInput represents the input parameters for getting the estimation for an issue on a board
type GetIssueEstimationForBoardInput struct {
	IssueIdOrKey string `json:"issueIdOrKey" jsonschema:"required,The ID or key of the issue"`
	BoardId      int64  `json:"boardId" jsonschema:"required,The ID of the board"`
}

// SetIssueEstimationForBoardInput represents the input parameters for setting the estimation for an issue on a board
type SetIssueEstimationForBoardInput struct {
	IssueIdOrKey string `json:"issueIdOrKey" jsonschema:"required,The ID or key of the issue"`
	BoardId      int64  `json:"boardId" jsonschema:"required,The ID of the board"`
	Value        string `json:"value" jsonschema:"required,The estimation value"`
}
