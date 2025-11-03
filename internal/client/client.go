package client

import (
	"atlassian-dc-mcp-go/internal/config"

	"github.com/hashicorp/go-retryablehttp"
)

type BaseClient struct {
	Config     *config.ClientConfig
	HTTPClient *retryablehttp.Client
	Name       string
}
