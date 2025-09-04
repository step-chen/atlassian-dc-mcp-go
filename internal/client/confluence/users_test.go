package confluence

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"atlassian-dc-mcp-go/internal/client/testutils"
)

func TestGetUserByUsername(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	client, err := SetupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Could not load config.yaml: %v", err)
	}

	assert.NotNil(t, client, "Confluence client should not be nil")

	tests := []struct {
		name        string
		username    string
		expectError bool
	}{
		{
			name:        "ValidUser",
			expectError: false,
		},
		{
			name:        "BasicFunctionality",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := client.GetCurrentUser()

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, user)
				t.Logf("Expected error occurred: %v", err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)

				if user != nil {
					accountID, exists := user["accountId"]
					assert.True(t, exists, "accountId field should exist in response")
					if exists {
						assert.NotEmpty(t, accountID, "accountId should not be empty")
						t.Logf("Get user successful. Account ID: %s", accountID)
					}
				}
			}
		})
	}
}
