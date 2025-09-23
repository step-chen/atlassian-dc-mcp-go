package jira

// SearchIssuesInput represents the input parameters for searching issues
type SearchIssuesInput struct {
	PaginationInput
	JQL            string   `json:"jql,omitempty" jsonschema:"The JQL query string"`
	ProjectKeyOrId string   `json:"projectKeyOrId,omitempty" jsonschema:"The project key or ID to filter by"`
	OrderBy        string   `json:"orderBy,omitempty" jsonschema:"The field to order results by"`
	Statuses       []string `json:"statuses,omitempty" jsonschema:"The statuses to filter by"`
	Fields         []string `json:"fields,omitempty" jsonschema:"The list of fields to return for each issue"`
}
