package confluence

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"atlassian-dc-mcp-go/internal/client/testutils"
)

func TestGetSpaces(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	testConfig, err := loadTestConfig()
	if err != nil {
		t.Fatalf("Could not load test configuration: %v", err)
	}

	client, err := SetupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Could not load config.yaml: %v", err)
	}

	tests := []struct {
		name        string
		limit       int
		start       int
		expectError bool
	}{
		{
			name:        "DefaultParams",
			limit:       testConfig.Spaces.Limit,
			start:       testConfig.Spaces.Start,
			expectError: false,
		},
		{
			name:        "CustomLimit",
			limit:       testConfig.Spaces.CustomLim,
			start:       testConfig.Spaces.Start,
			expectError: false,
		},
		{
			name:        "InvalidLimit",
			limit:       testConfig.Spaces.InvalidLimit,
			start:       0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetSpaces(tt.limit, tt.start)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if result != nil {
					sizeValue, exists := result["size"]
					assert.True(t, exists)
					if exists {
						size := int(sizeValue.(float64))
						assert.Greater(t, size, 0)
						t.Logf("Get spaces successful. Number of spaces: %d", size)
					}
				}
			}
		})
	}
}

func TestGetSpace(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	testConfig, err := loadTestConfig()
	if err != nil {
		t.Fatalf("Could not load test configuration: %v", err)
	}

	client, err := SetupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Could not load config.yaml: %v", err)
	}

	tests := []struct {
		name        string
		spaceKey    string
		expectError bool
	}{
		{
			name:        "Valid",
			spaceKey:    testConfig.Spaces.Key,
			expectError: false,
		},
		{
			name:        "NonExistent",
			spaceKey:    testConfig.Spaces.NonExistentKey,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			space, err := client.GetSpace(tt.spaceKey, nil)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, space)
				t.Logf("Expected error occurred: %v", err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, space)

				key, ok := space["key"]
				assert.True(t, ok)
				assert.Equal(t, tt.spaceKey, key)

				if ok {
					t.Logf("Successfully retrieved space: %s", key)
				}
			}
		})
	}

	t.Run("ValidWithExpand", func(t *testing.T) {
		space, err := client.GetSpace(testConfig.Spaces.Key, []string{"description", "homepage"})

		assert.NoError(t, err)
		assert.NotNil(t, space)

		key, ok := space["key"]
		assert.True(t, ok)
		assert.Equal(t, testConfig.Spaces.Key, key)

		t.Logf("Successfully retrieved space with expand: %s", key)
	})
}
