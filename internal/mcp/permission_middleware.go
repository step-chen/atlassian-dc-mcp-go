package mcp

import (
	"context"
	"fmt"

	"atlassian-dc-mcp-go/internal/config"

	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// PermissionType represents the type of permission (only write permission is checked)
type PermissionType int

const (
	// WritePermission represents write permissions
	WritePermission PermissionType = iota
)

// IsWriteOperation determines if a method is a write operation
func IsWriteOperation(method string) bool {
	writeMethods := []string{
		// Jira write operations
		"jira_create_issue",
		"jira_update_issue",
		"jira_create_subtask",
		"jira_add_comment",
		"jira_transition_issue",

		// Confluence write operations
		"confluence_create_content",
		"confluence_update_content",
		"confluence_delete_content",
		"confluence_create_content_comment",

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

// GetServiceFromMethod extracts the service name from a method name
func GetServiceFromMethod(method string) string {
	// Method names follow the pattern "service_method"
	switch {
	case len(method) >= 5 && method[:5] == "jira_":
		return "jira"
	case len(method) >= 11 && method[:11] == "confluence_":
		return "confluence"
	case len(method) >= 10 && method[:10] == "bitbucket_":
		return "bitbucket"
	default:
		return ""
	}
}

// CheckWritePermission checks if a service has write permission
func CheckWritePermission(method string, config *config.Config) bool {
	service := GetServiceFromMethod(method)
	
	switch service {
	case "jira":
		return config.Jira.Permissions.Write
	case "confluence":
		return config.Confluence.Permissions.Write
	case "bitbucket":
		return config.Bitbucket.Permissions.Write
	}

	return false
}

// CheckPermissionMiddleware creates a middleware that checks service permissions
func CheckPermissionMiddleware(config *config.Config) mcp.Middleware {
	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
			// Check if this is a write operation
			if IsWriteOperation(method) {
				// Check write permission
				hasPermission := CheckWritePermission(method, config)
				if !hasPermission {
					service := GetServiceFromMethod(method)
					return nil, fmt.Errorf("write operation not permitted for service: %s", service)
				}
			}

			// Continue with the next handler
			return next(ctx, method, req)
		}
	}
}