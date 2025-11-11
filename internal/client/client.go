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
func NewBaseClient(config *config.ClientConfig, name string, tokenKey ContextKey) (*BaseClient, error) {
	clientConfig := DefaultHTTPClientConfig()
	if config.Timeout > 0 {
		clientConfig.Timeout = time.Duration(config.Timeout) * time.Second
	}

	// Apply HTTP client connection pool settings from config
	clientConfig.MaxIdleConns = config.HTTP.MaxIdleConns
	clientConfig.MaxIdleConnsPerHost = config.HTTP.MaxIdleConnsPerHost
	clientConfig.IdleConnTimeout = time.Duration(config.HTTP.IdleConnTimeout) * time.Second

	httpClient := NewRetryableHTTPClient(clientConfig, &TokenAuthTransport{
		TokenKey: tokenKey,
	})

	return &BaseClient{
		Config:     config,
		HTTPClient: httpClient,
		Name:       name,
	}, nil
}
