// Package mcp provides the core server implementation for the Management Control Plane.
// It handles the initialization and management of various Atlassian services
// including Jira, Confluence, and Bitbucket through the Model Context Protocol (MCP).
package mcp

import (
	"context"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/config"
	bitbucketTools "atlassian-dc-mcp-go/internal/mcp/tools/bitbucket"
	confluenceTools "atlassian-dc-mcp-go/internal/mcp/tools/confluence"
	jiraTools "atlassian-dc-mcp-go/internal/mcp/tools/jira"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

type Server struct {
	config           *config.Config
	jiraClient       *jira.JiraClient
	confluenceClient *confluence.ConfluenceClient
	bitbucketClient  *bitbucket.BitbucketClient
	mcpServer        *mcp.Server
}

// NewServer creates a new MCP server instance with the provided configuration
func NewServer(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
	}
}

// Initialize sets up the server with clients for Jira, Confluence, and Bitbucket based on configuration
func (s *Server) Initialize(ctx context.Context) error {
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

// Start begins the MCP server using the stdio transport
func (s *Server) Start(ctx context.Context) error {
	return s.mcpServer.Run(ctx, &mcp.StdioTransport{})
}

// Stop gracefully stops the MCP server
func (s *Server) Stop() {
}

// addTools registers all available tools with the MCP server
func (s *Server) addTools() {
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