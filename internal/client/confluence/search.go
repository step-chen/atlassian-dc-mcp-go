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
	queryParams := url.Values{}
	client.SetQueryParam(queryParams, "cql", input.CQL, "")
	client.SetQueryParam(queryParams, "cqlcontext", input.CQLContext, "")
	client.SetQueryParam(queryParams, "excerpt", input.Excerpt, "")
	client.SetQueryParam(queryParams, "start", input.Start, 0)
	client.SetQueryParam(queryParams, "limit", input.Limit, 0)
	client.SetQueryParam(queryParams, "includeArchivedSpaces", strconv.FormatBool(input.IncludeArchivedSpaces), "")
	client.SetQueryParam(queryParams, "expand", input.Expand, []string{})

	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]string{"rest", "api", "search"},
		queryParams,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
