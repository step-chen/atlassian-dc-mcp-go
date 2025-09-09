# Atlassian Data Center MCP (Model Context Protocol) for Go

This project provides a Go-based Model Context Protocol (MCP) service for managing and interacting with Atlassian Data Center products including Jira, Confluence, and Bitbucket. It allows you to manage these products programmatically through a unified interface with configurable authentication and permissions.

## Features

- **Multi-Product Support**: Unified interface for Jira, Confluence, and Bitbucket
- **Model Context Protocol**: Exposes tools through the Model Context Protocol for all operations
- **Configuration Management**: Flexible configuration through files, environment variables, and hot reloading
- **Authentication**: Direct API token authentication for each service
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
# Server port (default is 8090 if not specified)
port: 8090

# Client timeout in seconds (default is 60 seconds if not specified)
client_timeout: 60

# Transport mode (stdio, sse, http)
# stdio: Standard input/output communication (default)
# sse: Server-Sent Events communication
# http: HTTP streaming communication on port 8090
transport: "stdio"

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

```

### Environment Variables


```bash
# Server configuration
export MCP_PORT=8090  # Server port (default: 8090)
export MCP_TRANSPORT="stdio"  # Transport mode (stdio, sse, http)

# Logging configuration
export MCP_LOGGING_DEVELOPMENT=true  # Enable human-friendly console output
export MCP_LOGGING_LEVEL="info"  # Log level (debug, info, warn, error, fatal)

# Service configurations
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

You can also use a `.env` file to manage environment variables:

```
# .env file
MCP_PORT=8090
IMAGE_NAME=github.com/step-chen/atlassian-dc-mcp
```

You can also use a `.env` file to manage environment variables:

```
# .env file
PORT=8090
IMAGE_NAME=github.com/step-chen/atlassian-dc-mcp
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


## Running the Application

### Direct Execution

```bash
# Run the server
go run cmd/server/main.go

# Run the client
go run cmd/client/main.go
```

### Running with Docker

First, create a config.yaml file by copying the example configuration:

```bash
cp config.yaml.example config.yaml
```

Then edit `config.yaml` with your settings before running the Docker container.

The application can be run in three different modes via Docker, depending on the transport mode configured:

#### Stdio Mode (Default)

For stdio mode, no ports need to be exposed as communication happens through standard input/output:

```bash
docker build -t atlassian-dc-mcp-go .
docker run -v $(pwd)/config.yaml:/app/config.yaml atlassian-dc-mcp-go
```

#### HTTP Mode

For HTTP mode, you need to expose the configured port:

```bash
# Using default port 8090
docker build -t atlassian-dc-mcp-go .
docker run -p 8090:8090 -v $(pwd)/config.yaml:/app/config.yaml atlassian-dc-mcp-go

# Using custom port
docker run -p 9000:9000 -e MCP_PORT=9000 -v $(pwd)/config.yaml:/app/config.yaml atlassian-dc-mcp-go
```

In your config.yaml, make sure to set:
```yaml
transport: "http"
port: 8090  # or your custom port
```

#### SSE Mode

For SSE mode, you also need to expose the configured port:

```bash
# Using default port 8090
docker build -t atlassian-dc-mcp-go .
docker run -p 8090:8090 -v $(pwd)/config.yaml:/app/config.yaml atlassian-dc-mcp-go
```

In your config.yaml, make sure to set:
```yaml
transport: "sse"
port: 8090
```

#### MCP Server Client Configuration Examples

When connecting to the MCP server as a client, you can use the following configurations based on your transport mode:

##### Stdio Mode Client Configuration

For stdio mode, the client communicates through standard input/output. This is the default mode:

```json
{
  "mcpServers": {
    "atlassian-dc-mcp-stdio": {
      "command": "docker",
      "args": [
        "run",
        "--rm",
        "-i",
        "--name",
        "atlassian-dc-mcp",
        "-v",
        "/path/to/your/config_stdio.yaml:/app/config.yaml:ro",
        "github.com/step-chen/atlassian-dc-mcp",
        "stdio"
      ]
    }
  }
}
```

##### HTTP Mode Client Configuration

For HTTP mode, first start the server using docker-compose, then configure the client to connect via HTTP URL:

```json
{
  "mcpServers": {
    "atlassian-dc-mcp": {
      "url": "http://localhost:8090/mcp"
    }
  }
}
```

To start the server in HTTP mode, use:

```bash
docker compose up
```

#### Docker Compose

A `docker-compose.yml` file is provided for easier deployment:

```bash
docker compose up
```

The docker-compose file supports configuration through environment variables. You can create a `.env` file to customize the deployment:

```env
PORT=8090
IMAGE_NAME=github.com/step-chen/atlassian-dc-mcp
```

When using docker compose, the configuration is the same as with direct Docker commands. 
Just make sure your `config.yaml` is properly configured for the desired transport mode.

## Cleanup Script

The project includes a [cleanup.sh](file:///home/stephen/workspace/atlassian-dc-mcp-go/cleanup.sh) script that stops and removes the container and deletes the associated Docker image:

```bash
# Stop and remove container, and delete image
./cleanup.sh
```

The script will:
1. Stop the running container
2. Remove the existing container
3. Remove the Docker image associated with the project

## Tools Documentation

### Capabilities Tool

Get information about what operations are supported. The response is filtered based on configured permissions.

### Health Check Tool

Check the health status of the services. The response includes the permission status of each service.

### Jira Tools

Tools for interacting with Jira:
- Get current user information
- Get issues
- Create issues
- And many more

### Confluence Tools

Tools for interacting with Confluence:
- Get current user information
- Get content
- Search content
- And more

### Bitbucket Tools

Tools for interacting with Bitbucket:
- Get current user information
- Get repositories
- Get commits
- And more
```
## Development

### Prerequisites

- Go 1.20 or higher
- Docker (for containerization)

### Building

```bash
go build -o atlassian-dc-mcp-go ./cmd/server
```

### Testing

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.