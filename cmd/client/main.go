package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/utils/logging"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

func main() {
	// Define command line flags
	configPath := flag.String("c", "", "Path to config file (optional)")
	flag.StringVar(configPath, "config", "", "Path to config file (optional)")
	help := flag.Bool("h", false, "Show help message")
	flag.BoolVar(help, "help", false, "Show help message")
	flag.Parse()

	if *help {
		fmt.Println("Atlassian Data Center MCP Client")
		fmt.Println("Usage: client [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Initialize logger
	logging.InitLogger(&logging.Config{
		Development: true,
		Level:       "info",
	})
	logger := logging.GetLogger()
	defer func() {
		_ = logger.Sync()
	}()

	// Load configuration
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	logger.Info("Atlassian Data Center MCP Client starting...")

	// Log configured services
	if cfg.Jira.URL != "" {
		logger.Info("Jira configuration", zap.String("url", cfg.Jira.URL))
	}
	if cfg.Confluence.URL != "" {
		logger.Info("Confluence configuration", zap.String("url", cfg.Confluence.URL))
	}
	if cfg.Bitbucket.URL != "" {
		logger.Info("Bitbucket configuration", zap.String("url", cfg.Bitbucket.URL))
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ClientTimeout)*time.Second)
	defer cancel()

	// Create an MCP client
	client := mcp.NewClient(&mcp.Implementation{
		Name:    "mcp-client",
		Version: "1.0.0",
	}, nil)

	// Connect to the server via stdio
	session, err := client.Connect(ctx, &mcp.StdioTransport{}, nil)
	if err != nil {
		logger.Fatal("Failed to connect to server", zap.Error(err))
	}
	defer session.Close()

	// List available tools
	tools, err := session.ListTools(ctx, &mcp.ListToolsParams{})
	if err != nil {
		logger.Error("Failed to list tools", zap.Error(err))
		return
	}

	fmt.Printf("Available tools (%d):\n", len(tools.Tools))
	for _, tool := range tools.Tools {
		fmt.Printf("- %s: %s\n", tool.Name, tool.Description)
	}

	// Example: Call a tool
	if len(tools.Tools) > 0 {
		fmt.Println("\nCalling capabilities tool...")
		result, err := session.CallTool(ctx, &mcp.CallToolParams{
			Name: "capabilities",
		})
		if err != nil {
			logger.Error("Failed to call capabilities tool", zap.Error(err))
			return
		}

		fmt.Printf("Capabilities result: %+v\n", result)
	}
}