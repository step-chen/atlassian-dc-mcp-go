package confluence

import (
	"net/url"

	"atlassian-dc-mcp-go/internal/utils"
)

// GetRelatedLabels retrieves labels related to a specific label.
//
// Parameters:
//   - labelName: The name of the label
//   - start: Starting index for pagination
//   - limit: Maximum number of results to return
//
// Returns:
//   - map[string]interface{}: The related labels data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetRelatedLabels(labelName string, start int, limit int) (map[string]interface{}, error) {

	params := url.Values{}
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)

	var labels map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "label", labelName, "related"}, params, nil, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}

// GetLabels retrieves labels based on various filters.
//
// Parameters:
//   - labelName: The name of the label to filter by
//   - owner: The owner of the labels
//   - namespace: The namespace of the labels
//   - spaceKey: The key of the space to filter by
//   - start: Starting index for pagination
//   - limit: Maximum number of results to return
//
// Returns:
//   - map[string]interface{}: The labels data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetLabels(labelName, owner, namespace, spaceKey string, start, limit int) (map[string]interface{}, error) {

	params := url.Values{}
	utils.SetQueryParam(params, "labelName", labelName, "")
	utils.SetQueryParam(params, "owner", owner, "")
	utils.SetQueryParam(params, "namespace", namespace, "")
	utils.SetQueryParam(params, "spaceKey", spaceKey, "")
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)

	var labels map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "label"}, params, nil, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}