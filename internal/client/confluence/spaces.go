package confluence

import (
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
//   - map[string]interface{}: The space data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetSpace(input GetSpaceInput) (map[string]interface{}, error) {
	if input.SpaceKey == "" {
		return nil, fmt.Errorf("spaceKey cannot be empty")
	}

	params := url.Values{}
	utils.SetQueryParam(params, "expand", input.Expand, []string{})

	var space map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "space", input.SpaceKey}, params, nil, &space); err != nil {
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
//   - map[string]interface{}: The contents data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentsInSpace(input GetContentsInSpaceInput) (map[string]interface{}, error) {
	if input.SpaceKey == "" {
		return nil, fmt.Errorf("spaceKey cannot be empty")
	}

	params := url.Values{}
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)
	utils.SetQueryParam(params, "expand", input.Expand, []string{})

	var contents map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "space", input.SpaceKey, "content"}, params, nil, &contents); err != nil {
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
//   - map[string]interface{}: The contents data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentsByType(input GetContentsByTypeInput) (map[string]interface{}, error) {
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

	var contents map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "space", input.SpaceKey, "content", input.ContentType}, params, nil, &contents); err != nil {
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
//   - map[string]interface{}: The spaces data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetSpacesByKey(input GetSpacesByKeyInput) (map[string]interface{}, error) {
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

	var spaces map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "space"}, params, nil, &spaces); err != nil {
		return nil, err
	}

	return spaces, nil
}
