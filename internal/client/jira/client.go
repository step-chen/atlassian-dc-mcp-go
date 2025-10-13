// Package jira provides a client for interacting with Jira Data Center APIs.
package jira

import (
	"fmt"
	"net/http"
	"net/url"

	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/utils"
)

// JiraClient represents a client for interacting with Jira Data Center APIs
type JiraClient struct {
	Config       *config.JiraConfig
	HTTPClient   *http.Client
	ClientConfig *utils.HTTPClientConfig
}

// NewJiraClient creates a new Jira client with the provided configuration.
func NewJiraClient(config *config.JiraConfig) *JiraClient {
	clientConfig := utils.DefaultHTTPClientConfig()
	httpClient := utils.NewHTTPClient(clientConfig)

	return &JiraClient{
		Config:       config,
		HTTPClient:   httpClient,
		ClientConfig: clientConfig,
	}
}

// executeRequest executes an HTTP request to the Jira API.
func (c *JiraClient) executeRequest(method string, pathSegments []string, queryParams url.Values, body []byte, result any, accept utils.Accept) error {
	req, err := utils.BuildHttpRequest(method, c.Config.URL, pathSegments, queryParams, body, c.Config.Token, accept)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	if err := utils.ExecuteHTTPRequestWithRetry(c.HTTPClient, req, "jira", result, c.ClientConfig); err != nil {
		return err
	}

	return nil
}
