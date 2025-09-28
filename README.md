# Atlassian Data Center MCP (Model Context Protocol)

This project provides a Go-based Model Context Protocol (MCP) service for managing and interacting with Atlassian Data Center products including Jira, Confluence, and Bitbucket. It allows you to manage these products programmatically through a unified interface with configurable authentication and permissions.

## Features

- **Multi-Product Support**: Unified interface for Jira, Confluence, and Bitbucket
- **Model Context Protocol**: Exposes tools through the Model Context Protocol for all operations
- **Configuration Management**: Flexible configuration through files, environment variables, and hot reloading
- **Authentication**: Direct API token authentication for each service
- **Permissions**: Fine-grained read/write permissions for each service
- **Health Monitoring**: Built-in health checks for all services
- **Docker Support**: Ready for containerized deployment

## Running the Application

You can run the application in multiple ways: directly using Go, by building binaries first, or using Docker.

### Direct Execution

```bash
# Run the server
go run cmd/server/main.go

# Run with custom config file path
go run cmd/server/main.go -c /path/to/your/config.yaml

# Or using long form
go run cmd/server/main.go --config /path/to/your/config.yaml

# Show help
go run cmd/server/main.go -h

# Or using long form
go run cmd/server/main.go --help
```

### Building and Running Binaries

The project uses a Makefile to simplify the build process. All binaries are placed in the `dist` directory.

```bash
# Show all available commands
make help

# Build the server binary for your current OS
make build

# Build a statically linked release binary
make release

# Clean the build directory
make clean
```

After building the project, you can run the binary directly from the `dist` directory:

```bash
# Run the server
./dist/atlassian-dc-mcp-server

# Run with custom config file path
./dist/atlassian-dc-mcp-server -c /path/to/your/config.yaml
```

### Running with Docker

First, create a config.yaml file by copying the example configuration:

```bash
cp config.yaml.example config.yaml
```

Then edit `config.yaml` with your settings before running the Docker container.

The application can be run in three different modes via Docker, depending on the transport mode configured:

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

#### Stdio Mode (Default)

For stdio mode, no ports need to be exposed as communication happens through standard input/output. This mode is typically used when the server is started by another process that communicates with it directly through stdio pipes:

```bash
# Build the binary first
make build

# Run in stdio mode (no need for Docker)
./dist/atlassian-dc-mcp-server

# Or run directly with go
make run-server
```

Note: When using stdio mode, the server should not be run in Docker with port mappings as it doesn't require network communication.

#### MCP Server Configuration

When running the MCP server, make sure your configuration file or environment variables are properly set for the desired transport mode.

#### Docker Compose

A `docker-compose.yml` file is provided for easier deployment:

```bash
docker compose up
```

The docker-compose file supports configuration through environment variables. You can create a `.env` file to customize the deployment:

```env
PORT=8090
IMAGE_NAME=atlassian-dc-mcp-go
```

When using docker compose, the configuration can be provided via environment variables or the config.yaml file.
Just make sure your configuration is properly set for the desired transport mode.

## Configuration

The application requires a configuration file to operate. By default, it looks for `config.yaml` in the current directory, but you can specify a different path using the `-c` or `--config` flag.

A sample configuration file is provided as `config.yaml.example`. Copy this file to `config.yaml` and edit it with your settings:

```bash
cp config.yaml.example config.yaml
```

The configuration file contains the following sections:

#### Server Configuration

- `port`: The port the server will listen on (default: 8090)
- `client_timeout`: Client timeout in seconds (default: 60)
- `transport`: Transport mode (stdio, sse, http)

#### Service Configuration

Each Atlassian service (Jira, Confluence, Bitbucket) has its own section with the following settings:

- `url`: The base URL of the service
- `token`: The API token for authentication
- `permissions`: A map of permissions for fine-grained access control

##### Permissions

Permissions allow you to control which operations can be performed on each service. Read permissions are always enabled, but write permissions can be selectively enabled or disabled.

###### Jira Permissions

- `jira_create_issue`: Create new issues
- `jira_update_issue`: Update existing issues
- `jira_transition_issue`: Transition issues between statuses
- `jira_add_comment`: Add comments to issues
- `jira_create_subtask`: Create subtasks
- `jira_set_issue_estimation_for_board`: Set issue estimation for a board
- `jira_add_worklogs`: Add worklogs to issues

###### Confluence Permissions

- `confluence_create_content`: Create new content
- `confluence_update_content`: Update existing content
- `confluence_delete_content`: Delete content
- `confluence_add_comment`: Add comments to content

###### Bitbucket Permissions

- `bitbucket_merge_pull_request`: Merge pull requests
- `bitbucket_decline_pull_request`: Decline pull requests
- `bitbucket_add_pull_request_comment`: Add comments to pull requests
- `bitbucket_create_attachment`: Create attachments
- `bitbucket_delete_attachment`: Delete attachments
- `bitbucket_update_pull_request_status`: Update pull request participant status (approve, request changes, or reset approval)

### Environment Variables

```bash
# Server configuration
export MCP_PORT=8090  # Server port (default: 8090)
export MCP_TRANSPORT="stdio"  # Transport mode (stdio, sse, http)

# Logging configuration
export MCP_LOGGING_DEVELOPMENT=true  # Enable human-friendly console output
export MCP_LOGGING_LEVEL="info"  # Log level (debug, info, warn, error, fatal)
export MCP_LOGGING_FILE_PATH="/var/log/atlassian-dc-mcp.log"  # Path to log file
export MCP_LOGGING_FILE_LEVEL="debug"  # Log level for file output

# Service configurations
export MCP_JIRA_URL="https://jira.example.com"
export MCP_JIRA_TOKEN="your-jira-api-token"
# Note: READ permissions are always enabled and cannot be disabled
# Jira write permissions:
export MCP_JIRA_PERMISSIONS_JIRA_CREATE_ISSUE=false
export MCP_JIRA_PERMISSIONS_JIRA_UPDATE_ISSUE=false
export MCP_JIRA_PERMISSIONS_JIRA_TRANSITION_ISSUE=false
export MCP_JIRA_PERMISSIONS_JIRA_ADD_COMMENT=false
export MCP_JIRA_PERMISSIONS_JIRA_CREATE_SUBTASK=false
export MCP_JIRA_PERMISSIONS_JIRA_SET_ISSUE_ESTIMATION_FOR_BOARD=false
export MCP_JIRA_PERMISSIONS_JIRA_ADD_WORKLOGS=false

export MCP_CONFLUENCE_URL="https://confluence.example.com"
export MCP_CONFLUENCE_TOKEN="your-confluence-api-token"
# Note: READ permissions are always enabled and cannot be disabled
# Confluence write permissions:
export MCP_CONFLUENCE_PERMISSIONS_CONFLUENCE_CREATE_CONTENT=false
export MCP_CONFLUENCE_PERMISSIONS_CONFLUENCE_UPDATE_CONTENT=false
export MCP_CONFLUENCE_PERMISSIONS_CONFLUENCE_DELETE_CONTENT=false
export MCP_CONFLUENCE_PERMISSIONS_CONFLUENCE_ADD_COMMENT=false

export MCP_BITBUCKET_URL="https://bitbucket.example.com"
export MCP_BITBUCKET_TOKEN="your-bitbucket-api-token"
# Note: READ permissions are always enabled and cannot be disabled
# Bitbucket write permissions:
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_MERGE_PULL_REQUEST=false
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_DECLINE_PULL_REQUEST=false
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_ADD_PULL_REQUEST_COMMENT=false
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_CREATE_ATTACHMENT=false
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_DELETE_ATTACHMENT=false
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_UPDATE_PULL_REQUEST_STATUS=false

```

You can also use a `.env` file to manage environment variables:

```
# .env file
MCP_PORT=8090
MCP_TRANSPORT="stdio"
MCP_CLIENT_TIMEOUT=60

# Logging configuration
MCP_LOGGING_DEVELOPMENT=true
MCP_LOGGING_LEVEL="info"

# Jira configuration
MCP_JIRA_URL="https://jira.example.com"
MCP_JIRA_TOKEN="your-jira-api-token"
# Note: READ permissions are always enabled and cannot be disabled
# Jira write permissions:
MCP_JIRA_PERMISSIONS_JIRA_CREATE_ISSUE=false
MCP_JIRA_PERMISSIONS_JIRA_UPDATE_ISSUE=false
MCP_JIRA_PERMISSIONS_JIRA_TRANSITION_ISSUE=false
MCP_JIRA_PERMISSIONS_JIRA_ADD_COMMENT=false
MCP_JIRA_PERMISSIONS_JIRA_CREATE_SUBTASK=false
MCP_JIRA_PERMISSIONS_JIRA_SET_ISSUE_ESTIMATION_FOR_BOARD=false
MCP_JIRA_PERMISSIONS_JIRA_ADD_WORKLOGS=false

# Confluence configuration
MCP_CONFLUENCE_URL="https://confluence.example.com"
MCP_CONFLUENCE_TOKEN="your-confluence-api-token"
# Note: READ permissions are always enabled and cannot be disabled
# Confluence write permissions:
MCP_CONFLUENCE_PERMISSIONS_CONFLUENCE_CREATE_CONTENT=false
MCP_CONFLUENCE_PERMISSIONS_CONFLUENCE_UPDATE_CONTENT=false
MCP_CONFLUENCE_PERMISSIONS_CONFLUENCE_DELETE_CONTENT=false
MCP_CONFLUENCE_PERMISSIONS_CONFLUENCE_ADD_COMMENT=false

# Bitbucket configuration
MCP_BITBUCKET_URL="https://bitbucket.example.com"
MCP_BITBUCKET_TOKEN="your-bitbucket-api-token"
# Note: READ permissions are always enabled and cannot be disabled
# Bitbucket write permissions:
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_MERGE_PULL_REQUEST=false
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_DECLINE_PULL_REQUEST=false
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_ADD_PULL_REQUEST_COMMENT=false
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_CREATE_ATTACHMENT=false
MCP_BITBUCKET_PERMISSIONS_BITBUCKET_DELETE_ATTACHMENT=false
```

Create .env file from example:
```bash
cp .env.example .env
```

Then edit `.env` file with your settings. Configuration changes will be automatically applied without restarting the service.

### Hot Reloading

The application supports hot reloading of configuration. When the config.yaml file is modified, 
the changes are automatically applied without restarting the service.

## Client Configuration

To use this service with an AI assistant that supports MCP, you need to configure the AI assistant to connect to the MCP server. Here's an example configuration for the MCP servers:


This configuration provides two ways to use the service with an AI assistant:

1. **stdio-based method** (`atlassian-dc-mcp-stdio`): 
   - The AI assistant will directly execute the server binary
   - No prior setup is required as the server starts on-demand
```json
{
  "mcpServers": {
    "atlassian-dc-mcp-stdio": {
      "command": "/path/to/atlassian-dc-mcp-go/dist/atlassian-dc-mcp-server",
      "args": [
        "-c",
        "/path/to/atlassian-dc-mcp-go/config_stdio.yaml"
      ]
    }
  }
}
```

2. **HTTP-based method** (`atlassian-dc-mcp-http`):
   - The AI assistant connects to a running HTTP server at the specified URL
   - You must start the HTTP server first before using this method
   - To start the HTTP server, you can use Docker:
```json
{
  "mcpServers": {
    "atlassian-dc-mcp-http": {
      "url": "http://localhost:8090/mcp"
    }
  }
}
```

   ```bash
   # Make sure you have configured config_http.yaml with your settings
   cp config.yaml.example config_http.yaml
   # Edit config_http.yaml to set transport: "http" and your service credentials
   
   # Start the HTTP server using Docker
   docker run -p 8090:8090 -v $(pwd)/config_http.yaml:/app/config.yaml atlassian-dc-mcp-go
   ```

   - Alternatively, you can start it directly:
   ```bash
   # Build the binary first
   make build
   
   # Run with HTTP config
   ./dist/atlassian-dc-mcp-server -c config_http.yaml
   ```

## Permissions

Each service (Jira, Confluence, Bitbucket) supports fine-grained permission controls:
- `read`: Allows read operations (always available, cannot be disabled)
- Individual write permissions for each operation (default: false for all)

Read operations are always available and cannot be disabled. Write operations must be explicitly enabled by setting the specific permission to `true` in the configuration.

Permissions are validated at startup and services with invalid configurations will be disabled.

For example, to allow creating Jira issues but not deleting them:
```yaml
jira:
  url: "https://jira.example.com"
  token: "your-jira-api-token"
  permissions:
    # Note: READ permissions are always enabled and cannot be disabled
    # Jira write permissions:
    jira_create_issue: true     # Allow creating issues
    jira_update_issue: false    # Don't allow updating issues
    jira_delete_issue: false    # Don't allow deleting issues
    jira_create_subtask: false  # Don't allow creating subtasks
    # ... other permissions
```

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

## Lingma Rules

This project includes predefined Lingma rules that demonstrate how to use the Atlassian Data Center MCP service for automated code review tasks. For detailed information on how to use these rules, please refer to the [Lingma Rules documentation](docs/lingma-rules.md).

The [code_review.md](.lingma/rules/code_review.md) rule file provides an example configuration showing how to leverage the Atlassian Data Center MCP service for automated Bitbucket PR code reviews. These rules help you:

- Automate code review processes using the MCP service
- Standardize code quality and security checks
- Improve interaction efficiency with AI assistants during code reviews

## Development

### Prerequisites

- Go 1.20 or higher
- Docker (for containerization)

### Building

It is recommended to use the Makefile to build the project, which ensures all build artifacts are placed in the unified `dist` directory:

```bash
# Build server binary to dist directory
make build
```

To build binaries for specific operating systems:
```bash
# Build for Linux
make build-linux

# Build for Windows
make build-windows

# Build for macOS
make build-macos
```

Benefits of using make commands:
- All build artifacts are placed in the `dist` directory
- Automatically handles cross-platform builds
- Ensures consistent build parameters

### Testing

```bash
make test
# or
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## References

- [Confluence REST API](https://developer.atlassian.com/server/confluence/rest/v1010/intro/#about)
- [Jira REST API](https://developer.atlassian.com/server/jira/platform/rest/v11000/intro/#gettingstarted)
- [Bitbucket REST API](https://developer.atlassian.com/server/bitbucket/rest/v1000/intro/#about)
- [Model Context Protocol Go SDK](https://github.com/modelcontextprotocol/go-sdk)
