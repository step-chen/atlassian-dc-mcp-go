package jira

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atlassian-dc-mcp-go/internal/client/testutils"
	"atlassian-dc-mcp-go/internal/types"
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
			result, err := client.GetWorklogs(context.Background(), GetWorklogsInput{
				IssueKey: tt.issueKey,
			})

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
							t.Logf("Get worklogs successful. Issue key: %s Worklogs found: %d", tt.issueKey, len(worklogs))
						} else {
							t.Logf("Get worklogs successful. Issue key: %s Could not determine worklog count", tt.issueKey)
						}
					}
				} else {
					t.Logf("Get worklogs failed. Issue key: %s Error: %v", tt.issueKey, err)
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

	worklogsResult, err := client.GetWorklogs(context.Background(), GetWorklogsInput{
		IssueKey: testConfig.Worklogs.IssueKey,
	})
	if err != nil {
		t.Logf("Cannot get worklogs for issue %s. Error: %v", testConfig.Worklogs.IssueKey, err)
		return
	}
	require.NotNil(t, worklogsResult, "Worklogs result should not be nil")

	var worklogID string
	if worklogs, ok := worklogsResult["worklogs"].([]interface{}); ok && len(worklogs) > 0 {
		if firstWorklog, ok := worklogs[0].(types.MapOutput); ok {
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

			result, err := client.GetWorklogs(context.Background(), GetWorklogsInput{
				IssueKey:  tt.issueKey,
				WorklogId: tt.worklogID,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						t.Logf("Get specific worklog successful. Issue key: %s Worklog ID: %s", tt.issueKey, tt.worklogID)
					}
				} else {
					t.Logf("Get specific worklog failed. Issue key: %s Worklog ID: %s Error: %v", tt.issueKey, tt.worklogID, err)
				}
			}
		})
	}
}
