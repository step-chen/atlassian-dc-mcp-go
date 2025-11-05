package confluence

import (
	"atlassian-dc-mcp-go/internal/client/testutils"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPage(t *testing.T) {
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
		pageID      string
		expand      []string
		expectError bool
	}{
		{
			name:        "ValidPage",
			pageID:      testConfig.Pages.ValidID,
			expand:      nil,
			expectError: false,
		},
		{
			name:        "InvalidPage",
			pageID:      testConfig.Pages.InvalidID,
			expand:      nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := GetContentByIDInput{
				ContentID: tt.pageID,
				Expand:    tt.expand,
			}
			result, err := client.GetContentByID(context.Background(), input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.NotEmpty(t, result)
			}
		})
	}
}

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
		pageID      string
		expectError bool
	}{
		{
			name:        "ValidPageWithComments",
			pageID:      testConfig.Pages.CommentsID,
			expectError: false,
		},
		{
			name:        "InvalidPage",
			pageID:      testConfig.Pages.InvalidID,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用新的结构体参数方式
			input := GetContentChildrenByTypeInput{
				ContentID: tt.pageID,
				ChildType: "comment",
			}
			result, err := client.GetContentChildrenByType(context.Background(), input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.NotEmpty(t, result)
			}
		})
	}
}
