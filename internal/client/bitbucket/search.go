package bitbucket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"atlassian-dc-mcp-go/internal/types"
	"atlassian-dc-mcp-go/internal/utils"
)

// SearchCode performs a code search in Bitbucket
func (c *BitbucketClient) SearchCode(input SearchCodeInput) (types.MapOutput, error) {
	var queryBuilder strings.Builder
	// Build the enhanced query string
	queryBuilder.WriteString(fmt.Sprintf("project:%s", input.ProjectKey))
	if input.RepoSlug != "" {
		queryBuilder.WriteString(fmt.Sprintf(" repo:%s", input.RepoSlug))
	}
	if input.FilePattern != nil && *input.FilePattern != "" {
		queryBuilder.WriteString(fmt.Sprintf(" path:%s", *input.FilePattern))
	}

	// Build smart search patterns
	smartQuery := c.buildSmartQuery(input.SearchQuery, input.SearchContext)
	queryBuilder.WriteString(fmt.Sprintf(" %s", smartQuery))

	// Prepare the request payload
	payload := BitbucketServerSearchRequest{
		Query: strings.TrimSpace(queryBuilder.String()),
		Entities: BitbucketSearchRequestEntities{
			Code: BitbucketCodeEntity{
				Start: input.Start,
				Limit: input.Limit,
			},
		},
	}

	// Marshal payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	var searchResult types.MapOutput
	if err := c.executeRequest(
		http.MethodPost,
		[]string{"rest", "search", "latest", "search"},
		nil,
		jsonPayload,
		&searchResult,
		utils.AcceptJSON,
	); err != nil {
		return nil, err
	}

	return searchResult, nil
}

// buildSmartQuery creates enhanced search queries based on context
func (c *BitbucketClient) buildSmartQuery(searchTerm, context string) string {
	patterns := c.buildContextualPatterns(searchTerm)

	var queryParts []string
	if context != "any" {
		if contextPatterns, ok := patterns[context]; ok {
			for _, pattern := range contextPatterns {
				queryParts = append(queryParts, pattern)
			}
		}
	} else {
		// For "any" context, include all patterns
		for _, contextPatterns := range patterns {
			queryParts = append(queryParts, contextPatterns...)
		}
	}

	if len(queryParts) == 0 {
		return searchTerm
	}

	return strings.Join(queryParts, " OR ")
}

// buildContextualPatterns creates contextual search patterns based on the search term
func (c *BitbucketClient) buildContextualPatterns(searchTerm string) map[string][]string {
	return map[string][]string{
		"assignment": {
			searchTerm + " =",           // Variable assignment
			searchTerm + ":",            // Object property, JSON key
			"= " + searchTerm,           // Right-hand assignment
		},
		"declaration": {
			searchTerm + " =",           // Variable definition
			searchTerm + ":",            // Object key, parameter definition
			"function " + searchTerm,    // Function declaration
			"class " + searchTerm,       // Class declaration
			"interface " + searchTerm,   // Interface declaration
			"const " + searchTerm,       // Const declaration
			"let " + searchTerm,         // Let declaration
			"var " + searchTerm,         // Var declaration
		},
		"usage": {
			"." + searchTerm + "(",      // Method call
			searchTerm + "(",            // Function call
			"(" + searchTerm + ")",      // Function parameter
		},
		"exact": {
			searchTerm,                  // Exact match
		},
	}
}