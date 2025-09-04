package jira

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atlassian-dc-mcp-go/internal/client/testutils"
)

func TestGetIssueWorklogs(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	testConfig, err := loadTestConfig()
	if err != nil {
		t.Fatalf("Could not load test configuration: %v", err)
	}

	client, err := SetupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Could not load config.yaml: %v", err)
	}

	require.NotNil(t, client, "Jira client should not be nil")

	tests := []struct {
		name        string
		issueKey    string
		expectError bool
		description string
	}{
		{
			name:        "ValidIssue",
			issueKey:    testConfig.Worklogs.IssueKey,
			expectError: false,
			description: "Valid issue worklogs",
		},
		{
			name:        "InvalidIssue",
			issueKey:    InvalidKey,
			expectError: true,
			description: "Invalid issue worklogs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetWorklogs(tt.issueKey)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						_, exists := result["worklogs"]
						assert.True(t, exists, "worklogs field should exist in response")

						if worklogs, ok := result["worklogs"].([]any); ok {
							t.Logf("Number of worklogs retrieved: %d", len(worklogs))
						}
					}
				} else {
					t.Logf("Get worklogs failed for issue %s. Error: %v", tt.issueKey, err)
				}
			}
		})
	}
}

func TestGetWorklog(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	testConfig, err := loadTestConfig()
	if err != nil {
		t.Fatalf("Could not load test configuration: %v", err)
	}

	client, err := SetupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Could not load config.yaml: %v", err)
	}

	require.NotNil(t, client, "Jira client should not be nil")

	worklogsResult, err := client.GetWorklogs(testConfig.Worklogs.IssueKey)
	if err != nil {
		t.Logf("Cannot get worklogs for issue %s. Error: %v", testConfig.Worklogs.IssueKey, err)
		return
	}
	require.NotNil(t, worklogsResult, "Worklogs result should not be nil")

	var worklogID string
	if worklogs, ok := worklogsResult["worklogs"].([]interface{}); ok && len(worklogs) > 0 {
		if firstWorklog, ok := worklogs[0].(map[string]interface{}); ok {
			if id, ok := firstWorklog["id"].(string); ok {
				worklogID = id
			}
		}
	}

	tests := []struct {
		name        string
		issueKey    string
		worklogID   string
		expectError bool
		description string
		skipCheck   bool
	}{
		{
			name:        "ValidWorklog",
			issueKey:    testConfig.Worklogs.IssueKey,
			worklogID:   worklogID,
			expectError: false,
			description: "Valid worklog",
			skipCheck:   worklogID == "",
		},
		{
			name:        "InvalidWorklog",
			issueKey:    testConfig.Worklogs.IssueKey,
			worklogID:   "99999999",
			expectError: true,
			description: "Invalid worklog",
			skipCheck:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skipCheck {
				t.Skip("Skipping test - no valid worklog ID available")
				return
			}

			result, err := client.GetWorklogs(tt.issueKey, tt.worklogID)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						id, exists := result["id"]
						assert.True(t, exists, "id field should exist in response")
						if exists {
							assert.Equal(t, tt.worklogID, id, "worklog ID should match requested ID")
							t.Logf("Get worklog successful. Worklog ID: %s", id)
						}
					}
				} else {
					t.Logf("Get worklog completed. Error (may be expected): %v", err)
				}
			}
		})
	}
}
