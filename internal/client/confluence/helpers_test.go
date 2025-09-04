package confluence

import (
	"testing"

	"atlassian-dc-mcp-go/internal/client/testutils"
	"atlassian-dc-mcp-go/internal/config"

	"github.com/stretchr/testify/require"
)

type TestPagesConfig struct {
	ValidID    string `json:"validID"`
	CommentsID string `json:"commentsID"`
	InvalidID  string `json:"invalidID"`
}

type TestSpacesConfig struct {
	Key            string `json:"key"`
	Limit          int    `json:"limit"`
	Start          int    `json:"start"`
	CustomLim      int    `json:"customLim"`
	InvalidLimit   int    `json:"invalidLimit"`
	NonExistentKey string `json:"nonExistentKey"`
}

type TestSearchQuery struct {
	Name   string   `json:"name"`
	CQL    string   `json:"cql"`
	Limit  int      `json:"limit"`
	Start  int      `json:"start"`
	Expand []string `json:"expand,omitempty"`
}

type TestConfig struct {
	Pages  TestPagesConfig   `json:"pages"`
	Spaces TestSpacesConfig  `json:"spaces"`
	Search []TestSearchQuery `json:"search"`
}

func loadTestConfig() (*TestConfig, error) {
	config := &TestConfig{}
	err := testutils.LoadTestConfig(config, "test_config.json")
	return config, err
}

func SetupIntegrationTest(t *testing.T) (*ConfluenceClient, error) {
	testutils.SkipIntegrationTest(t)

	_, client, err := newConfluenceIntegrationClient()
	if err != nil {
		return nil, err
	}

	require.NotNil(t, client, "Confluence client should not be nil. Confluence URL and Token must be configured in config.yaml.")

	return client, nil
}

func newConfluenceIntegrationClient() (*config.ConfluenceConfig, *ConfluenceClient, error) {
	confluenceConfig, client, err := testutils.NewIntegrationClient(func(cfg *config.Config) (interface{}, interface{}, error) {
		if cfg.Confluence.URL == "" || cfg.Confluence.Token == "" {
			return nil, nil, nil
		}

		client := NewConfluenceClient(&cfg.Confluence)
		return &cfg.Confluence, client, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if confluenceConfig == nil || client == nil {
		return nil, nil, nil
	}

	return confluenceConfig.(*config.ConfluenceConfig), client.(*ConfluenceClient), nil
}
