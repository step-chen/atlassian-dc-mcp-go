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
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Error performing code search: %v", err),
				},
			},
			IsError: true,
		}, nil, nil
	}

	// Format the result
	resultText := formatCodeSearchResult(searchResult, input.SearchQuery, input.SearchContext, input.ProjectKey, input.RepoSlug, input.Limit)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: resultText,
			},
		},
	}, searchResult, nil
}

// formatCodeSearchResult formats the search result for display
func formatCodeSearchResult(
	result types.MapOutput,
	searchQuery, searchContext, projectKey, repoSlug string, limit int) string {

	var resultBuilder strings.Builder

	// Build header
	resultBuilder.WriteString(fmt.Sprintf("Code search results for \"%s\"", searchQuery))
	if searchContext != "any" && searchContext != "" {
		resultBuilder.WriteString(fmt.Sprintf(" (context: %s)", searchContext))
	}
	resultBuilder.WriteString(fmt.Sprintf(" in %s", projectKey))
	if repoSlug != "" {
		resultBuilder.WriteString(fmt.Sprintf("/%s", repoSlug))
	}

	// Add results
	codeResults, ok := result["code"].(map[string]interface{})
	if !ok {
		resultBuilder.WriteString("\n\nNo code results found.")
		return resultBuilder.String()
	}

	count, ok := codeResults["count"].(float64)
	if !ok {
		count = 0
	}
	values, ok := codeResults["values"].([]interface{})
	if !ok {
		values = []interface{}{}
	}

	if count == 0 || len(values) == 0 {
		resultBuilder.WriteString("\n\nNo matches found.")
		return resultBuilder.String()
	}

	resultBuilder.WriteString(fmt.Sprintf("\n\nFound %d matches:", int(count)))

	if limit == 0 {
		limit = 5 // Default limit
	}

	// Format each result
	for i, value := range values {
		if i >= limit {
			resultBuilder.WriteString(fmt.Sprintf("\n\n... (showing first %d results)", limit))
			break
		}

		if hit, ok := value.(map[string]interface{}); ok {
			resultBuilder.WriteString(formatCodeHit(hit))
		}
	}

	return resultBuilder.String()
}

// formatCodeHit formats a single code search hit
func formatCodeHit(hit map[string]interface{}) string {
	var result strings.Builder

	// Get repository info
	if repo, ok := hit["repository"].(map[string]interface{}); ok {
		if name, ok := repo["name"].(string); ok {
			result.WriteString(fmt.Sprintf("\n\nRepository: %s", name))
		}
	}

	// Get file info
	if file, ok := hit["file"].(map[string]interface{}); ok {
		if path, ok := file["path"].(string); ok {
			result.WriteString(fmt.Sprintf("\nFile: %s", path))
		}
	}

	// Get line numbers
	if content, ok := hit["content"].(map[string]interface{}); ok {
		if lineNumbers, ok := content["lineNumbers"].([]interface{}); ok && len(lineNumbers) > 0 {
			var lines []string
			for _, ln := range lineNumbers {
				if line, ok := ln.(float64); ok {
					lines = append(lines, fmt.Sprintf("%d", int(line)))
				}
			}
			result.WriteString(fmt.Sprintf("\nLines: %s", strings.Join(lines, ", ")))
		}
	}

	// Get code fragments
	if fragments, ok := hit["fragments"].([]interface{}); ok {
		for _, f := range fragments {
			if fragment, ok := f.(map[string]interface{}); ok {
				if text, ok := fragment["text"].(string); ok {
					result.WriteString(fmt.Sprintf("\nCode: %s", text))
				}
			}
		}
	}

	return result.String()
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
