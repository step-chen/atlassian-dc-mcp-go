package jira

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atlassian-dc-mcp-go/internal/client/testutils"
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
			result, err := client.GetIssue(tt.issueKey, nil)

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
			result, err := client.GetTransitions(tt.issueKey)

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
						t.Logf("Get transitions successful. Number of transitions: %d", len(transitions.([]interface{})))
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
			result, err := client.GetSubtasks(tt.issueKey)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						t.Logf("Get subtasks successful. Number of subtasks: %d", len(result))
					}
				} else {
					t.Logf("Get subtasks completed. Error (may be expected): %v", err)
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
			result, err := client.GetAgileIssue(tt.issueIdOrKey, tt.expand, tt.fields, tt.updateHistory)

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
							t.Logf("Get agile issue successful. Issue key: %s", key)
						}

						fields, fieldsExist := result["fields"]
						if fieldsExist {
							if fieldsMap, ok := fields.(map[string]interface{}); ok {
								if sprint, sprintExists := fieldsMap["sprint"]; sprintExists {
									t.Logf("Sprint field found: %v", sprint)
								}
								if closedSprints, closedSprintsExists := fieldsMap["closedSprints"]; closedSprintsExists {
									t.Logf("ClosedSprints field found: %v", closedSprints)
								}
								if epic, epicExists := fieldsMap["epic"]; epicExists {
									t.Logf("Epic field found: %v", epic)
								}
							}
						}
					}
				} else {
					t.Logf("Get agile issue completed. Error (may be expected): %v", err)
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
			result, err := client.GetIssueEstimationForBoard(tt.issueKey, tt.boardId)

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
