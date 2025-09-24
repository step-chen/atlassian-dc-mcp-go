package confluence

import "atlassian-dc-mcp-go/internal/types"

// GetCurrentUser retrieves details of the current user.
//
// Parameters:
//   - input: Empty input struct (no parameters required)
//
// Returns:
//   - types.MapOutput: The current user data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetCurrentUser() (types.MapOutput, error) {
	var user types.MapOutput
	if err := c.executeRequest("GET", []string{"rest", "api", "user", "current"}, nil, nil, &user); err != nil {
		return nil, err
	}

	return user, nil
}
