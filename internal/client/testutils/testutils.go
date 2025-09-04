// Package testutils provides testing utilities and mock implementations.
package testutils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"

	"atlassian-dc-mcp-go/internal/config"
)

func IsIntegrationTest() bool {
	return os.Getenv("INTEGRATION_TEST") != ""
}

func SkipIntegrationTest(t *testing.T) {
	if !IsIntegrationTest() {
		t.Skip("Skipping integration test. Set INTEGRATION_TEST environment variable to run.")
	}
}

func LoadTestConfig(config interface{}, configFileName string) error {

	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)

	configPath := filepath.Join(dir, "testdata", configFileName)

	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return err
	}

	return nil
}

func NewIntegrationClient(clientCreator func(*config.Config) (interface{}, interface{}, error)) (interface{}, interface{}, error) {

	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load config: %w", err)
	}

	configOut, clientOut, err := clientCreator(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create client: %w", err)
	}

	if configOut == nil && clientOut == nil {
		return nil, nil, fmt.Errorf("client not created: configuration not set or incomplete")
	}

	return configOut, clientOut, nil
}

func RequireNotNil(t *testing.T, client interface{}, message string) {
	require.NotNil(t, client, message)
}
