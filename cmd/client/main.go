package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/utils/logging"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

func main() {

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
	cfg, err := config.LoadConfig()
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

	// Connect to the server using appropriate transport based on config
	var session *mcp.ClientSession
	if cfg.Transport == "http" {
		// Use HTTP transport
		endpoint := fmt.Sprintf("http://localhost:%d", cfg.Port)
		logger.Info("Connecting to MCP server via HTTP", zap.String("endpoint", endpoint))
		session, err = client.Connect(ctx, &mcp.StreamableClientTransport{
			Endpoint: endpoint,
		}, nil)
	} else {
		// Default to stdio transport
		logger.Info("Connecting to MCP server via stdio")
		session, err = client.Connect(ctx, &mcp.StdioTransport{}, nil)
	}
	
	if err != nil {
		logger.Fatal("Failed to connect to MCP server", zap.Error(err))
	}
	defer func() {
		session.Close()
		logger.Info("MCP session closed")
	}()

	// Call the jira_get_issue tool with issue key HAD-10228
	logger.Info("Calling jira_get_issue tool for HAD-10228")
	result, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name: "jira_get_issue",
		Arguments: map[string]interface{}{
			"issueKey": "HAD-10228",
		},
	})

	if err != nil {
		logger.Error("Failed to call jira_get_issue tool", zap.Error(err))
		os.Exit(1)
	}

	// Print the result
	logger.Info("Successfully called jira_get_issue tool")
	fmt.Println("\n=== Jira Issue HAD-10228 ===")
	
	// Print the raw result
	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("Raw result: %s\n", resultJSON)
	
	// Process content if available
	if result.Content != nil {
		fmt.Println("\n--- Content ---")
		for _, content := range result.Content {
			// Try to convert to TextContent
			contentJSON, _ := json.Marshal(content)
			fmt.Printf("Content: %s\n", string(contentJSON))
		}
	}
	
	// Check for error in result
	if result.StructuredContent != nil {
		fmt.Println("\n--- Structured Content ---")
		structuredJSON, _ := json.MarshalIndent(result.StructuredContent, "", "  ")
		fmt.Printf("Structured content: %s\n", structuredJSON)
	}
}