package bitbucket

// CreateAttachmentInput represents the input parameters for creating an attachment
type CreateAttachmentInput struct {
	CommonInput
	PullRequestID  int    `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	FileName       string `json:"fileName" jsonschema:"required,The name of the file"`
	FileAttachment []byte `json:"fileAttachment" jsonschema:"required,The file attachment content"`
}

// GetAttachmentInput represents the input parameters for getting an attachment
type GetAttachmentInput struct {
	CommonInput
	AttachmentId string `json:"attachmentId" jsonschema:"required,The ID of the attachment to retrieve"`
}

// GetAttachmentMetadataInput represents the input parameters for getting attachment metadata
type GetAttachmentMetadataInput struct {
	CommonInput
	AttachmentId string `json:"attachmentId" jsonschema:"required,The ID of the attachment to retrieve metadata for"`
}

// DeleteAttachmentInput represents the input parameters for deleting an attachment
type DeleteAttachmentInput struct {
	CommonInput
	PullRequestId int `json:"pullRequestId" jsonschema:"required,The pull request ID"`
	AttachmentId  int `json:"attachmentId" jsonschema:"required,The attachment ID"`
}
