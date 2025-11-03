package confluence

import (
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
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
	client.SetQueryParam(params, "start", input.Start, 0)
	client.SetQueryParam(params, "limit", input.Limit, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "label", input.LabelName, "related"},
		params,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
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
	client.SetQueryParam(params, "labelName", input.LabelName, "")
	client.SetQueryParam(params, "owner", input.Owner, "")
	client.SetQueryParam(params, "namespace", input.Namespace, "")
	client.SetQueryParam(params, "spaceKey", input.SpaceKey, "")
	client.SetQueryParam(params, "start", input.Start, 0)
	client.SetQueryParam(params, "limit", input.Limit, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "label"},
		params,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
