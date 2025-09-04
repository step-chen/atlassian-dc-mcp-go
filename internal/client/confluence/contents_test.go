package confluence

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atlassian-dc-mcp-go/internal/client/testutils"
)

func TestGetPage(t *testing.T) {
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
		pageID      string
		expand      []string
		expectError bool
	}{
		{
			name:        "ValidPage",
			pageID:      testConfig.Pages.ValidID,
			expand:      nil,
			expectError: false,
		},
		{
			name:        "InvalidPage",
			pageID:      testConfig.Pages.InvalidID,
			expand:      nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetContentByID(tt.pageID, tt.expand)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if result != nil {
					id, exists := result["id"]
					assert.True(t, exists, "id field should exist in response")
					if exists {
						assert.NotEmpty(t, id, "id should not be empty")
						t.Logf("Get page successful. Page ID: %s", id)
					}
				}
			}
		})
	}
}

func TestGetPageComments(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	testConfig, err := loadTestConfig()
	if err != nil {
		t.Fatalf("Could not load test configuration: %v", err)
	}

	client, err := SetupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Could not load config.yaml: %v", err)
	}

	require.NotNil(t, client, "Confluence client should not be nil")

	tests := []struct {
		name        string
		pageID      string
		expectError bool
	}{
		{
			name:        "ValidPageWithComments",
			pageID:      testConfig.Pages.CommentsID,
			expectError: false,
		},
		{
			name:        "InvalidPage",
			pageID:      "invalid",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comments, err := client.GetComments(tt.pageID, nil, 0, 0)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, comments)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, comments)

				if comments != nil {
					results, exists := comments["results"]
					assert.True(t, exists, "results field should exist in response")

					size, exists := comments["size"]
					assert.True(t, exists, "size field should exist in response")

					if exists && results != nil {
						t.Logf("Successfully retrieved %d comments", size)
					}
				}
			}
		})
	}
}
