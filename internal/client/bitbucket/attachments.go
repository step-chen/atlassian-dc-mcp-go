package bitbucket

import (
	"context"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
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
func (c *BitbucketClient) GetAttachment(ctx context.Context, input GetAttachmentInput) ([]byte, error) {
	var content []byte
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "attachments", input.AttachmentId},
		nil,
		nil,
		client.AcceptJSON,
		&content,
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
func (c *BitbucketClient) GetAttachmentMetadata(ctx context.Context, input GetAttachmentMetadataInput) (types.MapOutput, error) {
	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "attachments", input.AttachmentId, "metadata"},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
func (c *BitbucketClient) DeleteAttachment(ctx context.Context, input DeleteAttachmentInput) error {
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodDelete,
		[]any{"rest", "attachment", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", input.PullRequestId, "attachments", input.AttachmentId},
		nil,
		nil,
		client.AcceptJSON,
		nil,
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
func (c *BitbucketClient) CreateAttachment(ctx context.Context, input CreateAttachmentInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "filename", input.FileName, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodPost,
		[]any{"rest", "attachment", "latest", "projects", input.ProjectKey, "repos", input.RepoSlug, "pull-requests", input.PullRequestID, "attachments"},
		queryParams,
		input.FileAttachment,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
