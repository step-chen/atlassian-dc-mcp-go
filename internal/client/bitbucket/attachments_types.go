package bitbucket

// CreateAttachmentInput represents the input parameters for creating an attachment
type CreateAttachmentInput struct {
	ProjectKey     string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug       string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PullRequestID  int    `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	FileName       string `json:"fileName" jsonschema:"required,The name of the file"`
	FileAttachment []byte `json:"fileAttachment" jsonschema:"required,The file attachment content"`
}

// GetAttachmentInput represents the input parameters for getting an attachment
type GetAttachmentInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	AttachmentId string `json:"attachmentId" jsonschema:"required,The ID of the attachment to retrieve"`
}

// GetAttachmentMetadataInput represents the input parameters for getting attachment metadata
type GetAttachmentMetadataInput struct {
	ProjectKey   string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug     string `json:"repoSlug" jsonschema:"required,The repository slug"`
	AttachmentId string `json:"attachmentId" jsonschema:"required,The ID of the attachment to retrieve metadata for"`
}

// DeleteAttachmentInput represents the input parameters for deleting an attachment
type DeleteAttachmentInput struct {
	ProjectKey     string `json:"projectKey" jsonschema:"required,The project key"`
	RepoSlug       string `json:"repoSlug" jsonschema:"required,The repository slug"`
	PullRequestId  int    `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	AttachmentId   int    `json:"attachmentId" jsonschema:"required,The attachment ID"`
}