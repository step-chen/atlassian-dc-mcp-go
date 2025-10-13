package bitbucket

import (
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
)

// GetAttachment retrieves an attachment from a repository.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch an attachment
// from a specific repository.
//
// Parameters:
//   - input: GetAttachmentInput containing the parameters for the request
//
// Returns:
//   - []byte: The attachment content as bytes
//   - error: An error if the request fails
func (c *BitbucketClient) GetAttachment(input GetAttachmentInput) ([]byte, error) {
	var content []byte
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "attachments", input.AttachmentId},
		nil,
		nil,
		&content,
		"application/json",
	); err != nil {
		return nil, err
	}

	return content, nil
}

// GetAttachmentMetadata retrieves metadata for an attachment.
//
// This function makes an HTTP GET request to the Bitbucket API to fetch metadata
// for an attachment from a specific repository.
//
// Parameters:
//   - input: GetAttachmentMetadataInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The attachment metadata retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) GetAttachmentMetadata(input GetAttachmentMetadataInput) (types.MapOutput, error) {
	var metadata types.MapOutput
	if err := c.executeRequest(
		http.MethodGet,
		[]string{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "attachments", input.AttachmentId, "metadata"},
		nil,
		nil,
		&metadata,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return metadata, nil
}

// DeleteAttachment deletes an attachment.
//
// This function makes an HTTP DELETE request to the Bitbucket API to delete an attachment.
//
// Parameters:
//   - input: DeleteAttachmentInput containing the parameters for the request
//
// Returns:
//   - error: An error if the request fails
func (c *BitbucketClient) DeleteAttachment(input DeleteAttachmentInput) error {
	if err := c.executeRequest(
		http.MethodDelete,
		[]string{"rest", "attachment", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", fmt.Sprintf("%d", input.PullRequestId), "attachments", fmt.Sprintf("%d", input.AttachmentId)},
		nil,
		nil,
		nil,
		utils.AcceptJSON,
	); err != nil {
		return err
	}

	return nil
}

// CreateAttachment creates an attachment.
//
// This function makes an HTTP POST request to the Bitbucket API to create an attachment.
//
// Parameters:
//   - input: CreateAttachmentInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The created attachment data retrieved from the API
//   - error: An error if the request fails
func (c *BitbucketClient) CreateAttachment(input CreateAttachmentInput) (types.MapOutput, error) {
	queryParams := make(url.Values)
	utils.SetQueryParam(queryParams, "filename", input.FileName, "")

	var attachment types.MapOutput
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "attachment", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", fmt.Sprintf("%d", input.PullRequestID), "attachments"},
		queryParams,
		input.FileAttachment,
		&attachment,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return attachment, nil
}
