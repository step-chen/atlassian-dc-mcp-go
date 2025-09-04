# Atlassian Data Center MCP (Management Control Plane) for Go

This project provides a Go-based REST API service for managing and interacting with Atlassian Data Center products including Jira, Confluence, and Bitbucket. It allows you to manage these products programmatically through a unified API interface with configurable authentication and permissions.

## Features

- **Multi-Product Support**: Unified interface for Jira, Confluence, and Bitbucket
- **RESTful API**: Exposes a clean REST API for all operations
- **Configuration Management**: Flexible configuration through files, environment variables, and hot reloading
- **Authentication**: Secure API key authentication
- **Permissions**: Fine-grained read/write permissions for each service
- **Health Monitoring**: Built-in health checks for all services
- **Docker Support**: Ready for containerized deployment

## Configuration

The application supports multiple configuration sources with the following priority (from highest to lowest):
1. Environment variables
2. Configuration file (config.yaml)
3. Default values

### Configuration File

Copy the example configuration file and edit it:

```bash
cp config.yaml.example config.yaml
```

Then edit `config.yaml` with your settings:

```yaml
# Server port (default is 8080 if not specified)
port: 8080

# Logging configuration
logging:
  # Development mode enables human-friendly console output
  development: true
  
  # Log level (debug, info, warn, error, fatal)
  level: "info"

jira:
  url: "https://jira.example.com"
  token: "your-jira-api-token"
  permissions:
    read: true
    write: false

confluence:
  url: "https://confluence.example.com"
  token: "your-confluence-api-token"
  permissions:
    read: true
    write: false

bitbucket:
  url: "https://bitbucket.example.com"
  token: "your-bitbucket-api-token"
  permissions:
    read: true
    write: false

# API keys for accessing the MCP service
# Add one or more keys that will be used to authenticate requests
apiKeys:
  - "example_api_key_1"
  - "example_api_key_2"
```

### Environment Variables

All configuration options can be set using environment variables with the prefix `MCP_`. 
Nested properties are separated by underscores:

```bash
export MCP_PORT=8080
export MCP_JIRA_URL="https://jira.example.com"
export MCP_JIRA_TOKEN="your-jira-api-token"
export MCP_JIRA_PERMISSIONS_READ=true
export MCP_JIRA_PERMISSIONS_WRITE=false
export MCP_CONFLUENCE_URL="https://confluence.example.com"
export MCP_CONFLUENCE_TOKEN="your-confluence-api-token"
export MCP_CONFLUENCE_PERMISSIONS_READ=true
export MCP_CONFLUENCE_PERMISSIONS_WRITE=false
export MCP_BITBUCKET_URL="https://bitbucket.example.com"
export MCP_BITBUCKET_TOKEN="your-bitbucket-api-token"
export MCP_BITBUCKET_PERMISSIONS_READ=true
export MCP_BITBUCKET_PERMISSIONS_WRITE=false
```

### Hot Reloading

The application supports hot reloading of configuration. When the config.yaml file is modified, 
the changes are automatically applied without restarting the service.

## Permissions

Each service (Jira, Confluence, Bitbucket) supports fine-grained permission controls:
- `read`: Allows read operations (default: true)
- `write`: Allows write operations (default: false)

When a service is configured with `read: false`, all endpoints for that service will be disabled.
When a service is configured with `write: false`, only read endpoints will be available.

Permissions are validated at startup and services with invalid configurations will be disabled.

## API Keys

When API keys are configured, all endpoints (except `/capabilities` and `/health`) require authentication.
Provide the API key in the Authorization header using the Bearer scheme:

```bash
curl -H "Authorization: Bearer your_api_key_here" http://localhost:8080/jira/user/current
```

## Running the Application

### Direct Execution

```bash
go run main.go
```

### Docker

First, create a config.yaml file by copying the example configuration:

```bash
cp config.yaml.example config.yaml
```

Then edit `config.yaml` with your settings before running the Docker container.

```bash
docker build -t atlassian-dc-mcp-go .
docker run -p 8080:8080 atlassian-dc-mcp-go
```

### Docker Compose

First, create a config.yaml file by copying the example configuration:

```bash
cp config.yaml.example config.yaml
```

Then edit `config.yaml` with your settings before running with docker-compose.

```bash
docker-compose up
```

## API Documentation

### Capabilities Endpoint

Get information about what operations are supported. The response is filtered based on configured permissions:

```bash
curl http://localhost:8080/capabilities
```

### Health Check Endpoint

Check the health status of the services. The response includes the permission status of each service:

```bash
curl http://localhost:8080/health
```

### Jira Endpoints

Get current user information:
```bash
curl -H "Authorization: Bearer your_api_key_here" http://localhost:8080/jira/user/current
```

### Confluence Endpoints

Get current user information:
```bash
curl -H "Authorization: Bearer your_api_key_here" http://localhost:8080/confluence/user/current
```

### Bitbucket Endpoints

Get current user information:
```bash
curl -H "Authorization: Bearer your_api_key_here" http://localhost:8080/bitbucket/user/current
```

## Development

### Prerequisites

- Go 1.20 or higher
- Docker (for containerization)

### Building

```bash
go build -o atlassian-dc-mcp-go
```

### Testing

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.