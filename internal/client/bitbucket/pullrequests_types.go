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

// GetPullRequestChangesInput represents the input parameters for getting pull request changes
type GetPullRequestChangesInput struct {
	CommonInput
	PaginationInput
	PullRequestID int `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	// Additional query parameters for pull request changes
	SinceId      string `json:"sinceId,omitempty" jsonschema:"The change ID from which to start retrieving changes"`
	ChangeScope  string `json:"changeScope,omitempty" jsonschema:"The scope of changes to retrieve (e.g. UNREVIEWED)"`
	UntilId      string `json:"untilId,omitempty" jsonschema:"The change ID until which to retrieve changes"`
	WithComments bool   `json:"withComments,omitempty" jsonschema:"Include comments in the response"`
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

// AddPullRequestCommentInput represents the enhanced input parameters for adding a pull request comment
// Supports general comments, replies, inline comments, and code suggestions
type AddPullRequestCommentInput struct {
	CommonInput
	PullRequestID     int     `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	CommentText       string  `json:"commentText" jsonschema:"required,The main comment text. For suggestions, this is the explanation before the code suggestion."`
	CodeSnippet       *string `json:"codeSnippet,omitempty" jsonschema:"Exact code text from the diff to find and comment on. Use this instead of line_number for auto-detection. Must match exactly including whitespace (optional)"`
	FilePath          *string `json:"filePath,omitempty" jsonschema:"File path for inline comment. Required for inline comments. Example: src/components/Button.js (optional)"`
	LineNumber        *int    `json:"lineNumber,omitempty" jsonschema:"Exact line number in the file. Use this OR code_snippet, not both. Required with file_path unless using code_snippet (optional)"`
	LineType          *string `json:"lineType,omitempty" jsonschema:"Type of line: ADDED (green/new lines), REMOVED (red/deleted lines), or CONTEXT (unchanged lines). Default: CONTEXT"`
	MatchStrategy     *string `json:"matchStrategy,omitempty" jsonschema:"How to handle multiple matches when using code_snippet. \"strict\": fail with detailed error showing all matches. \"best\": automatically pick the highest confidence match. Default: \"strict\""`
	ParentCommentID   *int    `json:"parentCommentId,omitempty" jsonschema:"ID of comment to reply to. Use this to create threaded conversations (optional)"`
	Suggestion        *string `json:"suggestion,omitempty" jsonschema:"Replacement code for a suggestion. Creates a suggestion block that can be applied in Bitbucket UI. Requires file_path and line_number. For multi-line, include newlines in the string (optional)"`
	SuggestionEndLine *int    `json:"suggestionEndLine,omitempty" jsonschema:"For multi-line suggestions: the last line number to replace. If not provided, only replaces the single line at line_number (optional)"`
	SearchContext     *string `json:"searchContext,omitempty" jsonschema:"Additional context lines to help locate the exact position when using code_snippet. Useful when the same code appears multiple times (optional)"`
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

// RestJiraIssue represents a Jira issue linked to a pull request
type RestJiraIssue struct {
	Key string `json:"key"`
	URL string `json:"url"`
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

// GetPullRequestDiffInput represents the input parameters for getting a diff within a pull request
type GetPullRequestDiffInput struct {
	CommonInput
	PullRequestID int    `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	Path          string `json:"path" jsonschema:"required,Path to the file"`
	SrcPath       string `json:"srcPath,omitempty" jsonschema:"Source path for comparison"`
	ContextLines  string `json:"contextLines,omitempty" jsonschema:"Number of context lines to include"`
	SinceId       string `json:"sinceId,omitempty" jsonschema:"Filter changes since a specific time"`
	UntilId       string `json:"untilId,omitempty" jsonschema:"Filter changes until a specific time"`
	Whitespace    string `json:"whitespace,omitempty" jsonschema:"Whitespace handling option"`
	WithComments  string `json:"withComments,omitempty" jsonschema:"Include comments in response"`
	DiffType      string `json:"diffType,omitempty" jsonschema:"Filter by diff type"`
	AvatarScheme  string `json:"avatarScheme,omitempty" jsonschema:"Avatar scheme"`
	AvatarSize    string `json:"avatarSize,omitempty" jsonschema:"Avatar size"`
}

// GetPullRequestDiffStreamInput represents the input parameters for streaming pull request diff
type GetPullRequestDiffStreamInput struct {
	CommonInput
	PullRequestID int    `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	ContextLines  int    `json:"contextLines,omitempty" jsonschema:"Number of context lines to include"`
	Whitespace    string `json:"whitespace,omitempty" jsonschema:"Whitespace handling option"`
}

// GetPullRequestMergeStatusInput represents the input parameters for getting pull request merge status
type TestPullRequestCanMergeInput struct {
	CommonInput
	PullRequestID int `json:"pullRequestId" jsonschema:"required,The pull request ID"`
}

// SearchContext represents additional context lines to help locate the exact position when using code_snippet
type SearchContext struct {
	Before []string `json:"before,omitempty" jsonschema:"Array of code lines that appear BEFORE the target line. Helps disambiguate when code_snippet appears multiple times"`
	After  []string `json:"after,omitempty" jsonschema:"Array of code lines that appear AFTER the target line. Helps disambiguate when code_snippet appears multiple times"`
}

// ResolvedLineInfo represents the information about a resolved line in a pull request diff
type ResolvedLineInfo struct {
	LineNumber int    `json:"lineNumber"`
	FilePath   string `json:"filePath"`
	LineType   string `json:"lineType"`
	Confidence int    `json:"confidence"`

	// Internal fields for processing, not for serialization
	hunkBody      []string `json:"-"`
	hunkLineIndex int      `json:"-"`
}

// ResolveLineFromCodeInput represents the input parameters for resolving a line number from a code snippet
type ResolveLineFromCodeInput struct {
	CommonInput
	PullRequestID int            `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	CodeSnippet   string         `json:"codeSnippet" jsonschema:"required,The code snippet to find"`
	FilePath      *string        `json:"filePath,omitempty" jsonschema:"File path to filter matches. If not provided, will return matches from any file"`
	LineType      *string        `json:"lineType,omitempty" jsonschema:"Type of line to match: ADDED, REMOVED, or CONTEXT. If not provided, will match any line type"`
	MatchStrategy *string        `json:"matchStrategy,omitempty" jsonschema:"How to handle multiple matches. \"strict\": fail with detailed error. \"best\": automatically pick the best match"`
	SearchContext *SearchContext `json:"searchContext,omitempty" jsonschema:"Additional context to help disambiguate matches"`
}

// CommentPayload defines the structure for the Bitbucket API to add a comment.
type CommentPayload struct {
	Text       string      `json:"text"`
	Parent     *ParentID   `json:"parent,omitempty"`
	Anchor     *Anchor     `json:"anchor,omitempty"`
	Suggestion *Suggestion `json:"suggestion,omitempty"`
}

// ParentID specifies the parent comment for a reply.
type ParentID struct {
	ID int `json:"id"`
}

// Anchor defines where an inline comment should be placed.
type Anchor struct {
	Path          string  `json:"path"`
	Line          *int    `json:"line,omitempty"`
	LineType      string  `json:"lineType,omitempty"`
	FileType      string  `json:"fileType,omitempty"`
	DiffType      string  `json:"diffType,omitempty"`
	Snippet       *string `json:"snippet,omitempty"`
	MatchStrategy string  `json:"matchStrategy,omitempty"`
}

// Suggestion defines a code suggestion.
type Suggestion struct {
	Content string `json:"content"`
	EndLine *int   `json:"endLine,omitempty"`
}
