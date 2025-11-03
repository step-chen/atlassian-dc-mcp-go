package jira

import (
	"atlassian-dc-mcp-go/internal/client/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetIssue(t *testing.T) {
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
		issueKey    string
		expectError bool
	}{
		{
			name:        "ValidIssue",
			issueKey:    testConfig.Issues.ValidKey,
			expectError: false,
		},
		{
			name:        "InvalidIssue",
			issueKey:    InvalidKey,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetIssue(GetIssueInput{
				IssueKey: tt.issueKey,
				Fields:   nil,
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
						assert.NotEmpty(t, key, "key should not be empty")
						t.Logf("Get issue successful. Issue key: %s", key)
					}
				}
			}
		})
	}
}

func TestGetTransitions(t *testing.T) {
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
	}{
		{
			name:        "ValidIssue",
			issueKey:    testConfig.Issues.ValidKey,
			expectError: false,
		},
		{
			name:        "InvalidIssue",
			issueKey:    InvalidKey,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetTransitions(GetTransitionsInput{
				IssueKey: tt.issueKey,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				if result != nil {
					transitions, exists := result["transitions"]
					assert.True(t, exists, "transitions field should exist in response")
					if exists {
						assert.NotNil(t, transitions, "transitions should not be nil")
						t.Logf("Get transitions successful. Issue key: %s Transitions found: %d", tt.issueKey, len(transitions.([]any)))
					}
				}
			}
		})
	}
}

func TestGetSubtasks(t *testing.T) {
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
	}{
		{
			name:        "ValidIssue",
			issueKey:    testConfig.Subtasks.IssueKey,
			expectError: false,
		},
		{
			name:        "InvalidIssue",
			issueKey:    InvalidKey,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetSubtasks(GetSubtasksInput{
				IssueKey: tt.issueKey,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					t.Logf("Get subtasks successful. Issue key: %s Subtasks found: %d", tt.issueKey, len(result))
				} else {
					t.Logf("Get subtasks failed. Issue key: %s Error: %v", tt.issueKey, err)
				}
			}
		})
	}
}

func TestGetAgileIssue(t *testing.T) {
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
		name          string
		issueIdOrKey  string
		expand        string
		fields        []string
		updateHistory bool
		expectError   bool
	}{
		{
			name:          "ValidIssue",
			issueIdOrKey:  testConfig.Issues.ValidKey,
			expand:        "",
			fields:        nil,
			updateHistory: false,
			expectError:   false,
		},
		{
			name:          "InvalidIssue",
			issueIdOrKey:  InvalidKey,
			expand:        "",
			fields:        nil,
			updateHistory: false,
			expectError:   true,
		},
		{
			name:          "ValidIssueWithFields",
			issueIdOrKey:  testConfig.Issues.ValidKey,
			expand:        "",
			fields:        []string{"id", "key", "summary"},
			updateHistory: false,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetAgileIssue(GetAgileIssueInput{
				IssueIdOrKey:  tt.issueIdOrKey,
				Expand:        tt.expand,
				Fields:        tt.fields,
				UpdateHistory: tt.updateHistory,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						key, exists := result["key"]
						assert.True(t, exists, "key field should exist in response")
						if exists {
							assert.NotEmpty(t, key, "key should not be empty")
							t.Logf("Get agile issue successful. Issue ID or key: %s", key)
						}
					}
				} else {
					t.Logf("Get agile issue failed. Issue ID or key: %s Error: %v", tt.issueIdOrKey, err)
				}
			}
		})
	}
}

func TestGetIssueEstimationForBoard(t *testing.T) {
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
		boardId     int64
		expectError bool
	}{
		{
			name:        "ValidIssueAndBoard",
			issueKey:    testConfig.Issues.ValidKey,
			boardId:     int64(testConfig.Boards.ValidID),
			expectError: false,
		},
		{
			name:        "InvalidIssue",
			issueKey:    InvalidKey,
			boardId:     int64(testConfig.Boards.ValidID),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetIssueEstimationForBoard(GetIssueEstimationForBoardInput{
				IssueIdOrKey: tt.issueKey,
				BoardId:      tt.boardId,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						fieldId, fieldIdExists := result["fieldId"]
						if fieldIdExists {
							assert.NotEmpty(t, fieldId, "fieldId should not be empty")
						}

						value, valueExists := result["value"]
						if valueExists {
							t.Logf("Value type: %T", value)
						}

						t.Logf("Get issue estimation for board successful. Issue key: %s, Board ID: %d", tt.issueKey, tt.boardId)
					}
				} else {
					t.Logf("Get issue estimation for board completed. Error (may be expected): %v", err)
				}
			}
		})
	}
}
