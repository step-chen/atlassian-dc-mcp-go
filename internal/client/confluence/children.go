package confluence

import (
	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
	"context"
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
func (c *ConfluenceClient) GetContentChildren(ctx context.Context, input GetContentChildrenInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})
	client.SetQueryParam(queryParams, "parentVersion", input.ParentVersion, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content", input.ContentID, "child"},
		queryParams,
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
func (c *ConfluenceClient) GetContentChildrenByType(ctx context.Context, input GetContentChildrenByTypeInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "orderBy", input.OrderBy, "")

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content", input.ContentID, "child", input.ChildType},
		queryParams,
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
func (c *ConfluenceClient) GetContentComments(ctx context.Context, input GetContentCommentsInput) (types.MapOutput, error) {
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})
	client.SetQueryParam(queryParams, "parentVersion", input.ParentVersion, "")
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "content", input.ContentID, "child", "comment"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
