package confluence

import (
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
)

// GetContentChildren retrieves children of a specific content.
//
// Parameters:
//   - input: GetContentChildrenInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The content children data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentChildren(input GetContentChildrenInput) (types.MapOutput, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "expand", input.Expand, []string{})
	utils.SetQueryParam(params, "parentVersion", input.ParentVersion, "")

	var children types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content", input.ContentID, "child"}, params, nil, &children, utils.AcceptJSON); err != nil {
		return nil, err
	}

	return children, nil
}

// GetContentChildrenByType retrieves children of a specific content by type.
//
// Parameters:
//   - input: GetContentChildrenByTypeInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The content children data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentChildrenByType(input GetContentChildrenByTypeInput) (types.MapOutput, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "expand", input.Expand, []string{})
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)
	utils.SetQueryParam(params, "orderBy", input.OrderBy, "")

	var children types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content", input.ContentID, "child", input.ChildType}, params, nil, &children, utils.AcceptJSON); err != nil {
		return nil, err
	}

	return children, nil
}

// GetContentComments retrieves comments of a specific content.
//
// Parameters:
//   - input: GetContentCommentsInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The content comments data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentComments(input GetContentCommentsInput) (types.MapOutput, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "expand", input.Expand, []string{})
	utils.SetQueryParam(params, "parentVersion", input.ParentVersion, "")
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)

	var comments types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "content", input.ContentID, "child", "comment"}, params, nil, &comments, utils.AcceptJSON); err != nil {
		return nil, err
	}

	return comments, nil
}
