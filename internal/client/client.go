package client

import (
	"time"

	"atlassian-dc-mcp-go/internal/config"

	"github.com/hashicorp/go-retryablehttp"
)

type BaseClient struct {
	Config     *config.ClientConfig
	HTTPClient *retryablehttp.Client
	Name       string
}

// NewBaseClient creates a new BaseClient with the provided configuration and name.
func NewBaseClient(config *config.ClientConfig, name string) *BaseClient {
	clientConfig := DefaultHTTPClientConfig()
	if config.Timeout > 0 {
		clientConfig.Timeout = time.Duration(config.Timeout) * time.Second
	}

	return &BaseClient{
		Config:     config,
		HTTPClient: NewRetryableHTTPClient(clientConfig),
		Name:       name,
	}
}