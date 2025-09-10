package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"atlassian-dc-mcp-go/internal/utils/logging"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Permissions struct {
	Read  bool `mapstructure:"read"`
	Write bool `mapstructure:"write"`
}

type JiraConfig struct {
	URL         string      `mapstructure:"url"`
	Token       string      `mapstructure:"token"`
	Permissions Permissions `mapstructure:"permissions"`
}

type ConfluenceConfig struct {
	URL         string      `mapstructure:"url"`
	Token       string      `mapstructure:"token"`
	Permissions Permissions `mapstructure:"permissions"`
}

type BitbucketConfig struct {
	URL         string      `mapstructure:"url"`
	Token       string      `mapstructure:"token"`
	Permissions Permissions `mapstructure:"permissions"`
}

type Config struct {
	Port         int             `mapstructure:"port"`
	Jira         JiraConfig      `mapstructure:"jira"`
	Confluence   ConfluenceConfig `mapstructure:"confluence"`
	Bitbucket    BitbucketConfig  `mapstructure:"bitbucket"`
	Logging      logging.Config   `mapstructure:"logging"`
	Transport    string          `mapstructure:"transport"`
	ClientTimeout int            `mapstructure:"client_timeout"`
}

// Validate checks that the configuration is valid
func (c *Config) Validate() error {
	// Validate port
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("invalid port: %d, must be between 1 and 65535", c.Port)
	}

	// Validate transport mode
	if c.Transport == "" {
		c.Transport = "stdio" // default to stdio
	}

	validTransports := map[string]bool{
		"stdio": true,
		"sse":   true,
		"http":  true,
	}

	if !validTransports[c.Transport] {
		return fmt.Errorf("invalid transport mode: %s, valid options are: stdio, sse, http", c.Transport)
	}

	// Validate client timeout
	if c.ClientTimeout <= 0 {
		c.ClientTimeout = 60 // default to 60 seconds
	}

	if c.Jira.URL != "" {
		if c.Jira.Token == "" {
			return fmt.Errorf("jira token must be set when jira url is configured")
		}
	}

	if c.Confluence.URL != "" {
		if c.Confluence.Token == "" {
			return fmt.Errorf("confluence token must be set when confluence url is configured")
		}
	}

	if c.Bitbucket.URL != "" {
		if c.Bitbucket.Token == "" {
			return fmt.Errorf("bitbucket token must be set when bitbucket url is configured")
		}
	}

	return nil
}

// LoadConfig loads the application configuration from various sources.
// It attempts to load from:
// 1. Current directory
// 2. Directory of the executable and its parent directories
// 3. Current working directory and its parent directories
//
// Returns a pointer to the loaded Config and nil error if successful,
// or nil and an error if configuration loading fails.
func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigType("yaml")

	// If a config path is provided, use it directly
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")

		if execPath, err := os.Executable(); err == nil {
			execDir := filepath.Dir(execPath)
			viper.AddConfigPath(execDir)

			for i := 0; i < 3; i++ {
				execDir = filepath.Dir(execDir)
				viper.AddConfigPath(execDir)
			}
		}

		if wd, err := os.Getwd(); err == nil {
			viper.AddConfigPath(wd)

			for i := 0; i < 3; i++ {
				wd = filepath.Dir(wd)
				viper.AddConfigPath(wd)
			}
		}
	}

	viper.SetDefault("jira.permissions.read", true)
	viper.SetDefault("confluence.permissions.read", true)
	viper.SetDefault("bitbucket.permissions.read", true)
	viper.SetDefault("port", 8090)
	viper.SetDefault("logging.development", false)
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("client_timeout", 60) // Default client timeout in seconds

	viper.SetEnvPrefix("MCP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &config, nil
}

// WatchConfigOnChange sets up a callback for when the config file changes
func WatchConfigOnChange(run func()) {
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)

		var newConfig Config
		if err := viper.Unmarshal(&newConfig); err != nil {
			fmt.Printf("Error unmarshaling updated config: %v\n", err)
			return
		}

		if err := newConfig.Validate(); err != nil {
			fmt.Printf("Error validating updated config: %v\n", err)
			return
		}

		run()
	})
	viper.WatchConfig()
}