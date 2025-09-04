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

type Config struct {
	Jira       JiraConfig       `mapstructure:"jira"`
	Confluence ConfluenceConfig `mapstructure:"confluence"`
	Bitbucket  BitbucketConfig  `mapstructure:"bitbucket"`
	Logging    logging.Config   `mapstructure:"logging"`
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

// Validate checks that the configuration is valid
func (c *Config) Validate() error {
	if c.Jira.URL == "" && c.Confluence.URL == "" && c.Bitbucket.URL == "" {
		return fmt.Errorf("at least one of jira, confluence, or bitbucket must be configured")
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
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

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

	viper.SetDefault("jira.permissions.read", true)
	viper.SetDefault("confluence.permissions.read", true)
	viper.SetDefault("bitbucket.permissions.read", true)
	viper.SetDefault("logging.development", false)
	viper.SetDefault("logging.level", "info")

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