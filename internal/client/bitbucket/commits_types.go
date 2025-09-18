package bitbucket

// GetCommitsInput represents the input parameters for getting commits
type GetCommitsInput struct {
	ProjectKey     string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug       string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Until          string `json:"until,omitempty" jsonschema:"The commit ID or ref to retrieve commits until"`
	Since          string `json:"since,omitempty" jsonschema:"The commit ID or ref to retrieve commits since"`
	Path           string `json:"path,omitempty" jsonschema:"Filter commits by file path"`
	Limit          int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
	Start          int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Merges         string `json:"merges,omitempty" jsonschema:"Filter merge commits"`
	FollowRenames  bool   `json:"followRenames,omitempty" jsonschema:"Follow file renames"`
	IgnoreMissing  bool   `json:"ignoreMissing,omitempty" jsonschema:"Ignore missing commits"`
	WithCounts     bool   `json:"withCounts,omitempty" jsonschema:"Include commit counts"`
}

// GetCommitInput represents the input parameters for getting a specific commit
type GetCommitInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitID   string `json:"commitId" jsonschema:"required,The ID of the commit to retrieve"`
	Path       string `json:"path,omitempty" jsonschema:"Filter commit details by file path"`
}

// GetCommitChangesInput represents the input parameters for getting commit changes
type GetCommitChangesInput struct {
	ProjectKey    string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug      string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitID      string `json:"commitId" jsonschema:"required,The ID of the commit to retrieve changes for"`
	Start         int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Limit         int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
	WithComments  string `json:"withComments,omitempty" jsonschema:"Include comments in response"`
	Since         string `json:"since,omitempty" jsonschema:"Filter changes since a specific time"`
}

// GetCommitDiffStatsSummaryInput represents the input parameters for getting commit diff statistics summary
type GetCommitDiffStatsSummaryInput struct {
	ProjectKey  string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug    string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitID    string `json:"commitId" jsonschema:"required,The ID of the commit"`
	Path        string `json:"path" jsonschema:"required,The file path"`
	SrcPath     string `json:"srcPath,omitempty" jsonschema:"Source path for comparison"`
	AutoSrcPath string `json:"autoSrcPath,omitempty" jsonschema:"Automatically determine source path"`
	Whitespace  string `json:"whitespace,omitempty" jsonschema:"Whitespace handling option"`
	Since       string `json:"since,omitempty" jsonschema:"Filter changes since a specific time"`
}

// GetDiffBetweenCommitsInput represents the input parameters for getting diff between commits
type GetDiffBetweenCommitsInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Path         string `json:"path,omitempty" jsonschema:"The file path"`
	From         string `json:"from,omitempty" jsonschema:"The source commit ID or ref"`
	To           string `json:"to,omitempty" jsonschema:"The target commit ID or ref"`
	ContextLines int    `json:"contextLines,omitempty" jsonschema:"Number of context lines to include"`
	SrcPath      string `json:"srcPath,omitempty" jsonschema:"Source path for comparison"`
	Whitespace   string `json:"whitespace,omitempty" jsonschema:"Whitespace handling option"`
	FromRepo     string `json:"fromRepo,omitempty" jsonschema:"The source repository"`
}

// GetDiffBetweenRevisionsInput represents the input parameters for getting diff between revisions
type GetDiffBetweenRevisionsInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitID     string `json:"commitId" jsonschema:"required,The commit ID"`
	Path         string `json:"path" jsonschema:"required,The file path"`
	ContextLines int    `json:"contextLines,omitempty" jsonschema:"Number of context lines to include"`
	Since        string `json:"since,omitempty" jsonschema:"Filter changes since a specific time"`
	SrcPath      string `json:"srcPath,omitempty" jsonschema:"Source path for comparison"`
	Whitespace   string `json:"whitespace,omitempty" jsonschema:"Whitespace handling option"`
	Filter       string `json:"filter,omitempty" jsonschema:"Filter option"`
	AutoSrcPath  string `json:"autoSrcPath,omitempty" jsonschema:"Automatically determine source path"`
	WithComments string `json:"withComments,omitempty" jsonschema:"Include comments in response"`
}

// GetCommitCommentInput represents the input parameters for getting a commit comment
type GetCommitCommentInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitID   string `json:"commitId" jsonschema:"required,The ID of the commit"`
	CommentID  int    `json:"commentId" jsonschema:"required,The ID of the comment to retrieve"`
}

// GetCommitCommentsInput represents the input parameters for getting commit comments
type GetCommitCommentsInput struct {
	ProjectKey string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug   string `json:"repoSlug" jsonschema:"required,The repository slug"`
	CommitID   string `json:"commitId" jsonschema:"required,The ID of the commit"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	Path       string `json:"path,omitempty" jsonschema:"Filter comments by file path"`
	Since      string `json:"since,omitempty" jsonschema:"Filter comments since a specific time"`
}

// GetJiraIssueCommitsInput represents the input parameters for getting Jira issue commits
type GetJiraIssueCommitsInput struct {
	IssueKey   string `json:"issueKey" jsonschema:"required,The Jira issue key"`
	Limit      int    `json:"limit,omitempty" jsonschema:"Maximum number of results to return (default: 25)"`
	Start      int    `json:"start,omitempty" jsonschema:"Starting index for pagination (default: 0)"`
	MaxChanges int    `json:"maxChanges,omitempty" jsonschema:"Maximum number of changes to include"`
}

// GetDiffBetweenRevisionsForPathInput represents the input parameters for getting diff between revisions for a path
type GetDiffBetweenRevisionsForPathInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	Path         string `json:"path" jsonschema:"required,The file path"`
	ContextLines int    `json:"contextLines,omitempty" jsonschema:"Number of context lines to include"`
	Since        string `json:"since,omitempty" jsonschema:"Filter changes since a specific time"`
	Until        string `json:"until,omitempty" jsonschema:"Filter changes until a specific time"`
	SrcPath      string `json:"srcPath,omitempty" jsonschema:"Source path for comparison"`
	Whitespace   string `json:"whitespace,omitempty" jsonschema:"Whitespace handling option"`
}