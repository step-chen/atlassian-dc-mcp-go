// Package tools provides utility functions for MCP tool operations.
package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// HandleToolError creates a standardized error response for tool operations
func HandleToolError(err error, operation string) (*mcp.CallToolResult, map[string]interface{}, error) {
	errorMessage := fmt.Sprintf("%s failed: %v", operation, err)
	result := &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: errorMessage},
		},
		IsError: true,
	}
	return result, nil, err
}

// CreateToolResult converts data to structured content and wraps it in a CallToolResult
func CreateToolResult(data interface{}) (*mcp.CallToolResult, error) {
	// If data is already a string, use it directly as text content
	if text, ok := data.(string); ok {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: text},
			},
		}, nil
	}

	// Try to use data as structured content if it's a map or slice
	switch v := data.(type) {
	case map[string]interface{}, []interface{}, []map[string]interface{}:
		// Use as structured content
		jsonData, err := json.Marshal(v)
		if err != nil {
			// Fall back to text content if marshaling fails
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Result: %+v", v)},
				},
			}, nil
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: string(jsonData)},
			},
			StructuredContent: v,
		}, nil
	}

	// For other types, marshal to JSON and use as both text and structured content if possible
	jsonData, err := json.Marshal(data)
	if err != nil {
		// Fall back to simple string representation
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Result: %+v", data)},
			},
		}, nil
	}

	// Try to unmarshal as generic JSON for structured content
	var structured interface{}
	if err := json.Unmarshal(jsonData, &structured); err == nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: string(jsonData)},
			},
			StructuredContent: structured,
		}, nil
	}

	// Fall back to text only
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(jsonData)},
		},
	}, nil
}

// GetStringArg retrieves a string value from the args map
func GetStringArg(args map[string]interface{}, key string) (string, bool) {
	if val, ok := args[key]; ok {
		if str, ok := val.(string); ok {
			return str, true
		}
	}
	return "", false
}

// GetIntArg retrieves an integer value from the args map with a default fallback
func GetIntArg(args map[string]interface{}, key string, defaultValue int) int {
	if val, ok := args[key]; ok {
		if num, ok := val.(float64); ok {
			return int(num)
		}
	}
	return defaultValue
}

// GetBoolArg retrieves a boolean value from the args map with a default fallback
func GetBoolArg(args map[string]interface{}, key string, defaultValue bool) bool {
	if val, ok := args[key]; ok {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return defaultValue
}

// GetStringSliceArg retrieves a string slice from the args map
func GetStringSliceArg(args map[string]interface{}, key string) []string {
	result := []string{}
	if val, ok := args[key].([]interface{}); ok {
		for _, v := range val {
			if str, ok := v.(string); ok {
				result = append(result, str)
			}
		}
	}
	return result
}

// HandleToolOperation simplifies tool operation handling by wrapping the common pattern
func HandleToolOperation(operationName string, operation func() (interface{}, error)) (*mcp.CallToolResult, map[string]interface{}, error) {
	return HandleToolOperationWithContext(context.Background(), operationName, operation)
}

// HandleToolOperationWithContext simplifies tool operation handling with context support and timing
func HandleToolOperationWithContext(ctx context.Context, operationName string, operation func() (interface{}, error)) (*mcp.CallToolResult, map[string]interface{}, error) {
	startTime := time.Now()

	data, err := operation()

	// Calculate operation duration
	duration := time.Since(startTime)

	if err != nil {
		return HandleToolError(fmt.Errorf("failed to %s after %v: %w", operationName, duration, err), operationName)
	}

	result, err := CreateToolResult(data)
	if err != nil {
		return HandleToolError(fmt.Errorf("failed to create result after %v: %w", duration, err), operationName)
	}

	// Convert data to map[string]interface{} if possible
	resultMap, _ := data.(map[string]interface{})

	return result, resultMap, nil
}