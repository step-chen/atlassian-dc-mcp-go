package confluence

import (
	"fmt"
	"net/url"
	"strings"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetSpaces retrieves a list of spaces.
//
// Parameters:
//   - limit: Maximum number of results to return
//   - start: Starting index for pagination
//
// Returns:
//   - map[string]interface{}: The spaces data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetSpaces(limit int, start int) (map[string]interface{}, error) {
	params := url.Values{}

	utils.SetQueryParam(params, "limit", limit, 0)
	utils.SetQueryParam(params, "start", start, 0)

	var spaces map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "space"}, params, nil, &spaces); err != nil {
		return nil, err
	}

	return spaces, nil
}

// GetSpace retrieves a specific space by its key.
//
// Parameters:
//   - spaceKey: The key of the space to retrieve
//   - expand: Fields to expand in the response
//
// Returns:
//   - map[string]interface{}: The space data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetSpace(spaceKey string, expand []string) (map[string]interface{}, error) {
	if spaceKey == "" {
		return nil, fmt.Errorf("spaceKey cannot be empty")
	}

	params := url.Values{}
	utils.SetQueryParam(params, "expand", expand, []string{})

	var space map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "space", spaceKey}, params, nil, &space); err != nil {
		return nil, err
	}

	return space, nil
}

// GetContentsInSpace retrieves contents in a specific space.
//
// Parameters:
//   - spaceKey: The key of the space
//   - start: Starting index for pagination
//   - limit: Maximum number of results to return
//   - expand: Fields to expand in the response
//
// Returns:
//   - map[string]interface{}: The contents data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentsInSpace(spaceKey string, start int, limit int, expand []string) (map[string]interface{}, error) {
	if spaceKey == "" {
		return nil, fmt.Errorf("spaceKey cannot be empty")
	}

	params := url.Values{}
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)
	utils.SetQueryParam(params, "expand", expand, []string{})

	var contents map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "space", spaceKey, "content"}, params, nil, &contents); err != nil {
		return nil, err
	}

	return contents, nil
}

// GetContentsByType retrieves contents of a specific type in a space.
//
// Parameters:
//   - spaceKey: The key of the space
//   - contentType: The type of content to retrieve
//   - start: Starting index for pagination
//   - limit: Maximum number of results to return
//   - expand: Fields to expand in the response
//
// Returns:
//   - map[string]interface{}: The contents data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetContentsByType(spaceKey, contentType string, start int, limit int, expand []string) (map[string]interface{}, error) {
	if spaceKey == "" {
		return nil, fmt.Errorf("spaceKey cannot be empty")
	}

	if contentType == "" {
		return nil, fmt.Errorf("contentType cannot be empty")
	}

	params := url.Values{}
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)
	utils.SetQueryParam(params, "expand", expand, []string{})

	var contents map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "space", spaceKey, "content", contentType}, params, nil, &contents); err != nil {
		return nil, err
	}

	return contents, nil
}

// GetSpacesByKey retrieves spaces based on various filters.
//
// Parameters:
//   - keys: Space keys to filter by
//   - start: Starting index for pagination
//   - limit: Maximum number of results to return
//   - expand: Fields to expand in the response
//   - spaceIds: Space IDs to filter by
//   - spaceKeys: Space keys string
//   - spaceId: Space ID array
//   - spaceKeySingle: Single space key
//   - typ: Type of spaces to retrieve
//   - status: Status of spaces to retrieve
//   - label: Labels to filter by
//   - contentLabel: Content labels to filter by
//   - favourite: Filter favourite spaces
//   - hasRetentionPolicy: Filter spaces with retention policy
//
// Returns:
//   - map[string]interface{}: The spaces data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetSpacesByKey(
	keys []string,
	start int,
	limit int,
	expand []string,
	spaceIds []string,
	spaceKeys string,
	spaceId []string,
	spaceKeySingle string,
	typ string,
	status string,
	label []string,
	contentLabel []string,
	favourite *bool,
	hasRetentionPolicy *bool,
) (map[string]interface{}, error) {
	if len(keys) == 0 && len(spaceIds) == 0 && spaceKeys == "" && len(spaceId) == 0 && spaceKeySingle == "" {
		return nil, fmt.Errorf("at least one space identifier parameter must be provided")
	}

	params := url.Values{}
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)
	utils.SetQueryParam(params, "expand", expand, []string{})
	utils.SetQueryParam(params, "spaceKeys", spaceKeys, "")
	utils.SetQueryParam(params, "spaceIds", strings.Join(spaceIds, ","), "")
	utils.SetQueryParam(params, "spaceId", spaceId, []string{})
	utils.SetQueryParam(params, "spaceKeySingle", spaceKeySingle, "")
	utils.SetQueryParam(params, "type", typ, "")
	utils.SetQueryParam(params, "status", status, "")
	utils.SetQueryParam(params, "label", label, []string{})
	utils.SetQueryParam(params, "contentLabel", contentLabel, []string{})
	utils.SetQueryParam(params, "favourite", favourite, (*bool)(nil))
	utils.SetQueryParam(params, "hasRetentionPolicy", hasRetentionPolicy, (*bool)(nil))
	utils.SetQueryParam(params, "spaceKey", keys, []string{})

	var spaces map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "space"}, params, nil, &spaces); err != nil {
		return nil, err
	}

	return spaces, nil
}