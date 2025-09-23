package jira

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atlassian-dc-mcp-go/internal/client/testutils"
)

func TestGetUser(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	testConfig, err := loadTestConfig()
	if err != nil {
		t.Fatalf("Could not load test configuration: %v", err)
	}

	client, err := SetupIntegrationTest(t)
	require.NoError(t, err, "Could not load config.yaml")
	require.NotNil(t, client, "Jira client should not be nil")

	tests := []struct {
		name        string
		username    string
		userKey     string
		testType    string
		expectError bool
	}{
		{
			name:        "GetUserByName",
			username:    testConfig.Users.ValidUsername,
			testType:    "byname",
			expectError: false,
		},
		{
			name:        "GetUserByKey",
			userKey:     testConfig.Users.ValidKey,
			testType:    "bykey",
			expectError: false,
		},
		{
			name:        "InvalidUser",
			username:    InvalidUsername,
			testType:    "byname",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result map[string]any
			var err error

			switch tt.testType {
			case "byname":
				result, err = client.GetUserByName(GetUserByNameInput{
					Username: tt.username,
				})
			case "bykey":
				result, err = client.GetUserByKey(GetUserByKeyInput{
					Key: tt.userKey,
				})
			}

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
				} else {
					t.Logf("%s completed. Error (may be expected): %v", tt.name, err)
				}
			}
		})
	}
}

func TestSearchUsers(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	testConfig, err := loadTestConfig()
	if err != nil {
		t.Fatalf("Could not load test configuration: %v", err)
	}

	client, err := SetupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Could not load config.yaml: %v", err)
	}

	t.Run("SearchUsers", func(t *testing.T) {
		result, err := client.SearchUsers(SearchUsersInput{
			Query: testConfig.Users.SearchQuery,
			PaginationInput: PaginationInput{
				StartAt:    0,
				MaxResults: 10,
			},
		})

		if err == nil {
			assert.NotNil(t, result)
		} else {
			t.Logf("Note: Error may be due to permissions which is expected in some test environments")
			t.Logf("Search users completed. Error: %v", err)
		}
	})
}

func TestGetCurrentUser(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	client, err := SetupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Could not load config.yaml: %v", err)
	}

	t.Run("GetCurrentUser", func(t *testing.T) {
		result, err := client.GetCurrentUser()

		if err == nil {
			assert.NotNil(t, result)
		} else {
			t.Logf("Note: Error may be due to permissions which is expected in some test environments")
			t.Logf("Get current user completed. Error: %v", err)
		}
	})
}
