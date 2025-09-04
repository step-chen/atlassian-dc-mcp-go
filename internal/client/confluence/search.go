package confluence

import (
	"net/url"
	"strconv"

	"atlassian-dc-mcp-go/internal/utils"
)

// Search searches for content based on CQL.
//
// Parameters:
//   - cql: The CQL query string
//   - cqlcontext: The context for the CQL query
//   - excerpt: The excerpt format
//   - expand: Fields to expand in the response
//   - start: Starting index for pagination
//   - limit: Maximum number of results to return
//   - includeArchivedSpaces: Whether to include archived spaces in the search
//
// Returns:
//   - map[string]interface{}: The search results
//   - error: An error if the request fails
func (c *ConfluenceClient) Search(cql, cqlcontext, excerpt string, expand []string, start, limit int, includeArchivedSpaces bool) (map[string]interface{}, error) {
	params := url.Values{}
	utils.SetQueryParam(params, "cql", cql, "")
	utils.SetQueryParam(params, "cqlcontext", cqlcontext, "")
	utils.SetQueryParam(params, "excerpt", excerpt, "")
	utils.SetQueryParam(params, "start", start, 0)
	utils.SetQueryParam(params, "limit", limit, 0)
	utils.SetQueryParam(params, "includeArchivedSpaces", strconv.FormatBool(includeArchivedSpaces), "")
	utils.SetQueryParam(params, "expand", expand, []string{})

	var result map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "search"}, params, nil, &result); err != nil {
		return nil, err
	}

	return result, nil
}