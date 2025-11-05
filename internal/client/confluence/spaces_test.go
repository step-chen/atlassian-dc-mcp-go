package confluence

import (
	"atlassian-dc-mcp-go/internal/client/testutils"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			input := GetSpaceInput{
				SpaceKey: tt.spaceKey,
			}
			result, err := client.GetSpace(context.Background(), input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.NotEmpty(t, result)
			}
		})
	}

	t.Run("ValidWithExpand", func(t *testing.T) {
		input := GetSpaceInput{
			SpaceKey: testConfig.Spaces.Key,
			Expand:   []string{"description", "homepage"},
		}
		result, err := client.GetSpace(context.Background(), input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result)
	})
}
