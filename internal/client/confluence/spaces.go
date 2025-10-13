package confluence

import (
	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
	"fmt"
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
	utils.SetQueryParam(params, "expand", input.Expand, []string{})

	var space types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "space", input.SpaceKey}, params, nil, &space, utils.AcceptJSON); err != nil {
		return nil, err
	}

	return space, nil
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
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)
	utils.SetQueryParam(params, "expand", input.Expand, []string{})

	var contents types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "space", input.SpaceKey, "content"}, params, nil, &contents, utils.AcceptJSON); err != nil {
		return nil, err
	}

	return contents, nil
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
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)
	utils.SetQueryParam(params, "expand", input.Expand, []string{})

	var contents types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "space", input.SpaceKey, "content", input.ContentType}, params, nil, &contents, utils.AcceptJSON); err != nil {
		return nil, err
	}

	return contents, nil
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
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)
	utils.SetQueryParam(params, "expand", input.Expand, []string{})
	utils.SetQueryParam(params, "spaceKeys", input.SpaceKeys, "")
	utils.SetQueryParam(params, "spaceIds", strings.Join(input.SpaceIds, ","), "")
	utils.SetQueryParam(params, "spaceId", input.SpaceId, []string{})
	utils.SetQueryParam(params, "spaceKeySingle", input.SpaceKeySingle, "")
	utils.SetQueryParam(params, "type", input.Type, "")
	utils.SetQueryParam(params, "status", input.Status, "")
	utils.SetQueryParam(params, "label", input.Label, []string{})
	utils.SetQueryParam(params, "contentLabel", input.ContentLabel, []string{})
	utils.SetQueryParam(params, "favourite", input.Favourite, (*bool)(nil))
	utils.SetQueryParam(params, "hasRetentionPolicy", input.HasRetentionPolicy, (*bool)(nil))
	utils.SetQueryParam(params, "spaceKey", input.Keys, []string{})

	var spaces types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "space"}, params, nil, &spaces, utils.AcceptJSON); err != nil {
		return nil, err
	}

	return spaces, nil
}
