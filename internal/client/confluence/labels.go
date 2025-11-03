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

	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "label", input.LabelName, "related"},
		queryParams,
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

	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "labelName", input.LabelName, "")
	client.SetQueryParam(queryParams, "owner", input.Owner, "")
	client.SetQueryParam(queryParams, "namespace", input.Namespace, "")
	client.SetQueryParam(queryParams, "spaceKey", input.SpaceKey, "")
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "label"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
