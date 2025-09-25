package bitbucket

// GetPullRequestsInput represents the input parameters for getting pull requests
type GetPullRequestsInput struct {
	CommonInput
	PaginationInput
	State          string `json:"state,omitempty" jsonschema:"Filter pull requests by state"`
	WithAttributes bool   `json:"withAttributes,omitempty" jsonschema:"Include attributes in response"`
	At             string `json:"at,omitempty" jsonschema:"The ref to retrieve pull requests from"`
	WithProperties bool   `json:"withProperties,omitempty" jsonschema:"Include properties in response"`
	Draft          string `json:"draft,omitempty" jsonschema:"Filter by draft status"`
	FilterText     string `json:"filterText,omitempty" jsonschema:"Filter by text"`
	Order          string `json:"order,omitempty" jsonschema:"Sort order"`
	Direction      string `json:"direction,omitempty" jsonschema:"Sort direction"`
}

// GetPullRequestInput represents the input parameters for getting a specific pull request
type GetPullRequestInput struct {
	CommonInput
	PullRequestID int `json:"pullRequestId" jsonschema:"required,The pull request ID"`
}

// GetPullRequestActivitiesInput represents the input parameters for getting pull request activities
type GetPullRequestActivitiesInput struct {
	CommonInput
	PaginationInput
	PullRequestID int    `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	FromType      string `json:"fromType,omitempty" jsonschema:"Filter activities by from type"`
	FromId        string `json:"fromId,omitempty" jsonschema:"Filter activities by from id"`
}

// GetPullRequestCommentsInput represents the input parameters for getting pull request comments
type GetPullRequestCommentsInput struct {
	CommonInput
	PaginationInput
	PullRequestID int    `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	Path          string `json:"path,omitempty" jsonschema:"Filter comments by path"`
	FromHash      string `json:"fromHash,omitempty" jsonschema:"Filter comments by from hash"`
	AnchorState   string `json:"anchorState,omitempty" jsonschema:"Filter comments by anchor state"`
	ToHash        string `json:"toHash,omitempty" jsonschema:"Filter comments by to hash"`
	State         string `json:"state,omitempty" jsonschema:"Filter comments by state"`
	DiffType      string `json:"diffType,omitempty" jsonschema:"Filter comments by diff type"`
	DiffTypes     string `json:"diffTypes,omitempty" jsonschema:"Filter comments by diff types"`
	States        string `json:"states,omitempty" jsonschema:"Filter comments by states"`
}

// AddPullRequestCommentInput represents the input parameters for adding a pull request comment
type AddPullRequestCommentInput struct {
	CommonInput
	PullRequestID int    `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	CommentText   string `json:"commentText" jsonschema:"required,The text of the comment to add"`
}

// DeclinePullRequestInput represents the input parameters for declining a pull request
type DeclinePullRequestInput struct {
	CommonInput
	PullRequestID int     `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	Version       int     `json:"version,omitempty" jsonschema:"The version of the pull request"`
	Comment       *string `json:"comment,omitempty" jsonschema:"Comment for declining the pull request"`
}

// MergePullRequestInput represents the input parameters for merging a pull request
type MergePullRequestInput struct {
	CommonInput
	PullRequestID int     `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	Version       int     `json:"version" jsonschema:"required,The version of the pull request"`
	AutoMerge     *bool   `json:"autoMerge,omitempty" jsonschema:"Whether to auto merge the pull request"`
	AutoSubject   *string `json:"autoSubject,omitempty" jsonschema:"Auto-generated merge commit subject"`
	Message       *string `json:"message,omitempty" jsonschema:"The merge commit message"`
	StrategyId    *string `json:"strategyId,omitempty" jsonschema:"The merge strategy to use"`
}

// GetPullRequestSuggestionsInput represents the input parameters for getting pull request suggestions
type GetPullRequestSuggestionsInput struct {
	ChangesSince string `json:"changesSince,omitempty" jsonschema:"The commit ID or ref to compare since"`
	Limit        int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
}

// GetPullRequestJiraIssuesInput represents the input parameters for getting Jira issues linked to a pull request
type GetPullRequestJiraIssuesInput struct {
	CommonInput
	PullRequestID int `json:"pullRequestId" jsonschema:"required,The pull request ID"`
}

// GetPullRequestsForUserInput represents the input parameters for getting pull requests for a specific user
type GetPullRequestsForUserInput struct {
	PaginationInput
	ClosedSince       string `json:"closedSince,omitempty" jsonschema:"Filter pull requests closed since a specific time"`
	Role              string `json:"role,omitempty" jsonschema:"Filter by user role"`
	ParticipantStatus string `json:"participantStatus,omitempty" jsonschema:"Filter by participant status"`
	State             string `json:"state,omitempty" jsonschema:"Filter pull requests by state"`
	User              string `json:"user,omitempty" jsonschema:"Filter by user"`
	Order             string `json:"order,omitempty" jsonschema:"Sort order"`
}

// GetPullRequestCommentInput represents the input parameters for getting a specific comment on a pull request
type GetPullRequestCommentInput struct {
	CommonInput
	PullRequestID int    `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	CommentID     string `json:"commentId" jsonschema:"required,The ID of the comment to retrieve"`
}

// UpdatePullRequestStatusInput represents the input parameters for updating pull request status
type UpdatePullRequestWithoutStatusInput struct {
	CommonInput
	PullRequestID      int    `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	UserSlug           string `json:"userSlug" jsonschema:"required,The user slug"`
	Version            *int   `json:"version,omitempty" jsonschema:"The version of the pull request participant"`
	LastReviewedCommit string `json:"lastReviewedCommit,omitempty" jsonschema:"The commit ID of the last reviewed commit"`
}

// UpdatePullRequestStatusInput represents the input parameters for updating pull request status
type UpdatePullRequestStatusInput struct {
	UpdatePullRequestWithoutStatusInput
	Status string `json:"status" jsonschema:"required,The status to set for the pull request participant. Valid values: UNAPPROVED, NEEDS_WORK, APPROVED"`
}
