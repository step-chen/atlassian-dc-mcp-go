// Package mcp provides the core server implementation for the Management Control Plane.
// It handles the initialization and management of various Atlassian services
// including Jira, Confluence, and Bitbucket through the Model Context Protocol (MCP).
package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/config"
	bitbucketTools "atlassian-dc-mcp-go/internal/mcp/tools/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools/common"
	confluenceTools "atlassian-dc-mcp-go/internal/mcp/tools/confluence"
	jiraTools "atlassian-dc-mcp-go/internal/mcp/tools/jira"
	"atlassian-dc-mcp-go/internal/utils/logging"

	"go.uber.org/zap"
)

const (
	BitbucketTokenHeader  = "Bitbucket-Token"
	JiraTokenHeader       = "Jira-Token"
	ConfluenceTokenHeader = "Confluence-Token"
)

// Server represents the MCP server instance
type Server struct {
	config           *config.Config
	authMode         string
	version          string
	startTime        time.Time
	jiraClient       *jira.JiraClient
	confluenceClient *confluence.ConfluenceClient
	bitbucketClient  *bitbucket.BitbucketClient
	mcpServer        *mcp.Server
	httpServer       *http.Server
	// WaitGroup to manage goroutines
	wg sync.WaitGroup
	// Channel to signal shutdown
	shutdownChan chan struct{}
}

// NewServer creates a new MCP server instance with the provided configuration and version
func NewServer(cfg *config.Config, authMode, version string) *Server {
	server := &Server{
		config:       cfg,
		authMode:     authMode,
		version:      version,
		startTime:    time.Now(),
		shutdownChan: make(chan struct{}),
	}

	// Initialize prune configuration
	client.InitPruneConfig(cfg.Prune)

	return server
}

// Initialize sets up the server with clients for Jira, Confluence, and Bitbucket based on configuration
func (s *Server) Initialize() error {
	var err error
	if s.config.Jira.URL != "" {
		s.jiraClient, err = jira.NewJiraClient(&s.config.Jira)
		if err != nil {
			return fmt.Errorf("failed to create Jira client: %w", err)
		}
	}

	if s.config.Confluence.URL != "" {
		s.confluenceClient, err = confluence.NewConfluenceClient(&s.config.Confluence)
		if err != nil {
			return fmt.Errorf("failed to create Confluence client: %w", err)
		}
	}

	if s.config.Bitbucket.URL != "" {
		s.bitbucketClient, err = bitbucket.NewBitbucketClient(&s.config.Bitbucket)
		if err != nil {
			return fmt.Errorf("failed to create Bitbucket client: %w", err)
		}
	}

	s.mcpServer = mcp.NewServer(&mcp.Implementation{
		Name:    "Atlassian Data Center MCP Server",
		Version: s.version,
	}, nil)

	// Add middleware for logging and error handling
	s.mcpServer.AddReceivingMiddleware(LoggingMiddleware(&s.config.Logging))
	s.mcpServer.AddReceivingMiddleware(ErrorMiddleware())

	s.addTools()

	return nil
}

// Start begins the MCP server using the configured transports
func (s *Server) Start(ctx context.Context) error {
	// Create MCP server factory function
	serverFactory := func(req *http.Request) *mcp.Server {
		return s.mcpServer
	}

	// Create a single HTTP mux for all HTTP-based transports
	mux := http.NewServeMux()

	// Register health and readiness check endpoints
	s.registerHealthEndpoints(mux)

	// Initialize all requested transport modes
	s.initTransports(ctx, mux, serverFactory)

	// Apply authentication middleware
	authMux := s.AuthMiddleware(mux)

	// If we have HTTP-based transports, start the HTTP server
	if s.hasHTTPTransports() {
		s.startHTTPServer(authMux)
	}

	// Wait for context cancellation or shutdown signal
	<-ctx.Done()

	// Return nil to indicate that the server was stopped by the context
	return nil
}

// requestLogger is a middleware that logs all incoming HTTP requests
func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		logging.GetLogger().Debug("HTTP Request",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("user_agent", r.UserAgent()),
		)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Stop gracefully stops the MCP server
func (s *Server) Stop(ctx context.Context) error {
	var httpErr error

	// Shutdown HTTP server if it exists
	if s.httpServer != nil {
		httpErr = s.httpServer.Shutdown(ctx)
	}

	// Close the shutdown channel to signal all goroutines
	close(s.shutdownChan)

	// Wait for all goroutines to finish
	s.wg.Wait()

	return httpErr
}

// GetConfig returns the server's configuration.
func (s *Server) GetConfig() *config.Config {
	return s.config
}

// GetJiraClient returns the Jira client instance.
func (s *Server) GetJiraClient() *jira.JiraClient {
	return s.jiraClient
}

// GetConfluenceClient returns the Confluence client instance.
func (s *Server) GetConfluenceClient() *confluence.ConfluenceClient {
	return s.confluenceClient
}

// GetBitbucketClient returns the Bitbucket client instance.
func (s *Server) GetBitbucketClient() *bitbucket.BitbucketClient {
	return s.bitbucketClient
}

// addTools registers all available tools with the MCP server
func (s *Server) addTools() {
	common.AddHealthCheckTool(s.mcpServer, s)
	common.AddCapabilitiesTool(s.mcpServer)

	if s.jiraClient != nil {
		s.addJiraTools()
	}

	if s.confluenceClient != nil {
		s.addConfluenceTools()
	}

	if s.bitbucketClient != nil {
		s.addBitbucketTools()
	}
}

// AuthMiddleware injects the authentication token into the request context
// based on the server's configured authentication mode.
func (s *Server) AuthMiddleware(next http.Handler) http.Handler {
	type serviceConfig struct {
		headerKey string
		cfgToken  func() string
		ctxKey    client.ContextKey
		url       *string
	}

	services := []serviceConfig{
		{
			headerKey: BitbucketTokenHeader,
			cfgToken:  func() string { return s.config.Bitbucket.Token },
			ctxKey:    client.BitbucketTokenKey,
			url:       &s.config.Bitbucket.URL,
		},
		{
			headerKey: JiraTokenHeader,
			cfgToken:  func() string { return s.config.Jira.Token },
			ctxKey:    client.JiraTokenKey,
			url:       &s.config.Jira.URL,
		},
		{
			headerKey: ConfluenceTokenHeader,
			cfgToken:  func() string { return s.config.Confluence.Token },
			ctxKey:    client.ConfluenceTokenKey,
			url:       &s.config.Confluence.URL,
		},
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tokenValues := make(map[client.ContextKey]string)

		for _, svc := range services {
			if s.authMode == "header" {
				tokenValues[svc.ctxKey] = r.Header.Get(svc.headerKey)
				continue
			}

			token := svc.cfgToken()
			if token != "" {
				tokenValues[svc.ctxKey] = token
			}
		}

		for key, token := range tokenValues {
			if token != "" {
				ctx = context.WithValue(ctx, key, token)
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// registerHealthEndpoints registers health and readiness check endpoints
func (s *Server) registerHealthEndpoints(mux *http.ServeMux) {
	// Add a simple health check endpoint that doesn't require authentication
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		uptime := time.Since(s.startTime).String()
		response := fmt.Sprintf(`{"status": "ok", "version": "%s", "uptime": "%s"}`, s.version, uptime)
		_, _ = w.Write([]byte(response))
	})

	// Add a readiness check endpoint that verifies service dependencies
	mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		// Check if clients are initialized
		readiness := map[string]interface{}{
			"status": "ok",
		}

		issues := []string{}

		if s.config.Jira.URL != "" && s.jiraClient == nil {
			issues = append(issues, "Jira client not initialized")
		}

		if s.config.Confluence.URL != "" && s.confluenceClient == nil {
			issues = append(issues, "Confluence client not initialized")
		}

		if s.config.Bitbucket.URL != "" && s.bitbucketClient == nil {
			issues = append(issues, "Bitbucket client not initialized")
		}

		if len(issues) > 0 {
			readiness["status"] = "not ready"
			readiness["issues"] = issues
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		// Convert map to JSON and write to response
		jsonData, _ := json.Marshal(readiness)
		_, _ = w.Write(jsonData)
	})
}

// initTransports initializes all requested transport modes
func (s *Server) initTransports(ctx context.Context, mux *http.ServeMux, serverFactory func(req *http.Request) *mcp.Server) {
	// Start all requested transports
	for _, transport := range s.config.Transport.Modes {
		switch transport {
		case "stdio":
			s.wg.Add(1)
			go func() {
				defer s.wg.Done()
				// Create a new context that can be cancelled independently
				stdioCtx, cancel := context.WithCancel(ctx)
				defer cancel()

				if err := s.mcpServer.Run(stdioCtx, &mcp.StdioTransport{}); err != nil {
					fmt.Printf("Stdio transport error: %v\n", err)
				}
			}()
		case "sse":
			sseHandler := mcp.NewSSEHandler(serverFactory, &mcp.SSEOptions{})
			ssePath := s.config.Transport.SSE.Path
			if ssePath == "" {
				ssePath = "/sse"
			}
			mux.Handle(ssePath, sseHandler)
		case "http":
			handler := mcp.NewStreamableHTTPHandler(serverFactory, nil)
			httpPath := s.config.Transport.HTTP.Path
			if httpPath == "" {
				httpPath = "/mcp"
			}
			mux.Handle(httpPath, handler)
		}
	}
}

// hasHTTPTransports checks if HTTP-based transports (http or sse) are configured
func (s *Server) hasHTTPTransports() bool {
	for _, t := range s.config.Transport.Modes {
		if t == "http" || t == "sse" {
			return true
		}
	}
	return false
}

// startHTTPServer starts the HTTP server with the provided handler
func (s *Server) startHTTPServer(authMux http.Handler) {
	addr := fmt.Sprintf(":%d", s.config.Port)
	// Only use requestLogger in debug mode
	var handler http.Handler
	if s.config.Logging.Level == "debug" {
		handler = requestLogger(authMux)
	} else {
		handler = authMux
	}

	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	// Start HTTP server in a goroutine
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %v\n", err)
		}
	}()
}

// addJiraTools registers all Jira-related tools with the MCP server
func (s *Server) addJiraTools() {
	// Check if Jira write permission is enabled
	permissions := s.config.Jira.Permissions

	jiraTools.AddIssueTools(s.mcpServer, s.jiraClient, permissions)
	jiraTools.AddBoardTools(s.mcpServer, s.jiraClient, permissions)
	jiraTools.AddProjectTools(s.mcpServer, s.jiraClient, permissions)
	jiraTools.AddCommentTools(s.mcpServer, s.jiraClient, permissions)
	jiraTools.AddIssueTypeTools(s.mcpServer, s.jiraClient, permissions)
	jiraTools.AddPriorityTools(s.mcpServer, s.jiraClient, permissions)
	jiraTools.AddTransitionTools(s.mcpServer, s.jiraClient, permissions)
	jiraTools.AddUserTools(s.mcpServer, s.jiraClient, permissions)
	jiraTools.AddWorklogTools(s.mcpServer, s.jiraClient, permissions)
	jiraTools.AddSubtaskTools(s.mcpServer, s.jiraClient, permissions)
}

// addConfluenceTools registers all Confluence-related tools with the MCP server
func (s *Server) addConfluenceTools() {
	// Check if Confluence write permission is enabled
	permissions := s.config.Confluence.Permissions

	confluenceTools.AddContentTools(s.mcpServer, s.confluenceClient, permissions)
	confluenceTools.AddSpaceTools(s.mcpServer, s.confluenceClient, permissions)
	confluenceTools.AddChildrenTools(s.mcpServer, s.confluenceClient, permissions)
	confluenceTools.AddLabelTools(s.mcpServer, s.confluenceClient, permissions)
	confluenceTools.AddUserTools(s.mcpServer, s.confluenceClient, permissions)
}

// addBitbucketTools registers all Bitbucket-related tools with the MCP server
func (s *Server) addBitbucketTools() {
	permissions := s.config.Bitbucket.Permissions

	bitbucketTools.AddUserTools(s.mcpServer, s.bitbucketClient, permissions)
	bitbucketTools.AddProjectTools(s.mcpServer, s.bitbucketClient, permissions)
	bitbucketTools.AddBranchTools(s.mcpServer, s.bitbucketClient, permissions)
	bitbucketTools.AddCommitTools(s.mcpServer, s.bitbucketClient, permissions)
	bitbucketTools.AddPullRequestTools(s.mcpServer, s.bitbucketClient, permissions)
	bitbucketTools.AddAttachmentTools(s.mcpServer, s.bitbucketClient, permissions)
	bitbucketTools.AddTagTools(s.mcpServer, s.bitbucketClient, permissions)
	bitbucketTools.AddRepositoryTools(s.mcpServer, s.bitbucketClient, permissions)
	bitbucketTools.AddSearchTools(s.mcpServer, s.bitbucketClient, permissions)
}
