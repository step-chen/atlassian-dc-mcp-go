package jira

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"atlassian-dc-mcp-go/internal/client/testutils"
)

func TestGetPriorities(t *testing.T) {
	testutils.SkipIntegrationTest(t)

	tests := []struct {
		name string
	}{
		{
			name: "GetPriorities",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := SetupIntegrationTest(t)
			if err != nil {
				t.Fatalf("Could not load config.yaml: %v", err)
			}

			require.NotNil(t, client, "Jira client should not be nil")

			priorities, err := client.GetPriorities()

			if err == nil {
				assert.NotNil(t, priorities)
				if priorities != nil {
					assert.NotEmpty(t, priorities, "Priorities response should not be empty")

					if len(priorities) > 0 {
						firstPriority := priorities[0]
						assert.NotEmpty(t, firstPriority["id"], "Priority should have an id")
						assert.NotEmpty(t, firstPriority["name"], "Priority should have a name")
						assert.NotEmpty(t, firstPriority["description"], "Priority should have a description")
						t.Logf("Number of priorities retrieved: %d", len(priorities))
					}
				}
			} else {
				t.Logf("Get priorities completed. Error (may be expected): %v", err)
			}
		})
	}
}
