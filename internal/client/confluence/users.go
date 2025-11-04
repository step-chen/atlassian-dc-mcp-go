package confluence

import (
	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
	"net/http"
)

// GetCurrentUser retrieves details of the current user.
//
// Parameters:
//   - input: Empty input struct (no parameters required)
//
// Returns:
//   - types.MapOutput: The current user data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetCurrentUser() (types.MapOutput, error) {
	var output types.MapOutput
	if err := client.ExecuteRequest(
		c.BaseClient,
		http.MethodGet,
		[]any{"rest", "api", "user", "current"},
		nil,
		nil,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}
