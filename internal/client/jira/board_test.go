package jira

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atlassian-dc-mcp-go/internal/client/testutils"
	"atlassian-dc-mcp-go/internal/types"
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
			result, err := client.GetBoards(GetBoardsInput{
				PaginationInput: PaginationInput{
					StartAt:    tt.startAt,
					MaxResults: tt.maxResults,
				},
				Name:           tt.boardName,
				ProjectKeyOrId: tt.projectKeyOrId,
				BoardType:      tt.boardType,
			})

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
							if firstBoard, ok := boardsList[0].(types.MapOutput); ok {
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
			result, err := client.GetBoardBacklog(GetBoardBacklogInput{
				BoardId: tt.boardID,
				PaginationInput: PaginationInput{
					StartAt:    tt.startAt,
					MaxResults: tt.maxResults,
				},
				JQL:           tt.jql,
				ValidateQuery: tt.validateQuery,
				Fields:        tt.fields,
				Expand:        tt.expand,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if result != nil {
					issues, exists := result["issues"]
					assert.True(t, exists, "issues field should exist in response")
					if exists {
						assert.NotNil(t, issues, "issues should not be nil")
						t.Logf("Get board backlog successful. Board ID: %d Issues found: %d", tt.boardID, len(issues.([]any)))
					}
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
			result, err := client.GetBoard(GetBoardInput{
				Id: tt.boardID,
			})

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
							assert.NotNil(t, id, "id should not be nil")
							t.Logf("Get board successful. Board ID: %v", id)
						}
					}
				} else {
					t.Logf("Get board failed. Board ID: %d Error: %v", tt.boardID, err)
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
			result, err := client.GetBoardEpics(GetBoardEpicsInput{
				BoardId: tt.boardID,
				PaginationInput: PaginationInput{
					StartAt:    tt.startAt,
					MaxResults: tt.maxResults,
				},
				Done: tt.done,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if result != nil {
					values, exists := result["values"]
					assert.True(t, exists, "values field should exist in response")
					if exists {
						assert.NotNil(t, values, "values should not be nil")
						t.Logf("Get board epics successful. Board ID: %d Epics found: %d", tt.boardID, len(values.([]any)))
					}
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
			result, err := client.GetSprintIssues(GetSprintIssuesInput{
				SprintId: tt.sprintID,
				PaginationInput: PaginationInput{
					StartAt:    tt.startAt,
					MaxResults: tt.maxResults,
				},
				JQL:           tt.jql,
				ValidateQuery: true,
				Fields:        tt.fields,
				Expand:        tt.expand,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if result != nil {
					issues, exists := result["issues"]
					assert.True(t, exists, "issues field should exist in response")
					if exists {
						assert.NotNil(t, issues, "issues should not be nil")
						t.Logf("Get sprint issues successful. Sprint ID: %d Issues found: %d", tt.sprintID, len(issues.([]any)))
					}
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
			result, err := client.GetSprint(GetSprintInput{
				SprintId: tt.sprintID,
			})

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
							assert.NotNil(t, id, "id should not be nil")
							t.Logf("Get sprint successful. Sprint ID: %v", id)
						}
					}
				} else {
					t.Logf("Get sprint failed. Sprint ID: %d Error: %v", tt.sprintID, err)
				}
			}
		})
	}
}
