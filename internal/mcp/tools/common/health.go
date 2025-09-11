package common

import (
	"context"
	"sync"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/mcp/tools"

	"github.com/google/jsonschema-go/jsonschema"
	mcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// AppServer is an interface that represents the main application server,
// providing access to clients and configuration.
type AppServer interface {
	GetConfig() *config.Config
	GetJiraClient() *jira.JiraClient
	GetConfluenceClient() *confluence.ConfluenceClient
	GetBitbucketClient() *bitbucket.BitbucketClient
}

// getPermissionsList converts Permissions struct to a list of permission strings
func getPermissionsList(permissions config.Permissions) []string {
	var perms []string
	if permissions.Write {
		perms = append(perms, "write")
	}
	return perms
}

// checkJiraHealth checks the health of the Jira service
func checkJiraHealth(client *jira.JiraClient, permissions config.Permissions) map[string]interface{} {
	result := make(map[string]interface{})
	
	if client != nil {
		result["permissions"] = getPermissionsList(permissions)
		_, err := client.GetCurrentUser()
		if err != nil {
			result["status"] = "error"
			result["message"] = err.Error()
		} else {
			result["status"] = "ok"
		}
	} else {
		result["status"] = "disabled"
	}
	
	return result
}

// checkConfluenceHealth checks the health of the Confluence service
func checkConfluenceHealth(client *confluence.ConfluenceClient, permissions config.Permissions) map[string]interface{} {
	result := make(map[string]interface{})
	
	if client != nil {
		result["permissions"] = getPermissionsList(permissions)
		_, err := client.GetCurrentUser()
		if err != nil {
			result["status"] = "error"
			result["message"] = err.Error()
		} else {
			result["status"] = "ok"
		}
	} else {
		result["status"] = "disabled"
	}
	
	return result
}

// checkBitbucketHealth checks the health of the Bitbucket service
func checkBitbucketHealth(client *bitbucket.BitbucketClient, permissions config.Permissions) map[string]interface{} {
	result := make(map[string]interface{})
	
	if client != nil {
		result["permissions"] = getPermissionsList(permissions)
		_, err := client.GetCurrentUser()
		if err != nil {
			result["status"] = "error"
			result["message"] = err.Error()
		} else {
			result["status"] = "ok"
		}
	} else {
		result["status"] = "disabled"
	}
	
	return result
}

// performHealthCheck executes health checks for all services and returns the status
func performHealthCheck(cfg *config.Config, jiraClient *jira.JiraClient, confluenceClient *confluence.ConfluenceClient, bitbucketClient *bitbucket.BitbucketClient) interface{} {
	var wg sync.WaitGroup
	
	status := struct {
		Jira       map[string]interface{} `json:"jira"`
		Confluence map[string]interface{} `json:"confluence"`
		Bitbucket  map[string]interface{} `json:"bitbucket"`
	}{
		Jira:       make(map[string]interface{}),
		Confluence: make(map[string]interface{}),
		Bitbucket:  make(map[string]interface{}),
	}

	// Check Jira
	wg.Add(1)
	go func() {
		defer wg.Done()
		status.Jira = checkJiraHealth(jiraClient, cfg.Jira.Permissions)
	}()

	// Check Confluence
	wg.Add(1)
	go func() {
		defer wg.Done()
		status.Confluence = checkConfluenceHealth(confluenceClient, cfg.Confluence.Permissions)
	}()

	// Check Bitbucket
	wg.Add(1)
	go func() {
		defer wg.Done()
		status.Bitbucket = checkBitbucketHealth(bitbucketClient, cfg.Bitbucket.Permissions)
	}()

	wg.Wait()
	
	return status
}

// healthCheckHandler handles the health check tool call.
func healthCheckHandler(appServer AppServer) mcp.ToolHandler {
	return func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		result, _, err := tools.HandleToolOperation("health check", func() (interface{}, error) {
			cfg := appServer.GetConfig()
			jiraClient := appServer.GetJiraClient()
			confluenceClient := appServer.GetConfluenceClient()
			bitbucketClient := appServer.GetBitbucketClient()
			
			status := performHealthCheck(cfg, jiraClient, confluenceClient, bitbucketClient)
			
			return status, nil
		})
		return result, err
	}
}

// AddHealthCheckTool registers the health check tool with the MCP server.
func AddHealthCheckTool(server *mcp.Server, appServer AppServer) {
	server.AddTool(&mcp.Tool{
		Name:        "health_check",
		Description: "Check the health status of the configured services (Jira, Confluence, Bitbucket).",
		InputSchema: &jsonschema.Schema{
			Type: "object",
		},
	}, healthCheckHandler(appServer))
}
