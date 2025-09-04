// Package confluence provides a client for interacting with Confluence Data Center APIs.
package confluence

import (
	"encoding/json"
	"fmt"
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetContent retrieves content based on various filters.
func (c *ConfluenceClient) GetContent(
	typeParam string,
	spaceKey string,
	title string,
	status []string,
	postingDay string,
	expand []string,
	start int,
	limit int,
) (map[string]any, error) {
	params := url.Values{}

	utils.SetQueryParam(params, "type", typeParam, "")
	utils.SetQueryParam(params, "spaceKey", spaceKey, "")
	utils.SetQueryParam(params, "title", title, "")
	utils.SetQueryParam(params, "status", status, []string{})
	utils.SetQueryParam(params, "postingDay", postingDay, "")
	utils.SetQueryParam(params, "expand", expand, []string{})
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)

	var content map[string]any
	if err := c.executeRequest("GET", []string{"rest", "api", "content"}, params, nil, &content); err != nil {
		return nil, err
	}

	return content, nil
}

// GetContentByID retrieves content by its ID.
func (c *ConfluenceClient) GetContentByID(contentID string, expand []string) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "expand", expand, []string{})

	var content map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", contentID}, params, nil, &content); err != nil {
		return nil, err
	}

	return content, nil
}

// SearchContent searches for content based on CQL.
func (c *ConfluenceClient) SearchContent(cql, cqlcontext string, start, limit int, expand []string) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "cql", cql, "")
	utils.SetQueryParam(params, "cqlcontext", cqlcontext, "")
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)
	utils.SetQueryParam(params, "expand", expand, []string{})

	var result map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", "search"}, params, nil, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateContent creates new content.
func (c *ConfluenceClient) CreateContent(contentType, title string, space map[string]interface{}, body map[string]interface{}, ancestors []map[string]interface{}, metadata map[string]interface{}) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"type":  contentType,
		"title": title,
		"space": space,
		"body":  body,
	}

	if len(ancestors) > 0 {
		payload["ancestors"] = ancestors
	}

	if len(metadata) > 0 {
		payload["metadata"] = metadata
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var content map[string]interface{}
	if err := c.executeRequest("POST", []string{"rest", "api", "content"}, nil, jsonPayload, &content); err != nil {
		return nil, err
	}

	return content, nil
}

// UpdateContent updates existing content.
func (c *ConfluenceClient) UpdateContent(contentID string, contentData map[string]interface{}) (map[string]interface{}, error) {
	jsonPayload, err := json.Marshal(contentData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var updatedContent map[string]interface{}
	if err := c.executeRequest("PUT", []string{"rest", "api", "content", contentID}, nil, jsonPayload, &updatedContent); err != nil {
		return nil, err
	}

	return updatedContent, nil
}

// DeleteContent deletes content by its ID.
func (c *ConfluenceClient) DeleteContent(contentID string) error {
	if err := c.executeRequest("DELETE", []string{"rest", "api", "content", contentID}, nil, nil, nil); err != nil {
		return err
	}

	return nil
}

// GetContentHistory retrieves the history of content.
func (c *ConfluenceClient) GetContentHistory(contentID string, expand []string) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "expand", expand, []string{})

	var history map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", contentID, "history"}, params, nil, &history); err != nil {
		return nil, err
	}

	return history, nil
}

// GetComments retrieves comments for content.
func (c *ConfluenceClient) GetComments(contentID string, expand []string, start, limit int) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "expand", expand, []string{})
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)

	var comments map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", contentID, "child", "comment"}, params, nil, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

// AddComment adds a comment to content.
func (c *ConfluenceClient) AddComment(contentID, commentBody string) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"type": "comment",
		"container": map[string]string{
			"id":   contentID,
			"type": "page",
		},
		"body": map[string]interface{}{
			"storage": map[string]string{
				"value":          commentBody,
				"representation": "storage",
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var comment map[string]interface{}
	if err := c.executeRequest("POST", []string{"rest", "api", "content"}, nil, jsonPayload, &comment); err != nil {
		return nil, err
	}

	return comment, nil
}

// GetAttachments retrieves attachments for content.
func (c *ConfluenceClient) GetAttachments(contentID string, expand []string, start, limit int, filename string, mediaType string) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "expand", expand, []string{})
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)
	utils.SetQueryParam(params, "filename", filename, "")
	utils.SetQueryParam(params, "mediaType", mediaType, "")

	var attachments map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", contentID, "child", "attachment"}, params, nil, &attachments); err != nil {
		return nil, err
	}

	return attachments, nil
}

// GetExtractedText retrieves extracted text from an attachment.
func (c *ConfluenceClient) GetExtractedText(contentID, attachmentID string) (map[string]interface{}, error) {
	var extractedText map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", contentID, "child", "attachment", attachmentID, "extractedText"}, nil, nil, &extractedText); err != nil {
		return nil, err
	}

	return extractedText, nil
}

// GetContentLabels retrieves labels for content.
func (c *ConfluenceClient) GetContentLabels(contentID string) (map[string]interface{}, error) {
	var labels map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", contentID, "label"}, nil, nil, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}

// ScanContentBySpaceKey scans content by space key.
func (c *ConfluenceClient) ScanContentBySpaceKey(
	typeParam string,
	spaceKey string,
	status []string,
	postingDay string,
	expand []string,
	cursor string,
	limit int,
) (map[string]interface{}, error) {
	params := url.Values{}

	utils.SetQueryParam(params, "type", typeParam, "")
	utils.SetQueryParam(params, "spaceKey", spaceKey, "")
	utils.SetQueryParam(params, "status", status, []string{})
	utils.SetQueryParam(params, "postingDay", postingDay, "")
	utils.SetQueryParam(params, "expand", expand, []string{})
	utils.SetQueryParam(params, "cursor", cursor, "")
	utils.SetQueryParam(params, "limit", limit, 0)

	var content map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", "scan"}, params, nil, &content); err != nil {
		return nil, err
	}

	return content, nil
}