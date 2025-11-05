package jira

import (
	"testing"

	"atlassian-dc-mcp-go/internal/client/testutils"
	"atlassian-dc-mcp-go/internal/config"

	"github.com/stretchr/testify/require"
)

type TestIssuesConfig struct {
	ValidKey string `json:"validKey"`
}

type TestProjectsConfig struct {
	ValidKey   string `json:"validKey"`
	SearchName string `json:"searchName"`
	OrderBy    string `json:"orderBy"`
}

type TestUsersConfig struct {
	ValidUsername string `json:"validUsername"`
	ValidKey      string `json:"validKey"`
	SearchQuery   string `json:"searchQuery"`
}

type TestCommentsConfig struct {
	IssueKey string `json:"issueKey"`
}

type TestWorklogsConfig struct {
	IssueKey string `json:"issueKey"`
}

type TestSearchConfig struct {
	ValidJQL   string `json:"validJQL"`
	MaxResults int    `json:"maxResults"`
	ProjectKey string `json:"projectKey"`
}

type TestSubtasksConfig struct {
	IssueKey string `json:"issueKey"`
}

type TestBoardsConfig struct {
	ValidID int `json:"validID"`
}

type TestSprintsConfig struct {
	ValidBoardID  int `json:"validBoardID"`
	ValidSprintID int `json:"validSprintID"`
}

type TestConfig struct {
	Issues   TestIssuesConfig   `json:"issues"`
	Projects TestProjectsConfig `json:"projects"`
	Users    TestUsersConfig    `json:"users"`
	Search   TestSearchConfig   `json:"search"`
	Comments TestCommentsConfig `json:"comments"`
	Worklogs TestWorklogsConfig `json:"worklogs"`
	Boards   TestBoardsConfig   `json:"boards"`
	Subtasks TestSubtasksConfig `json:"subtasks"`
	Sprints  TestSprintsConfig  `json:"sprints"`
}

func loadTestConfig() (*TestConfig, error) {
	config := &TestConfig{}
	err := testutils.LoadTestConfig(config, "test_config.json")
	return config, err
}

func SetupIntegrationTest(t *testing.T) (*JiraClient, error) {
	testutils.SkipIntegrationTest(t)

	_, client, err := newJiraIntegrationClient()
	if err != nil {
		return nil, err
	}

	require.NotNil(t, client, "Jira client should not be nil. Jira URL and Token must be configured in config.yaml.")

	return client, nil
}

func newJiraIntegrationClient() (*config.ClientConfig, *JiraClient, error) {
	jiraConfig, client, err := testutils.NewIntegrationClient(func(cfg *config.Config) (interface{}, interface{}, error) {
		if cfg.Jira.URL == "" || cfg.Jira.Token == "" {
			return nil, nil, nil
		}

		client, err := NewJiraClient(&cfg.Jira)
		if err != nil {
			return nil, nil, err
		}
		return &cfg.Jira, client, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if jiraConfig == nil || client == nil {
		return nil, nil, nil
	}

	return jiraConfig.(*config.ClientConfig), client.(*JiraClient), nil
}
