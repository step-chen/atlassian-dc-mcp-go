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

type StatusCategory struct {
	Name string `json:"name,omitempty"`
}

type Status struct {
	Name           string          `json:"name,omitempty"`
	StatusCategory *StatusCategory `json:"statusCategory,omitempty"`
}

type Priority struct {
	Name string `json:"name,omitempty"`
}

type IssueType struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Subtask     bool   `json:"subtask,omitempty"`
}

type ProjectCategory struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type User struct {
	Active      bool   `json:"active,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type SubtaskType struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Inward  string `json:"inward,omitempty"`
	Outward string `json:"outward,omitempty"`
}

type OutwardIssue struct {
	ID     string        `json:"id,omitempty"`
	Key    string        `json:"key,omitempty"`
	Self   string        `json:"self,omitempty"`
	Fields SubtaskFields `json:"fields,omitempty"`
}

type SubtaskFields struct {
	SubTaskStatus Status `json:"status,omitempty"`
}

type SubtaskStatus struct {
	Name string `json:"name,omitempty"`
}

type Subtask struct {
	ID           string       `json:"id,omitempty"`
	Name         string       `json:"name,omitempty"`
	Type         SubtaskType  `json:"type,omitempty"`
	OutwardIssue OutwardIssue `json:"outwardIssue,omitempty"`
}

type Attachment struct {
	ID       string `json:"id,omitempty"`
	Author   User   `json:"author,omitempty"`
	Filename string `json:"filename,omitempty"`
	Size     int    `json:"size,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
	Content  string `json:"content,omitempty"`
}

type Comments struct {
	Total    int       `json:"total,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

type Comment struct {
	ID           string `json:"id,omitempty"`
	Author       *User  `json:"author,omitempty"`
	UpdateAuthor *User  `json:"updateAuthor,omitempty"`
	Body         string `json:"body,omitempty"`
	Created      string `json:"created,omitempty"`
	Updated      string `json:"updated,omitempty"`
}

type Worklog struct {
	Total    int            `json:"total,omitempty"`
	Worklogs []WorklogEntry `json:"worklogs,omitempty"`
}

type WorklogEntry struct {
	ID               string `json:"id,omitempty"`
	Author           *User  `json:"author,omitempty"`
	UpdateAuthor     *User  `json:"updateAuthor,omitempty"`
	Comment          string `json:"comment,omitempty"`
	TimeSpent        string `json:"timeSpent,omitempty"`
	TimeSpentSeconds int    `json:"timeSpentSeconds,omitempty"`
	Created          string `json:"created,omitempty"`
	Updated          string `json:"updated,omitempty"`
	Started          string `json:"started,omitempty"`
}

type Version struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Archived    bool   `json:"archived,omitempty"`
	Released    bool   `json:"released,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
	Overdue     bool   `json:"overdue,omitempty"`
}

type IssueLinkType struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Inward  string `json:"inward,omitempty"`
	Outward string `json:"outward,omitempty"`
}

type LinkedIssueFields struct {
	Summary   string     `json:"summary,omitempty"`
	Status    *Status    `json:"status,omitempty"`
	Priority  *Priority  `json:"priority,omitempty"`
	IssueType *IssueType `json:"issuetype,omitempty"`
}

type LinkedIssue struct {
	ID     string             `json:"id,omitempty"`
	Key    string             `json:"key,omitempty"`
	Self   string             `json:"self,omitempty"`
	Fields *LinkedIssueFields `json:"fields,omitempty"`
}

type IssueLink struct {
	ID           string         `json:"id,omitempty"`
	Type         *IssueLinkType `json:"type,omitempty"`
	OutwardIssue *LinkedIssue   `json:"outwardIssue,omitempty"`
	InwardIssue  *LinkedIssue   `json:"inwardIssue,omitempty"`
}

type ProgressInfo struct {
	Progress int `json:"progress,omitempty"`
	Total    int `json:"total,omitempty"`
	Percent  int `json:"percent,omitempty"`
}

type TimeTracking struct {
	OriginalEstimate         string `json:"originalEstimate,omitempty"`
	RemainingEstimate        string `json:"remainingEstimate,omitempty"`
	TimeSpent                string `json:"timeSpent,omitempty"`
	OriginalEstimateSeconds  int    `json:"originalEstimateSeconds,omitempty"`
	RemainingEstimateSeconds int    `json:"remainingEstimateSeconds,omitempty"`
	TimeSpentSeconds         int    `json:"timeSpentSeconds,omitempty"`
}

type Fields struct {
	Summary      string       `json:"summary,omitempty"`
	Description  string       `json:"description,omitempty"`
	Status       *Status      `json:"status,omitempty"`
	Priority     *Priority    `json:"priority,omitempty"`
	IssueType    *IssueType   `json:"issuetype,omitempty"`
	Project      *Project     `json:"project,omitempty"`
	Reporter     *User        `json:"reporter,omitempty"`
	Assignee     *User        `json:"assignee,omitempty"`
	Creator      *User        `json:"creator,omitempty"`
	Created      string       `json:"created,omitempty"`
	Duedate      string       `json:"duedate,omitempty"`
	Updated      string       `json:"updated,omitempty"`
	Progress     ProgressInfo `json:"progress,omitempty"`
	Labels       []string     `json:"labels,omitempty"`
	Subtasks     []Subtask    `json:"subtasks,omitempty"`
	Attachments  []Attachment `json:"attachment,omitempty"`
	Comments     *Comments    `json:"comment,omitempty"`
	Worklog      *Worklog     `json:"worklog,omitempty"`
	TimeTracking TimeTracking `json:"timetracking,omitempty"`
	FixVersions  []Version    `json:"fixVersions,omitempty"`
	Versions     []Version    `json:"versions,omitempty"`
	Components   []Component  `json:"components,omitempty"`
	IssueLinks   []IssueLink  `json:"issuelinks,omitempty"`
}

type Issue struct {
	ID     string `json:"id,omitempty"`
	Key    string `json:"key,omitempty"`
	Self   string `json:"self,omitempty"`
	Fields Fields `json:"fields,omitempty"`
}

type Issues struct {
	MaxResults int     `json:"maxResults,omitempty"`
	StartAt    int     `json:"startAt,omitempty"`
	Total      int     `json:"total,omitempty"`
	Issues     []Issue `json:"issues,omitempty"`
}
