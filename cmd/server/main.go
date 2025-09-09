package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/mcp"
	"atlassian-dc-mcp-go/internal/utils/logging"

	"go.uber.org/zap"
)

func main() {
	logging.InitLogger(&logging.Config{
		Development: true,
		Level:       "info",
	})
	logger := logging.GetLogger()
	defer func() {
		_ = logger.Sync()
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	config.WatchConfigOnChange(func() {
		logger.Info("Configuration reloaded")
	})

	mcpServer := mcp.NewServer(cfg)

	if err := mcpServer.Initialize(); err != nil {
		logger.Fatal("Failed to initialize MCP server", zap.Error(err))
	}

	logger.Info("Atlassian Data Center MCP (Model Context Protocol) server starting...")

	// Create a context that can be cancelled
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Start the server in a goroutine
	serverErr := make(chan error, 1)
	go func() {
		if err := mcpServer.Start(ctx); err != nil {
			serverErr <- err
		}
	}()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Wait for either server error or signal
	select {
	case err := <-serverErr:
		logger.Error("MCP server error", zap.Error(err))
	case <-sigChan:
		logger.Info("Received interrupt signal, shutting down...")
		// Cancel the context to signal the server to stop
		cancel()
		// Call the server's stop method
		mcpServer.Stop()
		logger.Info("Server stopped gracefully")
	}
}