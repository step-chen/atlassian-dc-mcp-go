package jira

import (
	"atlassian-dc-mcp-go/internal/client/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetProject(t *testing.T) {
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
		projectKey  string
		expectError bool
	}{
		{
			name:        "ValidProject",
			projectKey:  testConfig.Projects.ValidKey,
			expectError: false,
		},
		{
			name:        "InvalidProject",
			projectKey:  InvalidKey,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetProject(GetProjectInput{
				ProjectKey: tt.projectKey,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if result != nil {
					key, exists := result["key"]
					assert.True(t, exists, "key field should exist in response")
					if exists {
						assert.Equal(t, tt.projectKey, key, "project key should match")
						t.Logf("Get project successful. Project key: %s", key)
					}
				}
			}
		})
	}
}

func TestGetAllProjects(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	client, err := SetupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Could not load config.yaml: %v", err)
	}

	require.NotNil(t, client, "Jira client should not be nil")

	t.Run("GetAllProjects", func(t *testing.T) {
		result, err := client.GetAllProjects(GetAllProjectsInput{
			Expand:          "",
			Recent:          0,
			IncludeArchived: false,
			BrowseArchive:   false,
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		if result != nil {
			assert.Greater(t, len(result), 0, "should have at least one project")
			if len(result) > 0 {
				key, exists := result[0]["key"]
				assert.True(t, exists, "key field should exist in first project response")
				assert.NotEmpty(t, key, "key should not be empty")
				t.Logf("Get all projects successful. Projects found: %d First project key: %s", len(result), key)
			}
		}
	})
}
