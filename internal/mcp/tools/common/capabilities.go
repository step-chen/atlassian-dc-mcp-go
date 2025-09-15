package common

import (
	"context"

	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// CapabilitiesInput represents the input for the capabilities tool
type CapabilitiesInput struct{}

// CapabilitiesOutput represents the output of the capabilities tool
type CapabilitiesOutput struct {
	Capabilities map[string]interface{} `json:"capabilities"`
}

// capabilitiesHandler returns basic information about server capabilities.
func capabilitiesHandler(ctx context.Context, req *mcp.CallToolRequest, input CapabilitiesInput) (*mcp.CallToolResult, CapabilitiesOutput, error) {
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

	if err != nil {
		return nil, CapabilitiesOutput{}, err
	}

	// Convert result to CapabilitiesOutput
	capabilitiesMap, _ := result.StructuredContent.(map[string]interface{})
	capabilities, _ := capabilitiesMap["capabilities"].(map[string]interface{})

	return result, CapabilitiesOutput{Capabilities: capabilities}, nil
}

// AddCapabilitiesTool registers the capabilities tool with the MCP server.
func AddCapabilitiesTool(server *mcp.Server) {
	mcp.AddTool[CapabilitiesInput, CapabilitiesOutput](server, &mcp.Tool{
		Name:        "capabilities",
		Description: "Get detailed information about what tools and operations are supported by this server, including Jira, Confluence, and Bitbucket integrations.",
		InputSchema: &jsonschema.Schema{
			Type: "object",
		},
	}, capabilitiesHandler)
}
