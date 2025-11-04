package confluence

import (
	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// GetSpace retrieves a specific space by its key.
//
// Parameters:
//   - input: GetSpaceInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The space data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetSpace(input GetSpaceInput) (types.MapOutput, error) {
	if input.SpaceKey == "" {
		return nil, fmt.Errorf("spaceKey cannot be empty")
	}

	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "space", input.SpaceKey},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetContentsInSpace retrieves contents in a specific space.
//
// Parameters:
//   - input: GetContentsInSpaceInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The contents data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentsInSpace(input GetContentsInSpaceInput) (types.MapOutput, error) {
	if input.SpaceKey == "" {
		return nil, fmt.Errorf("spaceKey cannot be empty")
	}

	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "space", input.SpaceKey, "content"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetContentsByType retrieves contents of a specific type in a space.
//
// Parameters:
//   - input: GetContentsByTypeInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The contents data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentsByType(input GetContentsByTypeInput) (types.MapOutput, error) {
	if input.SpaceKey == "" {
		return nil, fmt.Errorf("spaceKey cannot be empty")
	}

	if input.ContentType == "" {
		return nil, fmt.Errorf("contentType cannot be empty")
	}

	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "space", input.SpaceKey, "content", input.ContentType},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// GetSpacesByKey retrieves spaces based on various filters.
//
// Parameters:
//   - input: GetSpacesByKeyInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The spaces data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetSpacesByKey(input GetSpacesByKeyInput) (types.MapOutput, error) {
	if len(input.Keys) == 0 && len(input.SpaceIds) == 0 && input.SpaceKeys == "" && len(input.SpaceId) == 0 && input.SpaceKeySingle == "" {
		return nil, fmt.Errorf("at least one space identifier parameter must be provided")
	}

	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})
	client.SetQueryParam(queryParams, "spaceKeys", input.SpaceKeys, "")
	client.SetQueryParam(queryParams, "spaceIds", strings.Join(input.SpaceIds, ","), "")
	client.SetQueryParam(queryParams, "spaceId", input.SpaceId, []string{})
	client.SetQueryParam(queryParams, "spaceKeySingle", input.SpaceKeySingle, "")
	client.SetQueryParam(queryParams, "type", input.Type, "")
	client.SetQueryParam(queryParams, "status", input.Status, "")
	client.SetQueryParam(queryParams, "label", input.Label, []string{})
	client.SetQueryParam(queryParams, "contentLabel", input.ContentLabel, []string{})
	client.SetQueryParam(queryParams, "favourite", input.Favourite, (*bool)(nil))
	client.SetQueryParam(queryParams, "hasRetentionPolicy", input.HasRetentionPolicy, (*bool)(nil))
	client.SetQueryParam(queryParams, "spaceKey", input.Keys, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "space"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
