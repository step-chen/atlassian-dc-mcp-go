// Package confluence provides a client for interacting with Confluence Data Center APIs.
package confluence

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// GetContent retrieves content based on various filters.
//
// Parameters:
//   - input: GetContentInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The content data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContent(input GetContentInput) (types.MapOutput, error) {
	queryParams := url.Values{}

	client.SetQueryParam(queryParams, "type", input.TypeParam, "")
	client.SetQueryParam(queryParams, "spaceKey", input.SpaceKey, "")
	client.SetQueryParam(queryParams, "title", input.Title, "")
	client.SetQueryParam(queryParams, "status", input.Status, []string{})
	client.SetQueryParam(queryParams, "postingDay", input.PostingDay, "")
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetContentByID retrieves content by its ID.
//
// Parameters:
//   - input: GetContentByIDInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The content data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentByID(input GetContentByIDInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content", input.ContentID},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// SearchContent searches for content based on CQL.
//
// Parameters:
//   - input: SearchContentInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The search results
//   - error: An error if the request fails
func (c *ConfluenceClient) SearchContent(input SearchContentInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "cql", input.CQL, "")
	client.SetQueryParam(queryParams, "cqlcontext", input.CQLContext, "")
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content", "search"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// CreateContent creates new content.
//
// Parameters:
//   - input: CreateContentInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The created content data
//   - error: An error if the request fails
func (c *ConfluenceClient) CreateContent(input CreateContentInput) (types.MapOutput, error) {
	payload := types.MapOutput{}
	client.SetRequestBodyParam(payload, "type", input.Type)
	client.SetRequestBodyParam(payload, "title", input.Title)
	client.SetRequestBodyParam(payload, "space", input.Space)
	client.SetRequestBodyParam(payload, "body", input.Body)
	client.SetRequestBodyParam(payload, "ancestors", input.Ancestors)
	client.SetRequestBodyParam(payload, "metadata", input.Metadata)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodPost,
		[]any{"rest", "api", "content"},
		nil,
		jsonPayload,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// UpdateContent updates existing content.
//
// Parameters:
//   - input: UpdateContentInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The updated content data
//   - error: An error if the request fails
func (c *ConfluenceClient) UpdateContent(input UpdateContentInput) (types.MapOutput, error) {
	jsonPayload, err := json.Marshal(input.ContentData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodPut,
		[]any{"rest", "api", "content", input.ContentID},
		nil,
		jsonPayload,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// DeleteContent deletes content by its ID.
//
// Parameters:
//   - input: DeleteContentInput containing the parameters for the request
//
// Returns:
//   - error: An error if the request fails
func (c *ConfluenceClient) DeleteContent(input DeleteContentInput) error {
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodDelete,
		[]any{"rest", "api", "content", input.ContentID},
		nil,
		nil,
		client.AcceptJSON,
		nil,
	); err != nil {
		return err
	}

	return nil
}

// GetContentHistory retrieves the history of content.
//
// Parameters:
//   - input: GetContentHistoryInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The content history data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentHistory(input GetContentHistoryInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content", input.ContentID, "history"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// AddComment adds a comment to content.
//
// Parameters:
//   - input: AddCommentInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The added comment data
//   - error: An error if the request fails
func (c *ConfluenceClient) AddComment(input AddCommentInput) (types.MapOutput, error) {
	payload := types.MapOutput{
		"type": "comment",
		"container": map[string]string{
			"id":   input.ContentID,
			"type": "page",
		},
		"body": types.MapOutput{
			"storage": map[string]string{
				"value":          input.CommentBody,
				"representation": "storage",
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodPost,
		[]any{"rest", "api", "content"},
		nil,
		jsonPayload,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetAttachments retrieves attachments for content.
//
// Parameters:
//   - input: GetAttachmentsInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The attachments data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetAttachments(input GetAttachmentsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "filename", input.Filename, "")
	client.SetQueryParam(queryParams, "mediaType", input.MediaType, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content", input.ContentID, "child", "attachment"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetExtractedText retrieves extracted text from an attachment.
//
// Parameters:
//   - input: GetExtractedTextInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The extracted text data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetExtractedText(input GetExtractedTextInput) (types.MapOutput, error) {
	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content", input.ContentID, "child", "attachment", input.AttachmentID, "extractedText"},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetContentLabels retrieves labels for content.
//
// Parameters:
//   - input: GetContentLabelsInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The labels data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentLabels(input GetContentLabelsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content", input.ContentID, "label"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// ScanContentBySpaceKey scans content by space key.
//
// Parameters:
//   - input: ScanContentBySpaceKeyInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The scanned content data
//   - error: An error if the request fails
func (c *ConfluenceClient) ScanContentBySpaceKey(input ScanContentBySpaceKeyInput) (types.MapOutput, error) {
	queryParams := url.Values{}

	client.SetQueryParam(queryParams, "type", input.TypeParam, "")
	client.SetQueryParam(queryParams, "spaceKey", input.SpaceKey, "")
	client.SetQueryParam(queryParams, "status", input.Status, []string{})
	client.SetQueryParam(queryParams, "postingDay", input.PostingDay, "")
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})
	client.SetQueryParam(queryParams, "cursor", input.Cursor, "")
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content", "scan"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
