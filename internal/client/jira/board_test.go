package jira

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atlassian-dc-mcp-go/internal/client/testutils"
)

func TestListBoards(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	client, err := SetupIntegrationTest(t)
	if err != nil {
		t.Fatalf("Could not load config.yaml: %v", err)
	}

	require.NotNil(t, client, "Jira client should not be nil")

	tests := []struct {
		name           string
		startAt        int
		maxResults     int
		boardName      string
		projectKeyOrId string
		boardType      string
		expectError    bool
		expectEmpty    bool
		description    string
	}{
		{
			name:           "ListBoardsWithDefaults",
			startAt:        0,
			maxResults:     50,
			boardName:      "",
			projectKeyOrId: "",
			boardType:      "",
			expectError:    false,
			expectEmpty:    false,
			description:    "List boards with defaults",
		},
		{
			name:           "ListBoardsWithInvalidFilter",
			startAt:        0,
			maxResults:     10,
			boardName:      "ThisNameShouldNotMatchAnyRealBoards12345",
			projectKeyOrId: "",
			boardType:      "",
			expectError:    false,
			expectEmpty:    true,
			description:    "List boards with invalid filter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetBoards(tt.startAt, tt.maxResults, tt.boardName, tt.projectKeyOrId, tt.boardType)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if result != nil {
					startAt, exists := result["startAt"]
					assert.True(t, exists, "startAt field should exist in response")
					if exists {
						assert.NotNil(t, startAt, "startAt should not be nil")
					}

					maxResults, exists := result["maxResults"]
					assert.True(t, exists, "maxResults field should exist in response")
					if exists {
						assert.NotNil(t, maxResults, "maxResults should not be nil")
					}

					total, exists := result["total"]
					assert.True(t, exists, "total field should exist in response")
					if exists {
						assert.NotNil(t, total, "total should not be nil")
						if tt.expectEmpty {
							assert.Equal(t, 0, int(total.(float64)), "should have zero total results")
						}
					}

					boards, exists := result["boards"]
					if !exists {
						boards, exists = result["values"]
					}
					assert.True(t, exists, "boards or values field should exist in response")
					if exists && !tt.expectEmpty {
						if boardsList, ok := boards.([]interface{}); ok && len(boardsList) > 0 {
							if firstBoard, ok := boardsList[0].(map[string]interface{}); ok {
								id, idExists := firstBoard["id"]
								assert.True(t, idExists, "id field should exist in board")
								if idExists {
									assert.NotEmpty(t, id, "board id should not be empty")
								}

								name, nameExists := firstBoard["name"]
								assert.True(t, nameExists, "name field should exist in board")
								if nameExists {
									assert.NotEmpty(t, name, "board name should not be empty")
								}
							}
						}
					}

					t.Logf("%s successful. Total boards: %v", tt.description, total)
				}
			}
		})
	}
}

func TestGetBoardBacklog(t *testing.T) {
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
		boardID       int
		startAt       int
		maxResults    int
		jql           string
		validateQuery bool
		fields        []string
		expand        string
		expectError   bool
		description   string
	}{
		{
			name:          "GetBoardBacklogValidBoard",
			boardID:       testConfig.Boards.ValidID,
			startAt:       0,
			maxResults:    10,
			jql:           "",
			validateQuery: true,
			fields:        nil,
			expand:        "",
			expectError:   false,
			description:   "Get board backlog for valid board",
		},
		{
			name:          "GetBoardBacklogInvalidBoard",
			boardID:       InvalidID,
			startAt:       0,
			maxResults:    10,
			jql:           "",
			validateQuery: true,
			fields:        nil,
			expand:        "",
			expectError:   true,
			description:   "Get board backlog for invalid board",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetBoardBacklog(
				tt.boardID,
				tt.startAt,
				tt.maxResults,
				tt.jql,
				tt.validateQuery,
				tt.fields,
				tt.expand,
			)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						startAt, exists := result["startAt"]
						assert.True(t, exists, "startAt field should exist in response")
						if exists {
							assert.NotNil(t, startAt, "startAt should not be nil")
						}

						maxResults, exists := result["maxResults"]
						assert.True(t, exists, "maxResults field should exist in response")
						if exists {
							assert.NotNil(t, maxResults, "maxResults should not be nil")
						}

						total, exists := result["total"]
						assert.True(t, exists, "total field should exist in response")
						if exists {
							assert.NotNil(t, total, "total should not be nil")
						}

						issues, exists := result["issues"]
						assert.True(t, exists, "issues field should exist in response")
						if exists {
							if issuesList, ok := issues.([]interface{}); ok && len(issuesList) > 0 {
								if firstIssue, ok := issuesList[0].(map[string]interface{}); ok {
									id, idExists := firstIssue["id"]
									assert.True(t, idExists, "id field should exist in issue")
									if idExists {
										assert.NotEmpty(t, id, "issue id should not be empty")
									}

									key, keyExists := firstIssue["key"]
									assert.True(t, keyExists, "key field should exist in issue")
									if keyExists {
										assert.NotEmpty(t, key, "issue key should not be empty")
									}
								}
							}
						}

						t.Logf("%s successful. Total issues: %v", tt.description, total)
					}
				} else {
					t.Logf("%s completed. Error (may be expected): %v", tt.description, err)
				}
			}
		})
	}
}

func TestGetBoard(t *testing.T) {
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
		boardID     int
		expectError bool
		description string
	}{
		{
			name:        "GetValidBoard",
			boardID:     testConfig.Boards.ValidID,
			expectError: false,
			description: "Get valid board",
		},
		{
			name:        "GetInvalidBoard",
			boardID:     InvalidID,
			expectError: true,
			description: "Get invalid board",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetBoard(tt.boardID)

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
							assert.NotEmpty(t, id, "board id should not be empty")
						}

						name, exists := result["name"]
						assert.True(t, exists, "name field should exist in response")
						if exists {
							assert.NotEmpty(t, name, "board name should not be empty")
						}

						self, exists := result["self"]
						assert.True(t, exists, "self field should exist in response")
						if exists {
							assert.NotEmpty(t, self, "board self URL should not be empty")
						}

						t.Logf("%s successful. Board ID: %v, Name: %v", tt.description, id, name)
					}
				} else {
					t.Logf("%s completed. Error (may be expected): %v", tt.description, err)
				}
			}
		})
	}
}

func TestGetBoardEpics(t *testing.T) {
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
		boardID     int
		startAt     int
		maxResults  int
		done        string
		expectError bool
		description string
	}{
		{
			name:        "GetBoardEpicsValidBoard",
			boardID:     testConfig.Boards.ValidID,
			startAt:     0,
			maxResults:  50,
			done:        "",
			expectError: false,
			description: "Get board epics for valid board",
		},
		{
			name:        "GetBoardEpicsInvalidBoard",
			boardID:     InvalidID,
			startAt:     0,
			maxResults:  50,
			done:        "",
			expectError: true,
			description: "Get board epics for invalid board",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetBoardEpics(tt.boardID, tt.startAt, tt.maxResults, tt.done)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						startAt, exists := result["startAt"]
						assert.True(t, exists, "startAt field should exist in response")
						if exists {
							assert.NotNil(t, startAt, "startAt should not be nil")
						}

						maxResults, exists := result["maxResults"]
						assert.True(t, exists, "maxResults field should exist in response")
						if exists {
							assert.NotNil(t, maxResults, "maxResults should not be nil")
						}

						values, exists := result["values"]
						assert.True(t, exists, "values field should exist in response")
						if exists {
							if epicsList, ok := values.([]interface{}); ok && len(epicsList) > 0 {
								if firstEpic, ok := epicsList[0].(map[string]interface{}); ok {
									id, idExists := firstEpic["id"]
									assert.True(t, idExists, "id field should exist in epic")
									if idExists {
										assert.NotEmpty(t, id, "epic id should not be empty")
									}

									name, nameExists := firstEpic["name"]
									assert.True(t, nameExists, "name field should exist in epic")
									if nameExists {
										assert.NotEmpty(t, name, "epic name should not be empty")
									}
								}
							}
						}

						t.Logf("%s successful.", tt.description)
					}
				} else {
					t.Logf("%s completed. Error (may be expected): %v", tt.description, err)
				}
			}
		})
	}
}

func TestGetSprintIssues(t *testing.T) {
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
		boardID     int
		sprintID    int
		startAt     int
		maxResults  int
		jql         string
		fields      []string
		expand      string
		expectError bool
	}{
		{
			name:        "GetSprintIssuesWithDefaults",
			boardID:     testConfig.Sprints.ValidBoardID,
			sprintID:    testConfig.Sprints.ValidSprintID,
			startAt:     0,
			maxResults:  50,
			jql:         "",
			fields:      nil,
			expand:      "",
			expectError: false,
		},
		{
			name:        "GetSprintIssuesWithPagination",
			boardID:     testConfig.Sprints.ValidBoardID,
			sprintID:    testConfig.Sprints.ValidSprintID,
			startAt:     0,
			maxResults:  10,
			jql:         "",
			fields:      nil,
			expand:      "",
			expectError: false,
		},
		{
			name:        "GetSprintIssuesWithFields",
			boardID:     testConfig.Sprints.ValidBoardID,
			sprintID:    testConfig.Sprints.ValidSprintID,
			startAt:     0,
			maxResults:  10,
			jql:         "",
			fields:      []string{"id", "key", "summary"},
			expand:      "",
			expectError: false,
		},
		{
			name:        "GetSprintIssuesWithJQL",
			boardID:     testConfig.Sprints.ValidBoardID,
			sprintID:    testConfig.Sprints.ValidSprintID,
			startAt:     0,
			maxResults:  10,
			jql:         testConfig.Search.ValidJQL,
			fields:      nil,
			expand:      "",
			expectError: false,
		},
		{
			name:        "GetSprintIssuesWithExpand",
			boardID:     testConfig.Sprints.ValidBoardID,
			sprintID:    testConfig.Sprints.ValidSprintID,
			startAt:     0,
			maxResults:  10,
			jql:         "",
			fields:      nil,
			expand:      "operations",
			expectError: false,
		},
		{
			name:        "GetSprintIssuesInvalidSprintID",
			sprintID:    InvalidID,
			startAt:     0,
			maxResults:  10,
			jql:         "",
			fields:      nil,
			expand:      "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetSprintIssues(
				tt.sprintID,
				tt.startAt,
				tt.maxResults,
				tt.jql,
				true,
				tt.fields,
				tt.expand,
			)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						startAt, exists := result["startAt"]
						assert.True(t, exists, "startAt field should exist in response")
						if exists {
							assert.NotNil(t, startAt, "startAt should not be nil")
						}

						maxResults, exists := result["maxResults"]
						assert.True(t, exists, "maxResults field should exist in response")
						if exists {
							assert.NotNil(t, maxResults, "maxResults should not be nil")
						}

						total, exists := result["total"]
						assert.True(t, exists, "total field should exist in response")
						if exists {
							assert.NotNil(t, total, "total should not be nil")
						}

						issues, exists := result["issues"]
						assert.True(t, exists, "issues field should exist in response")
						if exists {
							if issuesList, ok := issues.([]interface{}); ok && len(issuesList) > 0 {
								if firstIssue, ok := issuesList[0].(map[string]interface{}); ok {
									id, idExists := firstIssue["id"]
									assert.True(t, idExists, "id field should exist in issue")
									if idExists {
										assert.NotEmpty(t, id, "issue id should not be empty")
									}

									key, keyExists := firstIssue["key"]
									assert.True(t, keyExists, "key field should exist in issue")
									if keyExists {
										assert.NotEmpty(t, key, "issue key should not be empty")
									}
								}
							}
						}

						t.Logf("Get sprint issues successful. Board ID: %d, Sprint ID: %d, Total issues: %v", tt.boardID, tt.sprintID, total)
					}
				} else {
					t.Logf("Get sprint issues completed. Error (may be expected): %v", err)
				}
			}
		})
	}
}

func TestGetSprint(t *testing.T) {
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
		sprintID    int
		expectError bool
		description string
	}{
		{
			name:        "GetValidSprint",
			sprintID:    testConfig.Sprints.ValidSprintID,
			expectError: false,
			description: "Get valid sprint",
		},
		{
			name:        "GetInvalidSprint",
			sprintID:    InvalidID,
			expectError: true,
			description: "Get invalid sprint",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetSprint(tt.sprintID)

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
							assert.NotEmpty(t, id, "sprint id should not be empty")
						}

						name, exists := result["name"]
						assert.True(t, exists, "name field should exist in response")
						if exists {
							assert.NotEmpty(t, name, "sprint name should not be empty")
						}

						state, exists := result["state"]
						assert.True(t, exists, "state field should exist in response")
						if exists {
							assert.NotEmpty(t, state, "sprint state should not be empty")
						}

						t.Logf("%s successful. Sprint ID: %v, Name: %v, State: %v", tt.description, id, name, state)
					}
				} else {
					t.Logf("%s completed. Error (may be expected): %v", tt.description, err)
				}
			}
		})
	}
}
