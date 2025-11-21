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

// HTTPClientConfig represents the HTTP client configuration for connection pooling
type HTTPClientConfig struct {
	MaxIdleConns        int `mapstructure:"max_idle_conns"`
	MaxIdleConnsPerHost int `mapstructure:"max_idle_conns_per_host"`
	IdleConnTimeout     int `mapstructure:"idle_conn_timeout"`
}

type Permissions map[string]bool

type ClientConfig struct {
	URL         string         `mapstructure:"url"`
	Token       string         `mapstructure:"token"`
	Permissions Permissions    `mapstructure:"permissions"`
	Timeout     int            `mapstructure:"timeout"`
	HTTP        HTTPClientConfig `mapstructure:"http"`
}

type TransportConfig struct {
	Modes []string `mapstructure:"modes"`

	HTTP struct {
		Path string `mapstructure:"path"`
	} `mapstructure:"http"`

	SSE struct {
		Path string `mapstructure:"path"`
	} `mapstructure:"sse"`

	Stdio struct {
		Enabled bool `mapstructure:"enabled"`
	} `mapstructure:"stdio"`
}

type Config struct {
	Port          int             `mapstructure:"port"`
	Jira          ClientConfig    `mapstructure:"jira"`
	Confluence    ClientConfig    `mapstructure:"confluence"`
	Bitbucket     ClientConfig    `mapstructure:"bitbucket"`
	Logging       logging.Config  `mapstructure:"logging"`
	Transport     TransportConfig `mapstructure:"transport"`
	ClientTimeout int             `mapstructure:"client_timeout"`
	Prune         PruneConfig     `mapstructure:"prune"`
}

// Validate checks that the configuration is valid
func (c *Config) Validate(authMode string) error {
	// Validate port
	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("invalid port: %d, must be between 1 and 65535", c.Port)
	}

	// Validate transport modes
	validTransports := map[string]bool{
		"stdio": true,
		"sse":   true,
		"http":  true,
	}

	if len(c.Transport.Modes) == 0 {
		c.Transport.Modes = []string{"http", "sse"}
	}

	for _, transport := range c.Transport.Modes {
		if !validTransports[transport] {
			return fmt.Errorf("invalid transport mode: %s, valid options are: stdio, sse, http", transport)
		}
	}

	// Validate client timeout
	if c.ClientTimeout <= 0 {
		c.ClientTimeout = 60 // default to 60 seconds
	}

	// Apply global client timeout to individual services if not set
	if c.Jira.Timeout <= 0 {
		c.Jira.Timeout = c.ClientTimeout
	}

	if c.Confluence.Timeout <= 0 {
		c.Confluence.Timeout = c.ClientTimeout
	}

	if c.Bitbucket.Timeout <= 0 {
		c.Bitbucket.Timeout = c.ClientTimeout
	}

	// Set default HTTP client config values if not specified
	if c.Jira.HTTP.MaxIdleConns <= 0 {
		c.Jira.HTTP.MaxIdleConns = 100
	}
	if c.Jira.HTTP.MaxIdleConnsPerHost <= 0 {
		c.Jira.HTTP.MaxIdleConnsPerHost = 20
	}
	if c.Jira.HTTP.IdleConnTimeout <= 0 {
		c.Jira.HTTP.IdleConnTimeout = 90
	}

	if c.Confluence.HTTP.MaxIdleConns <= 0 {
		c.Confluence.HTTP.MaxIdleConns = 100
	}
	if c.Confluence.HTTP.MaxIdleConnsPerHost <= 0 {
		c.Confluence.HTTP.MaxIdleConnsPerHost = 20
	}
	if c.Confluence.HTTP.IdleConnTimeout <= 0 {
		c.Confluence.HTTP.IdleConnTimeout = 90
	}

	if c.Bitbucket.HTTP.MaxIdleConns <= 0 {
		c.Bitbucket.HTTP.MaxIdleConns = 100
	}
	if c.Bitbucket.HTTP.MaxIdleConnsPerHost <= 0 {
		c.Bitbucket.HTTP.MaxIdleConnsPerHost = 20
	}
	if c.Bitbucket.HTTP.IdleConnTimeout <= 0 {
		c.Bitbucket.HTTP.IdleConnTimeout = 90
	}

	// Set default prune config if not specified
	if len(c.Prune.FuzzyKeys) == 0 && len(c.Prune.RemovePaths) == 0 {
		c.Prune = DefaultPruneConfig()
	}

	if authMode != "header" {
		if c.Jira.URL != "" && c.Jira.Token == "" {
			return fmt.Errorf("jira token must be set when jira url is configured")
		}

		if c.Confluence.URL != "" && c.Confluence.Token == "" {
			return fmt.Errorf("confluence token must be set when confluence url is configured")
		}

		if c.Bitbucket.URL != "" && c.Bitbucket.Token == "" {
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
func LoadConfig(configPath string, authMode string) (*Config, error) {
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

	viper.SetDefault("port", 8090)
	viper.SetDefault("logging.development", false)
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("client_timeout", 60) // Default client timeout in seconds
	viper.SetDefault("transport.modes", []string{"http", "sse"})
	viper.SetDefault("transport.http.path", "/mcp")
	viper.SetDefault("transport.sse.path", "/sse")
	viper.SetDefault("jira.http.max_idle_conns", 100)
	viper.SetDefault("jira.http.max_idle_conns_per_host", 20)
	viper.SetDefault("jira.http.idle_conn_timeout", 90)
	viper.SetDefault("confluence.http.max_idle_conns", 100)
	viper.SetDefault("confluence.http.max_idle_conns_per_host", 20)
	viper.SetDefault("confluence.http.idle_conn_timeout", 90)
	viper.SetDefault("bitbucket.http.max_idle_conns", 100)
	viper.SetDefault("bitbucket.http.max_idle_conns_per_host", 20)
	viper.SetDefault("bitbucket.http.idle_conn_timeout", 90)

	// Set default prune config
	defaultPrune := DefaultPruneConfig()
	viper.SetDefault("prune.fuzzy_keys", defaultPrune.FuzzyKeys)
	viper.SetDefault("prune.remove_paths", defaultPrune.RemovePaths)

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

	if err := config.Validate(authMode); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &config, nil
}

// WatchConfigOnChange sets up a callback for when the config file changes
func WatchConfigOnChange(run func(), authMode string) {
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)

		var newConfig Config
		if err := viper.Unmarshal(&newConfig); err != nil {
			fmt.Printf("Error unmarshaling updated config: %v\n", err)
			return
		}

		if err := newConfig.Validate(authMode); err != nil {
			fmt.Printf("Error validating updated config: %v\n", err)
			return
		}

		run()
	})
	viper.WatchConfig()
}