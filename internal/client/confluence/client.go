// Package confluence provides a client for interacting with Confluence Data Center APIs.
package confluence

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

// ConfluenceClient represents a client for interacting with Confluence Data Center APIs
type ConfluenceClient struct {
	Config       *config.ConfluenceConfig
	HTTPClient   *retryablehttp.Client
	ClientConfig *utils.HTTPClientConfig
	Logger       *zap.Logger
}

// NewConfluenceClient creates a new Confluence client with the provided configuration.
func NewConfluenceClient(cfg *config.ConfluenceConfig) *ConfluenceClient {
	clientConfig := utils.DefaultHTTPClientConfig()
	if cfg.Timeout > 0 {
		clientConfig.Timeout = time.Duration(cfg.Timeout) * time.Second
	}
	httpClient := utils.NewRetryableHTTPClient(clientConfig)

	return &ConfluenceClient{
		Config:       cfg,
		HTTPClient:   httpClient,
		ClientConfig: clientConfig,
		Logger:       logging.GetLogger(),
	}
}

// executeRequest executes an HTTP request to the Confluence API.
func (c *ConfluenceClient) executeRequest(method string, pathSegments []string, queryParams url.Values, body []byte, result any, accept utils.Accept) error {
	req, err := utils.BuildHttpRequest(method, c.Config.URL, pathSegments, queryParams, body, c.Config.Token, accept)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	if err := utils.ExecuteHTTPRequestWithRetry(c.HTTPClient, req, "confluence", result); err != nil {
		return err
	}

	return nil
}
