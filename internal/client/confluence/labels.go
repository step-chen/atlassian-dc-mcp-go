package confluence

import (
	"net/url"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
)

// GetRelatedLabels retrieves labels related to a specific label.
//
// Parameters:
//   - input: GetRelatedLabelsInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The related labels data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetRelatedLabels(input GetRelatedLabelsInput) (types.MapOutput, error) {

	params := url.Values{}
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)

	var labels types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "label", input.LabelName, "related"}, params, nil, &labels, utils.AcceptJSON); err != nil {
		return nil, err
	}

	return labels, nil
}

// GetLabels retrieves labels based on various filters.
//
// Parameters:
//   - input: GetLabelsInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The labels data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetLabels(input GetLabelsInput) (types.MapOutput, error) {

	params := url.Values{}
	utils.SetQueryParam(params, "labelName", input.LabelName, "")
	utils.SetQueryParam(params, "owner", input.Owner, "")
	utils.SetQueryParam(params, "namespace", input.Namespace, "")
	utils.SetQueryParam(params, "spaceKey", input.SpaceKey, "")
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)

	var labels types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "label"}, params, nil, &labels, utils.AcceptJSON); err != nil {
		return nil, err
	}

	return labels, nil
}
