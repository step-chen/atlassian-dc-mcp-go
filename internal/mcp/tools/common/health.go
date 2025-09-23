package common

import (
	"context"
	"sync"

	"atlassian-dc-mcp-go/internal/client/bitbucket"
	"atlassian-dc-mcp-go/internal/client/confluence"
	"atlassian-dc-mcp-go/internal/client/jira"
	"atlassian-dc-mcp-go/internal/config"
	"atlassian-dc-mcp-go/internal/mcp/utils"

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

// Handler provides a struct for storing the app server
type Handler struct {
	appServer AppServer
}

// NewHandler creates a new Handler with the provided app server
func NewHandler(appServer AppServer) *Handler {
	return &Handler{
		appServer: appServer,
	}
}

// HealthCheckInput represents the input for the health check tool
type HealthCheckInput struct{}

// ServiceStatus represents the status of a service
type ServiceStatus struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

// HealthCheckOutput represents the output of the health check tool
type HealthCheckOutput struct {
	Jira       ServiceStatus `json:"jira"`
	Confluence ServiceStatus `json:"confluence"`
	Bitbucket  ServiceStatus `json:"bitbucket"`
}

// checkJiraHealth checks the health of the Jira service
func checkJiraHealth(client *jira.JiraClient) map[string]interface{} {
	result := make(map[string]interface{})

	if client != nil {
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
func checkConfluenceHealth(client *confluence.ConfluenceClient) map[string]interface{} {
	result := make(map[string]interface{})

	if client != nil {
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
func checkBitbucketHealth(client *bitbucket.BitbucketClient) map[string]interface{} {
	result := make(map[string]interface{})

	if client != nil {
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
func performHealthCheck(jiraClient *jira.JiraClient, confluenceClient *confluence.ConfluenceClient, bitbucketClient *bitbucket.BitbucketClient) HealthCheckOutput {
	var wg sync.WaitGroup

	status := HealthCheckOutput{
		Jira:       ServiceStatus{},
		Confluence: ServiceStatus{},
		Bitbucket:  ServiceStatus{},
	}

	// Check Jira
	wg.Add(1)
	go func() {
		defer wg.Done()
		jiraStatus := checkJiraHealth(jiraClient)
		status.Jira = ServiceStatus{
			Status:  jiraStatus["status"].(string),
			Message: jiraStatus["message"].(string),
		}
	}()

	// Check Confluence
	wg.Add(1)
	go func() {
		defer wg.Done()
		confluenceStatus := checkConfluenceHealth(confluenceClient)
		status.Confluence = ServiceStatus{
			Status:  confluenceStatus["status"].(string),
			Message: confluenceStatus["message"].(string),
		}
	}()

	// Check Bitbucket
	wg.Add(1)
	go func() {
		defer wg.Done()
		bitbucketStatus := checkBitbucketHealth(bitbucketClient)
		status.Bitbucket = ServiceStatus{
			Status:  bitbucketStatus["status"].(string),
			Message: bitbucketStatus["message"].(string),
		}
	}()

	wg.Wait()

	return status
}

// healthCheckHandler handles the health check tool call using the new generic API.
func (h *Handler) healthCheckHandler(ctx context.Context, req *mcp.CallToolRequest, input HealthCheckInput) (*mcp.CallToolResult, HealthCheckOutput, error) {
	status := performHealthCheck(
		h.appServer.GetJiraClient(),
		h.appServer.GetConfluenceClient(),
		h.appServer.GetBitbucketClient(),
	)
	return nil, status, nil
}

// AddHealthCheckTool registers the health check tool with the MCP server using the new generic API.
func AddHealthCheckTool(server *mcp.Server, appServer AppServer) {
	handler := NewHandler(appServer)
	utils.RegisterTool[HealthCheckInput, HealthCheckOutput](server, "health_check", "Check the health status of the configured services (Jira, Confluence, Bitbucket).", handler.healthCheckHandler)
}
