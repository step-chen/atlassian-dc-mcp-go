package confluence

import (
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetContentChildren retrieves children of a specific content.
//
// Parameters:
//   - contentID: The ID of the content
//   - expand: Fields to expand in the response
//   - parentVersion: The version of the parent content
//
// Returns:
//   - map[string]interface{}: The content children data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentChildren(contentID string, expand []string, parentVersion string) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "expand", expand, []string{})
	utils.SetQueryParam(params, "parentVersion", parentVersion, "")

	var children map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", contentID, "child"}, params, nil, &children); err != nil {
		return nil, err
	}

	return children, nil
}

// GetContentChildrenByType retrieves children of a specific content by type.
//
// Parameters:
//   - contentID: The ID of the content
//   - childType: The type of child content to retrieve
//   - expand: Fields to expand in the response
//   - start: Starting index for pagination
//   - limit: Maximum number of results to return
//   - orderBy: Field to order results by
//
// Returns:
//   - map[string]interface{}: The content children data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentChildrenByType(contentID string, childType string, expand []string, start, limit int, orderBy string) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "expand", expand, []string{})
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)
	utils.SetQueryParam(params, "orderBy", orderBy, "")

	var children map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", contentID, "child", childType}, params, nil, &children); err != nil {
		return nil, err
	}

	return children, nil
}

// GetContentComments retrieves comments of a specific content.
//
// Parameters:
//   - contentID: The ID of the content
//   - expand: Fields to expand in the response
//   - parentVersion: The version of the parent content
//   - start: Starting index for pagination
//   - limit: Maximum number of results to return
//
// Returns:
//   - map[string]interface{}: The content comments data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentComments(contentID string, expand []string, parentVersion string, start, limit int) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "expand", expand, []string{})
	utils.SetQueryParam(params, "parentVersion", parentVersion, "")
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)

	var comments map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "content", contentID, "child", "comment"}, params, nil, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}