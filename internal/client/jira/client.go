// Package jira provides a client for interacting with Jira Data Center APIs.
package jira

import (
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"

	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/utils"
	"atlassian-dc-mcp-go/internal/utils/logging"

	"go.uber.org/zap"
)

// JiraClient represents a client for interacting with Jira Data Center APIs
type JiraClient struct {
	Config       *config.JiraConfig
	HTTPClient   *retryablehttp.Client
	ClientConfig *utils.HTTPClientConfig
	Logger       *zap.Logger
}

// NewJiraClient creates a new Jira client with the provided configuration.
func NewJiraClient(cfg *config.JiraConfig) *JiraClient {
	clientConfig := utils.DefaultHTTPClientConfig()
	if cfg.Timeout > 0 {
		clientConfig.Timeout = time.Duration(cfg.Timeout) * time.Second
	}
	httpClient := utils.NewRetryableHTTPClient(clientConfig)

	return &JiraClient{
		Config:       cfg,
		HTTPClient:   httpClient,
		ClientConfig: clientConfig,
		Logger:       logging.GetLogger(),
	}
}

// executeRequest executes an HTTP request to the Jira API.
func (c *JiraClient) executeRequest(method string, pathSegments []string, queryParams url.Values, body []byte, result any, accept utils.Accept) error {
	req, err := utils.BuildHttpRequest(method, c.Config.URL, pathSegments, queryParams, body, c.Config.Token, accept)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	if err := utils.ExecuteHTTPRequestWithRetry(c.HTTPClient, req, "jira", result); err != nil {
		return err
	}

	return nil
}