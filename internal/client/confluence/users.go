package confluence

// GetCurrentUser retrieves details of the current user.
//
// Returns:
//   - map[string]interface{}: The current user data
//   - error: An error if the request fails
func (c *ConfluenceClient) GetCurrentUser() (map[string]interface{}, error) {
	var user map[string]interface{}
	if err := c.executeRequest("GET", []string{"rest", "api", "user", "current"}, nil, nil, &user); err != nil {
		return nil, err
	}

	return user, nil
}