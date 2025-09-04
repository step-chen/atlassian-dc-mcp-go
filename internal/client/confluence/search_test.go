package confluence

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"atlassian-dc-mcp-go/internal/client/testutils"
)

func TestSearch(t *testing.T) {
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
		cqlQuery    string
		limit       int
		start       int
		expand      []string
		expectError bool
	}{}

	for _, searchQuery := range testConfig.Search {
		tests = append(tests, struct {
			name        string
			cqlQuery    string
			limit       int
			start       int
			expand      []string
			expectError bool
		}{
			name:        searchQuery.Name,
			cqlQuery:    searchQuery.CQL,
			limit:       searchQuery.Limit,
			start:       searchQuery.Start,
			expand:      searchQuery.Expand,
			expectError: strings.Contains(searchQuery.Name, "Invalid"),
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := client.Search(tt.cqlQuery, "", "", tt.expand, tt.start, tt.limit, false)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				if result != nil {
					sizeValue, exists := result["size"]
					assert.True(t, exists, "size field should exist in response")
					if exists {
						size := int(sizeValue.(float64))
						assert.Greater(t, size, 0, "size should be greater than zero")
						t.Logf("Search successful. Size from API response: %d", size)
					}
				}
			}
		})
	}
}
