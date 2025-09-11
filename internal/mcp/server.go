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

// PermissionType represents the type of permission required
type PermissionType string

const (
	ReadPermission  PermissionType = "read"
	WritePermission PermissionType = "write"
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

	// Add permission checking middleware
	s.mcpServer.AddReceivingMiddleware(s.checkPermissionMiddleware())

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

// checkPermissionMiddleware creates a middleware that checks service permissions
func (s *Server) checkPermissionMiddleware() mcp.Middleware {
	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
			// Check if this is a write operation
			isWriteOperation := s.isWriteOperation(method)
			
			if isWriteOperation {
				// Determine service from method name
				service := s.getServiceFromMethod(method)
				
				// Check write permission
				hasPermission := s.checkPermission(service, WritePermission)
				if !hasPermission {
					return nil, fmt.Errorf("write operation not permitted for service: %s", service)
				}
			}
			
			// Continue with the next handler
			return next(ctx, method, req)
		}
	}
}

// isWriteOperation determines if a method is a write operation
func (s *Server) isWriteOperation(method string) bool {
	writeMethods := []string{
		// Jira write operations
		"jira_create_issue",
		"jira_update_issue",
		"jira_add_comment",
		"jira_transition_issue",
		
		// Confluence write operations
		"confluence_create_content",
		"confluence_update_content",
		"confluence_delete_content",
		"confluence_add_comment",
		
		// Bitbucket write operations
		"bitbucket_create_repository",
		"bitbucket_update_repository",
		"bitbucket_delete_repository",
		"bitbucket_create_pull_request",
		"bitbucket_add_pull_request_comment",
		"bitbucket_merge_pull_request",
		"bitbucket_decline_pull_request",
	}
	
	for _, writeMethod := range writeMethods {
		if method == writeMethod {
			return true
		}
	}
	
	return false
}

// getServiceFromMethod extracts the service name from a method name
func (s *Server) getServiceFromMethod(method string) string {
	// Method names follow the pattern "service_method"
	switch {
	case len(method) >= 4 && method[:4] == "jira":
		return "jira"
	case len(method) >= 11 && method[:11] == "confluence_":
		return "confluence"
	case len(method) >= 9 && method[:9] == "bitbucket":
		return "bitbucket"
	default:
		return ""
	}
}

// checkPermission checks if a service has the required permission
func (s *Server) checkPermission(service string, permission PermissionType) bool {
	switch service {
	case "jira":
		switch permission {
		case ReadPermission:
			return s.config.Jira.Permissions.Read
		case WritePermission:
			return s.config.Jira.Permissions.Write
		}
	case "confluence":
		switch permission {
		case ReadPermission:
			return s.config.Confluence.Permissions.Read
		case WritePermission:
			return s.config.Confluence.Permissions.Write
		}
	case "bitbucket":
		switch permission {
		case ReadPermission:
			return s.config.Bitbucket.Permissions.Read
		case WritePermission:
			return s.config.Bitbucket.Permissions.Write
		}
	}
	
	return false
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
	jiraTools.AddIssueTools(s.mcpServer, s.jiraClient)
	jiraTools.AddBoardTools(s.mcpServer, s.jiraClient)
	jiraTools.AddProjectTools(s.mcpServer, s.jiraClient)
	jiraTools.AddCommentTools(s.mcpServer, s.jiraClient)
	jiraTools.AddIssueTypeTools(s.mcpServer, s.jiraClient)
	jiraTools.AddPriorityTools(s.mcpServer, s.jiraClient)
	jiraTools.AddTransitionTools(s.mcpServer, s.jiraClient)
	jiraTools.AddUserTools(s.mcpServer, s.jiraClient)
	jiraTools.AddWorklogTools(s.mcpServer, s.jiraClient)
}

// addConfluenceTools registers all Confluence-related tools with the MCP server
func (s *Server) addConfluenceTools() {
	confluenceTools.AddContentTools(s.mcpServer, s.confluenceClient)
	confluenceTools.AddChildrenTools(s.mcpServer, s.confluenceClient)
	confluenceTools.AddLabelTools(s.mcpServer, s.confluenceClient)
	confluenceTools.AddSpaceTools(s.mcpServer, s.confluenceClient)
	confluenceTools.AddUserTools(s.mcpServer, s.confluenceClient)
}

// addBitbucketTools registers all Bitbucket-related tools with the MCP server
func (s *Server) addBitbucketTools() {
	bitbucketTools.AddUserTools(s.mcpServer, s.bitbucketClient)
	bitbucketTools.AddProjectTools(s.mcpServer, s.bitbucketClient)
	bitbucketTools.AddBranchTools(s.mcpServer, s.bitbucketClient)
	bitbucketTools.AddCommitTools(s.mcpServer, s.bitbucketClient)
	bitbucketTools.AddPullRequestTools(s.mcpServer, s.bitbucketClient)
	bitbucketTools.AddAttachmentTools(s.mcpServer, s.bitbucketClient)
	bitbucketTools.AddTagTools(s.mcpServer, s.bitbucketClient)
	bitbucketTools.AddRepositoryTools(s.mcpServer, s.bitbucketClient)
}