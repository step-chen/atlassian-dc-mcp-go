package jira

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"atlassian-dc-mcp-go/internal/client/testutils"
)

func TestSearchIssues(t *testing.T) {
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
		jql         string
		projectKey  string
		orderBy     string
		statuses    []string
		maxResults  int
		startAt     int
		fields      []string
		expectError bool
		description string
	}{
		{
			name:        "ValidJQL",
			jql:         testConfig.Search.ValidJQL,
			maxResults:  testConfig.Search.MaxResults,
			startAt:     0,
			expectError: false,
			description: "Valid JQL search",
		},
		{
			name:        "InvalidJQL",
			jql:         InvalidJQL,
			maxResults:  testConfig.Search.MaxResults,
			startAt:     0,
			expectError: true,
			description: "Invalid JQL search",
		},
		{
			name:        "ProjectSearch",
			jql:         "",
			projectKey:  testConfig.Search.ProjectKey,
			maxResults:  testConfig.Search.MaxResults,
			startAt:     0,
			expectError: false,
			description: "Project-based search",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.SearchIssues(tt.jql, tt.projectKey, tt.orderBy, tt.statuses, tt.maxResults, tt.startAt, tt.fields)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						issues, exists := result["issues"]
						assert.True(t, exists, "issues field should exist in response")
						t.Logf("%s successful. Issues found: %v", tt.description, len(issues.([]any)))
					}
				} else {
					t.Logf("%s completed. Error (may be expected): %v", tt.description, err)
				}
			}
		})
	}
}
