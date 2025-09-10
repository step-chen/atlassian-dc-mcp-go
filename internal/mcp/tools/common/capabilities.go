package common

import (
	"context"

	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// capabilitiesHandler returns basic information about server capabilities.
func capabilitiesHandler() mcp.ToolHandler {
	return func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		result, _, err := tools.HandleToolOperation("get capabilities", func() (interface{}, error) {
			// Return general information about the server capabilities
			capabilities := map[string]interface{}{
				"description": "Atlassian DC MCP Server",
				"features": []string{
					"health_check",
					"jira_integration",
					"confluence_integration",
					"bitbucket_integration",
				},
				"version": "1.0.0",
			}

			return map[string]interface{}{"capabilities": capabilities}, nil
		})
		return result, err
	}
}

// AddCapabilitiesTool registers the capabilities tool with the MCP server.
func AddCapabilitiesTool(server *mcp.Server) {
	server.AddTool(&mcp.Tool{
		Name:        "capabilities",
		Description: "Get detailed information about what tools and operations are supported by this server, including Jira, Confluence, and Bitbucket integrations.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
		},
	}, capabilitiesHandler())
}