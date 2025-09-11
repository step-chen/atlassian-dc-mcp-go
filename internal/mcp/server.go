// Package mcp provides the core server implementation for the Management Control Plane.
// It handles the initialization and management of various Atlassian services
// including Jira, Confluence, and Bitbucket through the Model Context Protocol (MCP).
package mcp

import (
	"context"
	"fmt"
	"net/http"

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

type Server struct {
	config           *config.Config
	jiraClient       *jira.JiraClient
	confluenceClient *confluence.ConfluenceClient
	bitbucketClient  *bitbucket.BitbucketClient
	mcpServer        *mcp.Server
	httpServer       *http.Server
}

// NewServer creates a new MCP server instance with the provided configuration
func NewServer(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
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
	s.mcpServer.AddReceivingMiddleware(CheckPermissionMiddleware(s.config))

	s.addTools()

	return nil
}

// Start begins the MCP server using the configured transport
func (s *Server) Start(ctx context.Context) error {
	switch s.config.Transport {
	case "stdio":
		return s.mcpServer.Run(ctx, &mcp.StdioTransport{})
	case "sse":
		return s.mcpServer.Run(ctx, &mcp.SSEServerTransport{})
	case "http":
		handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
			return s.mcpServer
		}, nil)

		addr := fmt.Sprintf(":%d", s.config.Port)
		s.httpServer = &http.Server{
			Addr:    addr,
			Handler: handler,
		}
		return s.httpServer.ListenAndServe()
	default:
		return s.mcpServer.Run(ctx, &mcp.StdioTransport{})
	}
}

// Stop gracefully stops the MCP server
func (s *Server) Stop(ctx context.Context) error {
	if s.httpServer != nil {
		// For HTTP transport, we gracefully stop the HTTP server
		return s.httpServer.Shutdown(ctx)
	}
	// For stdio and sse, the server is stopped by canceling the context passed to Start().
	return nil
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
	hasWritePermission := s.config.Jira.Permissions.Write
	
	jiraTools.AddIssueTools(s.mcpServer, s.jiraClient, hasWritePermission)
	jiraTools.AddBoardTools(s.mcpServer, s.jiraClient)
	jiraTools.AddProjectTools(s.mcpServer, s.jiraClient)
	jiraTools.AddCommentTools(s.mcpServer, s.jiraClient, hasWritePermission)
	jiraTools.AddIssueTypeTools(s.mcpServer, s.jiraClient)
	jiraTools.AddPriorityTools(s.mcpServer, s.jiraClient)
	jiraTools.AddTransitionTools(s.mcpServer, s.jiraClient, hasWritePermission)
	jiraTools.AddUserTools(s.mcpServer, s.jiraClient)
	jiraTools.AddWorklogTools(s.mcpServer, s.jiraClient, hasWritePermission)
}

// addConfluenceTools registers all Confluence-related tools with the MCP server
func (s *Server) addConfluenceTools() {
	// Check if Confluence write permission is enabled
	hasWritePermission := s.config.Confluence.Permissions.Write
	
	confluenceTools.AddContentTools(s.mcpServer, s.confluenceClient, hasWritePermission)
	confluenceTools.AddSpaceTools(s.mcpServer, s.confluenceClient, hasWritePermission)
	confluenceTools.AddChildrenTools(s.mcpServer, s.confluenceClient, hasWritePermission)
	confluenceTools.AddLabelTools(s.mcpServer, s.confluenceClient, hasWritePermission)
	confluenceTools.AddUserTools(s.mcpServer, s.confluenceClient)
}

// addBitbucketTools registers all Bitbucket-related tools with the MCP server
func (s *Server) addBitbucketTools() {
	hasWritePermission := s.config.Bitbucket.Permissions.Write
	
	bitbucketTools.AddUserTools(s.mcpServer, s.bitbucketClient, hasWritePermission)
	bitbucketTools.AddProjectTools(s.mcpServer, s.bitbucketClient, hasWritePermission)
	bitbucketTools.AddBranchTools(s.mcpServer, s.bitbucketClient, hasWritePermission)
	bitbucketTools.AddCommitTools(s.mcpServer, s.bitbucketClient, hasWritePermission)
	bitbucketTools.AddPullRequestTools(s.mcpServer, s.bitbucketClient, hasWritePermission)
	bitbucketTools.AddAttachmentTools(s.mcpServer, s.bitbucketClient, hasWritePermission)
	bitbucketTools.AddTagTools(s.mcpServer, s.bitbucketClient, hasWritePermission)
	bitbucketTools.AddRepositoryTools(s.mcpServer, s.bitbucketClient, hasWritePermission)
}
