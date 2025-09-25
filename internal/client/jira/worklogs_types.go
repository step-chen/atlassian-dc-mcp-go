package jira

// GetWorklogsInput represents the input parameters for getting worklogs
type GetWorklogsInput struct {
	IssueKey   string `json:"issueKey" jsonschema:"required,The key of the issue"`
	WorklogId  string `json:"worklogId,omitempty" jsonschema:"Optional worklog ID to retrieve a specific worklog"`
}

// AddWorklogInput represents the input parameters for adding a worklog
type AddWorklogInput struct {
	IssueKey       string `json:"issueKey" jsonschema:"required,The key of the issue"`
	TimeSpent      string `json:"timeSpent" jsonschema:"required,The time spent on the worklog"`
	Comment        string `json:"comment,omitempty" jsonschema:"Comment for the worklog"`
	Started        string `json:"started,omitempty" jsonschema:"The date and time the worklog started"`
	NewEstimate    string `json:"newEstimate,omitempty" jsonschema:"New estimate for the issue"`
	AdjustEstimate string `json:"adjustEstimate,omitempty" jsonschema:"How to adjust the estimate (new, leave, manual, auto)"`
	ReduceBy       string `json:"reduceBy,omitempty" jsonschema:"Amount to reduce the estimate by"`
}