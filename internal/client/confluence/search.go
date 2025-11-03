package confluence

import (
	"net/http"
	"net/url"
	"strconv"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// Search searches for content based on CQL.
//
// Parameters:
//   - input: SearchInput containing the parameters for the request
//
// Returns:
//   - types.MapOutput: The search results
//   - error: An error if the request fails
func (c *ConfluenceClient) Search(input SearchInput) (types.MapOutput, error) {
	params := url.Values{}
	client.SetQueryParam(params, "cql", input.CQL, "")
	client.SetQueryParam(params, "cqlcontext", input.CQLContext, "")
	client.SetQueryParam(params, "excerpt", input.Excerpt, "")
	client.SetQueryParam(params, "start", input.Start, 0)
	client.SetQueryParam(params, "limit", input.Limit, 0)
	client.SetQueryParam(params, "includeArchivedSpaces", strconv.FormatBool(input.IncludeArchivedSpaces), "")
	client.SetQueryParam(params, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "search"},
		params,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
