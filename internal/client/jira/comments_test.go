package jira

import (
	"atlassian-dc-mcp-go/internal/client/testutils"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetComments(t *testing.T) {
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
		startAt     int
		maxResults  int
		orderBy     string
		expand      string
		expectError bool
		description string
	}{
		{
			name:        "ValidIssue",
			issueKey:    testConfig.Comments.IssueKey,
			startAt:     0,
			maxResults:  0,
			orderBy:     "",
			expand:      "",
			expectError: false,
			description: "Valid issue comments",
		},
		{
			name:        "InvalidIssue",
			issueKey:    InvalidKey,
			startAt:     0,
			maxResults:  0,
			orderBy:     "",
			expand:      "",
			expectError: true,
			description: "Invalid issue comments",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.GetComments(context.Background(), GetCommentsInput{
				IssueKey: tt.issueKey,
				PaginationInput: PaginationInput{
					StartAt:    tt.startAt,
					MaxResults: tt.maxResults,
				},
				Expand:  tt.expand,
				OrderBy: tt.orderBy,
			})

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				if err == nil {
					assert.NotNil(t, result)
					if result != nil {
						comments, exists := result["comments"]
						assert.True(t, exists, "comments field should exist in response")
						if exists {
							assert.NotNil(t, comments, "comments should not be nil")
							t.Logf("Get comments successful. Issue key: %s Comments found: %d", tt.issueKey, len(comments.([]any)))
						}
					}
				} else {
					t.Logf("Get comments failed. Issue key: %s Error: %v", tt.issueKey, err)
				}
			}
		})
	}
}
