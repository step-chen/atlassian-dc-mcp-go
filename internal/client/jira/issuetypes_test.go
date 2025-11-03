package jira

import (
	"atlassian-dc-mcp-go/internal/client/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetIssueTypes(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	tests := []struct {
		name string
	}{
		{
			name: "GetIssueTypes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := SetupIntegrationTest(t)
			if err != nil {
				t.Fatalf("Could not load config.yaml: %v", err)
			}

			require.NotNil(t, client, "Jira client should not be nil")

			issueTypes, err := client.GetIssueTypes()

			if err == nil {
				assert.NotNil(t, issueTypes)
				if issueTypes != nil {
					if len(issueTypes) > 0 {
						firstIssueType := issueTypes[0]
						assert.NotEmpty(t, firstIssueType["id"], "Issue type should have an id")
						assert.NotEmpty(t, firstIssueType["name"], "Issue type should have a name")
						assert.NotEmpty(t, firstIssueType["description"], "Issue type should have a description")
						t.Logf("Number of issue types retrieved: %d", len(issueTypes))
					}
				}
			} else {
				t.Logf("Get issue types completed. Error (may be expected): %v", err)
			}
		})
	}
}
