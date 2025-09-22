package confluence

import (
	"net/url"
	"strconv"

	"atlassian-dc-mcp-go/internal/utils"
)

// Search searches for content based on CQL.
//
// Parameters:
//   - input: SearchInput containing the parameters for the request
//
// Returns:
//   - map[string]interface{}: The search results
//   - error: An error if the request fails
func (c *ConfluenceClient) Search(input SearchInput) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "cql", input.CQL, "")
	utils.SetQueryParam(params, "cqlcontext", input.CQLContext, "")
	utils.SetQueryParam(params, "excerpt", input.Excerpt, "")
	utils.SetQueryParam(params, "start", input.Start, 0)
	utils.SetQueryParam(params, "limit", input.Limit, 0)
	utils.SetQueryParam(params, "includeArchivedSpaces", strconv.FormatBool(input.IncludeArchivedSpaces), "")
	utils.SetQueryParam(params, "expand", input.Expand, []string{})

	var result map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "search"}, params, nil, &result); err != nil {
		return nil, err
	}

	return result, nil
}
