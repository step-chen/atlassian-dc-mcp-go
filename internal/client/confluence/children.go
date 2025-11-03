package confluence

import (
	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
	"net/http"
	"net/url"
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
	client.SetQueryParam(params, "expand", input.Expand, []string{})
	client.SetQueryParam(params, "parentVersion", input.ParentVersion, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "content", input.ContentID, "child"},
		params,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
	client.SetQueryParam(params, "expand", input.Expand, []string{})
	client.SetQueryParam(params, "start", input.Start, 0)
	client.SetQueryParam(params, "limit", input.Limit, 0)
	client.SetQueryParam(params, "orderBy", input.OrderBy, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "content", input.ContentID, "child", input.ChildType},
		params,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
	client.SetQueryParam(params, "expand", input.Expand, []string{})
	client.SetQueryParam(params, "parentVersion", input.ParentVersion, "")
	client.SetQueryParam(params, "start", input.Start, 0)
	client.SetQueryParam(params, "limit", input.Limit, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "content", input.ContentID, "child", "comment"},
		params,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
