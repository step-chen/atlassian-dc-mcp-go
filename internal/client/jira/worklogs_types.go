package jira

// GetWorklogsInput represents the input parameters for getting worklogs
type GetWorklogsInput struct {
	IssueKey   string `json:"issueKey" jsonschema:"required,The key of the issue"`
	WorklogId  string `json:"worklogId,omitempty" jsonschema:"Optional worklog ID to retrieve a specific worklog"`
}