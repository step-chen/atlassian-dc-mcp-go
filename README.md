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
  
  # Path to log file (optional)
  # If not specified, logs go to console only
  file_path: "/var/log/atlassian-dc-mcp.log"
  
  # Log level for file output (optional)
  # If not specified, uses the level set for console output
  file_level: "debug"

jira:
  url: "https://jira.example.com"
  token: "your-jira-api-token"
  permissions:
    # Enable/disable write operations (default: false)
    write: false

confluence:
  url: "https://confluence.example.com"
  token: "your-confluence-api-token"
  permissions:
    # Enable/disable write operations (default: false)
    write: false

bitbucket:
  url: "https://bitbucket.example.com"
  token: "your-bitbucket-api-token"
  permissions:
    # Enable/disable write operations (default: false)
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
export MCP_LOGGING_FILE_PATH="/var/log/atlassian-dc-mcp.log"  # Path to log file
export MCP_LOGGING_FILE_LEVEL="debug"  # Log level for file output

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
MCP_TRANSPORT="stdio"
MCP_CLIENT_TIMEOUT=60

# Logging configuration
MCP_LOGGING_DEVELOPMENT=true
MCP_LOGGING_LEVEL="info"

# Jira configuration
MCP_JIRA_URL="https://jira.example.com"
MCP_JIRA_TOKEN="your-jira-api-token"
MCP_JIRA_PERMISSIONS_READ=true
MCP_JIRA_PERMISSIONS_WRITE=false

# Confluence configuration
MCP_CONFLUENCE_URL="https://confluence.example.com"
MCP_CONFLUENCE_TOKEN="your-confluence-api-token"
MCP_CONFLUENCE_PERMISSIONS_READ=true
MCP_CONFLUENCE_PERMISSIONS_WRITE=false

# Bitbucket configuration
MCP_BITBUCKET_URL="https://bitbucket.example.com"
MCP_BITBUCKET_TOKEN="your-bitbucket-api-token"
MCP_BITBUCKET_PERMISSIONS_READ=true
MCP_BITBUCKET_PERMISSIONS_WRITE=false
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
- `write`: Allows write operations (default: false)

Read operations are always available and cannot be disabled. Write operations must be explicitly enabled by setting `write: true` in the configuration.

When a service is configured with `write: false`, only read endpoints will be available.
When a service is configured with `write: true`, both read and write endpoints will be available.

Permissions are validated at startup and services with invalid configurations will be disabled.

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