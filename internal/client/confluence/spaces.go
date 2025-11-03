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

	params := url.Values{}
	client.SetQueryParam(params, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "space", input.SpaceKey},
		params,
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

	params := url.Values{}
	client.SetQueryParam(params, "start", input.Start, 0)
	client.SetQueryParam(params, "limit", input.Limit, 0)
	client.SetQueryParam(params, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "space", input.SpaceKey, "content"},
		params,
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

	params := url.Values{}
	client.SetQueryParam(params, "start", input.Start, 0)
	client.SetQueryParam(params, "limit", input.Limit, 0)
	client.SetQueryParam(params, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "space", input.SpaceKey, "content", input.ContentType},
		params,
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

	params := url.Values{}
	client.SetQueryParam(params, "start", input.Start, 0)
	client.SetQueryParam(params, "limit", input.Limit, 0)
	client.SetQueryParam(params, "expand", input.Expand, []string{})
	client.SetQueryParam(params, "spaceKeys", input.SpaceKeys, "")
	client.SetQueryParam(params, "spaceIds", strings.Join(input.SpaceIds, ","), "")
	client.SetQueryParam(params, "spaceId", input.SpaceId, []string{})
	client.SetQueryParam(params, "spaceKeySingle", input.SpaceKeySingle, "")
	client.SetQueryParam(params, "type", input.Type, "")
	client.SetQueryParam(params, "status", input.Status, "")
	client.SetQueryParam(params, "label", input.Label, []string{})
	client.SetQueryParam(params, "contentLabel", input.ContentLabel, []string{})
	client.SetQueryParam(params, "favourite", input.Favourite, (*bool)(nil))
	client.SetQueryParam(params, "hasRetentionPolicy", input.HasRetentionPolicy, (*bool)(nil))
	client.SetQueryParam(params, "spaceKey", input.Keys, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "space"},
		params,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
