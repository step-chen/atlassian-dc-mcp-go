package bitbucket

import (
	"context"
	"fmt"
	"strings"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/mcp/utils"
	"atlassian-dc-mcp-go/internal/types"
)

// SearchCodeHandler handles code search requests with pre-parsed input
func (h *Handler) searchCodeHandler(ctx context.Context, req *mcp.CallToolRequest, input bitbucket.SearchCodeInput) (*mcp.CallToolResult, types.MapOutput, error) {
	// Perform the search
	searchResult, err := h.client.SearchCode(ctx, input)
	if err != nil {
		errorMessage := fmt.Sprintf("Error performing code search: %v", err)

		// Provide more helpful error message for query too long
		if strings.Contains(err.Error(), "QueryTooLongException") || strings.Contains(err.Error(), "maximum length of 250 characters") {
			errorMessage = "Error performing code search: Your search query is too long (exceeds 250 characters). Please try a shorter query. " +
				"Suggestions to reduce query length:\n" +
				"1. Use a more specific search term\n" +
				"2. Remove unnecessary filters like file patterns\n" +
				"3. Search in a specific repository instead of the entire project\n" +
				"4. Break complex searches into multiple simpler searches"
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: errorMessage,
				},
			},
			IsError: true,
		}, nil, nil
	}

	return nil, searchResult, nil
}

// AddSearchTools registers the search-related tools with the MCP server
func AddSearchTools(server *mcp.Server, client *bitbucket.BitbucketClient, permissions map[string]bool) {
	handler := NewHandler(client)

	utils.RegisterTool[bitbucket.SearchCodeInput, types.MapOutput](
		server,
		"bitbucket_search_code",
		"Search for code in Bitbucket repositories with enhanced contextual search capabilities",
		handler.searchCodeHandler)
}
