package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/mcp"
	"atlassian-dc-mcp-go/internal/utils/logging"

	"go.uber.org/zap"
)

// These variables are populated at build time by ldflags
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Define command line flags
	configPath := flag.String("c", "", "Path to config file (optional)")
	flag.StringVar(configPath, "config", "", "Path to config file (optional)")
	help := flag.Bool("h", false, "Show help message")
	flag.BoolVar(help, "help", false, "Show help message")
	versionFlag := flag.Bool("version", false, "Show version information")
	authMode := flag.String("auth-mode", "config", "Authentication mode. One of: config, header")
	flag.Parse()

	if *help {
		fmt.Println("Atlassian Data Center MCP Server")
		fmt.Println("Usage: server [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *versionFlag {
		fmt.Printf("Atlassian Data Center MCP Server Version: %s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Date: %s\n", date)
		os.Exit(0)
	}

	cfg, err := config.LoadConfig(*configPath, *authMode)
	if err != nil {
		// Use standard log since logger is not initialized yet
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger with configuration from file/env
	logging.InitLogger(&cfg.Logging)
	logger := logging.GetLogger()
	defer func() {
		// Flushes any buffered log entries
		_ = logger.Sync()
	}()

	logger.Info("Configuration loaded successfully",
		zap.String("version", version),
		zap.String("commit", commit),
		zap.String("date", date))

	config.WatchConfigOnChange(func() {
		logger.Info("Configuration reloaded")
	}, *authMode)

	mcpServer := mcp.NewServer(cfg, *authMode, version)

	if err := mcpServer.Initialize(); err != nil {
		logger.Fatal("Failed to initialize MCP server", zap.Error(err))
	}

	logger.Info("Atlassian Data Center MCP (Model Context Protocol) server starting...",
		zap.String("version", version))

	// Create a context that can be cancelled but without timeout
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the server in a goroutine
	serverErr := make(chan error, 1)
	go func() {
		// http.ErrServerClosed is expected on graceful shutdown, so we ignore it.
		if err := mcpServer.Start(ctx); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	// Print message that server has started, port and service paths
	httpPath := cfg.Transport.HTTP.Path
	if httpPath == "" {
		httpPath = "/mcp"
	}

	ssePath := cfg.Transport.SSE.Path
	if ssePath == "" {
		ssePath = "/sse"
	}

	logger.Info("Server started. ",
		zap.Int("port", cfg.Port),
		zap.String("http_path", httpPath),
		zap.String("sse_path", ssePath),
		zap.Strings("transport_modes", cfg.Transport.Modes))

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

		// Give the server some time to shut down gracefully
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		// Stop will handle the graceful shutdown of the HTTP server.
		// For other transports, this is a no-op.
		if err := mcpServer.Stop(shutdownCtx); err != nil {
			if err == context.DeadlineExceeded {
				logger.Warn("Server shutdown timeout, forcing exit")
			} else {
				logger.Warn("Server shutdown failed", zap.Error(err))
			}
		} else {
			logger.Info("Server stopped gracefully")
		}
	}
}
