package common

import (
	"context"

	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// CapabilitiesInput represents the input for the capabilities tool
type CapabilitiesInput types.EmptyInput

// CapabilitiesOutput represents the output of the capabilities tool
type CapabilitiesOutput struct {
	Capabilities types.MapOutput `json:"capabilities"`
}

// capabilitiesHandler handles getting server capabilities
func capabilitiesHandler(ctx context.Context, req *mcp.CallToolRequest, input CapabilitiesInput) (*mcp.CallToolResult, CapabilitiesOutput, error) {
	capabilities := types.MapOutput{
		"description": "Atlassian DC MCP Server",
		"features": []string{
			"health_check",
			"jira_integration",
			"confluence_integration",
			"bitbucket_integration",
		},
		"version": "1.0.0",
	}

	output := CapabilitiesOutput{Capabilities: capabilities}
	return nil, output, nil
}

// AddCapabilitiesTool registers the capabilities tool with the MCP server.
func AddCapabilitiesTool(server *mcp.Server) {
	utils.RegisterTool[CapabilitiesInput, CapabilitiesOutput](server, "capabilities", "Get detailed information about what tools and operations are supported by this server, including Jira, Confluence, and Bitbucket integrations.", capabilitiesHandler)
}
