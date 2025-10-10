// Package mcp provides the core server implementation for the Management Control Plane.
// It handles the initialization and management of various Atlassian services
// including Jira, Confluence, and Bitbucket through the Model Context Protocol (MCP).
package mcp

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/config"
	bitbucketTools "atlassian-dc-mcp-go/internal/mcp/tools/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/tools/common"
	confluenceTools "atlassian-dc-mcp-go/internal/mcp/tools/confluence"
	jiraTools "atlassian-dc-mcp-go/internal/mcp/tools/jira"
)

// Server represents the MCP server instance
type Server struct {
	config           *config.Config
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

// NewServer creates a new MCP server instance with the provided configuration
func NewServer(cfg *config.Config) *Server {
	return &Server{
		config:       cfg,
		shutdownChan: make(chan struct{}),
	}
}

// Initialize sets up the server with clients for Jira, Confluence, and Bitbucket based on configuration
func (s *Server) Initialize() error {
	if s.config.Jira.URL != "" && s.config.Jira.Token != "" {
		s.jiraClient = jira.NewJiraClient(&s.config.Jira)
	}

	if s.config.Confluence.URL != "" && s.config.Confluence.Token != "" {
		s.confluenceClient = confluence.NewConfluenceClient(&s.config.Confluence)
	}

	if s.config.Bitbucket.URL != "" && s.config.Bitbucket.Token != "" {
		s.bitbucketClient = bitbucket.NewBitbucketClient(&s.config.Bitbucket)
	}

	s.mcpServer = mcp.NewServer(&mcp.Implementation{
		Name:    "atlassian-dc-mcp",
		Version: "1.0.0",
	}, nil)

	s.mcpServer.AddReceivingMiddleware(LoggingMiddleware())

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

	// If we have HTTP-based transports, start the HTTP server
	hasHTTP := false
	for _, t := range s.config.Transport.Modes {
		if t == "http" || t == "sse" {
			hasHTTP = true
			break
		}
	}

	if hasHTTP {
		addr := fmt.Sprintf(":%d", s.config.Port)
		s.httpServer = &http.Server{
			Addr:    addr,
			Handler: mux,
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

	// Wait for context cancellation or shutdown signal
	<-ctx.Done()

	// Return nil to indicate that the server was stopped by the context
	return nil
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
}
