// Package confluence provides a client for interacting with Confluence Data Center APIs.
package confluence

import (
	"encoding/json"
	"fmt"
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
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
	params := url.Values{}

	utils.SetQueryParam(params, "type", input.TypeParam, "")
	utils.SetQueryParam(params, "spaceKey", input.SpaceKey, "")
	utils.SetQueryParam(params, "title", input.Title, "")
	utils.SetQueryParam(params, "status", input.Status, []string{})
	utils.SetQueryParam(params, "postingDay", input.PostingDay, "")
	utils.SetQueryParam(params, "expand", input.Expand, []string{})
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)

	var content types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content"}, params, nil, &content); err != nil {
		return nil, err
	}

	return content, nil
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
	params := url.Values{}
	utils.SetQueryParam(params, "expand", input.Expand, []string{})

	var content types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content", input.ContentID}, params, nil, &content); err != nil {
		return nil, err
	}

	return content, nil
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
	params := url.Values{}
	utils.SetQueryParam(params, "cql", input.CQL, "")
	utils.SetQueryParam(params, "cqlcontext", input.CQLContext, "")
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)
	utils.SetQueryParam(params, "expand", input.Expand, []string{})

	var result types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content", "search"}, params, nil, &result); err != nil {
		return nil, err
	}

	return result, nil
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
	payload := types.MapOutput{
		"type":  input.Type,
		"title": input.Title,
		"space": input.Space,
		"body":  input.Body,
	}

	if len(input.Ancestors) > 0 {
		payload["ancestors"] = input.Ancestors
	}

	if len(input.Metadata) > 0 {
		payload["metadata"] = input.Metadata
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var content types.MapOutput
	if err := c.executeRequest("POST", []string{"rest", "api", "content"}, nil, jsonPayload, &content); err != nil {
		return nil, err
	}

	return content, nil
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

	var updatedContent types.MapOutput
	if err := c.executeRequest("PUT", []string{"rest", "api", "content", input.ContentID}, nil, jsonPayload, &updatedContent); err != nil {
		return nil, err
	}

	return updatedContent, nil
}

// DeleteContent deletes content by its ID.
//
// Parameters:
//   - input: DeleteContentInput containing the parameters for the request
//
// Returns:
//   - error: An error if the request fails
func (c *ConfluenceClient) DeleteContent(input DeleteContentInput) error {
	if err := c.executeRequest("DELETE", []string{"rest", "api", "content", input.ContentID}, nil, nil, nil); err != nil {
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
	params := url.Values{}
	utils.SetQueryParam(params, "expand", input.Expand, []string{})

	var history types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content", input.ContentID, "history"}, params, nil, &history); err != nil {
		return nil, err
	}

	return history, nil
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

	var comment types.MapOutput
	if err := c.executeRequest("POST", []string{"rest", "api", "content"}, nil, jsonPayload, &comment); err != nil {
		return nil, err
	}

	return comment, nil
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
	params := url.Values{}
	utils.SetQueryParam(params, "expand", input.Expand, []string{})
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)
	utils.SetQueryParam(params, "filename", input.Filename, "")
	utils.SetQueryParam(params, "mediaType", input.MediaType, "")

	var attachments types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content", input.ContentID, "child", "attachment"}, params, nil, &attachments); err != nil {
		return nil, err
	}

	return attachments, nil
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
	var extractedText types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content", input.ContentID, "child", "attachment", input.AttachmentID, "extractedText"}, nil, nil, &extractedText); err != nil {
		return nil, err
	}

	return extractedText, nil
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
	params := url.Values{}
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)

	var labels types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content", input.ContentID, "label"}, params, nil, &labels); err != nil {
		return nil, err
	}

	return labels, nil
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
	params := url.Values{}

	utils.SetQueryParam(params, "type", input.TypeParam, "")
	utils.SetQueryParam(params, "spaceKey", input.SpaceKey, "")
	utils.SetQueryParam(params, "status", input.Status, []string{})
	utils.SetQueryParam(params, "postingDay", input.PostingDay, "")
	utils.SetQueryParam(params, "expand", input.Expand, []string{})
	utils.SetQueryParam(params, "cursor", input.Cursor, "")
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)

	var content types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content", "scan"}, params, nil, &content); err != nil {
		return nil, err
	}

	return content, nil
}
