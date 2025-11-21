package bitbucket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"atlassian-dc-mcp-go/internal/client"
	"atlassian-dc-mcp-go/internal/types"
)

// SearchCode performs a code search in Bitbucket
func (c *BitbucketClient) SearchCode(ctx context.Context, input SearchCodeInput) (types.MapOutput, error) {
	if input.ProjectKey == "" {
		return nil, fmt.Errorf("project key is required")
	}

	if input.Limit < 0 {
		return nil, fmt.Errorf("limit must be non-negative")
	}

	if input.Start < 0 {
		return nil, fmt.Errorf("start must be non-negative")
	}

	// Pre-check for query length - estimate the maximum possible query length
	maxProjectKeyLen := len("project:") + len(input.ProjectKey) + 2 // 2 for quotes if needed
	maxRepoSlugLen := 0
	if input.RepoSlug != "" {
		maxRepoSlugLen = len(" repo:") + len(input.RepoSlug) + 2
	}
	maxFilePatternLen := 0
	if input.FilePattern != nil && *input.FilePattern != "" {
		maxFilePatternLen = len(" path:") + len(*input.FilePattern) + 2
	}

	// Estimate space needed for search context wrapping
	maxContextLen := len(" ()") + len(input.SearchQuery)

	// If we know we'll exceed the limit, try to optimize
	estimatedLen := maxProjectKeyLen + maxRepoSlugLen + maxFilePatternLen + maxContextLen
	if estimatedLen > 250 {
		// If we're definitely going to exceed the limit, we need to take action
		// Priority order: project > repo > file pattern > search query
		if maxProjectKeyLen+maxRepoSlugLen > 240 {
			// Even project + repo is too long, this is an edge case
			return nil, fmt.Errorf("project and repository combination is too long for search")
		}
	}

	var queryBuilder strings.Builder
	// Build the enhanced query string
	queryBuilder.WriteString(fmt.Sprintf("project:%s", escapeSearchTerm(input.ProjectKey)))
	if input.RepoSlug != "" {
		queryBuilder.WriteString(fmt.Sprintf(" repo:%s", escapeSearchTerm(input.RepoSlug)))
	}
	if input.FilePattern != nil && *input.FilePattern != "" {
		queryBuilder.WriteString(fmt.Sprintf(" path:%s", escapeSearchTerm(*input.FilePattern)))
	}

	// Build smart search patterns only if we have space
	queryString := strings.TrimSpace(queryBuilder.String())
	availableSpace := 250 - len(queryString)

	var smartQuery string
	if availableSpace > 20 { // Need at least some space for the search term
		smartQuery = c.buildSmartQuery(input.SearchQuery, input.SearchContext)
	}

	if smartQuery != "" && len(smartQuery) < availableSpace-3 {
		queryBuilder.WriteString(" (" + smartQuery + ")")
	} else if input.SearchQuery != "" && len(input.SearchQuery) < availableSpace-3 {
		// Fallback to simple search term if smart query is too long or failed
		queryBuilder.WriteString(" (" + input.SearchQuery + ")")
	} else if input.SearchQuery != "" {
		// Even the simple search term is too long, truncate it
		maxSearchLen := availableSpace - 3 // 3 for the " ()"
		if maxSearchLen > 3 {              // Only truncate if we have meaningful space
			truncatedQuery := input.SearchQuery
			if len(truncatedQuery) > maxSearchLen {
				truncatedQuery = truncatedQuery[:maxSearchLen-3] + "..."
			}
			queryBuilder.WriteString(" (" + truncatedQuery + ")")
		}
	}

	queryString = strings.TrimSpace(queryBuilder.String())

	// Final hard limit check
	if len(queryString) > 250 {
		queryString = queryString[:250]
	}

	limit := input.Limit
	if limit == 0 {
		limit = 25
	}

	// Prepare the request payload
	payload := BitbucketServerSearchRequest{
		Query: queryString,
		Entities: BitbucketSearchRequestEntities{
			Code: BitbucketCodeEntity{
				Start: input.Start,
				Limit: limit,
			},
		},
	}

	// Marshal payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var output types.MapOutput
	if err := client.ExecuteRequest(
		ctx,
		c.BaseClient,
		http.MethodPost,
		[]any{"rest", "search", "latest", "search"},
		nil,
		jsonPayload,
		client.AcceptJSON,
		&output,
	); err != nil {
		return nil, err
	}

	return output, nil
}

// escapeSearchTerm 对搜索词进行转义处理
func escapeSearchTerm(term string) string {
	if strings.ContainsAny(term, " \t\n\r") {
		return fmt.Sprintf("\"%s\"", strings.ReplaceAll(term, "\"", "\\\""))
	}
	return term
}

// buildSmartQuery creates enhanced search queries based on context
func (c *BitbucketClient) buildSmartQuery(searchTerm, context string) string {
	escapedSearchTerm := escapeSearchTerm(searchTerm)

	patterns := c.buildContextualPatterns(escapedSearchTerm)

	var queryParts []string
	if context != "any" && context != "" {
		if contextPatterns, ok := patterns[context]; ok {
			queryParts = append(queryParts, contextPatterns...)
		}
	} else {
		// For "any" context, include all patterns
		for _, contextPatterns := range patterns {
			queryParts = append(queryParts, contextPatterns...)
		}
	}

	if len(queryParts) == 0 {
		return escapedSearchTerm
	}

	return strings.Join(queryParts, " OR ")
}

// buildContextualPatterns creates contextual search patterns based on the search term
func (c *BitbucketClient) buildContextualPatterns(searchTerm string) map[string][]string {
	return map[string][]string{
		"assignment": {
			searchTerm + " =", // Variable assignment
			searchTerm + ":",  // Object property, JSON key
			"= " + searchTerm, // Right-hand assignment
		},
		"declaration": {
			searchTerm + " =",         // Variable definition
			searchTerm + ":",          // Object key, parameter definition
			"function " + searchTerm,  // Function declaration
			"class " + searchTerm,     // Class declaration
			"interface " + searchTerm, // Interface declaration
			"const " + searchTerm,     // Const declaration
			"let " + searchTerm,       // Let declaration
			"var " + searchTerm,       // Var declaration
		},
		"usage": {
			"." + searchTerm + "(", // Method call
			searchTerm + "(",       // Function call
			"(" + searchTerm + ")", // Function parameter
		},
		"exact": {
			searchTerm, // Exact match
		},
	}
}
